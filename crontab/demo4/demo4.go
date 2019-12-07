package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main(){
	type Cronjob struct{
		expr *cronexpr.Expression
		nextTime time.Time
	}

	var(
		cronjob *Cronjob
		expr *cronexpr.Expression
		now time.Time
		scheduleTable map[string]*Cronjob
	)
	scheduleTable = make(map[string]*Cronjob)
	now = time.Now()

	//第一个任务
	expr = cronexpr.MustParse("*/3 * * * * * *")
	cronjob = &Cronjob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	scheduleTable["job1"] = cronjob

	//第二个任务
	expr = cronexpr.MustParse("*/8 * * * * * *")
	cronjob = &Cronjob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	scheduleTable["job2"] = cronjob

	//启动一个调度协程
	go func() {
		var(
			jobName string
			cronjob *Cronjob
			now time.Time
		)
		for {
			now = time.Now()
			for jobName,cronjob = range scheduleTable {
				//判断是否过期
				if cronjob.nextTime.Before(now) || cronjob.nextTime.Equal(now){
					//启动一个协程，执行这个任务
					//go func(jobName string) {
						fmt.Println("执行:",jobName)
					//}(jobName)

					//计算下次调度时间
					cronjob.nextTime = cronjob.expr.Next(now)
					fmt.Println(jobName,"下次执行时间:",cronjob.expr.Next(now))
				}
			}
			//select {
			//case<-time.NewTimer(100*time.Millisecond).C
			//}
			time.Sleep(500*time.Millisecond)
		}
	}()
	time.Sleep(100 * time.Second)

}
