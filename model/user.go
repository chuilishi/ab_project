package model

// User 用户结构体，面试用户
type User struct {
	Model
	StudentId  string `gorm:"primarykey;column:studentid;type:varchar(50);" json:"studentid"` //学号
	PassWord   string `gorm:"column:password;type:varchar(50);" json:"password"`              //密码不知道需不需要，但写了
	UserName   string `gorm:"column:username;type:varchar(50);" json:"username"`              //姓名
	Sex        string `gorm:"column:sex;type:varchar(2);" json:"sex"`                         //性别
	Grade      int    `gorm:"column:grade;type:int(100);" json:"grade"`                       //年级
	Profession string `gorm:"column:profession;type:varchar(50);" json:"profession"`          //专业
	Class      string `gorm:"column:profession;type:varchar(50);" json:"class"`               //班号
	Phone      string `gorm:"column:phone;type:varchar(13);" json:"phone"`                    //电话号
	WxId       string `gorm:"column:wxid;type:varchar(13);" json:"wxid"`                      //微信号
	WxOpenId   string `gorm:"column:wxOpenId;type:varchar(50);" json:"wxOpenId"`              //微信openid
	Status     string `gorm:"column:status;type:varchar(50);" json:"status"`                  //当前状态
	Remark     string `gorm:"column:remark;type:varchar(1000);" json:"remark"`                //备注
	One        string `gorm:"column:one;type:varchar(50);" json:"one"`                        //初试安排
	Two        string `gorm:"column:one;type:varchar(50);" json:"two"`                        //初试结论
	Three      string `gorm:"column:one;type:varchar(50);" json:"three"`                      //复试安排
	Four       string `gorm:"column:one;type:varchar(50);" json:"four"`                       //复试结论
}

// TableName 实现接口，使得数据库访问操作可以指定数据表
func (User) TableName() string {
	return "user"
}
