# Go Server Comparison

This is a comparison of Go web servers. The comparison is based on the following criteria:

- Creating a simple web server
- Routing
- Middleware
- Error handling
- Swagger generation
- Validation

The following web servers are compared:

- [StdLib](https://golang.org/pkg/net/http/)
- [Gin](https://github.com/gin-gonic/gin)
- [Fiber](https://github.com/gofiber/fiber)
- [WebRPC](https://github.com/webrpc/webrpc)

All three examples use swaggo for swagger generation:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

and generate the swagger documentation with:

```bash
swag init
```
