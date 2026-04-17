package golang

import "fmt"

// CreateKyc creates a new KYC verification record.
func (c *Client) CreateKyc(req KycCreateRequest) (*KycCreateResponse, error) {
	resp, err := c.doRequest("PUT", "/webapi/kyc", req)
	if err != nil {
		return nil, err
	}
	var result KycCreateResponse
	return &result, c.handleResponse(resp, &result)
}

// ListKyc returns all KYC verifications for the current user.
func (c *Client) ListKyc() ([]KycVerification, error) {
	resp, err := c.doRequest("GET", "/webapi/kyc", nil)
	if err != nil {
		return nil, err
	}
	var result KycListResponse
	return result.Items, c.handleResponse(resp, &result)
}

// StartKycSession starts a KYC verification session for the given UUID.
func (c *Client) StartKycSession(uuid string) (*KycParams, error) {
	resp, err := c.doRequest("POST", fmt.Sprintf("/webapi/kyc/user/%s", uuid), nil)
	if err != nil {
		return nil, err
	}
	var result KycParams
	return &result, c.handleResponse(resp, &result)
}

// UploadKycDocuments uploads user photos and document for verification.
func (c *Client) UploadKycDocuments(uuid string, req KycUploadRequest) error {
	resp, err := c.doRequest("PUT", fmt.Sprintf("/webapi/kyc/user/%s", uuid), req)
	if err != nil {
		return err
	}
	return c.handleResponse(resp, nil)
}

// GetKycDetails returns detailed KYC information including images.
func (c *Client) GetKycDetails(uuid string) (map[string]interface{}, error) {
	resp, err := c.doRequest("GET", fmt.Sprintf("/webapi/kyc/user/%s", uuid), nil)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	return result, c.handleResponse(resp, &result)
}

// GetKycParams returns the current session parameters for the given UUID.
func (c *Client) GetKycParams(uuid string) (*KycParams, error) {
	resp, err := c.doRequest("GET", fmt.Sprintf("/webapi/kyc/user/%s/params", uuid), nil)
	if err != nil {
		return nil, err
	}
	var result KycParams
	return &result, c.handleResponse(resp, &result)
}

// UpdateKycStep updates the current step in the verification process.
func (c *Client) UpdateKycStep(uuid string, req KycStepRequest) (*KycParams, error) {
	resp, err := c.doRequest("POST", fmt.Sprintf("/webapi/kyc/user/%s/step", uuid), req)
	if err != nil {
		return nil, err
	}
	var result KycParams
	return &result, c.handleResponse(resp, &result)
}

// CheckKycDocument runs a document data check for the given UUID.
func (c *Client) CheckKycDocument(uuid string, req KycCheckRequest) (*KycParams, error) {
	resp, err := c.doRequest("POST", fmt.Sprintf("/webapi/kyc/user/%s/check", uuid), req)
	if err != nil {
		return nil, err
	}
	var result KycParams
	return &result, c.handleResponse(resp, &result)
}

// StartKycLiveness generates liveness challenges and signs the challenge token.
func (c *Client) StartKycLiveness(uuid string) (*KycLivenessStartResponse, error) {
	resp, err := c.doRequest("POST", fmt.Sprintf("/webapi/kyc/user/%s/liveness/start", uuid), nil)
	if err != nil {
		return nil, err
	}
	var result KycLivenessStartResponse
	return &result, c.handleResponse(resp, &result)
}

// VerifyKycLiveness verifies completed liveness challenges and biometric data.
func (c *Client) VerifyKycLiveness(uuid string, req KycLivenessVerifyRequest) (*KycLivenessVerifyResponse, error) {
	resp, err := c.doRequest("POST", fmt.Sprintf("/webapi/kyc/user/%s/liveness/verify", uuid), req)
	if err != nil {
		return nil, err
	}
	var result KycLivenessVerifyResponse
	return &result, c.handleResponse(resp, &result)
}

// RecognizeKycDocument performs OCR on a document image.
func (c *Client) RecognizeKycDocument(uuid string, req KycRecognizeRequest) (*KycRecognizeResponse, error) {
	resp, err := c.doRequest("POST", fmt.Sprintf("/webapi/kyc/user/%s/recognize-document", uuid), req)
	if err != nil {
		return nil, err
	}
	var result KycRecognizeResponse
	return &result, c.handleResponse(resp, &result)
}
