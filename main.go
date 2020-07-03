package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"nodeWatch/def"
	"nodeWatch/helper"
	log2 "nodeWatch/log"
	service2 "nodeWatch/service"
	"os"

	"github.com/kardianos/service"
)

var ()

func main() {
	def.CurPath = helper.GetCurrPath()
	defer func() {

		if r := recover(); r != nil {
			str := fmt.Sprintf("捕获到的错误：%s\n", r)
			ioutil.WriteFile("D:/err.log", []byte(str), 0644)
		}
	}()
	//命令行解析
	//gcmd.BindHandle("config", config.Config)
	//gcmd.AutoRun()
	//日志初始化
	def.Log = log2.NewLogger() //实例化配置
	svcConfig := &service.Config{
		Name:        "nodesWatch",
		DisplayName: "nodesWatch",
		Description: "nodesWatch",
	}

	prg := &service2.Program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	service2.Logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) > 1 {
		if os.Args[1] == "help" {

			str := `
欢迎使用请安装服务后指定用户运行服务：
config 配置路径
install 安装服务
remove 移除服务
start 启动服务
stop 停止服务
restart 重启服务
			`
			fmt.Println(str)
			return
		}
		if os.Args[1] == "install" {
			err = s.Install()
			if err != nil {
				fmt.Println("服务安装失败:", err.Error())

			} else {

				fmt.Println("服务安装成功")
				err = s.Start()
				if err != nil {
					fmt.Println("服务启动失败:", err.Error())

				} else {
					s.Start()
					fmt.Println("服务启动成功")
				}
			}

			return
		}
		if os.Args[1] == "start" {

			err = s.Start()
			if err != nil {
				fmt.Println("服务启动失败:", err.Error())

			} else {
				fmt.Println("服务启动成功")
			}
			return
		}
		if os.Args[1] == "restart" {

			err = s.Restart()
			if err != nil {
				fmt.Println("服务重启失败:", err.Error())

			} else {
				fmt.Println("服务重启成功")
			}
			return
		}
		if os.Args[1] == "stop" {

			err = s.Stop()
			if err != nil {
				fmt.Println("服务停止失败:", err.Error())

			} else {
				fmt.Println("服务停止成功")
			}
			return
		}

		if os.Args[1] == "remove" {

			err = s.Uninstall()
			if err != nil {
				fmt.Println("服务卸载失败:", err.Error())
			} else {
				fmt.Println("服务卸载成功")
			}

			return
		}
	}
	err = s.Run()
	if err != nil {
		service2.Logger.Error(err)
	}
}
