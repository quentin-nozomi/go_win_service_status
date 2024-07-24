package main

import (
	"fmt"

	"golang.org/x/sys/windows/svc"

	"github.com/shirou/gopsutil/winservices"
)

func get() {
	service, err := winservices.NewService("Arc")
	if err != nil {
		fmt.Println("NewService error", err)
		return
	}
	status, err := service.QueryStatus()
	if err != nil {
		fmt.Println("QueryStatus error", err)
		return
	}

	switch status.State {
	case svc.Stopped:
		fmt.Println("Stopped")
	case svc.StartPending:
		fmt.Println("StartPending")
	case svc.StopPending:
		fmt.Println("StopPending")
	case svc.Running:
		fmt.Println("Running")
	case svc.ContinuePending:
		fmt.Println("ContinuePending")
	case svc.PausePending:
		fmt.Println("PausePending")
	case svc.Paused:
		fmt.Println("Paused")
	default:
		fmt.Println("unknown state")
	}
}
