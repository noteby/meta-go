package authservice

import (
	"errors"
	"log"
	"meta-go/db"
	"meta-go/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 用户注册
func Register(username, password string) error {
	return db.Conn().Transaction(func(tx *gorm.DB) error {
		// 判断用户名是否存在
		var user model.User
		err := tx.Where(&model.User{Username: username}).First(&user).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err)
			return errors.New("该用户名已存在")
		}
		// 创建用户
		hashPassword, err := GenHashPassword(password)
		if err != nil {
			return err
		}
		user = model.User{Username: username, HashPassword: hashPassword}
		return tx.Create(&user).Error
	})
}

// 用户登录
func Login(username, password string) (*model.User, error) {
	var user model.User
	err := db.Conn().Where(&model.User{Username: username}).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("该用户不存在")
		}
		return nil, err
	}
	// 密码验证
	if !CompareHashAndPassword(user.HashPassword, password) {
		return nil, errors.New("密码错误，请重新输入")
	}
	return &user, err
}

// 生成哈希密码
func GenHashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash), err
}

// 哈希密码验证
func CompareHashAndPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashPassword), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// 通过用户名获取用户
func GetUserByUsername(username string) model.User {
	var user model.User
	err := db.Conn().Where(model.User{Username: username}).First(&user).Error
	if err != nil {
		log.Println(err)
	}
	return user
}
