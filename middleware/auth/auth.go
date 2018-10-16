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
		decodedAuthHeader, _ := base64.StdEncoding.DecodeString(authHeader)
		authHeaderContent := new(AuthHeaderContent)
		json.Unmarshal(decodedAuthHeader, authHeaderContent)

		var signature []byte
		hex.Decode(signature, authHeaderContent.S)
		hashed := sha256.Sum256(authHeaderContent.D)

		err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
		if err != nil {
			// TODO handle errors
			ctx.EndRequest()
			panic("Signature is not valid")
		}
		payload := new(AuthPayload)
		json.Unmarshal(authHeaderContent.D, payload)
		if payload.exp == nil {
			// TODO: Handle
			ctx.EndRequest()
			panic("Payload expiration data was missing")
		}
		if time.Now().After(*payload.exp) {
			// TODO: Handle
			ctx.EndRequest()
			log.Println("Token was expired")
			return
		}
	}
}
