package helper

import (
	"fmt"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveAvatar(file *multipart.FileHeader, c *gin.Context) (string,error){
	dt := time.Now()
	res := fmt.Sprintf("%s-%s%s%s%s", dt.Format("02-Jan-2006"), strconv.Itoa(dt.Hour()), strconv.Itoa(dt.Minute()), strconv.Itoa(dt.Second()), file.Filename)
	fmt.Println("file name : "+ res)
	return res, c.SaveUploadedFile(file, "assets/"+res)
}

func Ifelse(param1 interface{}, param2 interface{}) interface{}{
	if param1 == 0 || param1 == ""{
		return param2
	}
	return param1
}