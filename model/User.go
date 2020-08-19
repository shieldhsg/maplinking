package model

import (
	"fmt"
	"maplinking/common"
)

type User struct {
	Id               int     `json:"id"`
	Account          string  `json:"account"`
	Phone            string  `json:"phone"`
	Name             string  `json:"name"`
	Password         string  `json:"password"`
	Birthday         string  `json:"birthday"`
	Headimg          string  `json:"headimg"`
	Sex              int     `json:"sex"`
	Country          string  `json:"country"`
	Area             string  `json:"area"`
	Religion         string  `json:"religion"`
	Fans             int     `json:"fans"`
	Level            int     `json:"level"`
	CreateTime       []uint8 `json:"create_time"`
	UpdateTime       []uint8 `json:"update_time"`
	First_Login_Time []uint8 `json:"first_login_time"`
	Last_Login_Time  []uint8 `json:":"last_login_time"`
}

//根据用户id获取用户详情
func (m *User) GetUserInfo(id int) bool {
	return common.Db().First(&m, id).RecordNotFound()
}

func (m *User) TableName() string {
	return "u_users"
}

//登录 todo 待修改 需要密码加密
func (m *User) GetUserByLogin(account string, password string) bool {
	fmt.Print(account)
	fmt.Print(password)
	return common.Db().Where("account = ? and password = ?", account, password).First(&m).RecordNotFound()
}

//注册 创建用户
func (m *User) CreateUser() bool {
	common.Db().Create(*m)
	if !common.Db().NewRecord(*m) {
		return true
	} else {
		return false
	}
}

//填充 修改用户
func (m *User) ModifyUser(id int, name string, password string, headimg string) bool {
	common.Db().First(&m, id)
	m.Name = name
	m.Password = password
	m.Headimg = headimg
	common.Db().Save(&m)
	return true
}

//获取用户创建的事件

//获取用户参与的事件
