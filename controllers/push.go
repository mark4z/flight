package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
	"time"
)

type PushController struct {
	beego.Controller
}

func (c *PushController) Post() {
	token := string(c.Ctx.Input.RequestBody)
	logs.Info("token "+ token)
	redis, err := cache.NewCache("redis", `{"key":"token","conn":"hx.anymre.top:6379","dbNum":"0"}`)
	if err != nil {
		fmt.Print(err)
		return
	}
	_ = redis.Put("devices", token, 240*time.Hour)
	logs.Info(redis.Get("devices"))

	c.Data["json"] = &token
	c.ServeJSON()
}
