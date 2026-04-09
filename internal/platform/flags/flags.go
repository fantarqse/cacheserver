package flags

import (
	"flag"
	"time"
)

type HTTPServer struct {
	Port int
}

type Cache struct {
	Address         string
	Port            int
	MaxCountOfItems int
	MaxItemSize     int
	GlobalCacheSize int
	TTL             time.Duration
}

type Logger struct{}

type Flags struct {
	HTTPServer HTTPServer
	Cache      Cache
	Logger     Logger
}

func Parse() Flags {
	var fs Flags

	// HTTP Server
	flag.IntVar(&fs.HTTPServer.Port, "http-port", 8080, "a port of an http server")

	// Cache
	flag.StringVar(&fs.Cache.Address, "cache-address", "localhost", "an address of redis")
	flag.IntVar(&fs.Cache.Port, "cache-port", 6379, "a port of redis")
	flag.IntVar(&fs.Cache.MaxCountOfItems, "cache-count-of-items", 100, "a maximum number of items in the storage")
	flag.IntVar(&fs.Cache.MaxItemSize, "cache-item-size", 100, "a maximum size of an item in the storage (in bytes)")
	flag.IntVar(&fs.Cache.GlobalCacheSize, "cache-global-size", 10000, "a maximum size of all items in the storage (in bytes)")
	flag.DurationVar(&fs.Cache.TTL, "cache-ttl", 1*time.Hour, "TTL")

	flag.Parse()
	return fs
}
