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
	task := toolbox.NewTask("task", "0 0 * * * *", func() error { service.Perform(); fmt.Println(time.Now()); return nil })
	err := task.Run()
	if err != nil {
		fmt.Println(err)
	}
	toolbox.AddTask("myTask", task)
	toolbox.StartTask()
	beego.Run()
}
