package config

type CORS struct {
	Mode      string          `mapstructure:"mode" json:"mode" yaml:"mode"`
	WhiteList []CORSWhiteList `mapstructure:"whiteList" json:"whiteList" yaml:"whiteList"`
}

type CORSWhiteList struct {
	AllowOrigin      string `mapstructure:"allow-origin" json:"allow-origin" yaml:"allow-origin"`
	AllowMethods     string `mapstructure:"allow-methods" json:"allow-methods" yaml:"allow-methods"`
	AllowHeaders     string `mapstructure:"allow-headers" json:"allow-headers" yaml:"allow-headers"`
	ExposeHeaders    string `mapstructure:"expose-headers" json:"expose-headers" yaml:"expose-headers"`
	AllowCredentials bool   `mapstructure:"allow-credentials" json:"allow-credentials" yaml:"allow-credentials"`
}
