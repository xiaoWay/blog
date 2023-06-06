package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoWay/blog/model/req"
	"github.com/xiaoWay/blog/utils"
	"github.com/xiaoWay/blog/utils/r"
)

type Message struct{}

func (*Message) Delete(c *gin.Context) {
	r.SendCode(c, messageService.Delete(utils.BindJson[[]int](c)))
}

// 修改审核
func (*Message) UpdateReview(c *gin.Context) {
	r.SendCode(c, messageService.UpdateReview(
		utils.BindValidJson[req.UpdateReview](c)))
}

// 条件查询列表(后台)
func (*Message) GetList(c *gin.Context) {
	var req = utils.BindValidQuery[req.GetMessages](c)
	utils.CheckQueryPage(&req.PageSize, &req.PageNum)
	r.SuccessData(c, messageService.GetList(req))
}
