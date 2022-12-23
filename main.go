package main

import (
	"InTScan/icmpcheck"
	"InTScan/mysqlscan"
	"InTScan/portscan"
	"fmt"
	"github.com/malfunkt/iprange"
	"os"
	"regexp"
	"strings"
)

func main() {
	var argv []string

	for _, s := range os.Args {
		argv = append(argv, s)
	}

	if len(os.Args) == 1 {
		fmt.Println(`'   __   __      ________      ___ __ __      ______       ______       __           ______    
'  /_/\ /_/\    /_______/\    /__//_//_/\    /_____/\     /_____/\     /_/\         /_____/\   
'  \:\ \\ \ \   \__.::._\/    \::\| \| \ \   \::::_\/_    \::::_\/_    \:\ \        \::::_\/_  
'   \:\ \\ \ \     \::\ \      \:.      \ \   \:\/___/\    \:\/___/\    \:\ \        \:\/___/\ 
'    \:\_/.:\ \    _\::\ \__    \:.\-/\  \ \   \_::._\:\    \::___\/_    \:\ \____    \:::._\/ 
'     \ ..::/ /   /__\::\__/\    \. \  \  \ \    /____\:\    \:\____/\    \:\/___/\    \:\ \   
'      \___/_(    \________\/     \__\/ \__\/    \_____\/     \_____\/     \_____\/     \_\/   
'                                                                                             `)
		fmt.Println("ServerScan for Port Scaner.\nVersion: v1.0.0\nBy: vimself\n" +
			"HOST -h Host to be scanned, supports four formats:\n\t\t192.168.1.1\n\t\t192.168.1.0/24\n" +
			"PORT -p Customize port list, separate with ',' example: 21,22,80-99,8000-8080 ...\n" +
			"Mysql_intruder -t IP:PORT example: 192.168.1.1:3306")
		os.Exit(0)
	}
	if argv[1] == "-t" {

		if argv[2] != "" {
			var des string = argv[2]
			split := strings.Split(des, ":")
			mysqlscan.MysqlScan(split[0], split[1])
		}
		os.Exit(0) //退出程序返回状态码0

	}
	if argv[1] == "-h" {
		defer os.Exit(0)

		hosts := argv[2]

		hostsPattern := `[\d]{1,3}?[.][\d]{1,3}[.][\d]{1,3}[.][\d]{1,3}([\/][\d]{2})?`
		hostsRegexp := regexp.MustCompile(hostsPattern)
		checkHost := hostsRegexp.MatchString(hosts)

		if hosts == "" || checkHost == false {
			fmt.Println("HOST  Host to be scanned, supports four formats:\n\t\t192.168.1.1\n\t\t192.168.1.0/24")
			os.Exit(0)
		}

		var hostLists []string
		hostlist, err := iprange.ParseList(hosts)
		if err == nil {
			hostsList := hostlist.Expand()
			for _, host := range hostsList {
				host := host.String()
				hostLists = append(hostLists, host)
			}
		} else {
			fmt.Println("HOST  Host to be scanned, supports four formats:\n\t\t192.168.1.1\n\t\t192.168.1.1/24")
			os.Exit(0)
		}

		if argv[3] == "" || argv[3] != "-p" {
			fmt.Println("Please input -parameter to achieve what you want to do,such -p 8080")
			os.Exit(0)
		}
		ports := argv[4]
		portsPattern := `^([0-9]|[1-9]\d|[1-9]\d{2}|[1-9]\d{3}|[1-5]\d{4}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])$|^\d+(-\d+)?(,\d+(-\d+)?)*$`
		portsRegexp := regexp.MustCompile(portsPattern)
		checkPort := portsRegexp.MatchString(ports)
		if ports != "" && checkPort == false {
			fmt.Println("PORT Error.  Customize port list, separate with ',' example: 21,22,80-99,8000-8080 ...")
			os.Exit(0)
		}

		AliveHosts := icmpcheck.ICMPRun(hostLists) //存活的ip地址集
		for _, host := range AliveHosts {
			fmt.Printf("(ICMP) Target '%s' is alive\n", host)
		}
		_, _ = portscan.TCPportScan(AliveHosts, ports, "icmp")
	}

}
