package model

// User 用户结构体，面试用户
type User struct {
	Model
	Studentid  string `gorm:"primarykey;column:studentid;type:varchar(13);" json:"studentid"` //学号
	UserName   string `gorm:"column:username;type:varchar(5);" json:"username"`               //姓名
	Sex        string `gorm:"column:sex;type:varchar(2);" json:"sex"`                         //性别
	Grade      int    `gorm:"column:grade;type:int(15);" json:"grade"`                        //年级
	Profession string `gorm:"column:profession;type:varchar(30);" json:"profession"`          //专业
	Class      string `gorm:"column:profession;type:varchar(10);" json:"class"`               //班号
	Phone      string `gorm:"column:phone;type:varchar(11);" json:"phone"`                    //电话号
	WxId       string `gorm:"column:wxid;type:varchar(30);" json:"wxid"`                      //微信号
	Direction  string `gorm:"column:direction;type:varchar(6);" json:"direction"`             //方向
	WxOpenId   string `gorm:"column:wxOpenId;type:varchar(50);" json:"wxOpenId"`              //微信openid
	Status     string `gorm:"column:status;type:varchar(10);" json:"status"`                  //面试状态
	Personalid string `gorm:"column:personalid;type:varchar(13);" json:"personalid"`          //唯一id
}

// TableName 实现接口，使得数据库访问操作可以指定数据表
func (User) TableName() string {
	return "sys_user"
}
