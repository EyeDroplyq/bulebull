package main

import (
	"bulebell/controller"
	"bulebell/dao/mysql"
	"bulebell/dao/redis"
	"bulebell/logger"
	"bulebell/pkg/snowflake"
	"bulebell/router"
	"bulebell/setting"
	"fmt"
)

// @title bulebell项目接口文档
// @version 1.0
// @description Go web开发进阶项目实战课程bulebell

// @host 127.0.0.1:8084
// @BasePath /api/v1
func main() {
	// 加载配置
	if err := setting.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	//初始化validator的错误信息翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans, err:%v\n", err)
		return
	}
	// 注册路由
	r := router.SetupRouter(setting.Conf.Mode)
	err := r.Run(fmt.Sprintf("127.0.0.1:%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
