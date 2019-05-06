package main

import (
	"fmt"
	"time"

)

func parseTime(t string)  {
	a,_:=time.Parse(time.ANSIC,t)
	fmt.Println(a.Date())


}
func parseDuration(t string) {
	d, err := time.ParseDuration(t)
	fmt.Printf("%v %v", d, err)
}
func sleep(t time.Duration) {
	time.Sleep(t)
}
func tick(ch <-chan time.Time)  {
	for t:=range ch{
		fmt.Println(t)
	}
}
func main() {
	now:=time.Now()
	fmt.Println("now:",now)
	fmt.Println("after 3 hours",now.Add(3*time.Hour))
	month:=now.Month()
	y,m,d := time.Now().Date()
	fmt.Println("y:",y,"m:",m,"d:",d)
	addDate:= time.Now().AddDate(1,1,1)
	fmt.Println("month:",month, "addDate:",addDate)

	parseTime("Mon May 8 16:19:35 2013")
    parseDuration("1h2m3s4ms5.8us6ns")
	sleep(100 * time.Millisecond)
	fmt.Println("go back")
	fmt.Println(time.Since(now))
//没有必要停止的定时器
//	for n1:= range time.Tick(1*time.Second){
//		fmt.Println("n=",n1)
//	}
	t1:=time.NewTicker(1*time.Second)
	go tick(t1.C)
	time.Sleep(3*time.Second)
	t1.Stop()

	timer := time.NewTimer(3*time.Second)
	t2:= <-timer.C
	fmt.Println("t2:",t2)
	fmt.Println(time.Now().Zone())
	fmt.Println(time.Now().ISOWeek())
	fmt.Println(time.Now().Clock())


}

