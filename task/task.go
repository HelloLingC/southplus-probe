package task

import (
	"fmt"
	"southplus-probe/network"
	"strings"
	"time"
)

func check(url *string) bool {
	html := network.Req(*url)
	if strings.Contains(string(html), "Member only knows") {
		fmt.Println("\033[31m注册未开放 :(")
		return false
	} else {
		fmt.Println("\033[32m注册已开放")
		return true
	}
}

func Start(url *string) {
	check(url)
	fmt.Println("\033[0m正在后台运行中... 将在24小时后进行下一次检测")
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Println("\033[0m开始测试...")
			if check(url) {
				ticker.Stop()
				notify_in_loop()
			}
		}
	}
}

func notify_in_loop() {
	notify()
	ticker := time.NewTicker(2 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			notify()
		}
	}
}
