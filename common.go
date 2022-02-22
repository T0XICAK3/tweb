package tweb

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"tweb/utils"
)

type ToxicWeb struct {
	postRouter map[string]func(*gin.Context)
	getRouter  map[string]func(*gin.Context)
	engine     *gin.Engine
	store      cookie.Store
}

func (tw *ToxicWeb) Start(address string) {
	for path, function := range tw.getRouter {
		tw.engine.GET(path, function)
	}
	for path, function := range tw.postRouter {
		tw.engine.POST(path, function)
	}
	err := tw.engine.Run(address) //:8080
	if err != nil {
		utils.ConsoleOut("TOXIC_WEB_RUN_ERR", err.Error(), utils.Red, utils.White+1)
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}

func NewToxicWeb(getRouter, postRouter map[string]func(*gin.Context), secret, sessionName, staticPath string) *ToxicWeb {
	if _, ok := getRouter["/"]; !ok {
		getRouter["/"] = func(c *gin.Context) {
			c.Header("Content-Type", "text/html; charset=utf-8")
			page := NewTextPage(200, "<br><br><br><center><img src=\"data:image/png;base64,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  \n            AABJRU5ErkJggg==\"/><br><br><head>Toxic Web Framework is READY!</head><br>Powered by gin</center>", gin.H{}, c)
			page.Show()
		}
	}
	tw := &ToxicWeb{
		getRouter:  getRouter,
		postRouter: postRouter,
	}
	tw.engine = gin.Default()
	tw.engine.Delims("{|", "|}")
	tw.store = cookie.NewStore([]byte(secret)) //secret
	tw.engine.Use(sessions.Sessions(sessionName, tw.store))
	//vw.engine.LoadHTMLFiles("static/index.html")
	if ok, _ := utils.PathExists(staticPath + "/html"); ok {
		tw.engine.LoadHTMLGlob(staticPath + "/html/*")
	} else if ok, _ := utils.PathExists(staticPath + "/css"); ok {
		tw.engine.Static("/css", staticPath+"/css")
	} else if ok, _ := utils.PathExists(staticPath + "/css"); ok {
		tw.engine.Static("/js", staticPath+"/js")
	} else if ok, _ := utils.PathExists(staticPath + "/favicon.ico"); ok {
		tw.engine.StaticFile("favicon.ico", staticPath+"/favicon.ico")
	}
	if ok, _ := utils.PathExists(staticPath + "/html/500.html"); ok {
		tw.engine.Use(errorPage)
	} else {
		tw.engine.Use(defaultErrorPage)
	}
	if ok, _ := utils.PathExists(staticPath + "/html/404.html"); ok {
		tw.engine.NoRoute(func(c *gin.Context) {
			c.HTML(http.StatusOK, "404.html", gin.H{
				"title": "404",
			})
		})
	} else {
		tw.engine.NoRoute(func(c *gin.Context) {
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(404, "<center><br><br><head>404 Not Found</head></center>")
		})
	}

	return tw
}
