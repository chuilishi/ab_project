package model

import (
	"database/sql"
	"time"
)

type MessageTemplate struct {
	ID        uint   `gorm:"primarykey;"`
	Message   string `gorm:"column:message;type:varchar(300);" json:"message"` //发送的消息
	Code      int    `gorm:"column:code;type:int;" json:"code"`                //发送消息的类型 面试or结果
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index;"`
}

func (MessageTemplate) TableName() string {
	return "sys_messagetemplate"
}
