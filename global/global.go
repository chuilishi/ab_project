package global

// mysql dsn配置
// const DSN = "ab_project:ab_project@tcp(123.207.73.185:3306)/ab_project?charset=utf8mb4&parseTime=True&loc=Local"
const JWTKey = "ab_project"
const DSN = "root:123456@tcp(localhost:3306)/ab_project?charset=utf8mb4&parseTime=True&loc=Local"

const ENV = "debug"

//const ENV = "main"

//const UserStatusList[] =

const MailFrom = "djy1349678887@163.com"
const MailTOManager = "djy1349678887@163.com"

const ManageID = "admin"
const ManagePassword = "admin"
const ManageName = "半塘"

var Status2int map[string]int
var Allstatus [16]string

func init() {
	Status2int = make(map[string]int)
	Status2int["草稿"] = 0
	Status2int["提交成功"] = 1
	Status2int["待筛选"] = 2
	Status2int["待安排初试"] = 3
	Status2int["待初试"] = 4
	Status2int["初试通过"] = 5
	Status2int["待安排复试"] = 6
	Status2int["待复试"] = 7
	Status2int["复试通过"] = 8
	Status2int["待安排终试"] = 9
	Status2int["待终试"] = 10
	Status2int["终试通过"] = 11

	Allstatus[0] = "草稿"
	Allstatus[1] = "提交成功"
	Allstatus[2] = "待筛选"
	//allstatus[3] = "筛选不通过"
	Allstatus[3] = "待安排初试"
	Allstatus[4] = "待初试"
	//allstatus[4] = "初试不通过"
	Allstatus[5] = "初试通过"
	Allstatus[6] = "待安排复试"
	Allstatus[7] = "待复试"
	//allstatus[8] = "复试不通过"
	Allstatus[8] = "复试通过"
	Allstatus[9] = "待安排终试"
	Allstatus[10] = "待终试"
	//allstatus[12] = "终试不通过"
	Allstatus[11] = "终试通过"
	Allstatus[13] = "待处理"
	Allstatus[14] = "挂起"
}
