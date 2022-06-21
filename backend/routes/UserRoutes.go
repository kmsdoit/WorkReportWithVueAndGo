package routes

import "github.com/kmsdoit/WorkReportWithVueAndGo/backend/services/User"

func UserRouter() {
	userApi := router.Group("/api/user")
	{
		userApi.POST("/register", User.UserRegister)
		userApi.POST("/login", User.UserLogin)
		userApi.GET("/accessTokenVerify", User.VerifyAccessToken)
		userApi.GET("/refreshTokenVerify", User.VerifyRefreshToken)
		userApi.GET("/reissuanceAccessToken", User.CreateReissuanceToken)
		userApi.GET("/findPassword", User.FindPassword)      //미완성
		userApi.POST("/updatePassword", User.UpdatePassword) //미완성
	}
}
