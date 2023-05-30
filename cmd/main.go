package main

import (
	"github.com/iot-master-contrib/alarm"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/web"
)

// @title 历史数据接口文档
// @version 1.0 版本
// @description API文档
// @BasePath /app/alarm/api/
// @query.collection.format multi
func main() {
	app := web.CreateEngine()

	//调用启动
	err := alarm.Startup(app)
	if err != nil {
		log.Fatal(err)
	}

	//注册静态页面
	fs := app.FileSystem()
	alarm.Static(fs)

	//监听HTTP
	app.Serve()
}
