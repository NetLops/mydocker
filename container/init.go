package container

import (
	"github.com/sirupsen/logrus"
	"os"
	"syscall"
)

func RunContainerInitProcess(command string, args []string) error {
	logrus.Infof("command %s", command)

	/**
	MS_NOEXEC: 在本文件系统中不允许运行其他程序
	MS_NOSUID：在本系统中运行程序的时候，不允许set-user-ID 或 set-group-ID
	MS_NODEV: 这个参数是自从Linux 2.4以来，所有mount的系统都会默认设定的参数
	*/
	defaultMountFlags := syscall.MS_NOEXEC |
		syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	argv := []string{command}
	/**
	其实最终调用了Kernel的int execve（const char*filename，char*const argv[]，char*const envp[]）；
	这个系统函数。它的作用是执行当前filename对应的程序。
	它会覆盖当前进程的镜像、数据和堆栈等信息，包括PID，这些都会被将要运行的进程覆盖掉。
	也就是说，调用这个方法，
	将用户指定的进程运行起来，把最初的init进程给替换掉，
	这样当进入到容器内部的时候，就会发现容器内的第一个程序就是我们指定的进程了。
	这其实也是目前Docker使用的容器引擎runC的实现方式之一。
	*/
	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		logrus.Errorf(err.Error())
	}
	return nil
}
