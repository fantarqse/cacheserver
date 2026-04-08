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

// TODO: config?
// I am considering using flags to configure the entire app.
/*
	- http server
		- port
	- storage
		- address, password, etc.
		- max count of items
		- size of one item in bytes
		- size of all items in bytes
		- TTL
	- logger
*/
func Parse() Flags {
	var fs Flags

	// HTTP Server
	flag.IntVar(&fs.HTTPServer.Port, "http-port", 8080, "a port of an http server")

	// Cache
	flag.StringVar(&fs.Cache.Address, "cache-address", "localhost", "an address of redis")
	flag.IntVar(&fs.Cache.Port, "cache-port", 6379, "a port of redis")

	flag.Parse()
	return fs
}
