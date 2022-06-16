package User

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

//로그인API 및 AccessToken 및 RefreshToken 부여
func UserLogin(c *gin.Context) {

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "bad request",
		})
	}

	inputPassword := user.Password

	result := dbConn.Debug().Find(user, "email = ?", user.Email)

	if result.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "DB can't find this user email",
		})
	}

	passwordCompare := ComparePassword(user.Password, inputPassword)

	if passwordCompare == false {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "허가하지 않는 계정입니다",
		})
	} else {
		accessToken, accessTokenErr := CreateJWT(user.Email)
		refreshToken, refreshTokenErr := CreateRefreshToken(user.Email)
		if accessTokenErr != nil && refreshTokenErr != nil {
			c.JSON(500, gin.H{
				"status":  500,
				"message": "server Error Pleaze check your url",
			})
		} else {
			c.SetCookie("access_token", accessToken, 60*60*24, "/", "localhost:8003", false, true)
			c.SetCookie("refresh_token", refreshToken, 60*60*24, "/", "localhost:8003", false, true)
			c.JSON(200, map[string]string{
				"message":       "토큰 발급 완료",
				"access_token":  accessToken,
				"refresh_token": refreshToken,
			})
		}
	}
}

// AccessToken 검증
func VerifyAccessToken(c *gin.Context) {
	cToken, err := c.Request.Cookie("access_token")
	if err != nil {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "Get Cookie failed",
		})
		c.Abort()
		return
	}

	tokenValue := cToken.Value

	if tokenValue == "" {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "token is not exist",
		})

		c.Abort()
		return
	}

	claims := jwt.MapClaims{}

	token, tokenErr := jwt.ParseWithClaims(tokenValue, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySignKey), nil
	})

	fmt.Printf("token : %v\n", token)

	if tokenErr != nil {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "토큰 인증 실패",
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "토큰 인증 완료",
		})
	}

	return
}

// RefreshToken 검증
func VerifyRefreshToken(c *gin.Context) {
	refreshToken, err := c.Request.Cookie("refresh-token")
	if err != nil {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "Get Cookie failed",
		})
		return
	}

	tokenValue := refreshToken.Value

	if tokenValue == "" {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "refresh token is None",
		})

		return
	}

	claims := jwt.MapClaims{}

	token, tokenErr := jwt.ParseWithClaims(tokenValue, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySignKeyRefresh), nil
	})

	if tokenErr != nil {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "refreshToken이 만료되었습니다, 다시 로그인해주세요",
		})
		return
	}

	fmt.Printf("token : %v\n", token)

	for key, val := range claims {
		fmt.Printf("key: %v, value : %v\n", key, val)
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "refreshToken 검증 완료",
	})
	return
}

// 새로운 AccessToken 부여
func CreateReissuanceToken(c *gin.Context) {
	accessToken, accessTokenErr := c.Request.Cookie("access_token")

	if accessTokenErr != nil {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "Get Cookie failed",
		})
		return
	}

	accessTokenValue := accessToken.Value

	if accessTokenValue == "" {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "accessToken is none",
		})
		return
	}

	claims := jwt.MapClaims{}

	token, _ := jwt.ParseWithClaims(accessTokenValue, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySignKey), nil
	})

	fmt.Println(token)

	for key, val := range claims {
		fmt.Printf("key: %v, value : %v\n", key, val)
	}

	Email := claims["email"].(string)

	newAccessToken, err := CreateJWT(Email)

	if err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "accessToken 생성중 에러",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":      200,
		"message":     "accessToken 재발급 완료",
		"accessToken": newAccessToken,
	})
	return

}
