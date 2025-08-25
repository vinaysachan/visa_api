package passport

// Token request types
type PasswordGrantRequest struct {
	GrantType    string `json:"grant_type" binding:"required"`
	ClientID     uint   `json:"client_id" binding:"required"`
	ClientSecret string `json:"client_secret" binding:"required"`
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Scope        string `json:"scope"`
}

type PersonalAccessTokenRequest struct {
	Name      string `json:"name" binding:"required"`
	Scopes    string `json:"scopes"`
	ExpiresIn int    `json:"expires_in"`
}
