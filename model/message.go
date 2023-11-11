package model

import (
	"database/sql"
	"time"
)

type Message struct {
	ID        uint `gorm:"primarykey;"`
	UserID    uint
	Message   string `gorm:"column:message;type:varchar(300)"` //发送的消息
	Code      int    `gorm:"column:code;type:int"`             //发送消息的类型
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index;"`
}

func (Message) TableName() string {
	return "sys_message"
}
