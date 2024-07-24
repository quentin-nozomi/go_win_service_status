package winservicemanager

import (
	"unsafe"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

type ManagedService struct {
	Name   string
	Config mgr.Config
	Status ServiceStatus
	Srv    *mgr.Service
}

// ServiceStatus combines State and Accepted commands to fully describe running service.
type ServiceStatus struct {
	State         svc.State
	Accepts       svc.Accepted
	Pid           uint32
	Win32ExitCode uint32
}

// An existing service managed by Windows
func NewManagedService(name string) (*ManagedService, error) {
	service, err := getService(name)
	if err != nil {
		return nil, err
	}
	return &ManagedService{
		Name: name,
		Srv:  service,
	}, nil
}

func (s *ManagedService) getServiceDetail() error {
	config, err := s.QueryServiceConfig()
	if err != nil {
		return err
	}
	s.Config = config

	status, err := s.QueryStatus()
	if err != nil {
		return err
	}
	s.Status = status

	return nil
}

// implement windows https://msdn.microsoft.com/en-us/library/windows/desktop/ms684932(v=vs.85).aspx
func (s *ManagedService) QueryServiceConfig() (mgr.Config, error) {
	return s.Srv.Config()
}

func (s *ManagedService) QueryStatus() (ServiceStatus, error) {
	var p *windows.SERVICE_STATUS_PROCESS
	var bytesNeeded uint32
	var buf []byte

	if err := windows.QueryServiceStatusEx(s.Srv.Handle, windows.SC_STATUS_PROCESS_INFO, nil, 0, &bytesNeeded); err != windows.ERROR_INSUFFICIENT_BUFFER {
		return ServiceStatus{}, err
	}

	buf = make([]byte, bytesNeeded)
	p = (*windows.SERVICE_STATUS_PROCESS)(unsafe.Pointer(&buf[0]))
	if err := windows.QueryServiceStatusEx(s.Srv.Handle, windows.SC_STATUS_PROCESS_INFO, &buf[0], uint32(len(buf)), &bytesNeeded); err != nil {
		return ServiceStatus{}, err
	}

	return ServiceStatus{
		State:         svc.State(p.CurrentState),
		Accepts:       svc.Accepted(p.ControlsAccepted),
		Pid:           p.ProcessId,
		Win32ExitCode: p.Win32ExitCode,
	}, nil
}
