package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	now := time.Now().Unix()
	//月、日、年
	endtime, _ := time.Parse("01/02/2006", "03/08/2019")
	//授权时间判断
	if endtime.Unix() <= now {
		fmt.Println("您的免费使用时间已过，请购买正式套餐正式使用!")
		//time.Sleep(100 * time.Minute)
		return
	}

	//目标文件位置判断
	disk := []string{"c", "d", "e"}
	program_file := []string{"Program Files", "Program Files (x86)"}
	//datapath := "C:/\"Program Files (x86)\"/Tencent/WeChat/WeChat.exe"
	base_path := "/Tencent/WeChat/WeChat.exe"
	var wechat_path, tmp_path string

	for _, dv := range disk {
		for _, pv := range program_file {
			tmp_path = dv + ":/" + pv + base_path
			_, err := os.Stat(tmp_path)
			//找到程序了
			if err == nil {
				wechat_path = dv + ":/\"" + pv + "\"" + base_path
			}
		}
	}

	if wechat_path == "" {
		fmt.Println("末检测到微信，请安装pc版本微信")
		time.Sleep(100 * time.Minute)
		return

	}

	//启动微信
	for a := 0; a < 2; a++ {
		cmd := exec.Command("cmd.exe", "/c", "start "+wechat_path)
		cmd.Run()
	}

	tip_cmd := exec.Command("echo", "双微信启动成完成!")
	buf, _ := tip_cmd.Output() // 错误处理略
  fmt.Println(string(buf))
}
