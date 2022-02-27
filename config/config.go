package config

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"Zap" yaml:"zap"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// auto
	AutoCode Autocode `mapstructure:"autoCode" json:"autoCode" yaml:"autoCode"`
	// gorm
	Mysql  MySql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql  Pgsql `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	DBList []DB  `mapstructure:"db-list" json:"db-list" yaml:"db-list"`

	// oss
	Local Local `mapstructure:"local" json:"local" yaml:"local"`

	Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`
	Timer Timer `mapstructure:"timer" json:"timer" yaml:"timer"`

	// CORS config
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
