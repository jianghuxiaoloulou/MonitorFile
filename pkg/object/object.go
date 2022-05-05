package object

import (
	"WowjoyProject/MonitorFile/global"
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"syscall"
)

func ReadProcessFile() (processName []string) {
	//os.Open是只读模式
	fileObj, err := os.Open(global.GeneralSetting.MonitorCfg)
	if err != nil {
		global.Logger.Error(err)
	}
	//一般情况下使用这种方式关闭文件
	defer fileObj.Close()

	//需要将文件对象传进去
	reader := bufio.NewReader(fileObj)
	for {
		//按行读取
		text, _, err := reader.ReadLine() //参数是字符，不是字符串
		if err == io.EOF {
			break
		}
		processName = append(processName, string(text))
	}
	return processName
}

// 获取进程ID
func GetPid(processName string) (int, error) {
	// 通过wmic process get name,processid | findstr server.exe 获取进程ID
	buf := bytes.Buffer{}
	cmd := exec.Command("wmic", "process", "get", "name,processid")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = &buf
	cmd.Run()

	cmd2 := exec.Command("findstr", processName)
	cmd2.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd2.Stdin = &buf
	data, _ := cmd2.CombinedOutput()
	if len(data) == 0 {
		return -1, errors.New("not find")
	}
	info := string(data)
	// 这里通过正则把进程id提取出来
	reg := regexp.MustCompile(`[0-9]+`)
	pid := reg.FindString(info)
	return strconv.Atoi(pid)
}

// 启动进程
func StartProcess(exePath string) error {
	global.Logger.Debug("开始启动进程：", exePath)
	cmd := exec.Command("cmd.exe", "/c", exePath)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	if err := cmd.Run(); err != nil {
		global.Logger.Error("启动程序失败", err)
		return err
	}
	return nil
}

func ProcessStart(proName, proPath string) {
	// 查找进程
	_, err := GetPid(proName)
	if err != nil {
		global.Logger.Error("进程不存在：", proName)
		// 重新启动进程
		go StartProcess(proPath)
	} else {
		global.Logger.Debug("进程存在：", proName)
	}
}
