# economia-awesome-api

This library provides Data Transfer Objects (DTOs) for handling `economia-awesome-api` information in a Go project. It defines structures for output data formats.

## Packages

### output

The `output` package is designed to facilitate the transfer of data between different layers of an application, ensuring a clear and consistent structure for currency information from `economia-awesome-api` in output context.

## Usage

### CurrencyInfoDTO

The `CurrencyInfoDTO` struct contains detailed information about an exchange rate. Each field is mapped to a JSON key, allowing easy serialization and deserialization.

```go
package main

import (
    "fmt"
    outputDTO "libs/services/acl/dtos/economia-awesome-api/output"
)

func main() {
    currencyInfo := outputDTO.CurrencyInfoDTO{
        Code:       "USD",
        CodeIn:     "BRL",
        Name:       "Dollar",
        High:       "5.30",
        Low:        "5.20",
        VarBid:     "0.05",
        PctChange:  "0.95",
        Bid:        "5.25",
        Ask:        "5.26",
        Timestamp:  "1625164800",
        CreateDate: "2021-07-01T12:00:00Z",
    }
    fmt.Printf("%+v\n", currencyInfo)
}
```

### CurrencyInfoMapDTO

The `CurrencyInfoMapDTO` is a map where the keys are currency codes and the values are `CurrencyInfoDTO` objects. This structure is useful for storing and retrieving exchange rate information for multiple currencies.

```go
package main

import (
    "fmt"
    outputDTO "libs/services/acl/dtos/economia-awesome-api/output"
)

func main() {
    currencyInfoMap := outputDTO.CurrencyInfoMapDTO{
        "USDBRL": outputDTO.CurrencyInfoDTO{
			Code:       "USD",
			CodeIn:     "BRL",
			Name:       "Dollar",
			High:       "5.30",
			Low:        "5.20",
			VarBid:     "0.05",
			PctChange:  "0.95",
			Bid:        "5.25",
			Ask:        "5.26",
			Timestamp:  "1625164800",
			CreateDate: "2021-07-01T12:00:00Z",
		},
        "EURBRL": outputDTO.CurrencyInfoDTO{
            Code:       "EUR",
            CodeIn:     "BRL",
            Name:       "Euro",
            High:       "6.20",
            Low:        "6.10",
            VarBid:     "0.03",
            PctChange:  "0.50",
            Bid:        "6.15",
            Ask:        "6.16",
            Timestamp:  "1625164800",
            CreateDate: "2021-07-01T12:00:00Z",
        },
    }
    fmt.Printf("%+v\n", currencyInfoMap)
}
```
