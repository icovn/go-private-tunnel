module icovn.dev/proxy_main

go 1.21.6

replace icovn.dev/client => ../client

replace icovn.dev/server => ../server

replace icovn.dev/network => ../network

require icovn.dev/network v0.0.0-00010101000000-000000000000

require golang.org/x/example/hello v0.0.0-20231013143937-1d6d2400d402 // indirect
