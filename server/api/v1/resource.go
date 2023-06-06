package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoWay/blog/model/req"
	"github.com/xiaoWay/blog/utils"
	"github.com/xiaoWay/blog/utils/r"
)

type Resource struct{}

// 获取资源列表 [树形结构, 不需要分页]
func (*Resource) GetTreeList(c *gin.Context) {
	r.SuccessData(c, resourceService.GetTreeList(utils.BindQuery[req.PageQuery](c)))
}

func (*Resource) SaveOrUpdate(c *gin.Context) {
	r.SendCode(c, resourceService.SaveOrUpdate(utils.BindJson[req.SaveOrUpdateResource](c)))
}

func (*Resource) Delete(c *gin.Context) {
	r.SendCode(c, resourceService.Delete(utils.GetIntParam(c, "id")))
}

// 修改资源允许匿名访问
func (*Resource) UpdateAnonymous(c *gin.Context) {
	r.SendCode(c, resourceService.UpdateAnonymous(utils.BindValidJson[req.UpdateAnonymous](c)))
}

// 角色管理模块界面选择树形数据
func (*Resource) GetOption(c *gin.Context) {
	r.SuccessData(c, resourceService.GetOptionList())
}
