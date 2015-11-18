# nchash
nginx consistent hash golang version
# Using

Importing ::

```go
import "github.com/bydsky/nchash"
```

Basic example usage ::

```go
nodes := []string{"192.168.0.1:8888",
                  "192.168.0.2:8888",
                  "192.168.0.3:8888"}

nc := nchash.New(nods)
server := nc.Get("test")
```
