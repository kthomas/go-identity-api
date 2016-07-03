package identity

import (
	"sync"
	"os"
	"github.com/kthomas/go-db-config"
)

type Config struct {
	ApplicationSecret	[]byte
	JwtSharedSecret		[]byte

	DbConfig		*dbconf.DbConfig
}

var configInstance *Config
var configOnce sync.Once

func GetConfig() (*Config) {
	configOnce.Do(func() {
		applicationSecret := os.Getenv("APPLICATION_SECRET")
		if applicationSecret == "" {
			applicationSecret = "thesecret"
		}

		jwtSharedSecret := os.Getenv("JWT_SHARED_SECRET")
		if jwtSharedSecret == "" {
			jwtSharedSecret = "thesecret"
		}

		configInstance = &Config{
			ApplicationSecret: []byte(applicationSecret),
			JwtSharedSecret: []byte(jwtSharedSecret),
			DbConfig: dbconf.GetDbConfig(),
		}
	})
	return configInstance
}

var config = GetConfig()
