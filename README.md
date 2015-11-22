# nchash
Nginx consistent hash golang version  
If you are using nginx consistent hash in upstream, this would give you the same result as nginx consistent hash

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

nc := nchash.New(nodes)
server := nc.Get("test")
```
