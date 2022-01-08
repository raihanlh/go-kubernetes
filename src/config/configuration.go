package config

type Configuration struct {
	Db DatabaseConfiguration `mapstructure:"db"`
}

type DatabaseConfiguration struct {
	Host      string `mapstructure:"host"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Name      string `mapstructure:"name"`
	Migration string `mapstructure:"migration"`
}
