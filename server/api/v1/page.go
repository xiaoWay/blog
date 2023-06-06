package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoWay/blog/model/req"
	"github.com/xiaoWay/blog/utils"
	"github.com/xiaoWay/blog/utils/r"
)

type Page struct{}

func (*Page) GetList(c *gin.Context) {
	r.SuccessData(c, pageService.GetList())
}

func (*Page) SaveOrUpdate(c *gin.Context) {
	r.SendCode(c, pageService.SaveOrUpdate(utils.BindJson[req.AddOrEditPage](c)))
}

func (*Page) Delete(c *gin.Context) {
	r.SendCode(c, pageService.Delete(utils.BindJson[[]int](c)))
}
