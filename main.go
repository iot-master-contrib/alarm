package alarm

import (
	"embed"
	"encoding/json"
	"github.com/iot-master-contrib/alarm/api"
	"github.com/iot-master-contrib/alarm/config"
	_ "github.com/iot-master-contrib/alarm/docs"
	"github.com/iot-master-contrib/alarm/internal"
	"github.com/iot-master-contrib/alarm/types"
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
}

func Startup(app *web.Engine) error {
	banner.Print("iot-master-plugin:alarm")
	build.Print()

	config.Load()

	err := log.Open()
	if err != nil {
		return err
	}

	//加载数据库
	err = db.Open()
	if err != nil {
		return err
	}
	//defer db.Close()

	//同步表结构
	err = db.Engine.Sync2(
		new(types.Alarm), new(types.Validator),
	)
	if err != nil {
		return err
	}

	//MQTT总线
	err = mqtt.Open()
	if err != nil {
		return err
	}
	//defer mqtt.Close()

	internal.SubscribeProperty(mqtt.Client)
	internal.SubscribeOffline()

	//加载验证器
	err = internal.LoadValidators()
	if err != nil {
		return err
	}

	//注册前端接口
	api.RegisterRoutes(app.Group("/app/alarm/api"))

	//注册接口文档
	web.RegisterSwaggerDocs(app.Group("/app/alarm"))

	return nil
}

func Register() error {
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
		Icon:    "/app/alarm/assets/alarm.svg",
	})
	return mqtt.Publish("master/register", payload, false, 0)
}

func Static(fs *web.FileSystem) {
	//前端静态文件
	fs.Put("/app/alarm", http.FS(wwwFiles), "", "index.html")
}

func Shutdown() error {

	//只关闭Web就行了，其他通过defer关闭

	return nil
}
