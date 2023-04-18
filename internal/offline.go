package internal

import (
	"alarm/types"
	"encoding/json"
	"fmt"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"strings"
)

func SubscribeOffline() {

	mqtt.Client.Subscribe("offline/+/+", 0, func(client paho.Client, message paho.Message) {
		topics := strings.Split(message.Topic(), "/")
		pid := topics[1]
		id := topics[2]

		alarm := types.Alarm{
			DeviceId: id, //TODO 配置化
			Type:     "掉线",
			Title:    "掉线",
			Level:    3,
		}
		_, err := db.Engine.Insert(&alarm)
		if err != nil {
			log.Error(err)
			//continue
		}

		topic := fmt.Sprintf("alarm/%s/%s", pid, id)
		payload, _ := json.Marshal(&alarm)
		err = mqtt.Publish(topic, payload, false, 0)
	})

}
