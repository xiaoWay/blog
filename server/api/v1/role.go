package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoWay/blog/model/req"
	"github.com/xiaoWay/blog/utils"
	"github.com/xiaoWay/blog/utils/r"
)

type Role struct{}

// 获取角色选项
func (*Role) GetOption(c *gin.Context) {
	r.SuccessData(c, roleService.GetOption())
}

func (*Role) GetTreeList(c *gin.Context) {
	r.SuccessData(c, roleService.GetTreeList(utils.BindPageQuery(c)))
}

func (*Role) SaveOrUpdate(c *gin.Context) {
	r.SendCode(c, roleService.SaveOrUpdate(utils.BindJson[req.SaveOrUpdateRole](c)))
}

func (*Role) Delete(c *gin.Context) {
	r.SendCode(c, roleService.Delete(utils.BindJson[[]int](c)))
}
