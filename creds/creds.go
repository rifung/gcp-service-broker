package creds

import (
	"os"
)

func GetDbHost() string {
	return getEnvDefault("DB_HOST", "")
}

func GetDbUserName() string {
	return getEnvDefault("DB_USERNAME", "")
}

func GetDbPassword() string {
	return getEnvDefault("DB_PASSWORD", "")
}

func GetRootCreds() string {
	return getEnvDefault("ROOT_SERVICE_ACCOUNT_JSON", "")
}

func GetSecurityUserName() string {
	return getEnvDefault("SECURITY_USER_NAME", "")
}

func GetSecurityUserPassword() string {
	return getEnvDefault("SECURITY_USER_PASSWORD", "")
}

func getEnvDefault(key string, def string) string {
	if e, ok := os.LookupEnv(key); ok {
		return e
	}
	return def
}
