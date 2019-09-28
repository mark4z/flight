package service

import (
	"encoding/json"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

const url = "http://www.juneyaoair.com/PriceCache"

type Forward struct {
	Id    int64     `orm:"auto"`
	Date  time.Time `orm:"type(datetime)"`
	Price float64
}
type Back struct {
	Id    int64     `orm:"auto"`
	Date  time.Time `orm:"type(datetime)"`
	Price float64
}

func init() {
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RegisterDataBase("default", "mysql", "root:123456@tcp(server.anymre.top:3306)/flight?charset=utf8&parseTime=true&charset=utf8&loc=Asia%2FShanghai")
	orm.RegisterModel(new(Forward), new(Back))
	_ = orm.RunSyncdb("default", false, true)
}

type OW struct {
	forward []Forward
	back    []Back
}

func Get() string {
	now := time.Now().Format("2006-01-02")
	req := httplib.Get(url)
	req.Param("flightType", "OW")
	req.Param("departureDate", now)
	req.Param("returnDate", now)
	req.Param("sendCode", "CIH")
	req.Param("arrCode", "NKG")
	req.Param("periodType", "Line")
	req.Param("_", strconv.FormatInt(time.Now().UnixNano()/1e6, 10))
	s, _ := req.String()
	return s
}

func Search() OW {
	str := Get()
	var sth [2]interface{}
	_ = json.Unmarshal([]byte(str), &sth)

	arr0 := sth[0].(map[string]interface{})
	arr1 := sth[1].(map[string]interface{})
	var forward []Forward
	var back []Back

	for i := range arr0 {
		v := arr0[i].(float64)
		flight := new(Forward)
		flight.Date = timeFormat(i)
		flight.Price = v
		forward = append(forward, *flight)
	}
	for i := range arr1 {
		v := arr0[i].(float64)
		flight := new(Back)
		flight.Date = timeFormat(i)
		flight.Price = v
		back = append(back, *flight)
	}
	return OW{forward, back}
}

func timeFormat(r string) time.Time {
	HourTemplate := "15:04:05"
	DayTemplate := "2006-01-02 15:04:05" //常规类型

	now := time.Now().Format(HourTemplate)
	result, _ := time.ParseInLocation(DayTemplate, r+" "+now, time.Local)
	return result
}

func Perform() {
	o := orm.NewOrm()
	r := Search()
	for e := range r.forward {
		_, _ = o.Insert(&r.forward[e])
	}
	for e := range r.back {
		_, _ = o.Insert(&r.back[e])
	}
}
