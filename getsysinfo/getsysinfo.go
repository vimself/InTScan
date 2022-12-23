package getsysinfo

import (
	"os"
	"os/user"
	"runtime"
)

type SystemInfo struct {
	OS          string //操作系统
	ARCH        string //处理器架构
	HostName    string //内核提供的主机名
	Groupid     string //初级组ID
	Userid      string //用户ID
	Username    string //用户名
	UserHomeDir string //用户的主目录
}

func GetSys() SystemInfo {
	var sysinfo SystemInfo

	sysinfo.OS = runtime.GOOS
	sysinfo.ARCH = runtime.GOARCH
	name, err := os.Hostname()
	if err == nil {
		sysinfo.HostName = name
	}

	u, err := user.Current()
	sysinfo.Groupid = u.Gid
	sysinfo.Userid = u.Uid
	sysinfo.Username = u.Username
	sysinfo.UserHomeDir = u.HomeDir

	return sysinfo
}
