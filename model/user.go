package model

import (
	"database/sql"
	"fmt"
	db "pet-paradise/model/common"
)

var (
	TIME_FORMAT  = "2006-01-02 15:04:05"
	DEFAULT_ROLE = "common"
	ROLES        = []string{"common", "admin", "vip"}
)

type userTable struct {
	db.Table
}

var UserTable = &userTable{db.Table{
	GetDB:     db.Conn,
	TableName: db.TBL_USER,
}}

type UserInfo struct {
	ID         int    `db:"id" json:"id"`
	Name       string `db:"user_name" json:"username" form:"username"`
	Password   string `db:"user_password" json:"-" form:"password"`
	Email      string `db:"user_email"json:"email" form:"email"`
	Phone      string `db:"user_phone" json:"phone" form:"phone"`
	Role       string `db:"role" json:"role" form:"role"`
	CreateTime string `db:"create_time" json:"create_time"`
	UpdateTime string `db:"update_time" json:"update_time"`
}

func (u *userTable) GetAllUserIds(whereCause string) ([]string, error) {
	query := "SELECT id FROM `" + u.TableName + "` WHERE is_deleted='0' " + whereCause
	var ids []string
	if err := u.Select(ids, query); err != nil {
		return nil, err
	}
	return ids, nil
}

func (u *userTable) GetOneByName(userName string) (*UserInfo, error) {
	return u.getOne("user_name", userName)
}

func (u *userTable) GetOneById(id string) (*UserInfo, error) {
	return u.getOne("id", id)
}

func (u *userTable) getOne(key, value interface{}) (*UserInfo, error) {
	query := "SELECT id, user_name, user_password, user_email, user_phone, role, create_time, update_time FROM `" + u.TableName + "` WHERE is_deleted='0' AND %s=?"
	info := &UserInfo{}
	if err := u.Get(info, fmt.Sprintf(query, key), value); err != nil {
		return nil, err
	}
	return info, nil
}

func (u *userTable) InsertNewUserInfo(userInfo UserInfo) (sql.Result, error) {
	m := make(map[string]interface{})
	m["user_name"] = userInfo.Name
	m["user_password"] = userInfo.Password
	m["user_email"] = userInfo.Email
	m["user_phone"] = userInfo.Phone
	m["role"] = DEFAULT_ROLE
	return u.Insert(m)
}

func (u *userTable) UpdateUserInfoById(userInfo UserInfo, id string) (sql.Result, error) {
	var userInfoMap = make(map[string]interface{})
	if userInfo.Role != "" {
		roleValid := false
		for _, v := range ROLES {
			if userInfo.Role == v {
				roleValid = true
			}
		}
		if !roleValid {
			return nil, fmt.Errorf("invalid role info, must in %v", ROLES)
		}
		userInfoMap["role"] = userInfo.Role
	}

	if userInfo.Email != "" {
		userInfoMap["user_email"] = userInfo.Email
	}

	if userInfo.Phone != "" {
		userInfoMap["user_phone"] = userInfo.Phone
	}

	if userInfo.Password != "" {
		userInfoMap["user_password"] = userInfo.Password
	}

	keys, values := _updateFiled(userInfoMap)
	return u.UpdateById(keys, []string{id}, values...)
}

func (u *userTable) DeleteUserInfoById(id []string) (sql.Result, error) {
	return u.DeleteById(id)
}

func _updateFiled(info map[string]interface{}) ([]string, []interface{}) {
	var keys []string
	var values []interface{}
	for k, v := range info {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}
