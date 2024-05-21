# go-request
`go-request` is a lightweight and flexible HTTP client for Go. It allows you to create and send HTTP requests with ease, providing support for various content types and custom headers.

## Features

- Build and send HTTP requests with custom headers, path parameters, and query parameters.
- Support for JSON, XML, and URL-encoded form bodies.
- Timeout handling for HTTP requests.
- Easy-to-use API with context support for request cancellation.

## Usage

Here's an example of how to use the `go-request` package to send an HTTP request:

```go
package main

import (
    "context"
    "fmt"
    "net/http"
    "time"

    gorequest "libs/shared/go-request"
)

func main() {
    client := &http.Client{}
    ctx := context.Background()

    headers := map[string]string{
        "Content-Type": "application/json",
    }

    queryParams := map[string]string{
        "query": "value",
    }

    body := map[string]interface{}{
        "key": "value",
    }

    req, err := gorequest.CreateRequest(ctx, "https://api.example.com", nil, queryParams, body, headers, "POST")
    if err != nil {
        fmt.Printf("Failed to create request: %v\n", err)
        return
    }

    var result map[string]interface{}
    err = gorequest.SendRequest(ctx, req, client, &result, 10*time.Second)
    if err != nil {
        fmt.Printf("Failed to send request: %v\n", err)
        return
    }

    fmt.Printf("Response: %+v\n", result)
}
```
