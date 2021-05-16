package config

import "anest.top/youli/api/proto/config"

func NewConfig() *config.Config {
	return &config.Config{
		ServerConf: &config.ServerConfig{
			Name: "MyApp",
			Host: ":8001",
		},
		MysqlConf: &config.MysqlConfig{
			Dsn:           "root:root@tcp(172.16.16.123:3306)/dpro?charset=utf8mb4&parseTime=True&loc=Local",
			MaxIdleConn:   5,
			MaxOpenConn:   30,
			MaxLifeMinute: 0,
			Debug:         true,
		},
		RedisConf: nil,
	}
}
