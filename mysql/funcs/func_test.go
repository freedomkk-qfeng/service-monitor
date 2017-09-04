package funcs

import (
	"fmt"
	"testing"
)

const (
	Addr     = "192.168.11.136"
	Port     = 3306
	Username = "root"
	Password = ""
)

func Test_mysql_stat(t *testing.T) {
	m := &MysqlIns{
		Host: Addr,
		Port: Port,
		Tag:  fmt.Sprintf("port=%d", Port),
	}
	data, err := MysqlStatus(m, Username, Password)
	t.Log(data)
	t.Error(err)
	version, err := MysqlVersion(m, Username, Password)
	t.Log(version)
	t.Error(err)

}
