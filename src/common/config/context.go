package config

type contextKey string

const (
	// DB関連コンテキストキー
	DBKey contextKey = "DB_KEY"
	TXKey contextKey = "TX_KEY"

	// ログ関連コンテキストキー
	LogKey contextKey = "LOG_KEY"
)
