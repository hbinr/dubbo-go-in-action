package conf

// AppConfig my app config
type AppConfig struct {
	MySQLConfig `yaml:"mysql"`
	RedisConfig `yaml:"redis"`
}

// MySQLConfig mysql config
type MySQLConfig struct {
	DSN          string `yaml:"dsn"`            // write data source name.
	LogMode      bool   `yaml:"log_mode"`       // whether to open the log
	MaxOpenConns int    `yaml:"max_open_conns"` // max open conns
	MaxIdleConns int    `yaml:"max_idle_conns"` // max idle conns
}

// RedisConfig redis config
type RedisConfig struct {
	Host         string `yaml:"host"`
	Password     string `yaml:"password"`
	Port         int    `yaml:"port"`
	DB           int    `yaml:"db"`
	PoolSize     int    `yaml:"pool_size"`
	MinIdleConns int    `yaml:"min_idle_conns"`
}
