package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity      string
	Name          string
	Password      string
	Phone         string
	Email         string
	ClientIp      string
	ClientPort    string
	LoginTime     uint64
	HeartbeatTime uint64
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "UserBasic"

}
