package blobovnicza

import (
	"bytes"
	"fmt"
	"sort"
	"time"

	objectSDK "github.com/nspcc-dev/neofs-api-go/pkg/object"
	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

// PutPrm groups the parameters of Put operation.
type PutPrm struct {
	addr *objectSDK.Address

	objData []byte
}

// PutRes groups resulting values of Put operation.
type PutRes struct {
}

// ErrFull is returned returned when trying to save an
// object to a filled blobovnicza.
var ErrFull = errors.New("blobovnicza is full")

var errNilAddress = errors.New("object address is nil")

// SetAddress sets address of saving object.
func (p *PutPrm) SetAddress(addr *objectSDK.Address) {
	p.addr = addr
}

// SetMarshaledObject sets binary representation of the object.
func (p *PutPrm) SetMarshaledObject(data []byte) {
	p.objData = data
}

// Put saves object in Blobovnicza.
//
// If binary representation of the object is not set,
// it is calculated via Marshal method.
//
// The size of the object MUST BE less that or equal to
// the size specified in WithObjectSizeLimit option.
//
// Returns any error encountered that
// did not allow to completely save the object.
//
// Returns ErrFull if blobovnicza is filled.
func (b *Blobovnicza) Put(prm *PutPrm) (*PutRes, error) {
	addr := prm.addr
	if addr == nil {
		return nil, errNilAddress
	}

	if b.full() {
		return nil, ErrFull
	}

	// calculate size
	sz := uint64(len(prm.objData))
	name := string(bucketForSize(sz))

	b.cacheMtx.Lock()
	old := b.cache[name]
	old = append(old, keyValue{
		key:   addressKey(addr),
		value: prm.objData,
		size:  sz,
	})
	b.cache[name] = old
	b.cacheMtx.Unlock()

	return nil, nil
}

func (b *Blobovnicza) batchLoop() {
	timer := time.NewTimer(time.Second * 4)
	for {
		select {
		case <-timer.C:
			if err := b.putBatch(); err != nil {
				b.log.Error("error on batch put", zap.Error(err))
			}
			timer.Stop()
			timer = time.NewTimer(time.Second * 4)
		}
	}
}

func (b *Blobovnicza) putBatch() error {
	b.cacheMtx.Lock()
	cache := b.cache
	b.cache = make(map[string][]keyValue)
	b.cacheMtx.Unlock()

	if len(cache) == 0 {
		return nil
	}

	for _, kv := range cache {
		sort.Slice(kv, func(i, j int) bool {
			ret := bytes.Compare(kv[i].key, kv[j].key)
			return ret == -1
		})
	}

	var retErr error
	sz := uint64(0)
	count := 0
	err := b.boltDB.Update(func(tx *bbolt.Tx) error {
		for name, kv := range cache {
			buck := tx.Bucket([]byte(name))
			if buck == nil {
				return fmt.Errorf("no bucket present: %x", name)
			}

			for i := range kv {
				err := buck.Put(kv[i].key, kv[i].value)
				if err != nil && retErr == nil {
					retErr = err
				} else if err == nil {
					sz += kv[i].size
					count += 1
				}
			}
		}
		return nil
	})

	b.log.Info("put in blobovnicza",
		zap.Uint64("size", sz),
		zap.Int("count", count))
	b.incSize(sz)

	if retErr == nil {
		retErr = err
	}
	return retErr
}

func addressKey(addr *objectSDK.Address) []byte {
	return []byte(addr.String())
}
