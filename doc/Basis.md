### 类型判断
>判断字符串 or 整型    希望自己通过例子能渗透其他
```
if _, err := strconv.Atoi(v); err == nil {
    fmt.Printf("%q looks like a number.\n", v)
}
```
>判断是否是 int64
```
if _, err := strconv.ParseInt(v,10,64); err == nil {
    fmt.Printf("%q looks like a number.\n", v)
}
```
### 查找字符串
```
crontabTime := "We are geek"
strings.Count(schedValue," ")
```
### 字符串分割
```
crontabTime := "We are geek"
Sched := strings.Split(crontabTime, " ")
//Sched[0]  :   We
//Sched[1]  :   are
//Sched[2]  :   geek
```
### 字符串按照位数切消
```
crontabTime := "Tgeek"
schedValue[1:len(sc)]
```
### 时区
>时区是时间运算非常重要的概念，特别强调与layout是两个完全不同的概念。go语言通过Location来作为时区的运行实例，同一时刻转换成为不同的时区，就需要通过不同的Location来进行。默认情况下，采用UTC（unix标准时间），而不是过去式的GMT（格林尼治标准时间）。 
 以下代码展示了UTC标准、北京、美国洛杉矶在同一时刻的转换：
 ```
     now := time.Now()
     local1, err1 := time.LoadLocation("") //等同于"UTC"
     if err1 != nil {
         fmt.Println(err1)
     }
     local2, err2 := time.LoadLocation("Local")//服务器设置的时区
     if err2 != nil {
         fmt.Println(err2)
     }
     local3, err3 := time.LoadLocation("America/Los_Angeles")
     if err3 != nil {
         fmt.Println(err3)
     }
 
     fmt.Println(now.In(local1))
     fmt.Println(now.In(local2))
     fmt.Println(now.In(local3))
     //output:
     //2016-12-04 07:39:06.270473069 +0000 UTC
     //2016-12-04 15:39:06.270473069 +0800 CST
     //2016-12-03 23:39:06.270473069 -0800 PST

    loc, _:= time.LoadLocation("Asia/Shanghai")
    fmt.Println(time.Now().In(loc))
 ```
 ### 关于string、int、int64之间的转换
 ```
#string到int
    int,err:=strconv.Atoi(string)
#string到int64
    int64, err := strconv.ParseInt(string, 10, 64)
#int到string
    string:=strconv.Itoa(int)
#int64到string
    string:=strconv.FormatInt(int64,10)
```

### 时间加减 time
```
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Add 时间相加
	now := time.Now()
	// ParseDuration parses a duration string.
	// A duration string is a possibly signed sequence of decimal numbers,
	// each with optional fraction and a unit suffix,
	// such as "300ms", "-1.5h" or "2h45m".
	//  Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// 10分钟前
	m, _ := time.ParseDuration("-1m")
	m1 := now.Add(m)
	fmt.Println(m1)

	// 8个小时前
	h, _ := time.ParseDuration("-1h")
	h1 := now.Add(8 * h)
	fmt.Println(h1)

	// 一天前
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(d)
	fmt.Println(d1)

	printSplit(50)

	// 10分钟后
	mm, _ := time.ParseDuration("1m")
	mm1 := now.Add(mm)
	fmt.Println(mm1)

	// 8小时后
	hh, _ := time.ParseDuration("1h")
	hh1 := now.Add(hh)
	fmt.Println(hh1)

	// 一天后
	dd, _ := time.ParseDuration("24h")
	dd1 := now.Add(dd)
	fmt.Println(dd1)

	printSplit(50)

	// Sub 计算两个时间差
	subM := now.Sub(m1)
	fmt.Println(subM.Minutes(), "分钟")

	sumH := now.Sub(h1)
	fmt.Println(sumH.Hours(), "小时")

	sumD := now.Sub(d1)
	fmt.Printf("%v 天\n", sumD.Hours()/24)

}

func printSplit(count int) {
	fmt.Println(strings.Repeat("#", count))
}
```
### 合理的使用range配合for可以有效循环map数组
```
    for i,v := range tmpArr {
   	fmt.Println(i)
   	fmt.Println(v)
    }
```           