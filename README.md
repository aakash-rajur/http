# http

[![Test](https://github.com/aakash-rajur/http/actions/workflows/test.yml/badge.svg)](https://github.com/aakash-rajur/sqlxgen/actions/workflows/test.yml)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/aakash-rajur/sqlxgen/main/LICENSE.md)

> yet another http routing library

## goal

1. provide minimum surface area for functioning a http library
2. understanding the underlying mechanisms of http routing
3. evaluate alternative approaches to various existing solutions
4. provide model for custom http routines

## usage

### routing
```go
import (
  h "github.com/aakash-rajur/http"
  "github.com/aakash-rajur/http/params"
  "net/http"
)

router := h.NewRouter()

router.GetFunc(
  "/health",
  func (w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    
    _, _ = w.Write([]byte("OK"))
  },
)

router.GetFunc(
  "/api/v2/books/:id",
  func(w http.ResponseWriter, r *http.Request) {
    p, ok := params.FromContext(r.Context())
    
    if !ok {
      http.Error(w, "unable to parse param", http.StatusInternalServerError)
      
      return
    }
    
    idString := p.Get("id", "")
    
    id, err := strconv.Atoi(idString)
    
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      
      return
    }
    
    book := books[id-1]
    
    buffer, err := json.Marshal(book)
    
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      
      return
    }
    
    w.Header().Set("Content-Type", "application/json")
    
    w.WriteHeader(http.StatusOK)
    
    _, _ = w.Write(buffer)
  },
)

server := &http.Server{
  Addr:    address,
  Handler: router,
}
```

### middleware
```go
import (
  h "github.com/aakash-rajur/http"
)

router := h.NewRouter()

router.Use(h.Logger(h.LoggerConfig{}))

router.Use(
  func(w http.ResponseWriter, r *http.Request, next h.Next) {
    next(r)
  },
)
```

### logging
```go
import (
  h "github.com/aakash-rajur/http"
)

router := h.NewRouter()

router.Use(h.Logger(h.LoggerConfig{}))

router.GetFunc(
  "/api/v2/books",
  func (w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    
    _, _ = w.Write([]byte("OK"))
  },
)
```

```shell
2024-01-07 14:21:33 | HTTP/2 |  200 |  237.083µs |                     text/plain |               application/json |       127.0.0.1 |     GET /api/v2/books 
2024-01-07 14:21:34 | HTTP/2 |  404 |   71.666µs |                     text/plain |      text/plain; charset=utf-8 |       127.0.0.1 |     GET /favicon.ico 
2024-01-07 14:21:36 | HTTP/2 |  200 |   46.125µs |                     text/plain |               application/json |       127.0.0.1 |     GET /api/v2/books 
2024-01-07 14:21:37 | HTTP/2 |  200 |  156.541µs |                     text/plain |               application/json |       127.0.0.1 |     GET /api/v2/books 
2024-01-07 14:21:37 | HTTP/2 |  200 |  105.125µs |                     text/plain |               application/json |       127.0.0.1 |     GET /api/v2/books 
2024-01-07 14:21:37 | HTTP/2 |  200 |   80.167µs |                     text/plain |               application/json |       127.0.0.1 |     GET /api/v2/books 
2024-01-07 14:21:38 | HTTP/2 |  200 |   78.458µs |                     text/plain |               application/json |       127.0.0.1 |     GET /api/v2/books 
2024-01-07 14:21:38 | HTTP/2 |  200 |   85.083µs |                     text/plain |               application/json |       127.0.0.1 |     GET /api/v2/books 
```

#### custom logger
```go
import (
  h "github.com/aakash-rajur/http"
)

router := h.NewRouter()

router.Use(
  h.Logger(
    h.LoggerConfig{
      LogFormatter: func(formatterParams h.LogFormatterParams) string {
        jsonPayload, err := json.Marshal(formatterParams)

        if err != nil {
          return ""
        }

        return fmt.Sprintf("%s\n", jsonPayload)
      },
    },
  ),
)

router.GetFunc(
  "/api/v2/books",
  func (w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    
    _, _ = w.Write([]byte("OK"))
  },
)
```

```shell
{"timestamp":"2024-01-07T14:34:19.03845+05:30","status_code":200,"latency":384750,"client_ip":"127.0.0.1","method":"GET","path":"/api/v2/books","query":{},"request_content_type":"text/plain","request_content_encoding":"identity","response_content_type":"application/json","response_content_encoding":"identity","protocol_version":2}
{"timestamp":"2024-01-07T14:34:20.315136+05:30","status_code":200,"latency":86334,"client_ip":"127.0.0.1","method":"GET","path":"/api/v2/books","query":{},"request_content_type":"text/plain","request_content_encoding":"identity","response_content_type":"application/json","response_content_encoding":"identity","protocol_version":2}
{"timestamp":"2024-01-07T14:34:20.871688+05:30","status_code":200,"latency":146875,"client_ip":"127.0.0.1","method":"GET","path":"/api/v2/books","query":{},"request_content_type":"text/plain","request_content_encoding":"identity","response_content_type":"application/json","response_content_encoding":"identity","protocol_version":2}
{"timestamp":"2024-01-07T14:34:23.677465+05:30","status_code":200,"latency":29875,"client_ip":"127.0.0.1","method":"GET","path":"/api/v2/books","query":{},"request_content_type":"text/plain","request_content_encoding":"identity","response_content_type":"application/json","response_content_encoding":"identity","protocol_version":2}
```

## mechanism

> search for incoming path in a pre-sorted list of registered routes using binary search costing O(log n) time
> where n is the number of registered routes

### registration
1. while registering we've always receive a path pattern and a method (along with handler).
2. create a new pattern of form `/method/pattern` and run sanity through the same.
3. split the pattern into segments (array of strings)
4. append segments to a list of segments
5. sort this list of segments with static segments taking precedence over path params.

### matching
1. prepend the method to the incoming path and run sanity through the same.
2. split the path into segments (array of strings)
3. search for the segments in the list of segments using binary search.
4. path params are matched to true for all values in that corresponding segment position.

## alternatives

### trie
1. a tree like data structure that is used to store strings.
2. more difficult to implement than a simple list of segments.
3. similar runtime complexity as a list of segments O(log n) where n is the number of registered routes.

### linear search
1. search for incoming path in a list of registered routes.
2. runtime complexity of O(n) where n is the number of registered routes.