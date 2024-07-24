package main

import (
	"fmt"
	"time"

	"golang.org/x/sys/windows/svc"

	"service_status/winservicemanager"
)

func getState(service *winservicemanager.ManagedService) (svc.State, error) {
	status, err := service.QueryStatus()
	if err != nil {
		return svc.Stopped, err
	}
	return status.State, nil
}

func get() {
	service, err := winservicemanager.NewManagedService("Arc")
	if err != nil {
		fmt.Println("NewService error", err)
		return
	}

	state, _ := getState(service)
	printState(state)

	fmt.Println("requesting stop")
	// https://learn.microsoft.com/en-us/windows/win32/api/winsvc/nf-winsvc-controlservice
	control, err := service.Srv.Control(svc.Stop)
	if err != nil {
		fmt.Println("stop error:", err)
	}
	fmt.Print("returned status (stop request): ")
	printState(control.State)

	// poll for 1 min, this is not interruptible
	for i := 0; i < 12; i++ {
		time.Sleep(5 * time.Second)
		state, _ = getState(service)
		printState(state)
		if state == svc.Stopped {
			break
		}
	}

	state, _ = getState(service)
	printState(state)

	fmt.Println("Request start")
	startErr := service.Srv.Start()
	if startErr != nil {
		fmt.Println("Start err:", err)
		return
	}
}

func printState(state svc.State) string {
	switch state {
	case svc.Stopped:
		return "Stopped"
	case svc.StartPending:
		return "StartPending"
	case svc.StopPending:
		return "StopPending"
	case svc.Running:
		return "Running"
	case svc.ContinuePending:
		return "ContinuePending"
	case svc.PausePending:
		return "PausePending"
	case svc.Paused:
		return "Paused"
	default:
		return "unknown state"
	}
}
