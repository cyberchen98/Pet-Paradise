package model

import (
	"fmt"
	"github.com/satori/go.uuid"
	db "pet-paradise/model/common"
)

const TIME_FORMAT = "2006-01-02 15:04:05"

type userTable struct {
	db.Table
}

var UserTable = &userTable{db.Table{
	GetDB:     db.Conn,
	TableName: db.TBL_USER,
}}

type UserInfo struct {
	ID         string `db:"id" json:"id"`
	UserID     string `db:"uid" json:"uid"`
	Name       string `db:"user_name" json:"name"`
	Password   string `db:"user_password" json:"password"`
	Email      string `db:"user_email"json:"email" `
	Phone      string `db:"user_phone" json:"phone"`
	Role       string `db:"role" json:"role"`
	CreateTime string `db:"create_time" json:"create_time"`
	UpdateTime string `db:"update_time" json:"update_time"`
}

func (u *userTable) GetAllUserIds(whereCause string) ([]string, error) {
	query := "SELECT id FROM `" + u.TableName + "` WHERE is_deleted=0 " + whereCause
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

func (u *userTable) getOne(key, value string) (*UserInfo, error) {
	query := "SELECT id, uid, user_name, user_email, user_phone, role, create_time, update_time FROM `" + u.TableName + "` WHERE %s=?"
	info := &UserInfo{}
	if err := u.Get(info, fmt.Sprintf(query, key), value); err != nil {
		return nil, err
	}
	return info, nil
}

func (u *userTable) InsertNewUserInfo(userInfo UserInfo) error {
	m := make(map[string]interface{})
	m["uid"] = uuid.NewV4().String()
	m["user_name"] = userInfo.Name
	m["user_password"] = userInfo.Password
	m["user_email"] = userInfo.Email
	m["user_phone"] = userInfo.Phone
	m["role"] = userInfo.Role
	if _, err := u.Insert(m); err != nil {
		return err
	}
	return nil
}

func (u *userTable) UpdateUserInfoById(userInfo map[string]interface{}, id string) error {
	keys, values := _updateFiled(userInfo)
	if _, err := u.UpdateById(keys, id, values); err != nil {
		return err
	}
	return nil
}

func (u *userTable) DeleteUserInfoById(id string) error {
	if err := u.DeleteById(id); err != nil {
		return err
	}
	return nil
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
