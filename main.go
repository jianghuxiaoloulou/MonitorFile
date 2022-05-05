package main

import (
	"WowjoyProject/MonitorFile/global"
	"WowjoyProject/MonitorFile/pkg/object"
	"strings"

	"github.com/robfig/cron"
)

func main() {
	global.Logger.Debug("......启动监控程序......")
	MyCron := cron.New()
	MyCron.AddFunc(global.GeneralSetting.CronSpec, func() {
		global.Logger.Info("开始执行定时监控任务")
		work()
	})
	MyCron.Start()
	defer MyCron.Stop()
	select {}
}

func work() {
	// 读取配置文件获取需要监控的进程
	processList := object.ReadProcessFile()
	// 判断进程是否存在，不存在，重启进程
	for i := 0; i < len(processList); i++ {
		global.Logger.Debug(processList[i])
		countSplit := strings.Split(processList[i], " ")
		global.Logger.Debug(countSplit, len(countSplit))
		if len(countSplit) >= 2 {
			object.ProcessStart(countSplit[0], countSplit[1])
		} else {
			global.Logger.Error("监控配置文件错误：", processList[i])
		}
	}
}
