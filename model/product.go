package model

import (
	"fmt"
	"github.com/satori/go.uuid"
	db "pet-paradise/model/common"
	"time"
)

type productTable struct {
	db.Table
}

var ProductTable = &userTable{db.Table{
	GetDB:     db.Conn,
	TableName: db.TBL_PRODUCT,
}}

type ProductInfo struct {
	ID                string `db:"id" json:"id"`
	ProductID         string `db:"pid" json:"pid"`
	ProductName       string `db:"product_name" json:"name"`
	ParentProductName string `db:"parent_product_name" json:"parent_product_name"`
	Price             string `db:"price" json:"price"`
	Count             int    `db:"count_remains" json:"count"`
	IsOnSale          bool   `db:"-" json:"is_on_sale"`
	IsOnDiscount      bool   `db:"-" json:"is_on_discount"`
	Details           string `db:"details" json:"details"`
	CreateTime        string `db:"create_time" json:"create_time"`
	UpdateTime        string `db:"update_time" json:"update_time"`
}

func (p *productTable) GetAllProductIds(whereCause string) ([]string, error) {
	query := "SELECT id FROM `" + p.TableName + "` WHERE is_deleted=0" + whereCause
	var ids []string
	if err := p.Select(ids, query); err != nil {
		return nil, err
	}
	return ids, nil
}

func (p *productTable) GetOneByName(productName string) (*ProductInfo, error) {
	return p.getOne("product_name", productName)
}

func (p *productTable) GetOneById(id string) (*ProductInfo, error) {
	return p.getOne("id", id)
}

func (p *productTable) getOne(key, value string) (*ProductInfo, error) {
	query := "SELECT id, pid, product_name, parent_product_name, price, count_remains, create_time, update_time FROM `" + p.TableName + "` WHERE %s=?"
	info := &ProductInfo{}
	if err := p.Get(info, fmt.Sprintf(query, key), value); err != nil {
		return nil, err
	}
	return info, nil
}

func (p *productTable) InsertNewProductInfo(productInfo ProductInfo) error {
	m := make(map[string]interface{})
	m["pid"] = uuid.NewV4().String()
	m["product_name"] = productInfo.ProductName
	m["parent_product_name"] = productInfo.ParentProductName
	m["price"] = productInfo.Price
	m["count"] = productInfo.Count
	m["update_time"] = time.Now().Format(TIME_FORMAT)
	m["details"] = productInfo.Details
	if _, err := p.Insert(m); err != nil {
		return err
	}
	return nil
}

func (p *productTable) AddProductCountById(id string, count int) error {
	if err := p.UpdateProductInfoById(map[string]interface{}{
		"count": count,
	}, id); err != nil {
		return err
	}
	return nil
}

func (p *productTable) UpdateProductInfoById(productInfo map[string]interface{}, id string) error {
	keys, values := _updateFiled(productInfo)
	if _, err := p.UpdateById(keys, id, values); err != nil {
		return err
	}
	return nil
}

func (p *productTable) DeleteProductInfoById(id string) error {
	if err := p.DeleteById(id); err != nil {
		return err
	}
	return nil
}
