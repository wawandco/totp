package auth

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

type Info struct {
	Secret string
	QR     string
}

func TOTPInfo(secret string, email string) (Info, error) {
	var info Info

	key, err := generateKey(email)
	if err != nil {
		return Info{}, err
	}

	qr, err := generateQR(key)

	if err != nil {
		return Info{}, err
	}

	if secret == "" {
		info.Secret = key.Secret()
		info.QR = qr
	}

	return info, nil
}

func generateKey(email string) (*otp.Key, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: email,
	})
	if err != nil {
		slog.Error(fmt.Sprintf("error generating secret key: %v", err))
		return key, err
	}

	return key, nil
}

func generateQR(key *otp.Key) (string, error) {
	// Convert TOTP key into a PNG
	var buf bytes.Buffer
	img, err := key.Image(150, 150)
	if err != nil {
		return "", nil
	}

	err = png.Encode(&buf, img)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), err
}
