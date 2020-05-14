package model

import (
	"database/sql"
	"fmt"
	db "pet-paradise/model/common"
	"time"
)

type productTable struct {
	db.Table
}

var ProductTable = &productTable{db.Table{
	GetDB:     db.Conn,
	TableName: db.TBL_PRODUCT,
}}

type ProductInfo struct {
	ID                int    `db:"id" json:"id"`
	ProductName       string `db:"product_name" json:"product_name" form:"product_name"`
	ParentProductName string `db:"parent_product_name" json:"parent_product_name" form:"parent_product_name"`
	Price             string `db:"price" json:"price" form:"price"`
	Describe          string `db:"description" json:"description" form:"description"`
	Count             int    `db:"count_remain" json:"count" form:"count"`
	IsOnSale          string `db:"is_on_sale" json:"is_on_sale" form:"is_on_sale"`
	IsOnDiscount      string `db:"is_on_discount" json:"is_on_discount" form:"is_on_discount"`
	Details           string `db:"details" json:"details" form:"details"`
	CreateTime        string `db:"create_time" json:"create_time"`
	UpdateTime        string `db:"update_time" json:"update_time"`
}

func (p *productTable) SelectByParentProductName(parentProductName string) ([]ProductInfo, error) {
	query := "SELECT id, product_name, parent_product_name, price, description, count_remain, is_on_sale, is_on_discount, create_time, update_time FROM `" + p.TableName + "` WHERE is_deleted='0' AND is_on_sale IN(0,1) AND parent_product_name=?"
	var info []ProductInfo
	if err := p.Select(&info, query, parentProductName); err != nil {
		return nil, err
	}
	return info, nil
}

func (p *productTable) GetOneByName(productName string) (*ProductInfo, error) {
	return p.getOne("product_name", productName)
}

func (p *productTable) GetOneById(id string) (*ProductInfo, error) {
	return p.getOne("id", id)
}

func (p *productTable) getOne(key, value interface{}) (*ProductInfo, error) {
	query := "SELECT id, product_name, parent_product_name, price, description, count_remains, create_time, update_time FROM `" + p.TableName + "` WHERE is_deleted='0' AND %s=?"
	info := &ProductInfo{}
	if err := p.Get(info, fmt.Sprintf(query, key), value); err != nil {
		return nil, err
	}
	return info, nil
}

func (p *productTable) InsertNewProductInfo(productInfo ProductInfo) (sql.Result, error) {
	m := make(map[string]interface{})
	m["product_name"] = productInfo.ProductName
	m["parent_product_name"] = productInfo.ParentProductName
	m["price"] = productInfo.Price
	m["description"] = productInfo.Describe
	m["count_remain"] = productInfo.Count
	m["update_time"] = time.Now().Format(TIME_FORMAT)
	m["details"] = productInfo.Details
	return p.Insert(m)
}

func (p *productTable) AddProductCountById(id string, count int) (sql.Result, error) {
	if info, err := p.GetOneById(id); err != nil {
		return nil, err
	} else {
		count += info.Count
	}
	return p.UpdateProductInfoById(ProductInfo{
		Count: count,
	}, id)
}

func (p *productTable) UpdateProductInfoById(productInfo ProductInfo, id string) (sql.Result, error) {
	var productInfoMap = make(map[string]interface{})

	if productInfo.ProductName != "" {
		productInfoMap["product_name"] = productInfo.ProductName
	}
	if productInfo.ParentProductName != "" {
		productInfoMap["parent_product_name"] = productInfo.ParentProductName
	}
	if productInfo.Describe != "" {
		productInfoMap["description"] = productInfo.Describe
	}
	if productInfo.Price != "" {
		productInfoMap["price"] = productInfo.Price
	}
	if productInfo.Details != "" {
		productInfoMap["details"] = productInfo.Details
	}
	if productInfo.Count != 0 {
		productInfoMap["count_remain"] = productInfo.Count
	}
	if productInfo.IsOnDiscount != "" {
		if productInfo.IsOnDiscount != "0" && productInfo.IsOnDiscount != "1" {
			productInfo.IsOnDiscount = "0"
		}
		productInfoMap["is_on_discount"] = productInfo.IsOnDiscount
	}
	if productInfo.IsOnSale != "" {
		if productInfo.IsOnSale != "0" && productInfo.IsOnSale != "1" {
			productInfo.IsOnSale = "1"
		}
		productInfoMap["is_on_sale"] = productInfo.IsOnSale
	}

	keys, values := _updateFiled(productInfoMap)
	return p.UpdateById(keys, []string{id}, values...)
}

func (p *productTable) GetParentProduct() ([]string, error) {
	query := "SELECT DISTINCT(parent_product_name) FROM product WHERE is_deleted='0'"
	var parentProducts []string
	if err := p.Select(&parentProducts, query); err != nil {
		return nil, err
	}
	return parentProducts, nil
}

func (p *productTable) DeleteParentProduct(parentProductName string) (sql.Result, error) {
	query := "UPDATE product SET is_deleted='1' WHERE parent_product_name=?"
	return p.GetDB().Exec(query, parentProductName)
}

func (p *productTable) DeleteProductInfoById(id []string) (sql.Result, error) {
	return p.DeleteById(id)
}
