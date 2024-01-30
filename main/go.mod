module icovn.dev/proxy_main

go 1.21.6

replace icovn.dev/client => ../client

replace icovn.dev/server => ../server

replace icovn.dev/network => ../network

require icovn.dev/network v0.0.0-00010101000000-000000000000

require (
	github.com/Allenxuxu/gev v0.5.0 // indirect
	github.com/Allenxuxu/ringbuffer v0.0.11 // indirect
	github.com/Allenxuxu/toolkit v0.0.1 // indirect
	github.com/RussellLuo/timingwheel v0.0.0-20220218152713-54845bda3108 // indirect
	github.com/libp2p/go-reuseport v0.4.0 // indirect
	golang.org/x/example/hello v0.0.0-20231013143937-1d6d2400d402
	golang.org/x/sys v0.16.0 // indirect
)
