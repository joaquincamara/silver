  <img  src="https://user-images.githubusercontent.com/26718123/127079860-deadbdab-1d60-49f7-9591-767e9eac48f8.png">

Silver is a minimalist router for building Go HTTP services.

This project started not as a way to reinvent the wheel but to know how the wheel works. The main goal is to create a complete, idiomatic and elegant Router to design and write REST API servers.

# Install

`go get github.com/joaquincamara/silver`

# usage

```golang
package main

import (
"log"
"net/http"
"github.com/joaquincamara/silver"
)

func main() {
      router := silver.AlchemyDoor()

      router.Use(middlewares.Recovery)

      router.GET("/", silver.AlchemyDoor)
      silver.Start("8080", router)
}
```

**Silver HTTP methods**

- GET: `router.GET("/MyRoute")`

- PUT: `router.PUT("/MyRoute")`

- DELETE: `router.DELETE("/MyRoute")`

- PUT: `router.PUT("/MyRoute")`

- PATCH: `router.PATCH("/MyRoute")`

# Dev goals

1. Add a middleware with the next feaures:

| Middleware handler | Description                                                             | Status |
| ------------------ | ----------------------------------------------------------------------- | ------ |
| BasicAuth          | Basic HTTP authentication                                               | ------ |
| Cors               | Can be used to enable CORS with various options.                        | ------ |
| Logger             | Logs the start and end of each request with the elapsed processing time | ------ |
| RealIP             | Sets a http.Request's RemoteAddr to either X-Real-IP or X-Forwarded-For | ------ |
| Recovery           | Gracefully absorb panics and prints the stack trace                     | Done   |



======================

**silver.go version: 0.1.1**

======================
