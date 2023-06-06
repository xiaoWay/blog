package service

import (
	"github.com/xiaoWay/blog/dao"
	"github.com/xiaoWay/blog/model"
	"github.com/xiaoWay/blog/model/req"
	"github.com/xiaoWay/blog/model/resp"
	"github.com/xiaoWay/blog/utils/r"
)

type OperationLog struct{}

func (*OperationLog) GetList(req req.PageQuery) resp.PageResult[[]model.OperationLog] {
	list, total := operationLogDao.GetList(req)
	return resp.PageResult[[]model.OperationLog]{
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		List:     list,
		Total:    total,
	}
}

func (*OperationLog) Delete(ids []int) (code int) {
	dao.Delete(model.OperationLog{}, "id in ?", ids)
	return r.OK
}
