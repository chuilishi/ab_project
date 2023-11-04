package service

import (
	"ab_project/mysqlDB"
	"ab_project/service/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

// UploadUserFileMessage 实现上传多个用户文件并存储用户文件信息
func UploadUserFileMessage(c *gin.Context) {
	wxopenid := c.Query("wxopenid")
	if wxopenid == "" {
		response.FailWithMessage("无法正确得到message", c)
		return
	}
	form, err := c.MultipartForm()
	if err != nil {
		response.FailWithMessage("得到文件失败"+err.Error(), c)
		return
	}
	files := form.File["file"]

	_, err = os.Stat("./userFile/" + wxopenid)
	if err != nil {
		err = os.Mkdir("./userFile", 0755)
		if err != nil {
			response.FailWithMessage("无法创建文件夹", c)
			return
		}
	}
	count := 0
	for _, file := range files {
		//保存文件到指定的路径
		if file.Size > 1024*1024*50 {
			response.FailWithMessage(file.Filename+"文件超出大小限制", c)
			continue
		}
		//上传到本地文件夹./userFile/wxopenid/下面
		err = c.SaveUploadedFile(file, "./userFile/"+wxopenid+"/"+file.Filename)
		if err != nil {
			response.FailWithMessage("保存文件"+file.Filename+"出错"+err.Error(), c)
			continue
		}
		count++
	}
	len := len(files)
	response.OkWithMessage(fmt.Sprintf("共上传%d个文件，成功%d个，失败%d个", len, count, len-count), c)

}

// DeleteUserFileMessage 实现在本地删除用户文件
func DeleteUserFileMessage(c *gin.Context) {
	wxopenid := c.Query("wxopenid")
	if wxopenid == "" {
		response.FailWithMessage("无法得到openid", c)
		return
	}
	filename := c.Query("file")
	if filename == "" {
		response.FailWithMessage("无法得到文件名", c)
		return
	}
	err := os.Remove("./userFile/" + wxopenid + "/" + filename)
	if err != nil {
		response.FailWithMessage("删除文件出错"+err.Error(), c)
	} else {
		response.OkWithMessage("删除文件成功", c)
	}
}

// ShowUserFileMessage 实现展示用户信息及其下载链接
func ShowUserFileMessage(c *gin.Context) {
	wxopenid := c.Query("wxopenid")
	if wxopenid == "" {
		response.FailWithMessage("无法得到openid", c)
		return
	}
	user := mysqlDB.IsUserHave(wxopenid)
	if user.ID == 0 {
		response.FailWithMessage("微信用户不存在", c)
		return
	}

	_, err := os.Stat("./userFile/" + wxopenid)

	if err != nil {
		err = os.Mkdir("./userFile/"+wxopenid, 0755)
		if err != nil {
			response.FailWithMessage("无法创建文件夹", c)
			return
		}
	}
	userFilePath := make(map[string]string)
	files, err := os.ReadDir("./userFile/" + wxopenid + "/")
	for _, file := range files {
		userFilePath[file.Name()] = "http://123.207.73.185:8090/userFile/" + wxopenid + "/" + file.Name()
	}
	response.OkWithDetailed(userFilePath, "获取文件信息成功！", c)
}
