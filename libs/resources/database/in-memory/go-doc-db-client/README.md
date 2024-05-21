# Go doc DB Client

The `go-doc-db-client` library provides a convenient interface to interact with the in-memory document database.

## Overview

The client package includes the following main components:
- `Client`: A struct that provides methods to perform CRUD operations on collections and documents within the in-memory document database.

## Features

The main functionalities provided by the package include:
- Creating and dropping collections.
- Inserting, finding, updating, and deleting documents in collections.
- Listing all collections in the database.
- Querying documents in collections based on specific criteria.

## Types

- **Client**: Provides an interface to interact with the in-memory document database.

## Functions

### Client Functions

- `NewClient(db *database.InMemoryDocBD) *Client`: Creates and returns a new `Client` instance.
- `CreateCollection(collectionName string) error`: Creates a new collection with the given name.
- `DropCollection(collectionName string) error`: Drops a collection by its name.
- `ListCollections() []string`: Lists the names of all collections in the database.
- `ConvertToDocument(document map[string]interface{}) (database.Document, error)`: Converts a map to a `Document` type.
- `InsertOne(collectionName string, document map[string]interface{}) error`: Inserts a single document into the specified collection.
- `FindOne(collectionName string, id string) (map[string]interface{}, error)`: Finds and returns a single document by its ID from the specified collection.
- `FindAll(collectionName string) ([]map[string]interface{}, error)`: Returns all documents from the specified collection.
- `Find(collectionName string, filter map[string]interface{}) ([]map[string]interface{}, error)`: Returns documents matching the given query from the specified collection.
- `UpdateOne(collectionName string, id string, update map[string]interface{}) error`: Updates a single document by its ID with the given update in the specified collection.
- `DeleteOne(collectionName string, id string) error`: Deletes a single document by its ID from the specified collection.
- `DeleteAll(collectionName string) error`: Deletes all documents from the specified collection.

## Usage

### Creating a New Client

```go
import (
	"libs/resources/database/in-memory/go-doc-db"
)

db := database.NewInMemoryDocBD("myDatabase")
client := client.NewClient(db)
```

### Creating a New Collection

```go
err := client.CreateCollection("myCollection")
if err != nil {
    log.Fatal(err)
}
```

### Inserting a Document

```go
document := map[string]interface{}{
    "_id": "12345",
    "name": "John Doe",
    "age": 30,
}

err = client.InsertOne("myCollection", document)
if err != nil {
    log.Fatal(err)
}
```

### Finding a Document by ID

```go
doc, err := client.FindOne("myCollection", "12345")
if err != nil {
    log.Fatal(err)
}
fmt.Println(doc)
```

### Finding All Documents

```go
docs, err := client.FindAll("myCollection")
if err != nil {
    log.Fatal(err)
}
fmt.Println(docs)
```

### Updating a Document

```go
update := map[string]interface{}{
    "age": 31,
}
err = client.UpdateOne("myCollection", "12345", update)
if err != nil {
    log.Fatal(err)
}
```

### Deleting a Document

```go
err = client.DeleteOne("myCollection", "12345")
if err != nil {
    log.Fatal(err)
}
```

### Dropping a Collection

```go
err = client.DropCollection("myCollection")
if err != nil {
    log.Fatal(err)
}
```
