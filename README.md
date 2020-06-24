# webfinger

Webfinger server & client library for Go

## Client

```go
u := "https://example.com" + webfinger.DefaultPath

q := webfinger.NewQuery("https://blog.example.com/article/id/314")

c, err := webfinger.NewClient(u)
if err != nil {
    panic(err)
}

r, err := c.Query(q)
if err != nil {
    panic(err)
}

// do something with r
```

## Server

```go
// implement DB interface or use webfinger.MemDB
func getDB() webfinger.DB {
    ...
}

h := webfinger.NewHandler(getDB(), webfinger.WithAllowOrigin("*"))

http.Handle(webfinger.DefaultPath, h)

http.ListenAndServe(":8080", nil)
```