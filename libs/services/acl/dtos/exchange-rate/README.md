# exchange-rate
This library provides Data Transfer Objects (DTOs) for handling exchange rate information in a Go project. It defines structures for both input and output data formats.


## Packages
### input

The `input` package is designed to facilitate the transfer of data between different layers of an application, ensuring a clear and consistent structure for currency information from `exchange-rate` api in input context.


### output

The `output` package is designed to facilitate the transfer of data between different layers of an application, ensuring a clear and consistent structure for currency information from `exchange-rate` api in input context.


### Usage
Here's an example of how to use the DTOs in your application:

#### Input
```go
import (
	inputDTO "libs/services/acl/dtos/exchange-rate/input"
)

func processInput(data inputDTO.ExchangeRateDTO) {
    // Process the input data
    fmt.Println(data.Name)
}
```

#### Output
```go
import (
	outputDTO "libs/services/acl/dtos/exchange-rate/output"
)

func generateOutput() outputdto.ExchangeRateDTO {
    return outputdto.ExchangeRateDTO{
        Code:       "USD",
        CodeIn:     "BRL",
        Name:       "Dollar",
        High:       5.40,
        Low:        5.20,
        VarBid:     0.02,
        PctChange:  0.37,
        Bid:        5.30,
        Ask:        5.32,
        Timestamp:  1612300800,
        CreateDate: "2024-05-21",
    }
}

func displayOutput(data outputdto.ExchangeRatesDTO) {
    for _, rate := range data {
        fmt.Printf("Currency: %s, Bid: %f, Ask: %f\n", rate.Name, rate.Bid, rate.Ask)
    }
}

```

