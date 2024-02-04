package service

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"strconv"
)

// FileUpload 文件上传
func FileUpload(c *gin.Context, file *multipart.FileHeader, uid uint, path, folder string) (string, error) {
	// 定义保存路径  ../static/imgs/avatar/user1/
	strUid := strconv.Itoa(int(uid))
	dst := ".." + path + folder + strUid + "/"

	// 判断文件夹是否存在，不存在则创建
	if err := os.MkdirAll(dst, 0777); err != nil {
		return "", err
	}

	// 将头像保存到指定路径
	if err := c.SaveUploadedFile(file, dst+file.Filename); err != nil {
		return "", err
	}
	return folder + strUid + "/" + file.Filename, nil
}
