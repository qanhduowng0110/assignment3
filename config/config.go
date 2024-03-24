package config

const (
	ROLE_ADMIN int = 0
	ROLE_USER  int = 1
)

type Configurations struct {
	Postgres PostgresConfig `mapstructure:"postgres"`
}

type PostgresConfig struct {
	ConnectionString string `mapstructure:"connectionString"`
}
