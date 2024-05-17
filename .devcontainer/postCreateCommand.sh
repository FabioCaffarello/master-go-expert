#!/bin/sh

# Setting node
npm install || { echo "npm install failed"; exit 1; }

# Setting husky
npm run prepare || { echo "npm run prepare failed"; exit 1; }

# Setting golang
## Dependency Injection
go install github.com/google/wire/cmd/wire@latest || { echo "go install failed"; exit 1; }
## Documentation
go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest