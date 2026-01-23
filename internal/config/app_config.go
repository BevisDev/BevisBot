package config

type appConfig struct {
	// Server
	Server Server `mapstructure:"server"`

	// Logger
	Logger Logger `mapstructure:"logger"`
}

type Server struct {
	Name           string   `mapstructure:"name"`
	Profile        string   `mapstructure:"profile"`
	TrustedProxies []string `mapstructure:"trustedProxies"`
	Port           string   `mapstructure:"port"`
	MasterKey      string   `mapstructure:"masterKey"`
	Version        string   `mapstructure:"version"`
}

type Logger struct {
	IsRotate   bool   `mapstructure:"isRotate"`
	DirName    string `mapstructure:"dirName"`
	FileName   string `mapstructure:"fileName"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge"`
}
