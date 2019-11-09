package service

import (
	"fmt"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
	"log"
	"strconv"
	"time"
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
	pushRedis(search())
}

func dateFormat(t time.Time) string {
	var shortForm = "2006-01-02 15:04:05"
	return t.Format(shortForm)[0:10]
}

func search() string {
	str := "Hi,Sir.\n Today: \n"
	f, b, err := Search("2020-01-28")
	if err != nil {
		return ""
	}
	str += "南京-长治:\n"
	for j := range b {
		if b[j].Date.After(time.Date(2020, 01, 22, 0, 0, 0, 0, time.Local)) && b[j].Date.Before(time.Date(2020, 01, 25, 0, 0, 0, 0, time.Local)) {
			content := dateFormat(b[j].Date) + " price: " + strconv.FormatFloat(float64(b[j].Price), 'f', 0, 64)
			str += content + "\n"
		}
	}
	str += "长治-南京:\n"
	for i := range f {
		if f[i].Date.After(time.Date(2020, 01, 29, 0, 0, 0, 0, time.Local)) && f[i].Date.Before(time.Date(2020, 02, 02, 0, 0, 0, 0, time.Local)) {
			content := dateFormat(f[i].Date) + " price: " + strconv.FormatFloat(float64(f[i].Price), 'f', 0, 64)
			str += content + "\n"
		}
	}

	return str
}
