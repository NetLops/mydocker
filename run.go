package main

import (
	"github.com/NetLops/mydocker/container"
	log "github.com/sirupsen/logrus"
	"os"
)

// 这里的Start方法是真正开始前面创建好的command的调用，它首先会clone出啦一个
// namespace隔离的进程，然后在子进程中，调用/proc/self/exe，也就是调用自己，
// 发送init参数，调用我们写的init方法
func Run(tty bool, command string) {
	parent := container.NewProcessProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}
