package task

import (
	"log"

	"gopkg.in/toast.v1"
)

func notify() {
	notification := toast.Notification{
		AppID:   "Southplus-probe",
		Title:   "Southplus now is opem for registeration",
		Message: ":)) Delete the exe file to stop the notification loop",
		Actions: []toast.Action{
			{"protocol", "Go", "https://south-plus.org/register.php"},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatal("Err pushing notification", err)
	}
}
