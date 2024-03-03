package config

type PgCounters struct {
	Host        string `env:"PG_COUNTERS_HOST, default=localhost"`
	Port        int    `env:"PG_COUNTERS_PORT, required"`
	User        string `env:"PG_COUNTERS_USER, required"`
	Password    string `env:"PG_COUNTERS_PASSWORD, required"`
	Database    string `env:"PG_COUNTERS_DATABASE, required"`
	SSLMode     string `env:"PG_COUNTERS_SSLMODE, default=verify-full"`
	ConnTimeout int    `env:"PG_COUNTERS_CONNTIMEOUT, default=5"`
	MaxConn     int    `env:"PG_COUNTERS_MAXCONN, default=5"`
}

type PgEmails01 struct {
	Host        string `env:"PG_EMAILS_01_HOST, default=localhost"`
	Port        int    `env:"PG_EMAILS_01_PORT, required"`
	User        string `env:"PG_EMAILS_01_USER, required"`
	Password    string `env:"PG_EMAILS_01_PASSWORD, required"`
	Database    string `env:"PG_EMAILS_01_DATABASE, required"`
	SSLMode     string `env:"PG_EMAILS_01_SSLMODE, default=verify-full"`
	ConnTimeout int    `env:"PG_EMAILS_01_CONNTIMEOUT, default=5"`
	MaxConn     int    `env:"PG_EMAILS_01_MAXCONN, default=5"`
}

type PgEmails02 struct {
	Host        string `env:"PG_EMAILS_02_HOST, default=localhost"`
	Port        int    `env:"PG_EMAILS_02_PORT, required"`
	User        string `env:"PG_EMAILS_02_USER, required"`
	Password    string `env:"PG_EMAILS_02_PASSWORD, required"`
	Database    string `env:"PG_EMAILS_02_DATABASE, required"`
	SSLMode     string `env:"PG_EMAILS_02_SSLMODE, default=verify-full"`
	ConnTimeout int    `env:"PG_EMAILS_02_CONNTIMEOUT, default=5"`
	MaxConn     int    `env:"PG_EMAILS_02_MAXCONN, default=5"`
}

type PgEmails03 struct {
	Host        string `env:"PG_EMAILS_03_HOST, default=localhost"`
	Port        int    `env:"PG_EMAILS_03_PORT, required"`
	User        string `env:"PG_EMAILS_03_USER, required"`
	Password    string `env:"PG_EMAILS_03_PASSWORD, required"`
	Database    string `env:"PG_EMAILS_03_DATABASE, required"`
	SSLMode     string `env:"PG_EMAILS_03_SSLMODE, default=verify-full"`
	ConnTimeout int    `env:"PG_EMAILS_03_CONNTIMEOUT, default=5"`
	MaxConn     int    `env:"PG_EMAILS_03_MAXCONN, default=5"`
}
