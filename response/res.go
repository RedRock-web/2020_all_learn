//@program: work
//@description: 
//@author: edte
//@create: 2020-05-30 18:09
package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 10000})
}

func FormError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 10001, "message": "request form error!"})
}

func OkWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": 10000, "message": "ok", "data": data})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
}
