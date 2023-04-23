package main

import (
	"alarm/api"
	"alarm/config"
	_ "alarm/docs"
	"alarm/internal"
	"alarm/types"
	"embed"
	"encoding/json"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/banner"
	"github.com/zgwit/iot-master/v3/pkg/build"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/web"
	"net/http"
)

//go:embed all:app/alarm
var wwwFiles embed.FS

// @title 历史数据接口文档
// @version 1.0 版本
// @description API文档
// @BasePath /app/alarm/api/
// @query.collection.format multi
func main() {
	banner.Print("iot-master-plugin:alarm")
	build.Print()

	config.Load()

	err := log.Open()
	if err != nil {
		log.Fatal(err)
	}

	//加载数据库
	err = db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//同步表结构
	err = db.Engine.Sync2(
		new(types.Alarm), new(types.Validator),
	)
	if err != nil {
		log.Fatal(err)
	}

	//MQTT总线
	err = mqtt.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer mqtt.Close()

	//注册应用
	//for _, v := range config.Config.Apps {
	payload, _ := json.Marshal(model.App{
		Id:   "alarm",
		Name: "报警管理",
		Entries: []model.AppEntry{{
			Path: "app/alarm/alarm",
			Name: "报警日志",
		}, {
			Path: "app/alarm/validator",
			Name: "报警检查",
		}},
		Type:    "tcp",
		Address: "http://localhost" + web.GetOptions().Addr,
	})
	_ = mqtt.Publish("master/register", payload, false, 0)
	//}

	internal.SubscribeProperty(mqtt.Client)
	internal.SubscribeOffline()

	//加载验证器
	err = internal.LoadValidators()
	if err != nil {
		log.Fatal(err)
	}

	app := web.CreateEngine()

	//注册前端接口
	api.RegisterRoutes(app.Group("/app/alarm/api"))

	//注册接口文档
	web.RegisterSwaggerDocs(app.Group("/app/alarm"))

	//前端静态文件
	app.RegisterFS(http.FS(wwwFiles), "", "app/alarm/index.html")

	//监听HTTP
	app.Serve()
}
