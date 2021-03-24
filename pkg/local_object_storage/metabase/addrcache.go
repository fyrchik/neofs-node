package meta

import (
	lru "github.com/hashicorp/golang-lru"
	"github.com/hashicorp/golang-lru/simplelru"
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/pkg/object"
)

type addrCache struct {
	simplelru.LRUCache
}

var cache = &addrCache{}

func init() {
	c, err := lru.New(1000)
	if err != nil {
		panic(err)
	}
	cache.LRUCache = c
}

func getOID(id *object.ID) string {
	oid := string(id.ToV2().GetValue())
	if res, ok := cache.Get(oid); ok {
		return res.(string)
	}

	res := id.String()
	cache.Add(oid, res)
	return res
}

func getCID(id *container.ID) string {
	cid := string(id.ToV2().GetValue())
	if res, ok := cache.Get(cid); ok {
		return res.(string)
	}

	res := id.String()
	cache.Add(cid, res)
	return res
}
