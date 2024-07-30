package totp

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"image/png"
	"log/slog"
)

const issuer = "Wawandco"

type service struct {
	key *otp.Key
}

func NewService() Authenticator {
	return &service{}
}

func (s *service) GenerateSecretKey(email string) (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: email,
	})
	if err != nil {
		slog.Error(fmt.Sprintf("error generating secret key: %v", err))
		return "", err
	}

	s.key = key

	return key.Secret(), nil
}

func (s *service) GenerateQR() (string, error) {
	// Convert info key into a PNG
	var buf bytes.Buffer
	img, err := s.key.Image(150, 150)
	if err != nil {
		return "", nil
	}

	err = png.Encode(&buf, img)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), err
}

func (s *service) Validate(code, secret string) bool {
	return totp.Validate(code, secret)
}
