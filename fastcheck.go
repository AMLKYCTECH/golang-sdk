package golang

import (
	"fmt"
	"net/url"
)

// ListFastChecks retrieves the list of fast checks
func (c *Client) ListFastChecks() ([]FastCheck, error) {
	resp, err := c.doRequest("GET", "/api/fast-check", nil)
	if err != nil {
		return nil, err
	}

	var listResp FastCheckListResponse
	if err := c.handleResponse(resp, &listResp); err != nil {
		return nil, err
	}

	return listResp.Data, nil
}

// CreateFastCheck creates a new fast check
func (c *Client) CreateFastCheck(address, currency, network string) (*FastCheck, error) {
	data := url.Values{}
	data.Set("address", address)
	data.Set("currency", currency)
	data.Set("network", network)

	resp, err := c.doRequest("PUT", "/api/fast-check", data)
	if err != nil {
		return nil, err
	}

	var fastCheck FastCheck
	if err := c.handleResponse(resp, &fastCheck); err != nil {
		return nil, err
	}

	return &fastCheck, nil
}

// GetFastCheck retrieves information about a specific fast check
func (c *Client) GetFastCheck(id int) (*FastCheck, error) {
	path := fmt.Sprintf("/api/fast-check/%d", id)

	resp, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var fastCheck FastCheck
	if err := c.handleResponse(resp, &fastCheck); err != nil {
		return nil, err
	}

	return &fastCheck, nil
}

// RunFastCheck starts the execution of a fast check
func (c *Client) RunFastCheck(id int) (*RunStatusResponse, error) {
	path := fmt.Sprintf("/api/fast-check/%d/run", id)

	resp, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var status RunStatusResponse
	if err := c.handleResponse(resp, &status); err != nil {
		return nil, err
	}

	return &status, nil
}
