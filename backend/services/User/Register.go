package User

import (
	"github.com/gin-gonic/gin"
	"log"
)

func UserRegister(c *gin.Context) {

	if err := c.BindJSON(&user); err == nil {
		log.Print(err)
		result := dbConn.Find(&user, "EmpId=?", user.EmpId)

		// 이미 이메일이 존재할 경우의 처리
		if result.RowsAffected != 0 {
			c.JSON(400, gin.H{
				"status":  400,
				"message": "existing email",
			})
			return
		}
		user.Password = GenerateHashPassword(user.Password)
		dbErr := dbConn.Debug().Create(&user).Error

		if dbErr == nil {
			dbConn.Debug().Create(&user)
			c.JSON(200, gin.H{
				"status":  "success",
				"message": "register is success"})
			return
		} else {
			c.JSON(500, gin.H{
				"status":  "fail",
				"message": "can not success this data"})
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
