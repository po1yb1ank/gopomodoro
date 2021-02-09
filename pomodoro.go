package main

import (
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"strconv"
	"time"
)

func alert(message string) {
	out, err := exec.Command("cmd", "/C", "msg", "%username%", message).Output()
	fmt.Println(string(out), err)
}
func main() {
	alert("Hello. This is pomodoro method. After 15 seconds the timer will start. ")
	time.Sleep(time.Second * 15)
	for {
		for i := 0; i < 4; i++ {
			alert("Work for 25 minutes")
			schedule(25)
			if i != 3 {
				alert("Have a rest for 5 minutes")
				schedule(5)
				continue
			}
			alert("Congratulations! You deserved a 20 minute long rest. Keep it up!")
			schedule(20)
		}
	}
}
func schedule(min int) {
	rand.Seed(time.Now().UnixNano())
	d, err := time.ParseDuration(strconv.Itoa(min) + "m")
	if err != nil {
		time.Sleep(time.Minute * 25)
		return
	}
	log.Println("Time left: ", d.String())
	go func(d time.Duration) {
		for range time.Tick(time.Second * 30) {
			if d <= time.Second*0 {
				break
			}
			d -= 30 * time.Second
			log.Println(d)
		}
	}(d)
	time.Sleep(d)
	return
}
