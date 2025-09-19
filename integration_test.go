package golang

import (
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

const (
	// Test data
	testAddress  = "TUebrU7t87ZBonfseZLah2RQ56L2ECaXPj"
	testCurrency = "USDT"
	testNetwork  = "trx"
)

// Integration tests - these will make real API calls
// Run with: go test -tags=integration

func getIntegrationClient(t *testing.T) *Client {
	apiKey := os.Getenv("AMLKYC_API_KEY")
	if apiKey == "" {
		apiKey = testAPIKey // fallback to hardcoded key
	}

	client := NewClient(apiKey)
	client.SetTimeout(60 * time.Second) // Longer timeout for integration tests
	return client
}

func TestIntegration_GetBalance(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := getIntegrationClient(t)

	balance, err := client.GetBalance()
	if err != nil {
		t.Fatalf("GetBalance failed: %v", err)
	}

	if balance == nil {
		t.Fatal("Balance response is nil")
	}

	t.Logf("Balance: %.2f, Investigations: %d, Fast Checks: %d",
		balance.Balance, balance.Investigations, balance.FastChecks)
}

func TestIntegration_GetAllowedFiatCurrencies(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := getIntegrationClient(t)

	currencies, err := client.GetAllowedFiatCurrencies()
	if err != nil {
		t.Fatalf("GetAllowedFiatCurrencies failed: %v", err)
	}

	if len(currencies) == 0 {
		t.Log("No allowed fiat currencies returned")
	} else {
		t.Logf("Found %d allowed fiat currencies", len(currencies))
		for _, currency := range currencies {
			t.Logf("Currency: %s, Min: %.2f, Max: %.2f",
				currency.Name, currency.Min, currency.Max)
		}
	}
}

func TestIntegration_SearchCryptocurrencies(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := getIntegrationClient(t)

	result, err := client.SearchCryptocurrencies(testAddress)
	if err != nil {
		t.Fatalf("SearchCryptocurrencies failed: %v", err)
	}

	if result == nil {
		t.Fatal("Search result is nil")
	}

	t.Logf("Search successful: %v, IsTx: %v", result.Success, result.IsTx)
	for network, currencies := range result.Items {
		t.Logf("Network %s supports: %v", network, currencies)
	}
}

func TestIntegration_CryptoAddress_Workflow(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := getIntegrationClient(t)

	// First, try to get existing address
	addr, err := client.GetCryptoAddress(testAddress, testCurrency, testNetwork)
	if err != nil {
		t.Logf("Address not found, will try to add it: %v", err)

		// Try to add the address
		addr, err = client.AddCryptoAddress(testAddress, testCurrency, testNetwork)
		if err != nil {
			t.Fatalf("AddCryptoAddress failed: %v", err)
		}
		t.Logf("Successfully added address with ID: %d", addr.ID)
	} else {
		t.Logf("Found existing address with ID: %d", addr.ID)
	}

	if addr.Address != testAddress {
		t.Errorf("Expected address %s, got %s", testAddress, addr.Address)
	}

	if strings.ToUpper(addr.Currency) != testCurrency {
		t.Errorf("Expected currency %s, got %s", testCurrency, addr.Currency)
	}
}

func TestIntegration_FastCheck_Workflow(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := getIntegrationClient(t)

	// List existing fast checks
	fastChecks, err := client.ListFastChecks()
	if err != nil {
		t.Fatalf("ListFastChecks failed: %v", err)
	}
	t.Logf("Found %d existing fast checks", len(fastChecks))

	// Create a new fast check
	newCheck, err := client.CreateFastCheck(testAddress, testCurrency, testNetwork)
	if err != nil {
		t.Fatalf("CreateFastCheck failed: %v", err)
	}
	t.Logf("Created fast check with ID: %d", newCheck.ID)

	// Get the fast check details
	checkDetails, err := client.GetFastCheck(newCheck.ID)
	if err != nil {
		t.Fatalf("GetFastCheck failed: %v", err)
	}
	t.Logf("Retrieved fast check details for ID: %d", checkDetails.ID)

	// Run the fast check
	runResult, err := client.RunFastCheck(newCheck.ID)
	if err != nil {
		t.Fatalf("RunFastCheck failed: %v", err)
	}
	t.Logf("Fast check run status: %s", runResult.Status)

	// Wait a bit and check the results
	time.Sleep(5 * time.Second)
	finalCheck, err := client.GetFastCheck(newCheck.ID)
	if err != nil {
		t.Fatalf("Final GetFastCheck failed: %v", err)
	}
	t.Logf("Final fast check status: %d", finalCheck.Status)
}

func TestIntegration_Investigation_Workflow(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := getIntegrationClient(t)

	// List existing investigations
	investigations, err := client.ListInvestigations()
	if err != nil {
		t.Fatalf("ListInvestigations failed: %v", err)
	}
	t.Logf("Found %d existing investigations", len(investigations))

	// Create a new investigation using address ID
	// First we need to ensure the address exists
	addr, err := client.GetCryptoAddress(testAddress, testCurrency, testNetwork)
	if err != nil {
		// Try to add the address if it doesn't exist
		addr, err = client.AddCryptoAddress(testAddress, testCurrency, testNetwork)
		if err != nil {
			t.Fatalf("Failed to add address for investigation: %v", err)
		}
	}

	// Create investigation using address ID
	target := strconv.Itoa(addr.ID) // Convert ID to string
	newInvestigation, err := client.CreateInvestigation(target, testCurrency)
	if err != nil {
		t.Fatalf("CreateInvestigation failed: %v", err)
	}
	t.Logf("Created investigation with ID: %d", newInvestigation.ID)

	// Get investigation details
	investigationDetails, err := client.GetInvestigation(newInvestigation.ID)
	if err != nil {
		t.Fatalf("GetInvestigation failed: %v", err)
	}
	t.Logf("Retrieved investigation details for ID: %d, Status: %d",
		investigationDetails.ID, investigationDetails.Status)

	// Run the investigation
	runResult, err := client.RunInvestigation(newInvestigation.ID)
	if err != nil {
		t.Fatalf("RunInvestigation failed: %v", err)
	}
	t.Logf("Investigation run status: %s", runResult.Status)

	// Wait a bit and check the results
	time.Sleep(10 * time.Second)
	finalInvestigation, err := client.GetInvestigation(newInvestigation.ID)
	if err != nil {
		t.Fatalf("Final GetInvestigation failed: %v", err)
	}
	t.Logf("Final investigation status: %d", finalInvestigation.Status)

	// Try to get graph data
	graph, err := client.GetInvestigationGraph(newInvestigation.ID, nil)
	if err != nil {
		t.Logf("GetInvestigationGraph failed (this might be expected): %v", err)
	} else {
		t.Logf("Graph retrieved successfully: Success=%v, IsMain=%v, NeedDeep=%v",
			graph.Success, graph.IsMain, graph.NeedDeep)
	}
}

func TestIntegration_End2End_Scenario(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := getIntegrationClient(t)

	t.Log("=== Starting End-to-End Integration Test ===")

	// Step 1: Check balance
	t.Log("Step 1: Checking balance...")
	balance, err := client.GetBalance()
	if err != nil {
		t.Fatalf("Failed to get balance: %v", err)
	}
	t.Logf("Current balance: %.2f, Available investigations: %d, Available fast checks: %d",
		balance.Balance, balance.Investigations, balance.FastChecks)

	// Step 2: Search for supported cryptocurrencies
	t.Log("Step 2: Searching supported cryptocurrencies...")
	searchResult, err := client.SearchCryptocurrencies(testAddress)
	if err != nil {
		t.Fatalf("Failed to search cryptocurrencies: %v", err)
	}
	t.Logf("Found supported networks and currencies: %v", searchResult.Items)

	// Step 3: Ensure address is in the system
	t.Log("Step 3: Ensuring address is in the system...")
	addr, err := client.GetCryptoAddress(testAddress, testCurrency, testNetwork)
	if err != nil {
		t.Log("Address not found, adding it...")
		addr, err = client.AddCryptoAddress(testAddress, testCurrency, testNetwork)
		if err != nil {
			t.Fatalf("Failed to add address: %v", err)
		}
		t.Logf("Successfully added address with ID: %d", addr.ID)
	} else {
		t.Logf("Address already exists with ID: %d", addr.ID)
	}

	// Step 4: Perform a fast check if we have available checks
	if balance.FastChecks > 0 {
		t.Log("Step 4: Performing fast check...")
		fastCheck, err := client.CreateFastCheck(testAddress, testCurrency, testNetwork)
		if err != nil {
			t.Fatalf("Failed to create fast check: %v", err)
		}
		t.Logf("Created fast check with ID: %d", fastCheck.ID)

		runResult, err := client.RunFastCheck(fastCheck.ID)
		if err != nil {
			t.Fatalf("Failed to run fast check: %v", err)
		}
		t.Logf("Fast check started with status: %s", runResult.Status)
	} else {
		t.Log("Step 4: Skipping fast check - no available checks")
	}

	// Step 5: Create an investigation if we have available investigations
	if balance.Investigations > 0 {
		t.Log("Step 5: Creating investigation...")
		investigation, err := client.CreateInvestigation(strconv.Itoa(addr.ID), testCurrency)
		if err != nil {
			t.Fatalf("Failed to create investigation: %v", err)
		}
		t.Logf("Created investigation with ID: %d", investigation.ID)

		runResult, err := client.RunInvestigation(investigation.ID)
		if err != nil {
			t.Fatalf("Failed to run investigation: %v", err)
		}
		t.Logf("Investigation started with status: %s", runResult.Status)
	} else {
		t.Log("Step 5: Skipping investigation - no available investigations")
	}

	t.Log("=== End-to-End Integration Test Completed Successfully ===")
}
