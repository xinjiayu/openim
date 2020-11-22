package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"openim/app/api/chat"
	"openim/app/service/middleware"
	"openim/app/web"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		s.BindMiddlewareDefault(middleware.CORS)
		//group.GET("/chat",chat.Chat)
		s.BindHandler("/index", web.Index)
		group.ALL("/ws", chat.Websocket)

		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.White, middleware.Auth)
			group.GET("/history", chat.GetHistory)
			group.GET("/newcount", chat.GetNewTopicCount)

		})

	})
}
