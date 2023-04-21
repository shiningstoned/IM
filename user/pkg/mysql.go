package pkg

import (
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Uuid     string
	Username string
	Password string
	Nickname string
	Email    string
}

func (u *User) BeforeCreate(_ *gorm.DB) error {
	if u.Uuid != "" {
		return nil
	}
	u.Uuid = uuid.New().String()
	return nil
}

type UserManager struct {
	db *gorm.DB
}

func NewUserManager(db *gorm.DB) *UserManager {
	m := db.Migrator()
	if !m.HasTable(&User{}) {
		err := m.CreateTable(&User{})
		if err != nil {
			klog.Fatalf("create table failed: %s", err)
		}
	}
	return &UserManager{
		db: db,
	}

}

func (u *UserManager) CreateUser(username, password string) error {
	var user User
	_, err := u.GetUserByName(username)
	if err == gorm.ErrRecordNotFound {
		user.Username = username
		user.Password = password
		err := u.db.Create(&user).Error
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}
	return errors.New("user exit")
}

func (u *UserManager) LoginCheck(username, password string) (bool, string, error) {
	user, err := u.GetUserByName(username)
	if err != nil {
		return false, "", err
	}
	if user.Password != password {
		return false, "", nil
	}
	return true, user.Uuid, nil
}

func (u *UserManager) GetUserByName(username string) (*User, error) {
	var user User
	err := u.db.Where("username=?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserManager) IsUserExist(uuid string) (bool, error) {
	err := u.db.Where("uuid=?", uuid).First(&User{}).Error
	if err == nil {
		return true, nil
	}
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return false, err
}
