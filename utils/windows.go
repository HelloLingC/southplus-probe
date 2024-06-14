//go:build windows

package utils

import (
	"log"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

func Set_auto_start() {
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

	err = k.SetStringValue("Southplus-probe", absPath)
	if err != nil {
		log.Fatalf("设置注册表值失败: %v", err)
	}

}
