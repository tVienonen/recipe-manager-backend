package config_provider

import (
	"github.com/kataras/iris/context"
)

type DBConfig struct {
	DB_USER string `json:"db_user"`
	DB_PASS string `json:"db_pass"`
	DB_NAME string `json:"db_name"`
	DB_HOST string `json:"db_host"`
	DB_PORT string `json:"db_port"`
}

type ConfigProvider struct {
	DBConfig *DBConfig
}

func New(dbConfig *DBConfig) context.Handler {
	provider := &ConfigProvider{
		DBConfig: dbConfig}
	return provider.ServeHTTP
}
func (c *ConfigProvider) ServeHTTP(ctx context.Context) {
	ctx.Values().SetImmutable("DB_USER", c.DBConfig.DB_USER)
	ctx.Values().SetImmutable("DB_PASS", c.DBConfig.DB_PASS)
	ctx.Values().SetImmutable("DB_HOST", c.DBConfig.DB_HOST)
	ctx.Values().SetImmutable("DB_PORT", c.DBConfig.DB_PORT)
	ctx.Values().SetImmutable("DB_NAME", c.DBConfig.DB_NAME)
	ctx.Next()
}
