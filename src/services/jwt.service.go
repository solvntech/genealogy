package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/duchai27798/demo_migrate/src/models/auth"
	"os"
	"time"
)

type IJWTService interface {
	GenericToken(claim *auth.JWTClaim) (string, error)
	VerifyToken(tokenString string, origin string) (bool, *auth.JWTClaim)
}

type JWTService struct {
	jwtPrivateKey string
}

// GenericToken generic jwt token
func (jwtService JWTService) GenericToken(claim *auth.JWTClaim) (string, error) {
	// expires duration is 10 mins
	claim.ExpiresAt = time.Now().UTC().Add(time.Duration(10) * time.Minute).Unix()
	claim.IssuedAt = time.Now().UTC().Unix()
	claim.Issuer = claim.Id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	fmt.Println(token)

	// sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(jwtService.jwtPrivateKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// VerifyToken verify token
func (jwtService JWTService) VerifyToken(tokenString string, origin string) (bool, *auth.JWTClaim) {
	claim := &auth.JWTClaim{}

	// check private key is valid or not
	if token, _ := getTokenFromString(tokenString, claim, jwtService.jwtPrivateKey); token.Valid {
		// check token is valid or not
		if e := claim.Valid(); e == nil {
			return true, claim
		}
	}

	// token is invalid
	return false, claim
}

func NewJWTService() IJWTService {
	jwtPrivateKey := os.Getenv("PRIVATE_KEY")
	return &JWTService{
		jwtPrivateKey,
	}
}

// ValidateToken
// - check signature is valid or not
// - encoded signature
func getTokenFromString(tokenString string, claim *auth.JWTClaim, jwtPrivateKey string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(jwtPrivateKey), nil
	})
}
