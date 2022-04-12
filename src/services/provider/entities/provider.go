package entities

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Provider struct {
	ID             string `json:"id" bson:"_id,omitempty"`
	Email          string `json:"email" bson:"email,omitempty"`
	Password       string `json:"password" bson:"password,omitempty"`
	FullName       string `json:"full_name" bson:"full_name,omitempty"`
	Phone          string `json:"phone" bson:"phone,omitempty"`
	Address        string `json:"address" bson:"address,omitempty"`
	CompanyName    string `json:"company_name" bson:"company_name,omitempty"`
	CompanyPhone   string `json:"company_phone" bson:"company_phone,omitempty"`
	CompanyEmail   string `json:"company_email" bson:"company_email,omitempty"`
	CompanyAddress string `json:"company_address" bson:"company_address,omitempty"`
	URL            string `json:"url" bson:"url,omitempty"`
}

func (this Provider) CreateJWT() (string, error) {
	jwtLifeTime, err := strconv.Atoi(os.Getenv("JWT_LIFE_TIME"))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    this.ID,
		"email": this.Email,
		"exp":   time.Now().Add(time.Hour * time.Duration(jwtLifeTime)).Unix(),
	})
	key := []byte(os.Getenv("SECRET_KEY"))
	return token.SignedString(key)
}

func (this Provider) CheckPassword(candidatePassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(this.Password), []byte(candidatePassword)) == nil
}
