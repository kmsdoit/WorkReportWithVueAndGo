package routes

import "github.com/kmsdoit/WorkReportWithVueAndGo/backend/services/User"

func UserRouter() {
	userApi := router.Group("/api/user")
	{
		userApi.POST("/register", User.UserRegister)
	}

}
