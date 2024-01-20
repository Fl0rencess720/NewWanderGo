package routers

import (
	"SparkForge/oss"
	ini "WanderGo/controller/init"
	pos "WanderGo/controller/position"
	user "WanderGo/controller/user"
	mid "WanderGo/middlewares"
	util "WanderGo/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
)

func Start() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("ruleOfPwd", util.RuleOfPwd)
	}
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	engine := gin.Default()
	engine.Use(mid.Cors())

	engine.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()
		userAccount := "get_user_account_somehow" // 从请求中获取用户账号或其他方式获取
		util.AddConnection(userAccount, conn)
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				util.RemoveConnection(userAccount)
				break
			}
		}
	})
	engine.POST("/register/captcha", user.SendEmail)
	engine.POST("/register", user.RegisterHandler)
	engine.POST("/login", user.LoginHandler)
	engine.POST("/exit", user.ExitHandler)
	engine.POST("/names/change", mid.LoginVerification(), user.ChangeNameHandler)
	engine.POST("/passwords/change", mid.LoginVerification(), user.ChangePwdHandler)
	engine.POST("/passwords/forget/captcha", user.ForgotPasswordGetCaptcha)
	engine.POST("/passwords/forget", user.ForgotPassword)
	engine.POST("/avatars/upload", mid.LoginVerification(), user.AvatarUpload)
	engine.POST("/avatars/change", mid.LoginVerification(), user.AvatarChange)
	engine.POST("/avatars/load", user.SendAvatarToFrontend)
	engine.POST("/comments/post", mid.LoginVerification(), user.PostComment)
	engine.POST("/comments/roam", mid.LoginVerification(), pos.Roaming)
	engine.POST("/comments/like", mid.LoginVerification(), user.LikeHandler)
	engine.POST("/test", user.TestComments)
	engine.POST("/places/mark", pos.MarkPlace)
	engine.POST("/sts/get", oss.GetSTS)
	engine.POST("/places/get", pos.PositionsHandler)
	engine.POST("/begin/user", ini.LoadPersonalInformation)
	engine.POST("/begin/places", ini.LoadPlacesInformation)
	engine.Run(":8080")
}
