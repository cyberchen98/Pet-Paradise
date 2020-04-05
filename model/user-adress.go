package model

import (
	db "pet-paradise/model/common"
)

type addressTable struct {
	db.Table
}

var AddressTable = &addressTable{db.Table{
	GetDB:     db.Conn,
	TableName: db.TBL_USER_ADDRESS,
}}

type UserAddressInfo struct {
	ID          int    `db:"id" json:"id"`
	UserID      int    `db:"uid" json:"uid"`
	Province    string `db:"province" json:"province" form:"province"`
	City        string `db:"city" json:"city" form:"city"`
	Details     string `db:"details" json:"details"`
	PhoneNumber string `db:"phone_number" json:"phone_number" form:"phone_number"`
	Receiver    string `db:"receiver" json:"receiver" form:"receiver"`
	PostCode    string `db:"post_code" json:"post_code" form:"post_code"`
}

func (a *addressTable) SelectAddressInfoByUserId(uid int) ([]UserAddressInfo, error) {
	query := "SELECT id, uid, province, city, details, phone_number, receiver, post_code FROM `" + a.TableName + "` WHERE uid=? AND is_deleted='0'"
	var infoSlice []UserAddressInfo
	if err := a.Select(&infoSlice, query, uid); err != nil {
		return nil, err
	}
	return infoSlice, nil
}

func (a *addressTable) GetOneById(id int) (*UserAddressInfo, error) {
	query := "SELECT id, uid, province, city, details, phone_number, receiver, post_code FROM `" + a.TableName + "` WHERE is_deleted='0' AND id=?"
	info := &UserAddressInfo{}
	if err := a.Get(info, query, id); err != nil {
		return nil, err
	}
	return info, nil
}

func (a *addressTable) SelectByUserId(uid int) ([]UserAddressInfo, error) {
	query := "SELECT id, uid, province, city, details, phone_number, receiver, post_code FROM `" + a.TableName + "` WHERE is_deleted='0' AND uid=?"
	var info []UserAddressInfo
	if err := a.Select(info, query, uid); err != nil {
		return nil, err
	}
	return info, nil
}

func (a *addressTable) InsertNewAddressInfo(addressInfo UserAddressInfo) error {
	m := make(map[string]interface{})

	m["uid"] = addressInfo.UserID
	m["province"] = addressInfo.Province
	m["city"] = addressInfo.City
	m["details"] = addressInfo.Details
	m["receiver"] = addressInfo.Receiver
	m["post_code"] = addressInfo.PostCode
	if _, err := a.Insert(m); err != nil {
		return err
	}
	return nil
}

func (a *addressTable) UpdateAddressInfoById(addressInfo UserAddressInfo, id int) error {
	var addressInfoMap = make(map[string]interface{})

	if addressInfo.Province != "" {
		addressInfoMap["province"] = addressInfo.Province
	}
	if addressInfo.City != "" {
		addressInfoMap["city"] = addressInfo.City
	}
	if addressInfo.Details != "" {
		addressInfoMap["details"] = addressInfo.Details
	}
	if addressInfo.Receiver != "" {
		addressInfoMap["receiver"] = addressInfo.Receiver
	}
	if addressInfo.PhoneNumber != "" {
		addressInfoMap["phone_number"] = addressInfo.PhoneNumber
	}
	if addressInfo.PostCode != "" {
		addressInfoMap["post_code"] = addressInfo.PostCode
	}

	keys, values := _updateFiled(addressInfoMap)
	if _, err := a.UpdateById(keys, id, values...); err != nil {
		return err
	}
	return nil
}

func (a *addressTable) DeleteAddressInfoById(id int) error {
	if err := a.DeleteById(id); err != nil {
		return err
	}
	return nil
}
