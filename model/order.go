package model

import (
	db "pet-paradise/model/common"
)

type orderTable struct {
	db.Table
}

var OrderTable = &orderTable{db.Table{
	GetDB:     db.Conn,
	TableName: db.TBL_ORDER,
}}

type OrderInfo struct {
	ID         int    `db:"id" json:"id"`
	UserID     int    `db:"uid" json:"uid"`
	ProductID  int    `db:"pid" json:"pid"`
	AddressID  int    `db:"aid" json:"aid"`
	Status     string `db:"status" json:"status"`
	Details    string `db:"details" json:"details"`
	CreateTime string `db:"create_time" json:"create_time"`
	UpdateTime string `db:"update_time" json:"update_time"`
}

func (o *orderTable) GetAllByUserId(uid int) ([]ProductInfo, error) {
	query := "SELECT id, uid, pid, aid, status, address, details, create_time, update_time FROM `" + o.TableName + "` WHERE uid=? AND is_deleted=0"
	var infoSlice []ProductInfo
	if err := o.Select(&infoSlice, query, uid); err != nil {
		return nil, err
	}
	return infoSlice, nil
}

func (o *orderTable) GetOneById(id int) (*OrderInfo, error) {
	query := "SELECT id, uid, pid, aid, status, details, create_time, update_time FROM `" + o.TableName + "` WHERE is_deleted='0' AND id=?"
	info := &OrderInfo{}
	if err := o.Get(info, query, id); err != nil {
		return nil, err
	}
	return info, nil
}

func (o *orderTable) InsertNewOrderInfo(orderInfo OrderInfo) error {
	m := make(map[string]interface{})
	m["aid"] = orderInfo.Address
	m["uid"] = orderInfo.UserID
	m["pid"] = orderInfo.ProductID
	m["status"] = orderInfo.Status
	m["address"] = orderInfo.Address
	m["details"] = orderInfo.Details
	if _, err := o.Insert(m); err != nil {
		return err
	}
	return nil
}

func (a *addressTable) UpdateOrderInfoById(addressInfo map[string]interface{}, id int) error {
	keys, values := _updateFiled(addressInfo)
	if _, err := a.UpdateById(keys, id, values...); err != nil {
		return err
	}
	return nil
}

func (a *addressTable) DeleteOrderInfoById(id int) error {
	if err := a.DeleteById(id); err != nil {
		return err
	}
	return nil
}
