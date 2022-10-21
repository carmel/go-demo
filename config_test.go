package demo

import (
	"encoding/json"
	"fmt"
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

func Init() {
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
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(dbconf)
	if err != nil {
		fmt.Println(err)
	}
}
