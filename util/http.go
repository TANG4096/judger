package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sta-golang/go-lib-utils/log"
)

const (
	SUCCESS int = 0 //操作成功
	FAILED  int = 1 //操作失败
)

type PostForm map[string]interface{}

func GetPostForm(c *gin.Context) (PostForm, error) {
	postForm := make(PostForm)
	bodyBuf, err := ioutil.ReadAll(c.Request.Body)
	log.Infof("%s", bodyBuf)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = json.Unmarshal(bodyBuf, &postForm)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return postForm, nil
}

func (f PostForm) GetValue2String(key string) string {
	return f[key].(string)
}

func (f PostForm) GetValue2Int(key string) int {
	an, _ := strconv.Atoi(f[key].(string))
	return an
}

func Query2Int(c *gin.Context, key string) int {
	an, _ := strconv.Atoi(c.Query(key))
	return an
}

func GetQuery2Int(c *gin.Context, key string) (ans int, ob bool) {
	s, ok := c.GetQuery(key)
	if ok {
		an, _ := strconv.Atoi(s)
		return an, true
	} else {
		return 0, ok
	}

}

func Query2Uint(c *gin.Context, key string) uint {
	an, _ := strconv.Atoi(c.Query(key))
	return uint(an)
}

func CatchErr(c *gin.Context, err error) {
	log.Error(err)
	Failed(c, err.Error())
}

func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"msg":  "success",
		"data": v,
	})
}

//请求失败的时候, 使用该方法返回信息
func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": FAILED,
		"data": nil,
		"msg":  v,
	})
}
