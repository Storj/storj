// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

// +build windows

package main

import (
	"log"
	"time"

	"golang.org/x/sys/windows/svc"
	"storj.io/storj/pkg/process"
)

func init() {
	interactive, err := svc.IsAnInteractiveSession()
	if err != nil {
		log.Fatalf("Failed to determine if session is interactive: %v", err)
	}

	if interactive {
		return
	}

	err = svc.Run("storagenode", &service{})
	if err != nil {
		log.Fatalf("Service failed: %v", err)
	}
}

type service struct{}

func (m *service) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown

	changes <- svc.Status{State: svc.StartPending}

	go func() {
		process.Exec(rootCmd)
	}()

	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

	for {
		select {
		case c := <-r:
			switch c.Cmd {
			case svc.Interrogate:
				log.Println("Interrogate request received.")
				changes <- c.CurrentStatus
				// Testing deadlock from https://code.google.com/p/winsvc/issues/detail?id=4
				time.Sleep(100 * time.Millisecond)
				changes <- c.CurrentStatus
			case svc.Stop, svc.Shutdown:
				log.Println("Stop/Shutdown request received.")
				changes <- svc.Status{State: svc.StopPending}
				// Cancel the command's root context to cleanup resources
				_, cancel := process.Ctx(runCmd)
				cancel()
				// Sleep some time to give chance for goroutines finish cleanup after cancelling the context
				time.Sleep(15 * time.Second)
				// After returning the Windows Service is stopped and the process terminates
				return
			default:
				log.Printf("Unexpected control request: %d\n", c)
			}
		}
	}
}
