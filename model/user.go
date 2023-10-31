package model

// User 用户结构体，面试用户
type User struct {
	Model
	Studentid    string `gorm:"primarykey;column:studentid;type:varchar(13);" json:"studentid"` //学号
	UserName     string `gorm:"column:username;type:varchar(5);" json:"username"`               //姓名
	Sex          string `gorm:"column:sex;type:varchar(2);" json:"sex"`                         //性别
	Grade        string `gorm:"column:grade;type:varchar(15);" json:"grade"`                    //年级
	Profession   string `gorm:"column:profession;type:varchar(30);" json:"profession"`          //专业
	Class        string `gorm:"column:profession;type:varchar(10);" json:"class"`               //班号
	Phone        string `gorm:"column:phone;type:varchar(11);" json:"phone"`                    //电话号
	WxId         string `gorm:"column:wxid;type:varchar(30);" json:"wxid"`                      //微信号
	Direction    string `gorm:"column:direction;type:varchar(6);" json:"direction"`             //方向
	WxOpenId     string `gorm:"column:wxopenid;type:varchar(50);" json:"wxopenid"`              //微信openid
	Status       string `gorm:"column:status;type:varchar(10);" json:"status"`                  //面试状态
	Personalid   string `gorm:"column:personalid;type:varchar(13);" json:"personalid"`          //唯一id
	Information1 string `gorm:"column:information1;type:varchar(300);" json:"information1"`     //个人信息1
	Information2 string `gorm:"column:information2;type:varchar(300);" json:"information2"`     //个人信息2
	Information3 string `gorm:"column:information3;type:varchar(300);" json:"information3"`     //个人信息3
	Award        string `gorm:"column:award;type:varchar(300);" json:"award"`                   //获奖经历
	Remark       string `gorm:"column:remark;type:varchar(300);" json:"remark"`                 //备注
	OK           int    `gorm:"column:ok;type:int;" json:"ok"`                                  //是否存完成投递
	First        string `gorm:"column:first;type:varchar(300);" json:"first"`                   //初试安排
	Second       string `gorm:"column:second;type:varchar(300);" json:"second"`                 //初试结论
	Third        string `gorm:"column:third;type:varchar(300);" json:"third"`                   //复试安排
	Fourth       string `gorm:"column:fourth;type:varchar(300);" json:"fourth"`                 //复试结论

}

// TableName 实现接口，使得数据库访问操作可以指定数据表
func (User) TableName() string {
	return "sys_users"
}
