package flags

import (
	"flag"
	"time"
)

type HTTPServer struct {
	Port int
}

type Redis struct {
	Address string
	Port    int
}

type Cache struct {
	MaxCountOfItems int
	MaxItemSize     int
	GlobalCacheSize int
	TTL             time.Duration
}

type Logger struct{}

type Flags struct {
	HTTPServer HTTPServer
	Redis      Redis
	Cache      Cache
	Logger     Logger
}

func Parse() Flags {
	var fs Flags

	// HTTP Server
	flag.IntVar(&fs.HTTPServer.Port, "http-port", 8080, "a port of an http server")

	// Redis
	flag.StringVar(&fs.Redis.Address, "redis-address", "localhost", "an address of redis")
	flag.IntVar(&fs.Redis.Port, "redis-port", 6379, "a port of redis")

	// Cache/Service
	flag.IntVar(&fs.Cache.MaxCountOfItems, "cache-count-of-items", 100, "a maximum number of items in the storage")
	flag.IntVar(&fs.Cache.MaxItemSize, "cache-item-size", 100, "a maximum size of an item in the storage (in bytes)")
	flag.IntVar(&fs.Cache.GlobalCacheSize, "cache-global-size", 10000, "a maximum size of all items in the storage (in bytes)")
	flag.DurationVar(&fs.Cache.TTL, "cache-ttl", 1*time.Hour, "TTL")

	flag.Parse()
	return fs
}
