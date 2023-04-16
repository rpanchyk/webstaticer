package models

type Config struct {
	Server  ServerConfig  `mapstructure:"server"`
	WebRoot WebRootConfig `mapstructure:"webroot"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type WebRootConfig struct {
	Dir  string `mapstructure:"dir"`
	File string `mapstructure:"file"`
}
