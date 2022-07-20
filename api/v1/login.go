package v1

import (
	"fmt"
	"gintest/model"
	"gintest/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginFront(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	var code int
	fmt.Println("formData.Username")
	fmt.Println(formData.Username)

	formData, code = model.CheckLoginFront(formData.Username, formData.Password)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    formData.Username,
		"id":      formData.ID,
		"message": errmsg.GetErrMsg(code),
	})
}
