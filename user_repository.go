package repositories

import (
	"github.com/social-mediam-users/config"
	"github.com/social-mediam-users/dto"
	"github.com/social-mediam-users/entity"
	"github.com/social-mediam-users/helpers"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB = config.SetupDatabase()
)

type UserRepository interface {
	Store(user dto.UserCreateDTO) dto.UserCreateDTO
	Login(user dto.UserLoginDTO) uint64
	IsDublicateUsername(username string) (ctx *gorm.DB)
	IsDublicatePhone(phone string) (ctx *gorm.DB)
	IsDublicateEmail(email string) (ctx *gorm.DB)
	IsLoginFromOtherDevice(userId uint64, agent string, clientIP string, randomToken string) bool
	Me(id uint64) entity.User
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (repository *userRepository) Store(user dto.UserCreateDTO) dto.UserCreateDTO {
	user.Password = helpers.PasswordHash(user.Password)
	DB.Create(user)
	return user
}

func (repository *userRepository) Login(user dto.UserLoginDTO) uint64 {
	password := user.Password
	DB.Where("username=?", user.Username).Or("email=?", user.Username).Or("phone=?", user.Username).First(&user)
	if user.ID != 0 {
		if helpers.PasswordCheckHash(password, user.Password) {
			return user.ID
		}
		return 0
	}
	return 0
}

func (repository *userRepository) IsDublicateUsername(username string) (ctx *gorm.DB) {
	var user entity.User
	return DB.Where("username = ?", username).Take(&user)
}

func (repository *userRepository) IsDublicatePhone(phone string) (ctx *gorm.DB) {
	var user entity.User
	return DB.Where("phone = ?", phone).Take(&user)
}
func (repository *userRepository) IsDublicateEmail(email string) (ctx *gorm.DB) {
	var user entity.User
	return DB.Where("email = ?", email).Take(&user)
}

func (repository *userRepository) IsLoginFromOtherDevice(userId uint64, agent string, clientIP string, randomToken string) bool {
	var token entity.Token
	DB.Where("agent=?", agent).Where("client_ip=?", clientIP).Where("user_id=?", userId).First(&token)
	if token.ID != 0 {
		token.Token = randomToken
		DB.Save(&token)
		return false
	}
	token.UserId = userId
	token.Agent = agent
	token.ClientIP = clientIP
	token.Token = randomToken
	DB.Create(&token)
	return true
}

func (repository *userRepository) Me(id uint64) entity.User {
	var user entity.User
	DB.Where("id=?", id).Find(&user)
	return user
}
