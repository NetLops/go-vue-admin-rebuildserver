package config

type Pgsql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Port         string `mapstructure:"port" json:"port" yaml:"port"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns int    `mapstructure:"max-dile-conns" json:"maxIdleConns" yaml:"max-dile-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      string `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	LogZap       bool   `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
}

//  Dsn: 基于配置文件获取dsn
func (p *Pgsql) Dsn() string {
	return "host=" + p.Path + " user=" + p.Username + " password=" + p.Password + " port=" + p.Port + " " + p.Config
}

//  LinkDsn: 根据 dbname 生成 dsn
func (p *Pgsql) LinkDsn(dbname string) string {
	return "host=" + p.Path + " user=" + p.Username + " password=" + p.Password + " dbname=" + dbname + " port=" + p.Port + " " + p.Config
}
