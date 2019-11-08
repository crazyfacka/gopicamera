package process

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"time"
)

var proc *os.Process
var running = false

func killProcess() error {
	var err error

	err = proc.Signal(os.Interrupt)
	if err != nil {
		_, err = proc.Wait()
	}

	if err == nil {
		running = false
		time.Sleep(5 * time.Second)
	}

	return err
}

func startProcess(cmd string, argv []string) error {
	var err error

	if running {
		log.Printf("%s already running\n", cmd)
		return errors.New("Process already running")
	}

	proc, err = os.StartProcess(cmd, argv, nil)

	if err != nil {
		log.Printf("Error starting process: %s\n", err.Error())
	}

	return err
}

func resetProcess(cmd string) error {
	err := exec.Command("killall " + cmd).Run()

	running = false
	time.Sleep(5 * time.Second)

	return err
}
