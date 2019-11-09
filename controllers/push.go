package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"time"
)

type PushController struct {
	beego.Controller
}

func (c *PushController) Post() {
	token := string(c.Ctx.Input.RequestBody)
	redis, err := cache.NewCache("memory", `{"key":"token","conn":"hx.anymre.top:6379","dbNum":"0"}`)
	if err != nil {
		fmt.Print(err)
		return
	}
	_ = redis.Put("devices", token, 240*time.Hour)
	fmt.Println(redis.Get("devices"))

	c.Data["json"] = &token
	c.ServeJSON()
}
