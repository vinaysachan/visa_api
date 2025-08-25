package interfaces

import (
	"github.com/google/uuid"
	"github.com/vinaysachan/visa_api/api/request"
)

type AuthAction interface {
	LoginUser(request.LoginUserRequest) (*request.LoginResponse, error)
	RevokeToken(string) (map[string]any, error)
	CsrfToken(uuid.UUID, string) (*request.CsrfTokenResponse, error)
}
