package utils

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) ([]byte, error) {
	passwordByte := []byte(password)
	encryptPassword, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	if err != nil {
		log.WithField("error", err).Errorf("生成加密密码失败")
		return nil, err
	}
	return encryptPassword, nil
}
