package tweb

import (
	"github.com/T0XICAK3/tweb/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

type Page struct {
	Templ      string
	StatusCode int
	Data       gin.H
	Session    sessions.Session
	c          *gin.Context
	pageType   string
}

func (pg *Page) SetSession(name, value string) {
	pg.Session.Set(name, value)
	err := pg.Session.Save()
	if err != nil {
		utils.ConsoleOut("SESSION_SAVE_ERR", err.Error(), utils.Red, utils.White+1)
		return
	}
}

func (pg *Page) GetSession(name string) string {
	return pg.Session.Get(name).(string)
}

func (pg *Page) SetData(name, value string) {
	pg.Data[name] = value
}

func (pg *Page) Redirect(statusCode int, location string) {
	pg.c.Redirect(statusCode, location)
}

func (pg *Page) Show() {
	switch pg.pageType {
	case "json":
		pg.c.JSON(pg.StatusCode, pg.Data)
	case "html":
		pg.c.HTML(pg.StatusCode, pg.Templ, pg.Data)
	case "text":
		pg.c.String(pg.StatusCode, pg.Templ)
	default:
		pg.c.String(pg.StatusCode, "PAGE_WRONG_TYPE")
	}
}

func NewPage(statusCode int, templ string, data gin.H, c *gin.Context) *Page {
	session := sessions.Default(c)
	return &Page{
		Templ:      templ,
		StatusCode: statusCode,
		Data:       data,
		Session:    session,
		c:          c,
		pageType:   "html",
	}
}

func errorPage(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic: %v\n", err)
			//debug.PrintStack()
			c.HTML(500, "500.html", gin.H{
				"error": err,
			})
		}
	}()
	c.Next()
}

func defaultErrorPage(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", err)
			//debug.PrintStack()
			//封装通用json返回
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(500, "<center><br><br><head>500 Server Error</head></center>")
		}
	}()
	c.Next()
}

func NewTextPage(statusCode int, text string, data gin.H, c *gin.Context) *Page {
	session := sessions.Default(c)
	return &Page{
		Templ:      text,
		StatusCode: statusCode,
		Data:       data,
		Session:    session,
		c:          c,
		pageType:   "text",
	}
}

func NewJsonPage(statusCode int, data gin.H, c *gin.Context) *Page {
	session := sessions.Default(c)
	return &Page{
		Templ:      "JSON_NO_TEMPL",
		StatusCode: statusCode,
		Data:       data,
		Session:    session,
		c:          c,
		pageType:   "json",
	}
}

/*
func p404(c *gin.Context){
	page:=NewPage(404,)
}
*/
