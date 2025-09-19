package golang

import (
	"net/url"
)

// GetCryptoAddress retrieves information about a cryptocurrency address
func (c *Client) GetCryptoAddress(address, currency, network string) (*CryptoAddress, error) {
	params := url.Values{}
	params.Set("address", address)
	params.Set("currency", currency)
	if network != "" {
		params.Set("network", network)
	}

	path := "/api/crypto/address?" + params.Encode()

	resp, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var cryptoAddress CryptoAddress
	if err := c.handleResponse(resp, &cryptoAddress); err != nil {
		return nil, err
	}

	return &cryptoAddress, nil
}

// AddCryptoAddress adds a new cryptocurrency address to the system
func (c *Client) AddCryptoAddress(address, currency, network string) (*CryptoAddress, error) {
	data := url.Values{}
	data.Set("address", address)
	data.Set("currency", currency)
	if network != "" {
		data.Set("network", network)
	}

	resp, err := c.doRequest("PUT", "/api/crypto/address", data)
	if err != nil {
		return nil, err
	}

	var cryptoAddress CryptoAddress
	if err := c.handleResponse(resp, &cryptoAddress); err != nil {
		return nil, err
	}

	return &cryptoAddress, nil
}

// SearchCryptocurrencies searches for supported cryptocurrencies by address
// Determines which cryptocurrencies can be associated with the specified address
func (c *Client) SearchCryptocurrencies(address string) (*CryptoSearchResponse, error) {
	params := url.Values{}
	params.Set("address", address)

	path := "/api/crypto/search?" + params.Encode()

	resp, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var searchResult CryptoSearchResponse
	if err := c.handleResponse(resp, &searchResult); err != nil {
		return nil, err
	}

	return &searchResult, nil
}
