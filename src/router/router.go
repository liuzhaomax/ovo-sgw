package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzhaomax/ovo-sgw/internal/middleware"
	"github.com/liuzhaomax/ovo-sgw/src/utils"
)

func Register(root *gin.RouterGroup, mw *middleware.Middleware) {
	root.GET("/login", mw.ReverseProxy.Redirect(utils.UserMicroserviceName))
	root.POST("/login", mw.ReverseProxy.Redirect(utils.UserMicroserviceName))

	root.Use(mw.Auth.ValidateToken())
	root.DELETE("/login", mw.ReverseProxy.Redirect(utils.UserMicroserviceName))
	routerUser := root.Group("/users")
	{
		routerUser.Any("/*url", mw.ReverseProxy.Redirect(utils.UserMicroserviceName))
	}
	routerFile := root.Group("/file")
	{
		routerFile.Any("/*url", mw.ReverseProxy.Redirect(utils.FileMicroserviceName))
	}
	routerStats := root.Group("/stats")
	{
		routerStats.Any("/*url", mw.ReverseProxy.Redirect(utils.StatsMicroserviceName))
	}
	routerKanban := root.Group("/kanban")
	{
		routerKanban.Any("/*url", mw.ReverseProxy.Redirect(utils.KanbanMicroserviceName))
	}
	routerDoc := root.Group("/doc")
	{
		routerDoc.Any("/*url", mw.ReverseProxy.Redirect(utils.DocMicroserviceName))
	}
}
