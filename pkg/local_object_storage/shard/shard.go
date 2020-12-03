package shard

import (
	"github.com/nspcc-dev/neofs-node/pkg/local_object_storage/blobstor"
	meta "github.com/nspcc-dev/neofs-node/pkg/local_object_storage/metabase/v2"
	"github.com/nspcc-dev/neofs-node/pkg/util/logger"
	"go.uber.org/atomic"
)

// Shard represents single shard of NeoFS Local Storage Engine.
type Shard struct {
	*cfg

	mode *atomic.Uint32

	writeCache *blobstor.BlobStor

	blobStor *blobstor.BlobStor

	metaBase *meta.DB
}

// Option represents Shard's constructor option.
type Option func(*cfg)

type cfg struct {
	useWriteCache bool

	info Info

	blobOpts []blobstor.Option

	metaOpts []meta.Option

	writeCacheOpts []blobstor.Option

	log *logger.Logger
}

func defaultCfg() *cfg {
	return new(cfg)
}

// New creates, initializes and returns new Shard instance.
func New(opts ...Option) *Shard {
	c := defaultCfg()

	for i := range opts {
		opts[i](c)
	}

	var writeCache *blobstor.BlobStor

	if c.useWriteCache {
		writeCache = blobstor.New(
			append(c.blobOpts, c.writeCacheOpts...)...,
		)
	}

	return &Shard{
		cfg:        c,
		mode:       atomic.NewUint32(0), // TODO: init with particular mode
		blobStor:   blobstor.New(c.blobOpts...),
		metaBase:   meta.New(c.metaOpts...),
		writeCache: writeCache,
	}
}

// WithID returns option to set shard identifier.
func WithID(id *ID) Option {
	return func(c *cfg) {
		c.info.ID = id
	}
}

// WithBlobStorOptions returns option to set internal BlobStor options.
func WithBlobStorOptions(opts ...blobstor.Option) Option {
	return func(c *cfg) {
		c.blobOpts = opts
	}
}

// WithMetaBaseOptions returns option to set internal metabase options.
func WithMetaBaseOptions(opts ...meta.Option) Option {
	return func(c *cfg) {
		c.metaOpts = opts
	}
}

// WithMetaBaseOptions returns option to set internal metabase options.
func WithWriteCacheOptions(opts ...blobstor.Option) Option {
	return func(c *cfg) {
		c.writeCacheOpts = opts
	}
}

// WithLogger returns option to set Shard's logger.
func WithLogger(l *logger.Logger) Option {
	return func(c *cfg) {
		c.log = l
	}
}

// WithWriteCache returns option to toggle write cache usage.
func WithWriteCache(use bool) Option {
	return func(c *cfg) {
		c.useWriteCache = use
	}
}

// hasWriteCache returns bool if write cache exists on shards.
func (s Shard) hasWriteCache() bool {
	return s.cfg.useWriteCache
}
