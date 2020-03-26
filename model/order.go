package model

import (
	"fmt"
	"math/rand"
	db "pet-paradise/model/common"
	"time"
)

type orderTable struct {
	db.Table
}

const ORDER_PREFIX_FORMAT = "20060102150405"

var OrderTable = &orderTable{db.Table{
	GetDB:     db.Conn,
	TableName: db.TBL_ORDER,
}}

type OrderInfo struct {
	ID         string `db:"id" json:"id"`
	OrderID    string `db:"oid" json:"oid"`
	UserID     string `db:"uid" json:"uid"`
	ProductID  string `db:"pid" json:"pid"`
	Status     string `db:"status" json:"status"`
	Address    string `db:"address" json:"address"`
	Details    string `db:"details" json:"details"`
	CreateTime string `db:"create_time" json:"create_time"`
	UpdateTime string `db:"update_time" json:"update_time"`
}

func (o *orderTable) GetAllByUserId(uid string) ([]string, error) {
	query := "SELECT id FROM `" + o.TableName + "` WHERE uid=? AND is_deleted=0"
	var ids []string
	if err := o.Select(ids, query, uid); err != nil {
		return nil, err
	}
	return ids, nil
}

func (o *orderTable) GetOneById(id string) (*OrderInfo, error) {
	query := "SELECT id, oid, uid, pid, status, address, details, create_time, update_time FROM `" + o.TableName + "` WHERE id=?"
	info := &OrderInfo{}
	if err := o.Get(info, query, id); err != nil {
		return nil, err
	}
	return info, nil
}

func (o *orderTable) InsertNewOrderInfo(orderInfo OrderInfo) error {
	m := make(map[string]interface{})
	m["oid"] = time.Now().Format(ORDER_PREFIX_FORMAT) + _generateRandomNumberInString()
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

func (a *addressTable) UpdateOrderInfoById(addressInfo map[string]interface{}, id string) error {
	keys, values := _updateFiled(addressInfo)
	if _, err := a.UpdateById(keys, id, values); err != nil {
		return err
	}
	return nil
}

func (a *addressTable) DeleteOrderInfoById(id string) error {
	if err := a.DeleteById(id); err != nil {
		return err
	}
	return nil
}

func _generateRandomNumberInString() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
