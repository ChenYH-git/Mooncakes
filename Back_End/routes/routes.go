package routes

import (
	"Mooncakes/controllers"
	"Mooncakes/logger"
	"Mooncakes/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	controllers.InitTrans("zh")

	r := gin.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.RateLimitMiddleware(500*time.Millisecond, 25))
	r.LoadHTMLGlob("./template/*")
	r.Static("/static", "./static")

	{
		r.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})

		r.GET("/sign", func(c *gin.Context) {
			c.HTML(http.StatusOK, "sign.html", nil)
		})

		r.GET("/prepare", func(c *gin.Context) {
			c.HTML(http.StatusOK, "prepare.html", nil)
		})

		r.GET("/ready", func(c *gin.Context) {
			c.HTML(http.StatusOK, "ready.html", nil)
		})

		r.GET("/start", func(c *gin.Context) {
			c.HTML(http.StatusOK, "start.html", nil)
		})

		r.GET("/dice", func(c *gin.Context) {
			c.HTML(http.StatusOK, "dice.html", nil)
		})

		r.GET("/result", func(c *gin.Context) {
			c.HTML(http.StatusOK, "result.html", nil)
		})

		r.GET("/rank", func(c *gin.Context) {
			c.HTML(http.StatusOK, "rank.html", nil)
		})
	}

	v2 := r.Group("api/v2")

	v2.POST("/signup", controllers.SignUpHandler)
	v2.POST("/login", controllers.LoginHandler)

	v2.Use(middleware.JWTAuthMiddleware())
	{
		v2.GET("/rank", controllers.GetRankHandler)
		v2.PUT("/dice", controllers.DiceResHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "404 NOT FOUND",
		})
	})
	return r
}
