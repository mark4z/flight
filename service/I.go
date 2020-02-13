package service

import (
	"fmt"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
	"log"
)

func pushMsg(msg, token string) {

	cert, err := certificate.FromP12File("service/cert.p12", "123456")
	if err != nil {
		log.Fatal("Cert Error:", err)
	}

	notification := &apns2.Notification{}
	notification.DeviceToken = token
	notification.Topic = "com.pluto.edith.Edith"
	notification.Payload = payload.NewPayload().Alert(msg)

	client := apns2.NewClient(cert).Production()
	res, err := client.Push(notification)

	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}

func pushRedis(msg string) {
	pushMsg(msg, "c75b5a3ffde63006ee08d3e46a77a2ba1b6790e84f74d16aef1543b7119401d3")
}

func Push() {
	pushRedis("test")
}
