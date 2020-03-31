package model

import (
	"fmt"
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
	ID                int    `db:"id" json:"id"`
	ProductName       string `db:"product_name" json:"name"`
	ParentProductName string `db:"parent_product_name" json:"parent_product_name"`
	Price             string `db:"price" json:"price"`
	Describe          string `db:"describe" json:"describe"`
	Count             int    `db:"count_remains" json:"count"`
	IsOnSale          string `db:"-" json:"is_on_sale"`
	IsOnDiscount      string `db:"-" json:"is_on_discount"`
	Details           string `db:"details" json:"details"`
	CreateTime        string `db:"create_time" json:"create_time"`
	UpdateTime        string `db:"update_time" json:"update_time"`
}

func (p *productTable) GetOneByName(productName string) (*ProductInfo, error) {
	return p.getOne("product_name", productName)
}

func (p *productTable) GetOneById(id int) (*ProductInfo, error) {
	return p.getOne("id", id)
}

func (p *productTable) getOne(key, value interface{}) (*ProductInfo, error) {
	query := "SELECT id, product_name, parent_product_name, price, describe, count_remains, create_time, update_time FROM `" + p.TableName + "` WHERE is_deleted='0' AND %s=?"
	info := &ProductInfo{}
	if err := p.Get(info, fmt.Sprintf(query, key), value); err != nil {
		return nil, err
	}
	return info, nil
}

func (p *productTable) InsertNewProductInfo(productInfo ProductInfo) error {
	m := make(map[string]interface{})
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

func (p *productTable) AddProductCountById(id int, count int) error {
	if info, err := p.GetOneById(id); err != nil {
		return err
	} else {
		count += info.Count
	}
	if err := p.UpdateProductInfoById(ProductInfo{
		Count: count,
	}, id); err != nil {
		return err
	}
	return nil
}

func (p *productTable) UpdateProductInfoById(productInfo ProductInfo, id int) error {
	var productInfoMap = make(map[string]interface{})

	if productInfo.ProductName != "" {
		productInfoMap["product_name"] = productInfo.ProductName
	}
	if productInfo.ParentProductName != "" {
		productInfoMap["parent_product_name"] = productInfo.ParentProductName
	}
	if productInfo.Describe != "" {
		productInfoMap["describe"] = productInfo.Describe
	}
	if productInfo.Price != "" {
		productInfoMap["price"] = productInfo.Price
	}
	if productInfo.Details != "" {
		productInfoMap["details"] = productInfo.Details
	}
	if productInfo.Count != 0 {
		productInfoMap["count_remains"] = productInfo.Count
	}
	if productInfo.IsOnDiscount != "" {
		productInfoMap["is_on_discount"] = productInfo.IsOnDiscount
	}
	if productInfo.IsOnSale != "" {
		productInfoMap["is_on_sale"] = productInfo.IsOnSale
	}

	keys, values := _updateFiled(productInfoMap)
	if _, err := p.UpdateById(keys, id, values...); err != nil {
		return err
	}
	return nil
}

func (p *productTable) DeleteProductInfoById(id int) error {
	if err := p.DeleteById(id); err != nil {
		return err
	}
	return nil
}
