package config

type appConfig struct {
	Server server `mapstructure:"server"`
	Logger logger `mapstructure:"logger"`
	TgBot  tgBot  `mapstructure:"tgBot"`
}

type server struct {
	Name    string `mapstructure:"name"`
	Port    string `mapstructure:"port"`
	Profile string `mapstructure:"profile"`
	Version string `mapstructure:"version"`
}

type logger struct {
	DirName  string `mapstructure:"dirname"`
	FileName string `mapstructure:"filename"`
}

type tgBot struct {
	Token string `mapstructure:"token"`
}
