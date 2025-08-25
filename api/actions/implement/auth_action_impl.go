package implement

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/vinaysachan/visa_api/api/actions/interfaces"
	"github.com/vinaysachan/visa_api/api/request"
	"github.com/vinaysachan/visa_api/base/packages/passport"
	"github.com/vinaysachan/visa_api/base/utils"
	"github.com/vinaysachan/visa_api/data/tasks"
)

type authActionImpl struct{}

func NewAuthAction() interfaces.AuthAction {
	return &authActionImpl{}
}

func (a *authActionImpl) LoginUser(req request.LoginUserRequest) (*request.LoginResponse, error) {

	user, err := tasks.FindActiveUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials (Email)")
	}

	if !utils.ComparePassword(user.Password, req.Password) {
		return nil, errors.New("invalid credentials (Password)")
	}

	// Generate access token
	accessToken, err := passport.GenerateToken()
	if err != nil {
		return nil, errors.New("could not generate token")
	}

	// Generate refresh token
	refreshToken, err := passport.GenerateToken()
	if err != nil {
		return nil, errors.New("could not generate refresh token")
	}

	expiresAt := time.Now().Add(time.Hour * 24)
	formattedExpiresAt := expiresAt.Format("2006-01-02 15:04:05") // Go's reference time format

	refreshExpiresAt := time.Now().Add(time.Hour * 24)
	formattedRefreshExpiresAt := refreshExpiresAt.Format("2006-01-02 15:04:05")

	fmt.Println("accessToken & refreshToken :", accessToken, refreshToken, formattedRefreshExpiresAt)

	// Create access token record
	tokenRecord := passport.OAuthAccessToken{
		ID:        accessToken,
		UserID:    user.ID,
		ClientID:  utils.GodotEnv("OAUTH_CLIENT_ID"),
		Revoked:   false,
		ExpiresAt: &expiresAt,
	}

	if err := passport.SaveAccessTokenInDB(&tokenRecord); err != nil {
		return nil, errors.New(err.Error() + " - could not save token in database")
	}

	// Create refresh token record
	refreshTokenRecord := passport.OAuthRefreshToken{
		ID:            refreshToken,
		AccessTokenID: accessToken,
		Revoked:       false,
		ExpiresAt:     &refreshExpiresAt,
	}

	if err := passport.SaveRefreshTokenInDB(&refreshTokenRecord); err != nil {
		return nil, errors.New(err.Error() + " - could not save refresh token in database")
	}

	passport.RevokeUserTokenInDB(user.ID, accessToken)

	fmt.Println("refreshToken :", refreshTokenRecord)

	return &request.LoginResponse{
		TokenType:    "Bearer",
		ExpiresIn:    int(expiresAt.Unix()),
		ExpiresAt:    formattedExpiresAt,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *authActionImpl) RevokeToken(token string) (map[string]any, error) {

	if err := passport.RevokeTokenInDB(token); err != nil {
		return nil, errors.New("could not revoke token")
	}

	return map[string]any{"message": "successfully logged out"}, nil
}

func (a *authActionImpl) CsrfToken(deviceId uuid.UUID, clientIp string) (*request.CsrfTokenResponse, error) {

	// Generate access token
	csrfToken, err := passport.GenerateToken()
	if err != nil {
		return nil, errors.New("could not generate token")
	}
	expiresAt := time.Now().Add(time.Hour * 24)
	formattedExpiresAt := expiresAt.Format("2006-01-02 15:04:05") // Go's reference time format

	//create csrf toke record :-
	OAuthCsrfClaim := passport.OAuthCsrfClaim{
		DeviceID: deviceId,
		IP:       clientIp,
		Token:    csrfToken,
		ExpireAt: &expiresAt,
	}

	if err := passport.SaveCsrfTokenInDB(&OAuthCsrfClaim); err != nil {
		return nil, errors.New(err.Error() + " - could not save csrf token in database")
	}

	return &request.CsrfTokenResponse{
		ExpiresAt: formattedExpiresAt,
		Token:     csrfToken,
	}, nil
}
