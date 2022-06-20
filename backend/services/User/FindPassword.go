package User

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func FindPassword(c *gin.Context) {
	var email = c.Request.URL.Query().Get("email")

	findEmail := dbConn.Where("email =?", email).Find(&user)

	if findEmail.Error != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "해당 사용자를 찾을 수 없습나다",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": user.Password,
	})

}

func UpdatePassword(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var data map[string]interface{}
	Jsonerr := json.Unmarshal([]byte(value), &data)
	
	if Jsonerr != nil {
		return
	}
	updatePassword := dbConn.Debug().Where("password = ?", data["password"]).Save(&user)

	if updatePassword.Error != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Bad Request",
		})
		return
	}

	findEmail := dbConn.Where("email =?", data["email"]).Find(&user)

	if findEmail.Error == nil {
		user.Password = GenerateHashPassword(user.Password)
		c.JSON(200, gin.H{
			"status":  200,
			"message": "비밀번호 변경 완료",
		})
	} else {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Bad Request",
		})
		return
	}

}
