package golang

import (
	"fmt"
	"net/url"
	"strconv"
)

// ListInvestigations retrieves the list of investigations
func (c *Client) ListInvestigations() ([]Investigation, error) {
	resp, err := c.doRequest("GET", "/api/investigations", nil)
	if err != nil {
		return nil, err
	}

	var listResp InvestigationListResponse
	if err := c.handleResponse(resp, &listResp); err != nil {
		return nil, err
	}

	return listResp.Data, nil
}

// CreateInvestigation creates a new investigation
func (c *Client) CreateInvestigation(target, currency string) (*Investigation, error) {
	data := url.Values{}
	data.Set("target", target)
	data.Set("currency", currency)

	resp, err := c.doRequest("PUT", "/api/investigations", data)
	if err != nil {
		return nil, err
	}

	var investigation Investigation
	if err := c.handleResponse(resp, &investigation); err != nil {
		return nil, err
	}

	return &investigation, nil
}

// GetInvestigation retrieves information about a specific investigation
func (c *Client) GetInvestigation(id int) (*Investigation, error) {
	path := fmt.Sprintf("/api/investigations/%d", id)

	resp, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var investigation Investigation
	if err := c.handleResponse(resp, &investigation); err != nil {
		return nil, err
	}

	return &investigation, nil
}

// RunInvestigation starts the execution of an investigation
func (c *Client) RunInvestigation(id int) (*RunStatusResponse, error) {
	path := fmt.Sprintf("/api/investigations/%d/run", id)

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

// GetInvestigationGraph retrieves graph data for an investigation
// If itemID is provided, it gets the graph for a specific address within the investigation
func (c *Client) GetInvestigationGraph(id int, itemID *int) (*InvestigationGraphResponse, error) {
	path := fmt.Sprintf("/api/investigations/%d/graph", id)

	// Add item parameter if provided
	if itemID != nil {
		path += "?item=" + strconv.Itoa(*itemID)
	}

	resp, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var graph InvestigationGraphResponse
	if err := c.handleResponse(resp, &graph); err != nil {
		return nil, err
	}

	return &graph, nil
}

// DeepInvestigation performs deep investigation on a specific address within an existing investigation
func (c *Client) DeepInvestigation(id int, address string) (*DeepInvestigationResponse, error) {
	path := fmt.Sprintf("/api/investigations/%d/deep", id)

	data := url.Values{}
	data.Set("address", address)

	resp, err := c.doRequest("POST", path, data)
	if err != nil {
		return nil, err
	}

	var result DeepInvestigationResponse
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
