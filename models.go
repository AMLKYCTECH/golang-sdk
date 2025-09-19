package golang

import "time"

// BalanceResponse represents the user balance response
type BalanceResponse struct {
	Success        bool    `json:"success"`
	Balance        float64 `json:"balance"`
	Investigations int     `json:"investigations"`
	FastChecks     int     `json:"fast-cheks"` // Note: API has typo in field name
}

// PaymentRequest represents the payment creation request
type PaymentRequest struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// PaymentResponse represents the payment creation response
type PaymentResponse struct {
	Success bool   `json:"success"`
	Link    string `json:"link"`
}

// AllowedFiatCurrency represents an allowed fiat currency for payment
type AllowedFiatCurrency struct {
	Name string  `json:"name"`
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
}

// CryptoAddress represents a cryptocurrency address
type CryptoAddress struct {
	ID        int       `json:"id"`
	Address   string    `json:"address"`
	Currency  string    `json:"currency"`
	Network   string    `json:"network"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CryptoAddressRequest represents a request to add a crypto address
type CryptoAddressRequest struct {
	Address  string `json:"address"`
	Currency string `json:"currency"`
	Network  string `json:"network"`
}

// CryptoSearchResponse represents the crypto search response
type CryptoSearchResponse struct {
	Success bool                `json:"success"`
	Items   map[string][]string `json:"items"`
	IsTx    bool                `json:"isTx"`
}

// FastCheck represents a fast check entity
type FastCheck struct {
	ID              int           `json:"id"`
	UserID          int           `json:"user_id"`
	CryptoAddressID int           `json:"crypto_address_id"`
	UUID            string        `json:"uuid"`
	Risk            *RiskData     `json:"risk,omitempty"`
	Analysis        *AnalysisData `json:"analysis,omitempty"`
	Info            *InfoData     `json:"info,omitempty"`
	Labels          []interface{} `json:"labels"`
	Status          int           `json:"status"`
	Balance         interface{}   `json:"balance"` // Can be string or number
	New             bool          `json:"new"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	Tracking        bool          `json:"tracking"`
	Address         string        `json:"address"`
	Currency        string        `json:"currency"`
	Network         string        `json:"network"`
	Price           float64       `json:"price"`
}

// FastCheckRequest represents a request to create a fast check
type FastCheckRequest struct {
	Address  string `json:"address"`
	Currency string `json:"currency"`
	Network  string `json:"network"`
}

// RunStatusResponse represents a status response for run operations
type RunStatusResponse struct {
	Status string `json:"status"`
}

// RiskData represents risk analysis data
type RiskData struct {
	Score      int                    `json:"score"`
	RiskDetail []RiskDetail           `json:"risk_detail"`
	Providers  map[string]interface{} `json:"providers"`
}

// RiskDetail represents detailed risk information
type RiskDetail struct {
	Label   string  `json:"label"`
	Type    string  `json:"type"`
	Volume  float64 `json:"volume"`
	Address string  `json:"address"`
	Percent float64 `json:"percent"`
	Risk    *string `json:"risk"`
}

// AnalysisData represents transaction analysis data
type AnalysisData struct {
	ReceivedTxs []TxAnalysis `json:"received_txs"`
	SpentTxs    []TxAnalysis `json:"spent_txs"`
}

// TxAnalysis represents transaction analysis information
type TxAnalysis struct {
	Action     string  `json:"action"`
	Count      int     `json:"count"`
	Proportion float64 `json:"proportion"`
	Amount     float64 `json:"amount"`
}

// InfoData represents general address information
type InfoData struct {
	Balance          string `json:"balance"`
	TxsCount         string `json:"txs_count"`
	FirstSeen        string `json:"first_seen"`
	LastSeen         string `json:"last_seen"`
	TotalReceived    string `json:"total_received"`
	TotalSpent       string `json:"total_spent"`
	ReceivedTxsCount string `json:"received_txs_count"`
	SpentTxsCount    string `json:"spent_txs_count"`
	ReceivedCount    string `json:"received_count"`
	SpentCount       string `json:"spent_count"`
	FirstTxTime      string `json:"first_tx_time"`
	LastTxTime       string `json:"last_tx_time"`
	TotalReceivedUSD string `json:"total_received_usd"`
	TotalSpentUSD    string `json:"total_spent_usd"`
	BalanceUSD       string `json:"balance_usd"`
	TxCount          string `json:"tx_count"`
}

// Investigation represents an investigation entity
type Investigation struct {
	ID              int           `json:"id"`
	UserID          int           `json:"user_id"`
	CryptoAddressID int           `json:"crypto_address_id"`
	Status          int           `json:"status"` // 0-новое, 1-в процессе, 2-завершено, 3-глубокое исследование
	Report          []interface{} `json:"report"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}

// InvestigationRequest represents a request to create an investigation
type InvestigationRequest struct {
	Target   string `json:"target"` // ID адреса или хеш транзакции
	Currency string `json:"currency"`
}

// InvestigationGraphResponse represents the investigation graph response
type InvestigationGraphResponse struct {
	Data     interface{} `json:"data"`
	Success  bool        `json:"success"`
	IsMain   bool        `json:"isMain"`
	NeedDeep bool        `json:"needDeep"`
}

// DeepInvestigationRequest represents a deep investigation request
type DeepInvestigationRequest struct {
	Address string `json:"address"`
}

// DeepInvestigationResponse represents a deep investigation response
type DeepInvestigationResponse struct {
	Success bool `json:"success"`
}

// ErrorResponse represents an error response from the API
type ErrorResponse struct {
	Status  bool   `json:"status,omitempty"`
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
}

// FastCheckListResponse represents the response from list fast checks API
type FastCheckListResponse struct {
	Data []FastCheck `json:"data"`
}

// InvestigationListResponse represents the response from list investigations API
type InvestigationListResponse struct {
	Data []Investigation `json:"data"`
}
