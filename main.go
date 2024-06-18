package main

import (
	"flag"
	"fmt"
	"southplus-probe/task"
	"southplus-probe/utils"
)

func main() {
	url := flag.String("url", "https://south-plus.org/register.php", "")
	noas := flag.Bool("noautostart", false, "Disable auto start")
	flag.Parse()
	utils.Set_auto_start(!*noas)
	if *noas {
		fmt.Println("已关闭开机自启")
		return
	}
	fmt.Println(*url)
	utils.Get_auto_start()
	task.Start(url)
}
