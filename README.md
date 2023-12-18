# http
> yet another http routing library

## goal

1. provide minimum surface area for functioning http library
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

