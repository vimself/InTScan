package mysqlscan

import (
	"InTScan/common"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

func MysqlScan(host, port string) (tmperr error) {
	starttime := time.Now().Unix()
	for _, user := range common.Userdict["mysql"] {

		for _, pass := range common.Passwords {

			pass = strings.Replace(pass, "{user}", user, -1)
			flag, err := MysqlConn(host, port, user, pass)
			if flag == true && err == nil {
				return err
			} else {
				errlog := fmt.Sprintf("[-] mysql %v:%v %v %v %v", host, port, user, pass, err)
				common.LogError(errlog)
				tmperr = err
				if common.CheckErrs(err) {
					return err
				}
				if time.Now().Unix()-starttime > (int64(len(common.Userdict["mysql"])*len(common.Passwords)) * common.Timeout) {
					return err
				}
			}
		}
	}
	return tmperr
}

func MysqlConn(host, port, user string, pass string) (flag bool, err error) {
	flag = false

	Host, Port, Username, Password := host, port, user, pass
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/mysql?charset=utf8&timeout=%v", Username, Password, Host, Port, time.Duration(common.Timeout)*time.Second)

	db, err := sql.Open("mysql", dataSourceName)

	if err == nil {
		db.SetConnMaxLifetime(time.Duration(common.Timeout) * time.Second)
		db.SetConnMaxIdleTime(time.Duration(common.Timeout) * time.Second)
		db.SetMaxIdleConns(0)
		defer db.Close()
		err = db.Ping()

		if err == nil {

			result := fmt.Sprintf("[+]succeed mysql:%v:%v ------> username:%v,password:%v", Host, Port, Username, Password)
			fmt.Println(result)
			flag = true
		}
	}

	return flag, err
}
