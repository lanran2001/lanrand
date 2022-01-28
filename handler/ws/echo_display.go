package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DisplayEcho(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "views/websockets.html")
}
