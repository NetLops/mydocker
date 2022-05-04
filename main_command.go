package main

import (
	"fmt"
	"github.com/NetLops/mydocker/container"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// 定义runCommand的Flags。类似于运行命令时使用 --指定参数
var runCommand = cli.Command{
	Name:  "run",
	Usage: `Create a container with namespace and cgroups limit mydocker run -it [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
	},
	//
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		// 判断参数是否包含command
		cmd := context.Args().Get(0)
		// 获取用户指定的command
		tty := context.Bool("ti")
		// 调用Run function 去准备启动容器
		Run(tty, cmd)
		return nil
	},
}

// 定义initCommand 具体操作，该操作为内部方法，禁止外部调用
var initCommand = cli.Command{
	Name: "init",
	Usage: "Init container process run user`s process in container ." +
		"Do not call it outside",
	Action: func(context *cli.Context) error {
		log.Infof("init come on")
		// 获取传递过来的command函数
		cmd := context.Args().Get(0)
		log.Infof("command %s", cmd)
		// 执行容器初始化操作
		container.RunContainerInitProcess(cmd, nil)
		return nil
	},
}
