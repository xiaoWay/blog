package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoWay/blog/model/req"
	"github.com/xiaoWay/blog/utils"
	"github.com/xiaoWay/blog/utils/r"
)

type Link struct{}

func (*Link) GetList(c *gin.Context) {
	r.SuccessData(c, linkService.GetList(utils.BindPageQuery(c)))
}

func (*Link) SaveOrUpdate(c *gin.Context) {
	r.SendCode(c, linkService.SaveOrUpdate(utils.BindValidJson[req.SaveOrUpdateLink](c)))
}

func (*Link) Delete(c *gin.Context) {
	r.SuccessData(c, linkService.Delete(utils.BindJson[[]int](c)))
}
