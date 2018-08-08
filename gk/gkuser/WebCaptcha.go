package gkuser

import (
	"github.com/gin-gonic/gin"

	"github.com/dchest/captcha"
	"github.com/ecdiy/itgeek/gk/ws"
)

func Captcha(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Cache-Control", "no-Cache, no-store, must-revalidate")
	c.Header("Pragma", "no-Cache")
	c.Header("Expires", "0")
	c.Header("Content-Type", "image/png")
	id := c.Query("t")
	captcha.WriteImage(c.Writer, id, captcha.StdWidth, captcha.StdHeight)
}

func CaptchaNew(web *ws.Web) {
	//c.JSON(200, ws.OK.Result(captcha.New()))
	web.Result(captcha.New())
}
