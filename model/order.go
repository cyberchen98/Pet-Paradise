package model

import (
	"database/sql"
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
	ProductID  int    `db:"pid" json:"pid" form:"pid"`
	AddressID  int    `db:"aid" json:"aid" form:"aid"`
	Status     string `db:"status" json:"status" form:"status"`
	Details    string `db:"details" json:"details" form:"details"`
	CreateTime string `db:"create_time" json:"create_time"`
	UpdateTime string `db:"update_time" json:"update_time"`
}

type OrderInfoDetail struct {
	OrderInfo
	Address UserAddressInfo `json:"address"`
	Product ProductInfo     `json:"product"`
}

func (o *orderTable) SelectOrderInfoByUserId(uid string) ([]OrderInfo, error) {
	return o.selectOrderInfo(uid, "uid")

}

func (o *orderTable) SelectOrderInfoByProductId(pid string) ([]OrderInfo, error) {
	return o.selectOrderInfo(pid, "pid")
}

func (o *orderTable) selectOrderInfo(id, info string) ([]OrderInfo, error) {
	query := "SELECT id, uid, pid, aid, status, details, create_time, update_time FROM `" + o.TableName + "` WHERE " + info + "=? AND is_deleted='0'"
	var infoSlice []OrderInfo
	if err := o.Select(&infoSlice, query, id); err != nil {
		return nil, err
	}
	return infoSlice, nil
}

func (o *orderTable) InsertNewOrderInfo(orderInfo OrderInfo) (sql.Result, error) {
	m := make(map[string]interface{})
	m["aid"] = orderInfo.AddressID
	m["uid"] = orderInfo.UserID
	m["pid"] = orderInfo.ProductID
	m["status"] = orderInfo.Status
	m["details"] = orderInfo.Details
	return o.Insert(m)
}

func (o *orderTable) UpdateOrderInfoById(orderInfo OrderInfo, id string) (sql.Result, error) {
	var orderInfoMap = make(map[string]interface{})

	if orderInfo.AddressID != 0 {
		orderInfoMap["aid"] = orderInfo.AddressID
	}
	if orderInfo.Details != "" {
		orderInfoMap["details"] = orderInfo.Details
	}
	if orderInfo.Status != "" {
		orderInfoMap["status"] = orderInfo.Status
	}

	keys, values := _updateFiled(orderInfoMap)
	return o.UpdateById(keys, []string{id}, values...)
}

func (o *orderTable) DeleteOrderInfoById(id []string) (sql.Result, error) {
	return o.DeleteById(id)
}
