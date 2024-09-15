package helpers

import (
	"log"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func GenerateOpt() *otp.Key {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Example.com",
		AccountName: "alice@example.com",
	})

	if err != nil {
		log.Fatal(err)
	}

	return key
}
