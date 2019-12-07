package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main(){
	var(
		expr *cronexpr.Expression
		err error
		now time.Time
		nextTime time.Time
	)
	now = time.Now()
	//每5分钟执行一次
	if expr,err = cronexpr.Parse("*/5 * * * * * *"); err != nil{
		fmt.Println(err.Error())
		return
	}
	nextTime = expr.Next(now)
	fmt.Println(now,nextTime)

	time.AfterFunc(nextTime.Sub(now),func(){
		fmt.Println("被调度了",nextTime)
	})
	time.Sleep(5 * time.Second)
	expr = expr

}