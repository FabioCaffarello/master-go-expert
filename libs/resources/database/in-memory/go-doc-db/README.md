# Go Doc DB

The `go-doc-db` library provides an in-memory document database with support for basic CRUD operations and simple query matching.

## Overview

This package includes the following main components:
- `DocumentID`: A type alias for a string representing a unique document identifier.
- `Document`: A type alias for a map representing a document with string keys and `interface{}` values.
- `Collection`: A struct representing a collection of documents with thread-safe operations.
- `InMemoryDocBD`: A struct representing an in-memory document database containing multiple collections.

## Features

The main functionalities provided by the package include:
- Creating a new collection.
- Inserting, finding, updating, and deleting documents in a collection.
- Listing all documents in a collection.
- Querying documents in a collection based on specific criteria.
- Managing collections in an in-memory document database.

## Types

- **DocumentID**: Represents the unique identifier for a document.
- **Document**: Represents a document as a map with string keys and `interface{}` values.

## Functions

### Collection Functions

- `NewCollection`: Creates and returns a new `Collection` instance.
- `InsertOne(document Document) error`: Inserts a single document into the collection.
- `FindOne(id string) (Document, error)`: Finds and returns a single document by its ID.
- `FindAll() []Document`: Returns all documents in the collection.
- `Find(query map[string]interface{}) []Document`: Finds and returns documents matching the given query.
- `DeleteOne(id string) error`: Deletes a single document by its ID.
- `UpdateOne(id string, update Document) error`: Updates a single document by its ID with the given update.
- `DeleteAll() error`: Deletes all documents in the collection.

### Utility Functions

- `matchesQuery(document, query map[string]interface{}) bool`: Checks if a document matches the query criteria.

### InMemoryDocBD Functions

- `NewInMemoryDocBD(name string) *InMemoryDocBD`: Creates and returns a new `InMemoryDocBD` instance with the given name.
- `GetCollection(collectionName string) (*Collection, error)`: Retrieves a collection by its name.
- `CreateCollection(collectionName string) error`: Creates a new collection with the given name.
- `DropCollection(collectionName string) error`: Drops a collection by its name.
- `ListCollections() []string`: Lists the names of all collections in the database.

## Usage
### Creating a New Database

```go

import (
	"libs/resources/database/in-memory/go-doc-db/database"
)

db := database.NewInMemoryDocBD("myDatabase")
```

### Creating a New Collection
```go
err := db.CreateCollection("myCollection")
if err != nil {
    log.Fatal(err)
}
```
### Inserting a Document
```go
collection, err := db.GetCollection("myCollection")
if err != nil {
    log.Fatal(err)
}

document := database.Document{
    "_id": "12345",
    "name": "John Doe",
    "age": 30,
}

err = collection.InsertOne(document)
if err != nil {
    log.Fatal(err)
}
```

### Finding a Document by ID
```go
doc, err := collection.FindOne("12345")
if err != nil {
    log.Fatal(err)
}
fmt.Println(doc)
```

### Finding All Documents
```go
update := database.Document{
    "age": 31,
}
err = collection.UpdateOne("12345", update)
if err != nil {
    log.Fatal(err)
}
```

### Deleting a Document

```go
err = collection.DeleteOne("12345")
if err != nil {
    log.Fatal(err)
}
```

### Dropping a Collection
```go
err = db.DropCollection("myCollection")
if err != nil {
    log.Fatal(err)
}
``` 
