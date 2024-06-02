# go-uuid

`go-uuid` is a Go package designed to generate unique identifiers (UUIDs) based on provided properties. This package serializes properties to JSON, hashes the resulting byte slice, and generates a UUID from the hash.

## Usage

Here's a basic example of how to use the `go-uuid` package to generate a unique ID:

```go
package main

import (
	"fmt"
	"log"
	gouuid "libs/shared/go-uuid"
)

func main() {
	properties := map[string]interface{}{
		"name": "example",
		"value": 12345,
	}

	id, err := gouuid.GetID(properties)
	if err != nil {
		log.Fatalf("Error generating ID: %v", err)
	}

	fmt.Printf("Generated ID: %s\n", id)
}
```
