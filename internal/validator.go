package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PaesslerAG/gval"
	"github.com/iot-master-contrib/alarm/types"
	"github.com/zgwit/iot-master/v3/pkg/calc"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"time"
)

type Validator struct {
	model      *types.Validator
	expression gval.Evaluable
	//again      bool
	start int64 //开始时间s
	total uint  //报警次数
}

type Device struct {
	values     map[string]interface{}
	validators []Validator
}

func (d *Device) Push(pid, id string, ctx map[string]interface{}) {
	for k, v := range ctx {
		d.values[k] = v
	}
	now := time.Now().Unix()

	for _, v := range d.validators {
		ret, err := v.expression.EvalBool(context.Background(), d.values)
		if err != nil {
			log.Error(err)
			continue
		}

		if ret {
			//约束OK，检查下一个
			v.total = 0
			v.start = 0
			continue
		}

		//now := time.Now().Unix()
		if v.start == 0 {
			v.start = now
		}

		cs := v.model

		//延迟报警
		if cs.Delay > 0 {
			if now < v.start+int64(cs.Delay) {
				continue
			}
		}

		//重复报警
		if cs.Again > 0 && v.total < cs.Total {
			if now < v.start+int64(cs.Again) {
				continue
			}

			//重置开始时间
			v.start = now // + int64(cs.Delay)
			v.total++
		}

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
			//continue
		}

		topic := fmt.Sprintf("alarm/%s/%s", pid, id)
		payload, _ := json.Marshal(&alarm)
		err = mqtt.Publish(topic, payload, false, 0)
		if err != nil {
			log.Error(err)
		}
	}
}

type Product struct {
	//validators []Validator
	validators []*types.Validator
	devices    lib.Map[Device]
}

func (p *Product) Push(pid, id string, ctx map[string]interface{}) {
	dev := p.devices.Load(id)
	if dev == nil {
		//加载设备
		//dev.validators
		for _, v := range p.validators {
			vv := &Validator{
				model:      v,
				expression: nil,
			}
			eval, err := calc.New(v.Expression) //重复编译
			if err != nil {
				log.Error(err)
				continue
			}
			vv.expression = eval
		}
		p.devices.Store(id, dev)
	}
	dev.Push(pid, id, ctx)
}

var products lib.Map[Product]

func Push(pid, id string, ctx map[string]interface{}) {
	p := products.Load(pid)
	if p == nil {
		//TODO 加载项目？？？？
		return
	}
	p.Push(pid, id, ctx)
}

func LoadValidator(validator *types.Validator) error {
	p := products.Load(validator.ProductId)
	if p == nil {
		p = &Product{}
		products.Store(validator.ProductId, p)
	}

	p.validators = append(p.validators, validator)
	//TODO 统一编译表达式
	//evaluable, err := calc.New(validator.Expression)

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
