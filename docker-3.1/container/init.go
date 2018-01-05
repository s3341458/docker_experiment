package container

import (
	"os"
	"syscall"
	"github.com/Sirupsen/logrus"
)

func RunContainerInitProcess(command string, args []string) error {
	logrus.Infof("command %s", command)

    syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, "")
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	argv := []string{command}
	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		logrus.Errorf(err.Error())
	}
	return nil
}
