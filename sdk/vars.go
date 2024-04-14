package sdk

import (
	"os"
	"time"
)

const (
	EnvPostgres = "POSTGRES"
	EnvHost     = "HOST"
	EnvSecret   = "SECRET"
)

const (
	HeaderToken = "token"

	QueryTagID           = "tag_id"
	QueryFeatureID       = "feature_id"
	QueryUseLastRevision = "use_last_revision"
	QueryOffset          = "offset"
	QueryLimit           = "limit"
)
const (
	MaxLimit     = 1000
	DefaultLimit = 100
)

var (
	DefaulCacheLifetime = 5 * time.Minute
)

var (
	Secret []byte
)

func InitSecret() {
	Secret = []byte(os.Getenv(EnvSecret))
}
