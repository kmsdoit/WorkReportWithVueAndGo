package User

import (
	"github.com/gin-gonic/gin"
	"log"
)

func UserRegister(c *gin.Context) {

	if err := c.BindJSON(&user); err == nil {
		log.Print(err)
		userEmailValidation := isEmailValid(user.Email)
		if userEmailValidation == true {
			user.Password = GenerateHashPassword(user.Password)

			dbErr := dbConn.Debug().Create(&user).Error
			if dbErr == nil {
				dbConn.Debug().Create(&user)
				c.JSON(200, gin.H{
					"status":  "success",
					"message": "register is success",
				})
			} else {
				c.JSON(500, gin.H{
					"status":  "fail",
					"message": "can not success this data",
				})
				return
			}
		} else {
			c.JSON(400, gin.H{
				"status":  "fail",
				"message": "is not Email form",
			})
			return
		}
	} else {
		c.JSON(400, gin.H{
			"status":  "fail",
			"message": "bad request",
		})
		return
	}
}
