package main

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Auth holds configuration for JWT authentication and cookies
// This struct centralizes all authentication-related settings
type Auth struct {
	Issuer        string        // The entity issuing the JWT tokens (typically your API name)
	Audience      string        // The intended recipient of the token (typically your frontend app)
	Secret        string        // Secret key used to sign JWT tokens
	TokenExpiry   time.Duration // How long access tokens remain valid
	RefreshExpiry time.Duration // How long refresh tokens remain valid
	CookieDomain  string        // Domain for which cookies are valid
	CookiePath    string        // Path restriction for cookies
	CookieName    string        // Name of the cookie storing the token
}

// jwtUser represents user information stored in JWT claims
// Contains minimal user data needed for authentication
type jwtUser struct {
	ID        int    `json:"id"`         // User's unique identifier
	FirstName string `json:"first_name"` // User's first name
	LastName  string `json:"last_name"`  // User's last name
}

// TokenPair contains both access and refresh tokens
// Returned to clients after successful authentication
type TokenPair struct {
	Token        string `json:"access_token"`  // Short-lived JWT for API access
	RefreshToken string `json:"refresh_token"` // Long-lived JWT for obtaining new access tokens
}

// TokenClaims extends the standard JWT claims
// Used for parsing and validating incoming JWTs
type TokenClaims struct {
	jwt.RegisteredClaims // Standard JWT claims (iss, sub, exp, etc.)
	// Custom claims can be added here as needed
}

//

func (j *Auth) GenerateTokenPair(user *jwtUser) (TokenPair, error) {
	// Create a token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.FirstName + " " + user.LastName
	claims["sub"] = user.ID
	claims["iat"] = time.Now().UTC().Unix()
	claims["typ"] = "JWT"

	//set Expiry FoR jwt
	claims["exp"] = time.Now().UTC().Add(j.TokenExpiry).Unix()
	//Signed Access Token
	signedAccessToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPair{}, err
	}

	//create Refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	//Refresh token Claims
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["sub"] = user.ID
	//Expiry For Refresh token
	refreshClaims["exp"] = time.Now().UTC().Add(j.RefreshExpiry).Unix()

	//Cretae A Signed JWT
	signedRefreshToken, err := refreshToken.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPair{}, err
	}

	return TokenPair{
		Token:        signedAccessToken,
		RefreshToken: signedRefreshToken,
	}, nil
}

func (j *Auth) GetRefreshCookie(refreshToken string) *http.Cookie {
	return &http.Cookie{
		Name:     j.CookieName,
		Value:    refreshToken,
		Path:     j.CookiePath,
		Expires:  time.Now().Add(j.RefreshExpiry),
		MaxAge:   int(j.RefreshExpiry.Seconds()),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
}

func (j *Auth) GetExpiredCookie() *http.Cookie {
	return &http.Cookie{
		Name:     j.CookieName,
		Value:    "",
		Path:     j.CookiePath,
		Expires:  time.Now().Add(-1000),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
}
