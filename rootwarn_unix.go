// +build !windows

package main

import (
	"log"
	"os"
	"time"
)

const interval = time.Hour

const warnInsecure = `

	************************************************************
	You are running fabio as root with the '-insecure' flag
	Please check the fabio wiki for alternatives
	************************************************************

`

const warn17behavior = `

	************************************************************
	You are running fabio as root without the '-insecure' flag
	This will stop working with fabio 1.7!
	************************************************************

`

func WarnIfRunAsRoot(allowRoot bool) {

	isRoot := os.Getuid() == 0
	if !isRoot {
		return
	}

	warn := warnInsecure
	if !allowRoot {
		warn = warn17behavior
	}

	go func() {
		for {
			log.Printf("[INFO] Running fabio as UID=%d EUID=%d GID=%d", os.Getuid(), os.Geteuid(), os.Getgid())
			log.Print("[WARN] ", warn)
			time.Sleep(interval)
		}
	}()
}
