package userControllers

import (
	"fmt"
	"net/http"

	"example.com/models"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/goonode/mogo"
)

func SignupUser(c *gin.Context) {

	var json models.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// var name = fmt.Sprintf(json.Name)
	// fmt.Println(name)
	user := mogo.NewDoc(models.User{}).(*models.User)

	errDB := user.FindOne(bson.M{"email": json.Email}, user)
	if errDB == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "user exists"})
		return
	}

	user.Email = json.Email
	user.Password = json.Password
	user.Name = json.Name

	err := user.Save()
	if vErr, ok := err.(*mogo.ValidationError); ok {
		fmt.Println("Validation errors are:", vErr.Errors)
		return
	} else {
		c.String(http.StatusOK, "User Created")
		return
	}

}
