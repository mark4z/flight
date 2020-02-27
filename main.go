package main

import (
	_ "flight/routers"
	"flight/service"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/astaxie/beego/logs"
	"time"
)

func main() {
	task := toolbox.NewTask("task", "0 0 * * * *", func() error { service.Perform(); fmt.Println("search:" + time.Now().String()); return nil })
	err := task.Run()
	if err != nil {
		fmt.Println(err)
	}
	logs.Info("0945")

	toolbox.AddTask("myTask", task)
	toolbox.StartTask()
	beego.Run()
}
