package service

import (
	"fmt"
	"nodeWatch/config"
	"nodeWatch/def"
	"nodeWatch/helper"
	"nodeWatch/watch"
	"time"

	"github.com/kardianos/service"
)

var Logger service.Logger

type Program struct{}

func (p *Program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *Program) run() {

	config.Init() //初始化配置
	user, err := config.GetUserInfo()
	if err != nil {
		def.Log.Error(err.Error())
		def.Log.Debug("配置读取错误,将使用当前目录 ")
		user = new(config.UserInfo)
		user.SyncPath = def.CurPath
	}
	//设置定时器定期更新

	// 设置终端执行参数
	//path := helper.GetCurrPath()
	def.Log.Debug("开始监控")

	helper.UserTicker(60000, cmdSyncGit, user)
	watch.WatchDir(user.SyncPath)

	// Do work here
}

func (p *Program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func cmdSyncSvn(user *config.UserInfo) {
	//var res string
	//var err error
	def.Log.Debug("定时器执行")
	if def.NeedSync == true {

		//svnUpdate := fmt.Sprintf(" svn update --username %s --password %s --no-auth-cache %s", user.Username, user.Password, user.SyncPath)
		//svnAdd := fmt.Sprintf(" cd /d %s && svn add . ", user.SyncPath)
		//svnCommit := fmt.Sprintf(" svn commit  -m \"auto commit\" --username %s --password %s --no-auth-cache %s", user.Username, user.Password, user.SyncPath)
		//def.Log.Debug(svnUpdate)
		//def.Log.Debug(svnAdd)
		//def.Log.Debug(svnCommit)
		AutoRm(user)
		//res, err = helper.RunCmd(svnAdd)
		//if err != nil {
		//	def.Log.Error("添加svn数据失败:", err.Error())
		//} else {
		//	def.Log.Debug("添加svn文件成功:", res)
		//}
		//res, err = helper.RunCmd(svnUpdate, 20)
		//if err != nil {
		//	def.Log.Error("更新svn失败:", err.Error())
		//} else {
		//	def.Log.Debug("更新svn文件成功:", res)
		//}
		//res, err = helper.RunCmd(svnCommit, 20)
		//if err != nil {
		//	def.Log.Error("提交svn失败:", err.Error())
		//} else {
		//	def.Log.Debug("提交svn文件成功:", res)
		//}
		//def.NeedSync = false
	}
}
func AutoRm(user *config.UserInfo) {
	cmdStr := fmt.Sprintf(" cd /d %s &&  svn status", user.SyncPath)
	str, err := helper.RunCmd(cmdStr)
	if err != nil {
		def.Log.Error("获取svn状态失败:", err.Error(), "结果", str, "命令", cmdStr)
		return
	}
	def.Log.Debug(str)

}
func cmdSyncGit(user *config.UserInfo) {
	var res string
	var err error
	//def.Log.Debug("定时器执行")
	if def.NeedSync == true {
		cmds := []string{
			fmt.Sprintf(" cd /d %s && git add .", user.SyncPath),
			fmt.Sprintf(" cd /d %s && git commit -m \"%d\"", user.SyncPath, time.Now().Unix()),
			fmt.Sprintf(" cd /d %s && git pull ", user.SyncPath),
			fmt.Sprintf(" cd /d %s && git push ", user.SyncPath),
		}
		for _, cmd := range cmds {
			res, err = helper.RunCmd(cmd, 30)
			if err != nil {
				def.Log.Error(fmt.Sprintf("执行命令(%s)错误:%s", cmd, err.Error()))
			} else {
				def.Log.Debug(fmt.Sprintf("执行命令(%s)成功:%s", cmd, res))
			}
		}

		def.NeedSync = false
	}
}
