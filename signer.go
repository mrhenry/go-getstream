package getstream

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"strings"

	"gopkg.in/dgrijalva/jwt-go.v3"
)

// Credits to https://github.com/hyperworks/go-getstream for the urlSafe and generateToken methods

// Signer is responsible for generating Tokens
type Signer struct {
	Secret string
}

func (s Signer) urlSafe(src string) string {
	src = strings.Replace(src, "+", "-", -1)
	src = strings.Replace(src, "/", "_", -1)
	src = strings.Trim(src, "=")
	return src
}

// generateToken will user the Secret of the signer and the message passed as an argument to generate a Token
func (s Signer) generateToken(message string) string {
	hash := sha1.New()
	hash.Write([]byte(s.Secret))
	key := hash.Sum(nil)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(message))
	digest := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return s.urlSafe(digest)
}

// GenerateFeedScopeToken returns a jwt
func (s Signer) GenerateFeedScopeToken(resource ScopeContext, action ScopeAction, feed Feed) (string, error) {

	claims := jwt.MapClaims{
		"resource": resource.Value(),
		"action":   action.Value(),
		// "aud":
		// "exp": time.Now().UTC().Add(time.Hour * 1),
		// "jti": uuid.New(),
		// "iat": time.Now(),
		// "iss":
		// "nbf": time.Now().Unix(),
		// "sub":
	}

	if feed != nil {
		claims["feed_id"] = feed.feedIDWithoutColon()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(s.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GenerateUserScopeToken returns a jwt
func (s Signer) GenerateUserScopeToken(resource ScopeContext, action ScopeAction, userID string) (string, error) {

	claims := jwt.MapClaims{
		"resource": resource.Value(),
		"action":   action.Value(),
		// "aud":
		// "exp": time.Now().UTC().Add(time.Hour * 1),
		// "jti": uuid.New(),
		// "iat": time.Now(),
		// "iss":
		// "nbf": time.Now().Unix(),
		// "sub":
	}

	if userID != "" {
		claims["user_id"] = userID
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(s.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
