package totp

type Authenticator interface {
	GenerateQR() (string, error)
	GenerateSecretKey(string) (string, error)
	Validate(string, string) bool
}
