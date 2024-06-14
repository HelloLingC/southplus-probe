package main

import (
	"flag"
	"fmt"
	"southplus-probe/task"
	"southplus-probe/utils"
)

func main() {
	url := flag.String("url", "https://south-plus.org/register.php", "")
	fmt.Println(*url)
	utils.Set_auto_start()
	task.Start(url)
}
