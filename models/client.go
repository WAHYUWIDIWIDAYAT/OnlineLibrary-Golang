package models

import (
	"html"
	"strings"
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"web/utils/token"
)

type Client struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
	Email string `gorm:"size:255;not null;" json:"email"`
}
func GetClientByID(uid uint) (Client,error) {

	var u Client

	if err := DB.First(&u,uid).Error; err != nil {
		return u,errors.New("User not found!")
	}

	u.PrepareGive()
	
	return u,nil

}

func (u *Client) PrepareGive(){
	u.Password = ""
}

func VerifyPasswordClient(password,hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginClientCheck(username string, password string) (string,error) {
	
	var err error

	u := Client{}

	err = DB.Model(Client{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPasswordClient(password, u.Password)
	

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token,err := token.GenerateToken(u.ID)

	if err != nil {
		return "",err
	}

	return token,nil
	
}

func (u *Client) SaveUser() (*Client, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &Client{}, err
	}
	return u, nil
}

func (u *Client) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username 
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}