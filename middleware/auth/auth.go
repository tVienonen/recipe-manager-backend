package auth

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"time"

	"encoding/base64"

	"github.com/kataras/iris"
)

type AuthHeaderContent struct {
	D []byte `json:"d"`
	S []byte `json:"s"`
}
type AuthPayload struct {
	exp *time.Time `json:"exp,omitempty"`
	// TODO: Define user data
}

func New(publicKey *rsa.PublicKey) iris.Handler {
	return func(ctx iris.Context) {
		authHeader := ctx.GetHeader("Authorization")
		decodedAuthHeader, err := base64.StdEncoding.DecodeString(authHeader)
		if err != nil {
			log.Println(err)
			rejectRequest(ctx)
		}
		authHeaderContent := new(AuthHeaderContent)
		json.Unmarshal(decodedAuthHeader, authHeaderContent)

		var signature []byte
		hex.Decode(signature, authHeaderContent.S)
		hashed := sha256.Sum256(authHeaderContent.D)

		if err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature); err != nil {
			log.Println(err)
			log.Println("Signature was invalid")
			rejectRequest(ctx)
			return
		}
		payload := new(AuthPayload)
		json.Unmarshal(authHeaderContent.D, payload)
		if payload.exp == nil {
			log.Println("Token missing expiration data")
			rejectRequest(ctx)
			return
		}
		if time.Now().After(*payload.exp) {
			log.Println("Token was expired")
			rejectRequest(ctx)
			return
		}
	}
}
func rejectRequest(ctx iris.Context) {
	ctx.StatusCode(401)
	ctx.EndRequest()
}
