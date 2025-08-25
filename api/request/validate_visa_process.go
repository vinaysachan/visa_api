package request

type CountryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PortResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// Visa Apply Form Data Response structure
type VisaApplyFormDataResponse struct {
	Countries []CountryResponse `json:"countries"`
	Ports     []PortResponse    `json:"ports"`
}
