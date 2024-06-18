//go:build windows

package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

func Get_auto_start() {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatalf("打开注册表键失败: %v", err)
	}
	defer k.Close()
	v, _, err := k.GetStringValue("Southplus-probe")
	if err != nil {
		fmt.Println("已关闭开机自启")
	} else {
		fmt.Println("已设置开机自启: " + v)
	}

}

func Set_auto_start(ac bool) {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("获取可执行文件路径失败: %v", err)
	}
	absPath, err := filepath.Abs(exePath)
	if err != nil {
		log.Fatalf("获取绝对路径失败: %v", err)
	}

	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatalf("打开注册表键失败: %v", err)
	}
	defer k.Close()
	if ac {
		err = k.SetStringValue("Southplus-probe", absPath)
	} else {
		err = k.DeleteValue("Southplus-probe")
	}
	if err != nil {
		log.Fatalf("设置注册表值失败: %v", err)
	}
}
