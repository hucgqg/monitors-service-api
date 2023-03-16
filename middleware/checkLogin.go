package middleware

import (
	"encoding/base64"
	"monitors-service-api/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		flag := doAuth(c)
		if flag {
			c.Next()
		} else {
			requireAuth(c)
			c.Abort()
		}
	}
}

func doAuth(c *gin.Context) bool {
	header := c.GetHeader("Authorization")
	s := strings.SplitN(header, " ", 2)
	if len(s) != 2 || s[0] != "Basic" {
		return false
	}
	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return false
	}
	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return false
	}

	return checkPassword(pair[0], pair[1])
}

func requireAuth(c *gin.Context) {
	resp := models.RespData{T: time.Now().Format("2006-01-02 15:04:05:000")}
	resp.Msg = "缺少 Basic auth"
	c.Header("WWW-Authenticate", "Basic realm=\"private\"")
	c.JSON(http.StatusUnauthorized, &resp)
}

func checkPassword(username, password string) bool {
	return username == "admin" && password == "123456"
}
