# Web Server

The `webserver` library provides a simple and flexible HTTP server using the Chi router.

## Overview

The `webserver` package includes the following main components:
- `Server`: A struct that provides methods to configure and run an HTTP server with middleware and routing support.

## Features

The main functionalities provided by the package include:
- Setting up default middleware.
- Registering custom middleware.
- Adding routes and route groups.
- Starting the HTTP server.

## Types

- **Server**: Represents an HTTP server with a router and address.

## Functions

### Server Functions

- `NewWebServer(addr string) *Server`: Creates and returns a new `Server` instance with the specified address.
- `ConfigureDefaults()`: Sets up the default middleware for the server.
- `RegisterMiddlewares(middlewares ...func(http.Handler) http.Handler)`: Adds multiple middlewares to the server.
- `RegisterRoute(method, pattern string, handler http.HandlerFunc, group ...string)`: Adds a new route with an HTTP method, pattern, and handler function.
- `RegisterRouteGroup(prefix string, routes func(r chi.Router))`: Registers a group of routes under a common prefix.
- `Start() error`: Runs the web server on the specified address.

## Usage

### Creating a New Server

```go
import (
	"libs/services/infrastructure/server/http/webserver"
)

server := webserver.NewWebServer(":8080")
```

### Configuring Default Middleware

```go
server.ConfigureDefaults()
```

### Registering Custom Middleware

```go
server.RegisterMiddlewares(customMiddleware1, customMiddleware2)
```

### Adding a Route

```go
server.RegisterRoute("GET", "/hello", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
})
```

### Adding a Route Group

```go
server.RegisterRouteGroup("/api", func(r chi.Router) {
    r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("User endpoint"))
    })
    r.Post("/user", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Create user endpoint"))
    })
})
```

### Starting the Server

```go
err := server.Start()
if err != nil {
    log.Fatal(err)
}
```
