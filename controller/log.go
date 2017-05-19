package controller

import (
	"fmt"
	"maid-data-log/model"
	"maid-data-log/model/schema"

	iris "gopkg.in/kataras/iris.v6"
)

type LogCtrl struct {
}

var logBean = &model.LogBean{}

func (logCtrl *LogCtrl) GetAll(ctx *iris.Context) {
	if list, err := logBean.FindAll(); err != nil {
		fmt.Println(err)
		ctx.SetStatusCode(500)
		ctx.Writef("查询数据失败")
		return
	} else {
		ctx.JSON(200, list)
	}
}

func (logCtrl *LogCtrl) Post(ctx *iris.Context) {
	data := &schema.Log{}
	if err := ctx.ReadJSON(data); err != nil {
		ctx.SetStatusCode(500)
		fmt.Println(err)
		ctx.Done()
		return
	}
	_, err := logBean.Add(data)
	if err != nil {
		fmt.Println("add error:", err)
		ctx.SetStatusCode(500)
	} else {
		ctx.SetStatusCode(200)
	}

	ctx.Done()
}
