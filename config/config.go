package config

import (
	"fmt"
	"io/ioutil"
	"nodeWatch/def"
	"nodeWatch/helper"
	"os"
	path2 "path"

	"github.com/gogf/gf/encoding/gjson"

	"github.com/gogf/gf/crypto/gaes"
)

var key16 = []byte("anjgkilnwhj81lng")
var confFile = "nodes.conf"

func Init() {
}

type UserInfo struct {
	Username string
	Password string
	SyncPath string
}

func GetUserInfo() (user *UserInfo, err error) {
	filePath := path2.Join(def.CurPath, confFile)
	helper.Prl("配置文件路径", filePath)
	con, err := ioutil.ReadFile(filePath)
	if err != nil {
		def.Log.Error("配置读取错误: ", err.Error(), con)
		return nil, err
	}

	str, err := gaes.Decrypt(con, key16)
	if err != nil {
		def.Log.Error("配置解密错误:%s", err.Error())
		return nil, err
	}
	user = new(UserInfo)
	err = gjson.DecodeTo(str, user)
	if err != nil {
		def.Log.Error("配置反序列化错误:%s", err.Error())
		return nil, err
	}
	//if user.Password == "" || user.Username == "" {
	//	def.Log.Error("用户名或者密码为空")
	//	return nil, errors.New("用户名或者密码为空")
	//}
	if user.SyncPath == "" { //默认使用当前目录路径
		user.SyncPath = def.CurPath
	}
	return user, nil
}
func Config() {
	user := new(UserInfo)
	//fmt.Println("请输入账号并输入回车:")
	//fmt.Scanln(&user.Username)
	//fmt.Println("请输入密码并输入回车:")
	//fmt.Scanln(&user.Password)
	fmt.Println("请输入监控路径并输入回车（默认当前路径）:")
	fmt.Scanln(&user.SyncPath)
	fmt.Println(fmt.Sprintf("您的账户密码同步路径是:%s-%s-%s", user.Username, user.Password))
	str, err := gjson.Encode(user)
	if err != nil {
		panic("转换配置错误:" + err.Error())
	}
	data, err := gaes.Encrypt(str, key16)
	if err != nil {
		panic("配置错误:" + err.Error())
	}
	file := path2.Join(def.CurPath, confFile)
	err = ioutil.WriteFile(file, data, 0644)
	if err != nil {
		panic("写入文件失败:" + err.Error())
	}
	fmt.Println("配置成功 请使用 install 安装服务")
	os.Exit(0)
}
