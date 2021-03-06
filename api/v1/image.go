package v1

import (
	"bufio"
	"ginserver/model"
	"ginserver/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"strings"
)

//var code int

func UploadImg(c *gin.Context) {
	var img model.Image
	//_ = c.ShouldBindJSON(&img)
	img.ImageName = c.Request.FormValue("imagename")
	img.ImageDesc = c.Request.FormValue("imagedesc")
	img.ImageTag = c.Request.FormValue("imagetag")

	//username := c.Request.FormValue("imageuser")
	// 在 jwt 验证时，已经设置了 username 参数，此处可以直接使用，不用前端传入 username 参数
	username := c.Keys["username"]
	// 断言，将 interface 类型转换为 string
	userid := model.GetUserId(username.(string))
	if userid < 0 {
		userid = 1
	}
	img.AuthID = int(userid)

	file, header, errfile := c.Request.FormFile("file")
	if errfile != nil {
		code = utils.UPLOAD_FILE_IS_EMPTY
	}

	// 上传文件的文件名
	filename := header.Filename
	filetype := strings.Split(filename, ".")[1]
	// 生成uuid
	fileuuid := utils.GenUUID()

	uuidname := strings.Join([]string{fileuuid, filetype}, ".")

	// 分片读取文件
	r := bufio.NewReader(file)
	var chunks []byte
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if n == 0 {
			break
		}
		chunks = append(chunks, buf...)
	}

	// 上传文件
	go utils.UploadImg(chunks, uuidname)

	img.ImagePath = utils.QiniuDomain + "/" + uuidname

	code = model.UploadImg(&img)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    img.ImagePath,
		"message": utils.GetErrMsg(code),
	})
}

// 获取当前用户下的 image
func GetImg(c *gin.Context) {
	//var img model.Image

	username := c.Keys["username"]
	userid := model.GetUserId(username.(string))

	data := model.GetImgByUserId(int(userid))
	code = utils.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": utils.GetErrMsg(code),
	})
}

// 删除 image
func DeleteImg(c *gin.Context) {
	username := c.Keys["username"]
	userid := model.GetUserId(username.(string))

	imageid, _ := strconv.Atoi(c.Request.FormValue("imageid"))

	img := model.GetImgById(imageid)

	if img.ID <= 0 {
		code = utils.IMAGE_NOT_FOUND
	} else if img.AuthID != int(userid) {
		code = utils.NO_PER_TO_DELETE_IMAGE
	} else {
		// 数据库逻辑删除
		code = model.DeleteImg(imageid)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    img.ID,
		"message": utils.GetErrMsg(code),
	})
}

// 获取已经被删除的图片
func GetDeletedImg(c *gin.Context) {
	username := c.Keys["username"]
	userid := model.GetUserId(username.(string))

	data := model.GetDeleted(int(userid))

	code = utils.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": utils.GetErrMsg(code),
	})
}

func RealDeleteImg(c *gin.Context) {
	username := c.Keys["username"]
	userid := model.GetUserId(username.(string))

	imageid, _ := strconv.Atoi(c.Request.FormValue("imageid"))

	img := model.GetImgById(imageid)

	if img.ID <= 0 {
		code = utils.IMAGE_NOT_FOUND
	} else if img.AuthID != int(userid) {
		code = utils.NO_PER_TO_DELETE_IMAGE
	} else {
		// 七牛云 oss 删除
		filename := strings.Split(img.ImagePath, "/")
		utils.DeleteImg(filename[len(filename)-1])
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    img.ID,
		"message": utils.GetErrMsg(code),
	})
}
