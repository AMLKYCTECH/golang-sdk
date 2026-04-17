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

// KycVerification represents a KYC verification entity
type KycVerification struct {
	ID                    int                    `json:"id"`
	UUID                  string                 `json:"uuid"`
	UserID                int                    `json:"user_id"`
	Email                 string                 `json:"email"`
	Status                int                    `json:"status"` // 0: new, 1: process, 2: moderate, 3: fail, 4: success
	Step                  int                    `json:"step"`
	StartTime             *time.Time             `json:"start_time"`
	EndTime               *time.Time             `json:"end_time"`
	DocType               *string                `json:"doc_type"`
	DocInfo               map[string]interface{} `json:"doc_info"`
	StepsInfo             map[string]interface{} `json:"steps_info"`
	ComparePhoto          map[string]interface{} `json:"compare_photo"`
	StatusCheckUser       *int                   `json:"status_check_user"`
	StatusCheckPhoto      *int                   `json:"status_check_photo"`
	CountryRestrictionMode string                `json:"country_restriction_mode"` // all, whitelist
	AllowedCountries      []string               `json:"allowed_countries"`
	AllowedDocumentTypes  []string               `json:"allowed_document_types"`
	LivenessScore         *float64               `json:"liveness_score"`
	LivenessStatus        int                    `json:"liveness_status"` // 0: pending, 1: pass, 2: fail
	LivenessAttempts      int                    `json:"liveness_attempts"`
	CreatedAt             time.Time              `json:"created_at"`
	UpdatedAt             time.Time              `json:"updated_at"`
}

// KycParams represents session parameters returned during KYC flow
type KycParams struct {
	Success      bool                   `json:"success"`
	UUID         string                 `json:"uuid"`
	StartTime    *time.Time             `json:"start_time"`
	Status       int                    `json:"status"`
	Step         int                    `json:"step"`
	SessionEnd   *time.Time             `json:"session_end"`
	Steps        map[string]interface{} `json:"steps"`
	Restrictions *KycRestrictions       `json:"restrictions"`
}

// KycRestrictions holds country/document restrictions for a KYC session
type KycRestrictions struct {
	CountryMode          string   `json:"countryMode"`
	AllowedCountries     []string `json:"allowedCountries"`
	AllowedDocumentTypes []string `json:"allowedDocumentTypes"`
}

// KycCreateRequest represents the request body for creating a KYC verification
type KycCreateRequest struct {
	Email                  string   `json:"email"`
	CountryRestrictionMode string   `json:"country_restriction_mode,omitempty"`
	AllowedCountries       []string `json:"allowed_countries,omitempty"`
	AllowedDocumentTypes   []string `json:"allowed_document_types,omitempty"`
}

// KycCreateResponse represents the response when creating a KYC verification
type KycCreateResponse struct {
	Success bool            `json:"success"`
	Item    KycVerification `json:"item"`
}

// KycListResponse represents the response for listing KYC verifications
type KycListResponse struct {
	Items []KycVerification `json:"items"`
}

// KycUploadRequest represents the request body for uploading documents
type KycUploadRequest struct {
	Cadrs        []string `json:"cadrs"`
	Passport     string   `json:"passport"`
	DocumentType string   `json:"document_type"`
	Country      string   `json:"country"`
}

// KycStepRequest represents the request body for updating a verification step
type KycStepRequest struct {
	Step int                    `json:"step"`
	Data map[string]interface{} `json:"data,omitempty"`
}

// KycCheckRequest represents the request body for document check
type KycCheckRequest struct {
	DocumentType   string `json:"document_type,omitempty"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Patronymic     string `json:"patronymic,omitempty"`
	Dob            string `json:"dob"`
	PassportDate   string `json:"passport_date,omitempty"`
	Seria          string `json:"seria,omitempty"`
	Number         string `json:"number,omitempty"`
	DocumentNumber string `json:"document_number,omitempty"`
	Residence      string `json:"residence,omitempty"`
	Citizenship    string `json:"citizenship,omitempty"`
	FullAddress    string `json:"fullAddress,omitempty"`
	ZipCode        string `json:"zipCode,omitempty"`
	City           string `json:"city,omitempty"`
	Country        string `json:"country,omitempty"`
}

// KycLivenessStartResponse represents the response for starting a liveness check
type KycLivenessStartResponse struct {
	Success        bool          `json:"success"`
	ChallengeToken string        `json:"challengeToken"`
	Challenges     []interface{} `json:"challenges"`
	TimeoutSeconds int           `json:"timeoutSeconds"`
}

// KycLivenessFrame represents a single frame submitted during liveness verification
type KycLivenessFrame struct {
	ChallengeIndex int                    `json:"challengeIndex"`
	Image          string                 `json:"image"`
	Landmarks      map[string]interface{} `json:"landmarks,omitempty"`
	Blendshapes    map[string]interface{} `json:"blendshapes,omitempty"`
	HeadPose       map[string]interface{} `json:"headPose,omitempty"`
	Timestamp      float64                `json:"timestamp"`
}

// KycLivenessVerifyRequest represents the request body for liveness verification
type KycLivenessVerifyRequest struct {
	ChallengeToken string                 `json:"challengeToken"`
	Frames         []KycLivenessFrame     `json:"frames"`
	AntiSpoofing   map[string]interface{} `json:"antiSpoofing"`
	Passport       string                 `json:"passport"`
	DocumentType   string                 `json:"documentType"`
	Country        string                 `json:"country"`
}

// KycLivenessVerifyResponse represents the response for liveness verification
type KycLivenessVerifyResponse struct {
	Success bool                   `json:"success"`
	Score   float64                `json:"score"`
	Passed  bool                   `json:"passed"`
	Details map[string]interface{} `json:"details"`
}

// KycRecognizeRequest represents the request body for document OCR recognition
type KycRecognizeRequest struct {
	Image        string `json:"image"`
	DocumentType string `json:"document_type"`
	Country      string `json:"country"`
}

// KycRecognizeResponse represents the response for document OCR recognition
type KycRecognizeResponse struct {
	Success    bool                   `json:"success"`
	Confidence float64                `json:"confidence"`
	Data       map[string]interface{} `json:"data"`
	Warnings   []string               `json:"warnings"`
}
