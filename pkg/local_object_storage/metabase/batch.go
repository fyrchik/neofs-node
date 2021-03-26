package meta

import (
	"bytes"
	"sort"
	"time"

	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

type keyValue struct {
	key   []byte
	value []byte
}

type Batch struct {
	boltB   *bbolt.Bucket
	buckets map[string]*Batch
	kv      []keyValue
}

func (db *DB) newBatch() *Batch {
	b := new(Batch)
	b.buckets = map[string]*Batch{}
	return b
}

func (b *Batch) Get(k []byte) []byte {
	for i := len(b.kv) - 1; i >= 0; i-- {
		if bytes.Equal(b.kv[i].key, k) {
			return b.kv[i].value
		}
	}
	return nil
}

func (b *Batch) Put(k, v []byte) error {
	var exact bool
	index := sort.Search(len(b.kv), func(i int) bool {
		ret := bytes.Compare(b.kv[i].key, k)
		if ret == 0 {
			exact = true
		}
		return ret != -1
	})

	if index < len(b.kv) {
		if exact {
			b.kv[index].value = v
			return nil
		}
		b.kv = append(b.kv, keyValue{})
		copy(b.kv[index+1:], b.kv[index:])
		b.kv[index].key = k
		b.kv[index].value = v
		return nil
	}
	b.kv = append(b.kv, keyValue{key: k, value: v})
	return nil
}

func (b *Batch) bucket(name []byte) *Batch {
	sn := string(name)
	buk, ok := b.buckets[sn]
	if ok {
		return buk
	}
	buk = &Batch{
		buckets: make(map[string]*Batch),
	}
	b.buckets[sn] = buk
	return buk
}

func (b *Batch) Bucket(name []byte) *Batch {
	return b.bucket(name)
}

func (b *Batch) CreateBucketIfNotExists(name []byte) *Batch {
	buk := b.bucket(name)
	return buk
}

func (b *Batch) merge(a *Batch) *Batch {
	b.kv = mergeKV(b.kv, a.kv)
	for name, subB := range b.buckets {
		subA, ok := a.buckets[name]
		if ok {
			subB.merge(subA)
			delete(a.buckets, name)
		}
	}
	for name, subA := range a.buckets {
		b.buckets[name] = subA
	}
	return b
}

func mergeKV(a, b []keyValue) []keyValue {
	return append(a, b...)
}

func sortKV(a []keyValue) []keyValue {
	sort.Slice(a, func(i, j int) bool {
		return bytes.Compare(a[i].key, a[j].key) == -1
	})
	return a
}

func (db *DB) addBucket(b *Batch) {
	db.batchMtx.Lock()
	defer db.batchMtx.Unlock()
	db.batch = db.batch.merge(b)
}

func (db *DB) batchLoop() {
	tick := time.NewTicker(time.Second)
	defer tick.Stop()
	for range tick.C {
		db.putBatch()
	}
}

func (db *DB) putBatch() {
	db.batchMtx.Lock()
	b := db.batch
	db.batch = db.newBatch()
	db.batchMtx.Unlock()

	cnt := 0
	err := db.boltDB.Update(func(tx *bbolt.Tx) error {
		for name, sub := range b.buckets {
			buk, err := tx.CreateBucketIfNotExists([]byte(name))
			if err != nil {
				return err
			}
			if err := putBucket(buk, sub); err != nil {
				return err
			}
			cnt += len(sub.kv)
		}
		return nil
	})
	if err != nil {
		db.log.Info("batch put error", zap.Error(err))
	}
	db.log.Info("put keys", zap.Int("count", cnt))
}

func putBucket(b *bbolt.Bucket, bt *Batch) error {
	for name, sub := range bt.buckets {
		buk, err := b.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return err
		}
		bt.kv = sortKV(bt.kv)
		for i := range bt.kv {
			if err := buk.Put(bt.kv[i].key, bt.kv[i].value); err != nil {
				return err
			}
		}
		if err := putBucket(buk, sub); err != nil {
			return err
		}
	}
	return nil
}
