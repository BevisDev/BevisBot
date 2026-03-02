package config

type appConfig struct {
	Server   Server   `mapstructure:"server"`
	Logger   Logger   `mapstructure:"logger"`
	Database Database `mapstructure:"database"`
	TgBot    TGBot    `mapstructure:"tgbot"`
	OpenAI   OpenAI   `mapstructure:"openai"`
}

type OpenAI struct {
	APIKey string `mapstructure:"apiKey"`
}

type Server struct {
	Name    string `mapstructure:"name"`
	Port    string `mapstructure:"port"`
	Profile string `mapstructure:"profile"`
	Version string `mapstructure:"version"`
}

type Database struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db"`
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
}

type Logger struct {
	DirName  string `mapstructure:"dirname"`
	FileName string `mapstructure:"filename"`
}

type TGBot struct {
	Token string `mapstructure:"token"`
}
