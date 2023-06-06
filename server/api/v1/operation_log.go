package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoWay/blog/utils"
	"github.com/xiaoWay/blog/utils/r"
)

type OperationLog struct{}

func (*OperationLog) GetList(c *gin.Context) {
	r.SuccessData(c, operationLogService.GetList(utils.BindPageQuery(c)))
}

func (*OperationLog) Delete(c *gin.Context) {
	r.SendCode(c, operationLogService.Delete(utils.BindJson[[]int](c)))
}
