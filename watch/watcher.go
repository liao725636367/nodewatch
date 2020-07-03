package watch

import (
	"nodeWatch/def"
	path2 "path"

	"github.com/gogf/gf/os/gfsnotify"
	"github.com/gogf/gf/os/glog"
)

func WatchDir(path string) {
	_, err := gfsnotify.Add(path, func(event *gfsnotify.Event) {
		//if event.IsCreate() {
		//	def.Log.Debug("创建文件 : ", event.Path)
		//}
		//if event.IsWrite() {
		//	def.Log.Debug("写入文件 : ", event.Path)
		//}
		//if event.IsRemove() {
		//	def.Log.Debug("删除文件 : ", event.Path)
		//}
		//if event.IsRename() {
		//	def.Log.Debug("重命名文件 : ", event.Path)
		//}
		//if event.IsChmod() {
		//	def.Log.Debug("修改权限 : ", event.Path)
		//}
		ext := path2.Ext(event.Path)

		if ext != ".log" {
			def.Log.Debug("修改文件后缀", ext)
			def.NeedSync = true
		}

		//def.Log.Debug(event)
	})
	gfsnotify.Remove(path2.Join(path, ".git"))
	// 移除对该path的监听
	// gfsnotify.Remove(path)

	if err != nil {
		glog.Fatal(err)
	} else {
		select {}
	}
}
func cmdSync() {

}
