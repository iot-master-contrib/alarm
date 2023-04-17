package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(app *gin.RouterGroup) {

	alarmRouter(app.Group("/alarm"))

	inspectorRouter(app.Group("/inspector"))

}
