package conf

import "os"

const (
	// EnvProd 本番環境
	EnvProd = "production"
	// EnvDev 開発環境
	EnvDev = "development"
	// EnvTest テスト環境
	EnvTest = "test"
)

// Env 環境名取得
func Env() string {
	switch os.Getenv("GO_ENV") {
	case EnvProd:
		return EnvProd
	case EnvDev:
		return EnvDev
	case EnvTest:
		return EnvTest
	}

	return EnvDev
}

// IsProd 本番環境かどうか
func IsProd() bool {
	return Env() == EnvProd
}

// IsDev 開発環境かどうか
func IsDev() bool {
	return Env() == EnvDev
}

// IsTest 本番環境かどうか
func IsTest() bool {
	return Env() == EnvTest
}
