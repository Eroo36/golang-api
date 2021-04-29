package test

import (
	"net/http"
	// "os"
	"github.com/gin-gonic/gin"
)

func Test(res *gin.Context) {
	res.String(http.StatusOK, "TEST")
}
