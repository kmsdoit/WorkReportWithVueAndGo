package User

import (
	"github.com/jinzhu/gorm"
	models "github.com/kmsdoit/WorkReportWithVueAndGo/backend/model"
	"github.com/olahol/go-imageupload"
)

var dbConn *gorm.DB
var user models.User
var currentImage *imageupload.Image // 이미지 할당을 위한 구조체
