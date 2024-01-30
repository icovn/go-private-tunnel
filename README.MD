## Business flows

### Main flow
- Step 1: Server listen and wait for TCP connections on port XXXX, HTTP connections on port YYYY
- Step 2: Client connect to server over TCP on port XXXX
- Step 3: Server accept the connection, and register a sub-domain that mapped with the connection
- Step 4: User send a HTTP request to the sub-domain
- Step 5: Server forward the HTTP request to the client through th mapped connection
- Step 6: Client processes the HTTP request then send the reponse to server
- Step 7: Server forward the HTTP response to user

## How to run
Start demo server
```shell
cd main
go run demo.go
```

Connect to TCP server
```shell
telnet localhost 1833
```

Connect to HTTP server
```shell
http://localhost:6060/abcHowAreYou
```

## Common commands

Init a module
```shell
go init module icovn.dev/web
```

Add a module to the workspace
```shell
go work use ./web
```

Run a main module
```shell
go run ./web
```

Build
```shell
go build
```

## Reference
- https://github.com/Allenxuxu/gev
