package api

import (
	"alarm/types"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
)

// @Summary 查询报警
// @Schemes
// @Description 查询报警
// @Tags alarm
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回报警信息
// @Router /alarm/count [post]
func noopAlarmCount() {}

// @Summary 查询报警
// @Schemes
// @Description 查询报警
// @Tags alarm
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Alarm] 返回报警信息
// @Router /alarm/search [post]
func noopAlarmSearch() {}

// @Summary 查询报警
// @Schemes
// @Description 查询报警
// @Tags alarm
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Alarm] 返回报警信息
// @Router /alarm/list [get]
func noopAlarmList() {}

// @Summary 删除报警
// @Schemes
// @Description 删除报警
// @Tags alarm
// @Param id path int true "报警ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Alarm] 返回报警信息
// @Router /alarm/{id}/delete [get]
func noopAlarmDelete() {}

// @Summary 阅读报警
// @Schemes
// @Description 阅读报警
// @Tags alarm
// @Param id path int true "报警ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Alarm] 返回报警信息
// @Router /alarm/{id}/read [get]
func noopAlarmRead() {}

func alarmRead(ctx *gin.Context) {
	alarm := types.Alarm{
		Read: true,
	}
	cnt, err := db.Engine.ID(ctx.GetInt64("id")).Cols("read").Update(alarm)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, cnt)
}

func alarmRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[types.Alarm]())

	app.POST("/search", curd.ApiSearchWith[types.Alarm](
		"alarm", []curd.Join{{
			Table:        "device",
			LocaleField:  "device_id",
			ForeignField: "id",
			Field:        "name",
			As:           "device",
		}},
	))

	app.GET("/list", curd.ApiList[types.Alarm]())

	app.GET("/:id", curd.ParseParamId, curd.ApiGet[types.Alarm]())

	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDelete[types.Alarm]())

	app.GET("/:id/read", curd.ParseParamId, alarmRead)
}
