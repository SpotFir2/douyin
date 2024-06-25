package service

import (
	"strconv"

	"github.com/SpotFir2/douyin/model"
	"github.com/SpotFir2/douyin/pkg/config"
	"github.com/SpotFir2/douyin/pkg/snowflake"
)

// 初始化config,model,snowflake
func Init() {
	//从config/config.yaml处加载配置文件,如失败则在config/config.example.yaml生成配置文件
	if err := config.LoadConfig("config/config.yaml"); err != nil {
		config.InitConfig("config/config.example.yaml")
		panic(err)
	}

	//从config中读取数据库信息并链接
	dsn := config.Confi.MySQL.User + ":" + config.Confi.MySQL.Password + "@tcp(" + config.Confi.MySQL.IP + ":" + strconv.Itoa(config.Confi.MySQL.Port) + ")/" + config.Confi.MySQL.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	if err := model.LoadDatabase(dsn); err != nil {
		panic(err)
	}

	//初始化雪花算法
	if err := snowflake.Init("2024-06-24 18:55:00", 1); err != nil {
		panic(err)
	}
}
