package model

// 用户信息结构体数据库
type Information struct {
	Model
	Studentid    string `gorm:"primarykey;column:studentid;type:varchar(13);" json:"studentid"` //学号
	Information1 string `gorm:"column:information1;type:varchar(300);" json:"information1"`     //个人信息1
	Information2 string `gorm:"column:information2;type:varchar(300);" json:"information2"`     //个人信息2
	Information3 string `gorm:"column:information3;type:varchar(300);" json:"information3"`     //个人信息3
	Award        string `gorm:"column:award;type:varchar(300);" json:"award"`                   //获奖经历
	Remark       string `gorm:"column:remark;type:varchar(300);" json:"remark"`                 //备注
	First        string `gorm:"column:first;type:varchar(300);" json:"first"`                   //初试安排
	Second       string `gorm:"column:second;type:varchar(300);" json:"second"`                 //初试结论
	Third        string `gorm:"column:third;type:varchar(300);" json:"third"`                   //复试安排
	Fourth       string `gorm:"column:fourth;type:varchar(300);" json:"fourth"`                 //复试结论
	Complete     string `gorm:"column:complete;type:bool;default:'0'" json:"complete"`          //是否提交完成
}

// TableName 实现接口，使得数据库访问操作可以指定数据表
func (Information) TableName() string {
	return "sys_information"
}
