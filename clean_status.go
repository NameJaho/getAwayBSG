package main

import (
	"fmt"
	"getAwayBSG/db"
	"os"
	"strings"
)

func main() {

	var choice int
	if strings.Index(os.Args[1], "lianjia") > -1 {
		choice = 1
	} else if strings.Index(os.Args[1], "zhilian") > -1 {
		choice = 2
	} else {
		fmt.Println("清除抓取状态（不清除状态的话爬虫会从上次停止位置继续抓取）")
		fmt.Println("请选择需要清除哪个爬虫的的状态数据：（输入数字）")
		fmt.Println("1.链家")
		fmt.Println("2.智联")
		fmt.Scanln(&choice)

	}

	if choice == 1 {
		db.SetLianjiaStatus(0)
		fmt.Println("Done!")
	} else if choice == 2 {
		db.SetZhilianStatus(0, 0)
		fmt.Println("Done!")
	} else {
		fmt.Println("选择错误！")
	}

}
