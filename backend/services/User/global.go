package User

import (
	"github.com/jinzhu/gorm"
	models "github.com/kmsdoit/WorkReportWithVueAndGo/backend/model"
)

var dbConn *gorm.DB
var user models.User
