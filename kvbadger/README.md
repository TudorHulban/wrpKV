# Interacting with KV stores: Badger DB
## How to use
Create a logger:
```go
l := log.NewLogger(log.DEBUG, os.Stderr, true)
```
Create a in memory store:
```go
inmemStore, err := NewBStoreInMem(l)
```
Close the store on exit:
```go
defer inmemStore.Close()
```
## Resources
```
https://github.com/dgraph-io/badger
```