package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"net/url"
	"time"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)


//var phones= []string{"18940037726","13386869011","19941497051","13342483502","13390242619","13390555760","13332497378"}
func main() {

	// 监听退出信号
	sigchan := make(chan os.Signal, 1)
	defer close(sigchan)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	tack:=time.Tick(time.Second*3)
	i:=0
	for  range tack{
		i++
		if i==5{
			break
			return
		}
		httpPost(getPhone())
		//i:=0
		//for {
		//	i++
		//	if i==5{
		//		break
		//	}
		//	go func() {
		//		getPhone()
		//	}()
		//}

	}

	<-sigchan
	{
	signal.Stop(sigchan)
	return
	}



}



func httpPost(phone string) {


	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}


var Line2 =[]string {"3","4","5","7","8"}
var Line33 =[]string {"0","1","2","3","4","5","6","7","8","9"}
var Line43 =[]string {"5","7","9"}
var Line53 =[]string {"0","1","2","3","4","6","7","8","9"}
var Line73 =[]string {"0","1","2","3","4","6","7","8"}
var Line83 =[]string {"0","1","2","3","4","5","6","7","8","9"}


func getPhone()string {

	 rand.Seed(time.Now().Unix())
	// 第二位数字
	second := Line2[rand.Intn(4)]

	third:=""
	if second=="3"{
		third=Line33[rand.Intn(len(Line33))]
	}
	if second=="4"{
		third=Line43[rand.Intn(len(Line43))]
	}
	if second=="5"{
		third=Line53[rand.Intn(len(Line53))]
	}
	if second=="7"{
		third=Line73[rand.Intn(len(Line73))]
	}
	if second=="8"{
		third=Line83[rand.Intn(len(Line83))]
	}


	ranPhone:=fmt.Sprintf("1%s%s%d",second,third,RandInt(9999999,100000000))
    fmt.Println(ranPhone)
	return ranPhone

}


func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}
