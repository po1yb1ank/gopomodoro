package main

import (
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"strconv"
	"time"
)
func alert(message string){
	out, err := exec.Command("cmd", "/C","msg", "%username%", "").Output()
	fmt.Println(string(out), err)
}
func main(){
	for i := 0; i < 4; i++{
		alert("MSG")
		Schedule()
	}

}
func Schedule(){
	rand.Seed(time.Now().UnixNano())
	d, err:= time.ParseDuration(strconv.Itoa(25)+"m")
	if err != nil{
		log.Println(err, ",will sleep for 5m")
		time.Sleep(time.Minute * 5)
		return
	}
	log.Println("Time until next commit: ", d.String())
	go func(d time.Duration) {
		for range time.Tick(time.Second * 2){
			if d <= time.Second * 0{
				break
			}
			d -= 2 * time.Second
			log.Println(d)
		}
	}(d)

	time.Sleep(d)
	log.Println("Done sleeping!")
	return
}
