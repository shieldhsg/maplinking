//配置文件
package common

type ConfigT struct {
	Db string
}

var Config = &ConfigT{}

func init() {
	Config = &ConfigT{
		Db: "root:3636592hsg@tcp(47.74.235.80:3306)/maplinking?charset=utf8",
	}
}
