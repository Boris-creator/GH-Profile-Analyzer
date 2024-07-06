```
cd proto
export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go_out=. --go-grpc_out=.  gh.proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative
```

```
cd analyzer
go generate .
```

```
cd server
air
```

```
cd service
air
```