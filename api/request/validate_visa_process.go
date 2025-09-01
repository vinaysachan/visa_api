package request

type CountryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ArrivalPortResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type VisaTypeResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ApplicationTypeResponse struct {
	ID        uint               `json:"id"`
	Name      string             `json:"name"`
	VisaTypes []VisaTypeResponse `json:"visa_types"`
}

// Visa Apply Form Data Response structure
type VisaApplyFormDataResponse struct {
	Countries        []CountryResponse         `json:"countries"`
	ArrivalPorts     []ArrivalPortResponse     `json:"arrival_ports"`
	ApplicationTypes []ApplicationTypeResponse `json:"application_types"`
}
