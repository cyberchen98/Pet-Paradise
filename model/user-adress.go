package model

import (
	"github.com/satori/go.uuid"
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
	ID          string `db:"id" json:"id"`
	AddressID   string `db:"aid" json:"aid"`
	UserID      string `db:"uid" json:"uid"`
	Province    string `db:"province" json:"province"`
	City        string `db:"city" json:"city"`
	Details     string `db:"details" json:"details"`
	PhoneNumber string `db:"phone_number" json:"phone_number"`
	Receiver    string `db:"receiver" json:"receiver"`
	PostCode    string `db:"post_code" json:"post_code"`
}

func (a *addressTable) GetAllIdsByUserId(uid string) ([]string, error) {
	query := "SELECT id FROM `" + a.TableName + "` WHERE uid=? AND is_deleted=0"
	var ids []string
	if err := a.Select(ids, query, uid); err != nil {
		return nil, err
	}
	return ids, nil
}

func (a *addressTable) GetOneById(id string) (*UserAddressInfo, error) {
	query := "SELECT id, aid, uid, province, city, details, phone_number, receiver, post_code FROM `" + a.TableName + "` WHERE id=?"
	info := &UserAddressInfo{}
	if err := a.Get(info, query, id); err != nil {
		return nil, err
	}
	return info, nil
}

func (a *addressTable) SelectById(uid string) ([]UserAddressInfo, error) {
	query := "SELECT id, aid, uid, province, city, details, phone_number, receiver, post_code FROM `" + a.TableName + "` WHERE uid=?"
	var info []UserAddressInfo
	if err := a.Select(info, query, uid); err != nil {
		return nil, err
	}
	return info, nil
}

func (a *addressTable) InsertNewAddressInfo(addressInfo UserAddressInfo) error {
	m := make(map[string]interface{})

	m["aid"] = uuid.NewV4().String()
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

func (a *addressTable) UpdateAddressInfoById(addressInfo map[string]interface{}, id string) error {
	keys, values := _updateFiled(addressInfo)
	if _, err := a.UpdateById(keys, id, values); err != nil {
		return err
	}
	return nil
}

func (a *addressTable) DeleteAddressInfoById(id string) error {
	if err := a.DeleteById(id); err != nil {
		return err
	}
	return nil
}
