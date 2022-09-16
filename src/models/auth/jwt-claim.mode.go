package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTClaim struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	RoleId string `json:"role_id"`
	jwt.StandardClaims
}

// Valid check token is valid or not
func (claim JWTClaim) Valid() error {
	var now = time.Now().UTC().Unix()

	// check expire and issuer
	if claim.VerifyExpiresAt(now, true) && claim.VerifyIssuer(claim.UserId, true) {
		return nil
	}

	// token is invalid
	return fmt.Errorf("token is invalid")
}
