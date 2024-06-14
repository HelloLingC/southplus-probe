package task

import (
	"log"

	"gopkg.in/toast.v1"
)

func notify() {
	notification := toast.Notification{
		AppID:   "Southplus-probe",
		Title:   "Southplus已开放注册",
		Message: "请前去网站进行注册吧",
		Actions: []toast.Action{
			{"protocol", "注册", ""},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatal("Err pushing notification", err)
	}
}
