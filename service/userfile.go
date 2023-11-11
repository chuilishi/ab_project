package service

import (
	"ab_project/mysqlDB"
	"ab_project/service/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

// UploadUserFileMessage 实现上传多个用户文件并存储用户文件信息
func UploadUserFileMessage(c *gin.Context) {
	//wxopenid := c.Query("wxopenid")
	ids := c.Query("studentid")
	if ids == "" {
		response.FailWithMessage("无法正确得到message", c)
		return
	}
	form, err := c.MultipartForm()
	if err != nil {
		response.FailWithMessage("得到文件失败"+err.Error(), c)
		return
	}
	files := form.File["file"]
	pictures := form.File["picture"]
	_, err = os.Stat("./userFile/" + ids)
	if err != nil {
		err = os.Mkdir("./userFile/"+ids, 0755)
		if err != nil {
			response.FailWithMessage("无法创建文件夹", c)
			return
		}
	}
	_, err = os.Stat("./userFile/" + ids)
	if err != nil {
		err = os.Mkdir("./userFile/"+ids, 0755)
		if err != nil {
			response.FailWithMessage("无法创建文件夹", c)
			return
		}
	}
	count := 0
	for _, file := range files {
		//保存文件到指定的路径
		if file.Size > 1024*1024*50 {
			response.FailWithMessage(file.Filename+"文件超出大小限制，50M", c)
			continue
		}
		//上传到本地文件夹./userFile/id/下面
		err = c.SaveUploadedFile(file, "./userFile/"+ids+"/"+file.Filename)
		if err != nil {
			response.FailWithMessage("保存文件"+file.Filename+"出错"+err.Error(), c)
			continue
		}
		count++
	}
	lens := len(files)
	response.OkWithMessage(fmt.Sprintf("共上传%d个文件，成功%d个，失败%d个", lens, count, lens-count), c)
	if len(pictures) == 0 {
		response.FailWithMessage("图片未上传", c)
		return
	}
	picture := pictures[0]

	err = c.SaveUploadedFile(picture, "./userFile/"+ids+"/照片"+path.Ext(picture.Filename))
	if err != nil {
		response.FailWithMessage("保存文件"+picture.Filename+"出错"+err.Error(), c)
		return
	}
	response.OkWithMessage("保存图片成功", c)

}

// DeleteUserFileMessage 实现在本地删除用户文件
func DeleteUserFileMessage(c *gin.Context) {
	//wxopenid := c.Query("wxopenid")
	ids := c.Query("studentid")
	if ids == "" {
		response.FailWithMessage("无法得到studentid", c)
		return
	}
	filename := c.Query("file")
	if filename == "" {
		response.FailWithMessage("无法得到文件名", c)
		return
	}
	err := os.Remove("./userFile/" + ids + "/" + filename)
	if err != nil {
		response.FailWithMessage("删除文件出错"+err.Error(), c)
	} else {
		response.OkWithMessage("删除文件成功", c)
	}
}

// ShowUserFileMessage 实现展示用户信息及其下载链接
func ShowUserFileMessage(c *gin.Context) {
	//wxopenid := c.Query("wxopenid")
	ids := c.Query("studentid")
	if ids == "" {
		response.FailWithMessage("无法得到studenid", c)
		return
	}
	user := mysqlDB.IsUserHaveByStudentId(ids)
	if user.ID == 0 {
		response.FailWithMessage("微信用户不存在", c)
		return
	}

	_, err := os.Stat("./userFile/" + ids)
	if err != nil {
		err = os.Mkdir("./userFile/"+ids, 0755)
		if err != nil {
			response.FailWithMessage("无法创建文件夹", c)
			return
		}
	}
	userFilePath := make(map[string]string)
	files, err := os.ReadDir("./userFile/" + ids + "/")
	for _, file := range files {
		userFilePath[file.Name()] = "http://123.207.73.185:8090/userFile/" + ids + "/" + file.Name()
	}
	response.OkWithDetailed(userFilePath, "获取文件信息成功！", c)
}
func DonLoadPictures(c *gin.Context) {

	response.OkWithMessage("http://123.207.73.185:8090/userFile/ppp/picture", c)
}

// UploadPicture 上传首页图片
func UploadPicture(c *gin.Context) {
	picture, err := c.FormFile("picture")
	if err != nil {
		response.FailWithMessage("得到文件失败"+err.Error(), c)
		return
	}
	if path.Ext(picture.Filename) != ".png" && path.Ext(picture.Filename) != ".jpg" && path.Ext(picture.Filename) != ".bmp" && path.Ext(picture.Filename) != ".gif" && path.Ext(picture.Filename) != ".svg" {
		response.FailWithMessage("文件格式不正确", c)
		return
	}
	err = os.RemoveAll("./userFile/ppp/")
	if err != nil {
		response.FailWithMessage("上传文件出错，请重新上传", c)
		return
	}
	_, err = os.Stat("./userFile/ppp")
	if err != nil {
		err = os.Mkdir("./userFile/ppp", 0755)
		if err != nil {
			response.FailWithMessage("无法创建文件夹", c)
			return
		}
	}

	err = c.SaveUploadedFile(picture, "./userFile/ppp/picture"+path.Ext(picture.Filename))
	if err != nil {
		response.FailWithMessage("保存图片失败", c)
		return
	}
	response.OkWithMessage("保存图片成功！", c)
	return
}
