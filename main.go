package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	_ "github.com/go-sql-driver/mysql"
	_ "i/routers"
	"i/service"
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
