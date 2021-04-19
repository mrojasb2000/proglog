# Proglog Project

### Running server 
```
$ go run cmd/server/main.go
```

### Testing functional

1. Create new record on log register
```
$ curl -X POST localhost:9090 -d '{"record":{"value":"QWERT"}}'
```

2. Read record from log register
```
$ curl -X GET localhost:9090 -d '{"offset":2}'
```
