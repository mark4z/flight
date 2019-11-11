package main

import (
	_ "flight/routers"
	"flight/service"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"time"
)

func main() {
	task := toolbox.NewTask("task", "0 0 * * * *", func() error { service.Perform(); fmt.Println("search:" + time.Now().String()); return nil })
	err := task.Run()
	if err != nil {
		fmt.Println(err)
	}

	push := toolbox.NewTask("push", "0 0 15 * * *", func() error { service.Push(); fmt.Println("push:" + time.Now().String()); return nil })
	_ = push.Run()

	toolbox.AddTask("myTask", task)
	toolbox.AddTask("push", push)
	toolbox.StartTask()
	beego.Run()
}
