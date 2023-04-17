package internal

import (
	"alarm/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/PaesslerAG/gval"
	"github.com/zgwit/iot-master/v3/pkg/calc"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
)

type Validator struct {
	model      *types.Validator
	expression gval.Evaluable
}

type Device struct {
	values     map[string]interface{}
	validators []Validator
}

func (d *Device) Push(pid, id string, ctx map[string]interface{}) {
	for k, v := range ctx {
		d.values[k] = v
	}
	for _, v := range d.validators {
		res, err := v.expression.EvalBool(context.Background(), d.values)
		if err != nil {
			log.Error(err)
			continue
		}
		if !res {
			alarm := types.Alarm{
				DeviceId: id,
				Type:     v.model.Type,
				Title:    v.model.Title,
				Level:    v.model.Level,
				//Message:  v.model.Message, //TODO 模板格式化
			}
			_, err = db.Engine.Insert(&alarm)
			if err != nil {
				log.Error(err)
				continue
			}
			topic := fmt.Sprintf("alarm/%s/%s", pid, id)
			payload, _ := json.Marshal(&alarm)
			err = mqtt.Publish(topic, payload, false, 0)
			if err != nil {
				log.Error(err)
			}
		}
	}
}

type Project struct {
	//validators []Validator
	validators []*types.Validator
	devices    lib.Map[Device]
}

func (p *Project) Push(pid, id string, ctx map[string]interface{}) {
	dev := p.devices.Load(pid)
	if dev == nil {
		//TODO 加载设备
	}
	dev.Push(pid, id, ctx)
}

var projects lib.Map[Project]

func LoadValidator(validator *types.Validator) error {
	//log.Info("load validator", validator.Id, validator.Name)
	v := &Validator{
		model: validator,
	}
	evaluable, err := calc.New(validator.Expression)
	if err != nil {
		return err
	}
	v.expression = evaluable
	//projects.Store(validator.Id, v)

	return nil
}

func LoadValidators() error {
	var vs []*types.Validator
	err := db.Engine.Find(&vs)
	if err != nil {
		return err
	}

	for _, p := range vs {
		err = LoadValidator(p)
		if err != nil {
			log.Error(err)
			//return err
		}
	}

	return nil
}
