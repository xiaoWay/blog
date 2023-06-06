package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoWay/blog/model/req"
	"github.com/xiaoWay/blog/utils"
	"github.com/xiaoWay/blog/utils/r"
)

type Tag struct{}

// 新增/编辑 标签
func (*Tag) SaveOrUpdate(c *gin.Context) {
	r.SendCode(c, tagService.SaveOrUpdate(utils.BindValidJson[req.AddOrEditTag](c)))
}

// 删除(批量)
func (*Tag) Delete(c *gin.Context) {
	r.SendCode(c, tagService.Delete(utils.BindJson[[]int](c)))
}

// 获取下拉框选项数据
func (*Tag) GetOption(c *gin.Context) {
	r.SuccessData(c, tagService.GetOption())
}

// 条件标签列表(后台)
func (*Tag) GetList(c *gin.Context) {
	r.SuccessData(c, tagService.GetList(utils.BindPageQuery(c)))
}

// 查询标签列表(前台)
func (*Tag) GetFrontList(c *gin.Context) {
	r.SuccessData(c, tagService.GetFrontList())
}
