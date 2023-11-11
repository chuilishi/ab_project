package model

// User 用户结构体，面试用户
type User struct {
	Model
	StudentId    string `gorm:"primarykey;column:studentid;type:varchar(13);default:'未知';" json:"studentid"`  //学号
	UserName     string `gorm:"column:username;type:varchar(5);default:'未知';" json:"username"`                //姓名
	Sex          string `gorm:"column:sex;type:varchar(2);default:'未知';" json:"sex"`                          //性别
	Grade        string `gorm:"column:grade;type:varchar(15);default:'未知';" json:"grade"`                     //年级
	Profession   string `gorm:"column:profession;type:varchar(30);default:'未知';" json:"profession"`           //专业
	Class        string `gorm:"column:class;type:varchar(10);default:'未知';" json:"class"`                     //班号
	Phone        string `gorm:"column:phone;type:varchar(11);default:'未知';" json:"phone"`                     //电话号
	WxId         string `gorm:"column:wxid;type:varchar(30);default:'未知';" json:"wxid"`                       //微信号
	Direction    string `gorm:"column:direction;type:varchar(6);default:'未知';" json:"direction"`              //方向
	WxOpenId     string `gorm:"column:wxopenid;type:varchar(50);default:'未知';" json:"wxopenid"`               //微信openid
	Status       string `gorm:"column:status;type:varchar(10);default:'未知';" json:"status"`                   //面试状态
	Introduction string `gorm:"column:introduction;type:varchar(300);default:'还没有自我介绍';" json:"introduction"` //自我介绍
	Reasons      string `gorm:"column:reasons;type:varchar(300);default:'未知';" json:"reasons"`                //加入ab的原因
	Experience   string `gorm:"column:experience;type:varchar(300);default:'未知';" json:"experience"`          //个人经历
	Award        string `gorm:"column:award;type:varchar(300);default:'未知';" json:"award"`                    //获奖经历
	Remark       string `gorm:"column:remark;type:varchar(300);default:'未知';" json:"remark"`                  //备注
	OK           int    `gorm:"column:ok;type:int;default:0;" json:"ok"`                                      //是否已经完成投递
	ISProblem    int    `gorm:"column:isproblem;type:int;default:0;" json:"isproblem"`                        //是否有异常挂起
	Problem      string `gorm:"problem:problem;type:varchar(300);default:'无异常信息';" json:"problem"`            //用户异常信息
	First        string `gorm:"column:first;type:varchar(300);default:'还没有安排';" json:"first"`                 //初试安排
	Second       string `gorm:"column:second;type:varchar(300);default:'还没有结论';" json:"second"`               //初试结论
	Third        string `gorm:"column:third;type:varchar(300);default:'还没有安排';" json:"third"`                 //复试安排
	Fourth       string `gorm:"column:fourth;type:varchar(300);default:'还没有结论';" json:"fourth"`               //复试结论
	//Etc          string `gorm:"type:json" json:"etc"` //如果有额外的信息mysql5.7.8原生支持json，服务器记得升级mysql
	Messages []Message
}

// TableName 实现接口，使得数据库访问操作可以指定数据表
func (User) TableName() string {
	return "sys_users"
}
