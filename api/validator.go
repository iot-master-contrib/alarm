package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iot-master-contrib/alarm/types"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

// @Summary 查询检查数量
// @Schemes
// @Description 查询检查数量
// @Tags validator
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回检查数量
// @Router /validator/count [post]
func noopValidatorCount() {}

// @Summary 查询检查
// @Schemes
// @Description 查询检查
// @Tags validator
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Validator] 返回检查信息
// @Router /validator/search [post]
func noopValidatorSearch() {}

// @Summary 查询检查
// @Schemes
// @Description 查询检查
// @Tags validator
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Validator] 返回检查信息
// @Router /validator/list [get]
func noopValidatorList() {}

// @Summary 创建检查
// @Schemes
// @Description 创建检查
// @Tags validator
// @Param search body types.Validator true "检查信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Validator] 返回检查信息
// @Router /validator/create [post]
func noopValidatorCreate() {}

// @Summary 修改检查
// @Schemes
// @Description 修改检查
// @Tags validator
// @Param id path int true "检查ID"
// @Param validator body types.Validator true "检查信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Validator] 返回检查信息
// @Router /validator/{id} [post]
func noopValidatorUpdate() {}

// @Summary 获取检查
// @Schemes
// @Description 获取检查
// @Tags validator
// @Param id path int true "检查ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Validator] 返回检查信息
// @Router /validator/{id} [get]
func noopValidatorGet() {}

// @Summary 删除检查
// @Schemes
// @Description 删除检查
// @Tags validator
// @Param id path int true "检查ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Validator] 返回检查信息
// @Router /validator/{id}/delete [get]
func noopValidatorDelete() {}

// @Summary 启用检查
// @Schemes
// @Description 启用检查
// @Tags validator
// @Param id path int true "检查ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Validator] 返回检查信息
// @Router /validator/{id}/enable [get]
func noopValidatorEnable() {}

// @Summary 禁用检查
// @Schemes
// @Description 禁用检查
// @Tags validator
// @Param id path int true "检查ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Validator] 返回检查信息
// @Router /validator/{id}/disable [get]
func noopValidatorDisable() {}

// @Summary 导出检查
// @Schemes
// @Description 导出检查
// @Tags product
// @Accept json
// @Produce octet-stream
// @Router /validator/export [get]
func noopValidatorExport() {}

// @Summary 导入检查
// @Schemes
// @Description 导入检查
// @Tags product
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回检查数量
// @Router /validator/import [post]
func noopValidatorImport() {}

func validatorRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[types.Validator]())
	app.POST("/search", curd.ApiSearch[types.Validator]())
	app.GET("/list", curd.ApiList[types.Validator]())
	app.POST("/create", curd.ApiCreateHook[types.Validator](curd.GenerateRandomId[types.Validator](8), nil))
	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[types.Validator]())
	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[types.Validator](nil, nil,
		"id", "product_id", "expression", "type", "title",
		"template", "level", "delay", "again", "total", "disabled"))
	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[types.Validator](nil, nil))
	app.GET("/export", curd.ApiExport("validator", "验证器"))
	app.POST("/import", curd.ApiImport("validator"))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisableHook[types.Validator](true, nil, nil))
	app.GET(":id/enable", curd.ParseParamStringId, curd.ApiDisableHook[types.Validator](false, nil, nil))
}
