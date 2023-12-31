// Code generated by hertz generator.

package Gateway

import (
	"context"
	"errors"
	"sync"

	Gateway "Gateway/biz/model/Gateway"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

var ServiceNameMap = make(map[string]Gateway.ServiceInfo)

var mapMutex = &sync.Mutex{}

// AddService .
// @router /add [POST]
func AddService(ctx context.Context, c *app.RequestContext) {
	var err error
	var req Gateway.ServiceInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(Gateway.SuccessResp)

	AddNewService(req)
	resp = &Gateway.SuccessResp{
		Success: true,
		Message: "Add " + req.ServiceName + " successfully!!",
	}

	c.JSON(consts.StatusOK, resp)
}

func AddNewService(service Gateway.ServiceInfo) {
	mapMutex.Lock() // 获取锁定
	defer mapMutex.Unlock()
	if _, ok := ServiceNameMap[service.ServiceName]; !ok {
		ServiceNameMap[service.ServiceName] = service
	} else {
		err := errors.New("service" + service.ServiceName + "has been added！")
		panic(err)
	}
}
