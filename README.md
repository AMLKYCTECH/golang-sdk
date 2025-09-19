# AML/KYC Go SDK

Go client library for the AML/KYC API that provides cryptocurrency address risk assessment, fast checks, and investigations.

## Features

- **Balance Management**: Check account balance and create payment invoices
- **Fast Checks**: Quick risk assessment of cryptocurrency addresses
- **Investigations**: Comprehensive analysis with transaction graphs
- **Crypto Address Management**: Add and retrieve cryptocurrency addresses
- **Multi-Network Support**: Support for various blockchain networks (TRX, ETH, etc.)

## Installation

```bash
go get github.com/your-repo/amlkyc-sdk
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    
    amlkyc "github.com/your-repo/amlkyc-sdk"
)

func main() {
    // Initialize client with your API key
    client := amlkyc.NewClient("7500|V7Yz2LcScGI2HtfbGNc08e2huNI6XOv6ppR7tc7L")
    
    // Check your account balance
    balance, err := client.GetBalance()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Balance: %d, Available investigations: %d, Available fast checks: %d\n",
        balance.Balance, balance.Investigations, balance.FastChecks)
}
```

## Usage Examples

### Balance Management

```go
// Get current balance
balance, err := client.GetBalance()
if err != nil {
    log.Fatal(err)
}

// Get allowed fiat currencies for payment
currencies, err := client.GetAllowedFiatCurrencies()
if err != nil {
    log.Fatal(err)
}

// Create payment invoice
payment, err := client.CreatePayment("300", "rub")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Payment link: %s\n", payment.Link)
```

### Cryptocurrency Address Operations

```go
// Search supported cryptocurrencies for an address
address := "TUebrU7t87ZBonfseZLah2RQ56L2ECaXPj"
searchResult, err := client.SearchCryptocurrencies(address)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Supported networks: %v\n", searchResult.Items)

// Add a cryptocurrency address
cryptoAddr, err := client.AddCryptoAddress(address, "USDT", "trx")
if err != nil {
    log.Fatal(err)
}

// Get cryptocurrency address information
cryptoAddr, err = client.GetCryptoAddress(address, "USDT", "trx")
if err != nil {
    log.Fatal(err)
}
```

### Fast Checks

```go
// Create a fast check
fastCheck, err := client.CreateFastCheck("TUebrU7t87ZBonfseZLah2RQ56L2ECaXPj", "USDT", "trx")
if err != nil {
    log.Fatal(err)
}

// Run the fast check
runResult, err := client.RunFastCheck(fastCheck.ID)
if err != nil {
    log.Fatal(err)
}

// Get fast check results
result, err := client.GetFastCheck(fastCheck.ID)
if err != nil {
    log.Fatal(err)
}

if result.Risk != nil {
    fmt.Printf("Risk Score: %d\n", result.Risk.Score)
    for _, detail := range result.Risk.RiskDetail {
        fmt.Printf("Risk Type: %s, Percent: %d%%\n", detail.Type, detail.Percent)
    }
}

// List all fast checks
fastChecks, err := client.ListFastChecks()
if err != nil {
    log.Fatal(err)
}
```

### Investigations

```go
// Create an investigation (using address ID or transaction hash)
investigation, err := client.CreateInvestigation("1", "USDT")
if err != nil {
    log.Fatal(err)
}

// Run the investigation
runResult, err := client.RunInvestigation(investigation.ID)
if err != nil {
    log.Fatal(err)
}

// Get investigation details
result, err := client.GetInvestigation(investigation.ID)
if err != nil {
    log.Fatal(err)
}

// Get investigation graph data
graph, err := client.GetInvestigationGraph(investigation.ID, nil)
if err != nil {
    log.Fatal(err)
}

// Perform deep investigation on a specific address
deepResult, err := client.DeepInvestigation(investigation.ID, "TUebrU7t87ZBonfseZLah2RQ56L2ECaXPj")
if err != nil {
    log.Fatal(err)
}

// List all investigations
investigations, err := client.ListInvestigations()
if err != nil {
    log.Fatal(err)
}
```

### Complete Example

```go
package main

import (
    "fmt"
    "log"
    "time"
    
    amlkyc "github.com/your-repo/amlkyc-sdk"
)

func main() {
    // Initialize client
    client := amlkyc.NewClient("your-api-key-here")
    client.SetTimeout(60 * time.Second)
    
    address := "TUebrU7t87ZBonfseZLah2RQ56L2ECaXPj"
    currency := "USDT"
    network := "trx"
    
    // Step 1: Check what cryptocurrencies are supported for this address
    fmt.Println("Searching supported cryptocurrencies...")
    searchResult, err := client.SearchCryptocurrencies(address)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Supported networks: %v\n", searchResult.Items)
    
    // Step 2: Add address to the system
    fmt.Println("Adding address to system...")
    cryptoAddr, err := client.AddCryptoAddress(address, currency, network)
    if err != nil {
        // Address might already exist, try to get it
        cryptoAddr, err = client.GetCryptoAddress(address, currency, network)
        if err != nil {
            log.Fatal(err)
        }
    }
    fmt.Printf("Address ID: %d\n", cryptoAddr.ID)
    
    // Step 3: Perform a fast check
    fmt.Println("Creating fast check...")
    fastCheck, err := client.CreateFastCheck(address, currency, network)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Running fast check...")
    _, err = client.RunFastCheck(fastCheck.ID)
    if err != nil {
        log.Fatal(err)
    }
    
    // Wait for results
    time.Sleep(10 * time.Second)
    
    result, err := client.GetFastCheck(fastCheck.ID)
    if err != nil {
        log.Fatal(err)
    }
    
    if result.Risk != nil {
        fmt.Printf("Fast Check Results - Risk Score: %d\n", result.Risk.Score)
        for _, detail := range result.Risk.RiskDetail {
            fmt.Printf("  %s: %d%% risk\n", detail.Type, detail.Percent)
        }
    }
    
    if result.Info != nil {
        fmt.Printf("Address Info - Balance: %s, Total Transactions: %s\n", 
            result.Info.Balance, result.Info.TxsCount)
    }
}
```

## Configuration

### Custom Base URL

```go
client := amlkyc.NewClientWithURL("your-api-key", "https://custom.api.url")
```

### Custom Timeout

```go
client := amlkyc.NewClient("your-api-key")
client.SetTimeout(120 * time.Second)
```

## Error Handling

The SDK provides detailed error messages for different scenarios:

```go
balance, err := client.GetBalance()
if err != nil {
    // Handle API errors
    fmt.Printf("API Error: %v\n", err)
    return
}
```

## Testing

### Unit Tests
```bash
go test
```

### Integration Tests (requires valid API key)
```bash
# Set your API key
export AMLKYC_API_KEY="your-api-key-here"

# Run integration tests
go test -v
```

### Run specific integration test
```bash
go test -v -run TestIntegration_GetBalance
```

## API Reference

### Client Methods

#### Balance Operations
- `GetBalance() (*BalanceResponse, error)` - Get account balance
- `CreatePayment(amount, currency string) (*PaymentResponse, error)` - Create payment invoice
- `GetAllowedFiatCurrencies() ([]AllowedFiatCurrency, error)` - Get allowed currencies

#### Cryptocurrency Operations  
- `SearchCryptocurrencies(address string) (*CryptoSearchResponse, error)` - Search supported cryptocurrencies
- `GetCryptoAddress(address, currency, network string) (*CryptoAddress, error)` - Get address info
- `AddCryptoAddress(address, currency, network string) (*CryptoAddress, error)` - Add new address

#### Fast Check Operations
- `ListFastChecks() ([]FastCheck, error)` - List all fast checks
- `CreateFastCheck(address, currency, network string) (*FastCheck, error)` - Create fast check
- `GetFastCheck(id int) (*FastCheck, error)` - Get fast check details  
- `RunFastCheck(id int) (*RunStatusResponse, error)` - Run fast check

#### Investigation Operations
- `ListInvestigations() ([]Investigation, error)` - List all investigations
- `CreateInvestigation(target, currency string) (*Investigation, error)` - Create investigation
- `GetInvestigation(id int) (*Investigation, error)` - Get investigation details
- `RunInvestigation(id int) (*RunStatusResponse, error)` - Run investigation
- `GetInvestigationGraph(id int, itemID *int) (*InvestigationGraphResponse, error)` - Get graph data
- `DeepInvestigation(id int, address string) (*DeepInvestigationResponse, error)` - Deep investigation

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For support and questions, please contact the AML/KYC API support team.