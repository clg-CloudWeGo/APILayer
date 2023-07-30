package handler

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/json"
)

func SolveUrl(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interface{}
	err = c.BindAndValidate(&req)
	if err != nil {
		panic(err)
	}
	serviceName := c.Param("service")
	methodName := c.Param("method")

	jsonReq, err := json.Marshal(req)
	if err != nil {
		fmt.Println("error:", err)
	}
	//写不下去了，后面要还没看懂

}
