package model

// Manager 管理员数据库
type Manager struct {
	Model
	Studentid string `gorm:"primarykey;column:studentid;type:varchar(13);" json:"studentid"` //学号
	Password  string `gorm:"column:passwords;type:varchar(30);" json:"passwords"`            //密码
	UserName  string `gorm:"column:username;type:varchar(5);" json:"username"`               //姓名
}

// TableName 实现接口，使得数据库访问操作可以指定数据表
func (Manager) TableName() string {
	return "sys_manager"
}
