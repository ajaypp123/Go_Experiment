# Microservice

## Web server

- `net/http` can be used to create web server and handler to handle website.
- Handler are event that will be triggered when anyone hits url.
- [Ref](https://golang.org/pkg/net/)
- Listen to server:
- `http.ListenAndServe("localhost:3000", handlers)

- Handle request:
    - `http.HandleFunc`
- Request:
    - read all parameter and data getting from request
    - `ioutil.ReadAll(r.Body)`
- Response:
    - `fprintf` is used to write data to response.

```
[root@node1 ~]# curl -d 'aa' localhost:3000
Response: aa
```

## ServeMux

- `ServeMux` Handle registers the handler for the given pattern.