package config

import (
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/web"
	"gopkg.in/yaml.v2"
	"os"
)

type Configure struct {
	Crontab  string       `json:"crontab"`
	Web      web.Options  `json:"web"`
	Database db.Options   `json:"database"`
	Mqtt     mqtt.Options `json:"mqtt"`
	Log      log.Options  `json:"log"`
	Apps     []model.App  `json:"apps"`
}

var Config = Configure{
	Crontab:  "0 * * * *", //默认 1h
	Web:      web.Default(),
	Database: db.Default(),
	Mqtt:     mqtt.Default(),
	Log:      log.Default(),
	Apps: []model.App{
		{
			Id:      "alarm",
			Name:    "报警",
			Address: "http://localhost:40007",
			Entries: []model.AppEntry{
				{Name: "报警记录", Path: "alarm"},
				{Name: "报警检查", Path: "validator"},
				{Name: "配置", Path: "config"},
			},
		},
	},
}

func init() {
	Config.Web.Addr = ":40007"
	//Config.Database.URL = "root:root@tcp(git.zgwit.com:3306)/modbus?charset=utf8"
	//TODO get imei sn
}

// Load 加载
func Load(filename string) error {
	// 如果没有文件，则使用默认信息创建
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return Store(filename)
	} else {
		y, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer y.Close()

		d := yaml.NewDecoder(y)
		err = d.Decode(&Config)
		if err != nil {
			log.Fatal(err)
			return err
		}

		return nil
	}
}

func Store(filename string) error {
	y, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755) //os.Create(filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer y.Close()

	e := yaml.NewEncoder(y)
	defer e.Close()

	err = e.Encode(&Config)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
