package config

type AppConfiguration struct {
	// Server
	Server Server `mapstructure:"server"`

	// Cron run
	Cron Cron `mapstructure:"cron"`
}

type Server struct {
	Name           string   `mapstructure:"name"`
	Profile        string   `mapstructure:"profile"`
	TrustedProxies []string `mapstructure:"trustedProxies"`
	Port           string   `mapstructure:"port"`
	MasterKey      string   `mapstructure:"masterKey"`
	Version        string   `mapstructure:"version"`
}

type Cron struct {
	IsDisabled bool `mapstructure:"IsDisabled"`

	// ReportDaily cron
	ReportDaily struct {
		IsDisable bool   `mapstructure:"IsDisabled"`
		Cron      string `mapstructure:"cron"`
	} `mapstructure:"reportDaily"`
}
