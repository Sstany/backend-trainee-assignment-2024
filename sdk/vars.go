package sdk

import "os"

const (
	EnvPostgres = "POSTGRES"
	EnvHost     = "HOST"
	EnvSecret   = "SECRET"
)

const (
	HeaderToken = "token"
)

var (
	Secret []byte
)

func InitSecret() {
	Secret = []byte(os.Getenv(EnvSecret))
}
