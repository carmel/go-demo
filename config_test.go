package core

import (
	"encoding/json"
	"os"
	"sync"
)

var once sync.Once

//数据库配置
type DBConfig struct {
	DiverName string
}

var (
	dbconf *DBConfig
)

func init() {
	GetDBConfig()
}

func GetDBConfig() *DBConfig {
	once.Do(func() {
		dbconf = &DBConfig{}
		setConfig("../conf/db", dbconf)
	})
	return dbconf
}

func setConfig(path string, conf interface{}) {
	file, err := os.Open(path)
	ErrorHandler(err, false)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(dbconf)
	ErrorHandler(err, false)
}
