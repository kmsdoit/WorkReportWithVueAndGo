package User

import (
	"os"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var mySignKey = []byte(os.Getenv("SECRET_KEY"))
var mySignKeyRefresh = []byte(os.Getenv("REFRESH_KEY"))

func GenerateHashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

func ComparePassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false
		}
		return false
	}

	return true
}

func CreateJWT(Email string) (string, error) {

	aToken := jwt.New(jwt.SigningMethodHS256)
	claims := aToken.Claims.(jwt.MapClaims)
	claims["Email"] = Email
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	token, err := aToken.SignedString(mySignKey)
	if err != nil {
		return "", err
	} else {
		return token, nil
	}
}

func CreateRefreshToken(Email string) (string, error) {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	claims := refreshToken.Claims.(jwt.MapClaims)
	claims["Email"] = Email
	claims["exp"] = time.Now().Add(time.Hour * 720).Unix()

	token, err := refreshToken.SignedString(mySignKey)
	if err != nil {
		return "", err
	} else {
		return token, nil
	}
}

func SetDB(db *gorm.DB) {
	dbConn = db
	dbConn.AutoMigrate(&user)
}

func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}
