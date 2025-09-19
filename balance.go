package golang

import (
	"net/url"
)

// GetBalance retrieves the user's current balance and available services
func (c *Client) GetBalance() (*BalanceResponse, error) {
	resp, err := c.doRequest("GET", "/api/user/balance", nil)
	if err != nil {
		return nil, err
	}

	var balance BalanceResponse
	if err := c.handleResponse(resp, &balance); err != nil {
		return nil, err
	}

	return &balance, nil
}

// CreatePayment creates a payment invoice for balance top-up
func (c *Client) CreatePayment(amount, currency string) (*PaymentResponse, error) {
	data := url.Values{}
	data.Set("amount", amount)
	data.Set("currency", currency)

	resp, err := c.doRequest("PUT", "/api/user/balance", data)
	if err != nil {
		return nil, err
	}

	var payment PaymentResponse
	if err := c.handleResponse(resp, &payment); err != nil {
		return nil, err
	}

	return &payment, nil
}

// GetAllowedFiatCurrencies retrieves the list of allowed fiat currencies for payment
func (c *Client) GetAllowedFiatCurrencies() ([]AllowedFiatCurrency, error) {
	resp, err := c.doRequest("GET", "/api/user/balance/allowed_fiat", nil)
	if err != nil {
		return nil, err
	}

	var currencies []AllowedFiatCurrency
	if err := c.handleResponse(resp, &currencies); err != nil {
		return nil, err
	}

	return currencies, nil
}
