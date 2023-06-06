package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoWay/blog/model/req"
	"github.com/xiaoWay/blog/utils"
	"github.com/xiaoWay/blog/utils/r"
)

type Comment struct{}

func (*Comment) Delete(c *gin.Context) {
	r.SendCode(c, commentService.Delete(utils.BindJson[[]int](c)))
}

// 修改审核
func (*Comment) UpdateReview(c *gin.Context) {
	r.SendCode(c, commentService.UpdateReview(utils.BindValidJson[req.UpdateReview](c)))
}

// TODO: 完善
func (*Comment) GetList(c *gin.Context) {
	var req = utils.BindValidQuery[req.GetComments](c)
	utils.CheckQueryPage(&req.PageSize, &req.PageNum)
	r.SuccessData(c, commentService.GetList(req))
}
