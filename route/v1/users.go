package v1

import (
	"NASDAQ_Slot_Machine/controller"
	_ "NASDAQ_Slot_Machine/docs"
	"NASDAQ_Slot_Machine/middleware"
	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"time"
)


func RequestIDMiddleware(ctx *gin.Context) {
	uuidV4 := uuid.NewV4()
	ctx.Header("X-Request-Id", uuidV4.String())

	ctx.Next()
}



func RouteUsers(r *gin.Engine, m *persist.RedisStore) {
	posts := r.Group("/users")
	posts.Use(RequestIDMiddleware)
	{
		posts.POST("/register/", controller.NewUsersController().RegisterOne)
		posts.POST("/login/", controller.OldUsersController().LoginOne)
	}

	auth := r.Group("/users/auth")
	auth.Use(middleware.JWTAuth())
	auth.Use(RequestIDMiddleware)
	{
		auth.GET("/:id", cache.CacheByRequestURI(m, 2*time.Second), controller.NewUsersController().GetOne)
	}
}
