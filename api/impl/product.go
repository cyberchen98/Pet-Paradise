package impl

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"pet-paradise/log"
	"pet-paradise/model"
	"pet-paradise/utils"
)

func GetParentProduct(ctx *gin.Context) {
	log.Logger().Info("[GetParentProduct] %s", ctx.ClientIP())

	parentProduct, err := model.ProductTable.GetParentProduct()
	if err == sql.ErrNoRows {
		utils.Success(ctx, "none", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", parentProduct)
}

func GetProductInfoByParentName(ctx *gin.Context) {
	log.Logger().Info("[GetProductInfoByParentName] %s", ctx.ClientIP())

	parentProductName := ctx.Param("parentProductName")

	productInfo, err := model.ProductTable.SelectByParentProductName(parentProductName)
	if err == sql.ErrNoRows {
		utils.Success(ctx, "none", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}
	utils.Success(ctx, "ok", productInfo)
}

func GetProductInfoByName(ctx *gin.Context) {
	log.Logger().Info("[GetProductInfoByName] %s", ctx.ClientIP())

	productName := ctx.Query("product_name")
	productInfo, err := model.ProductTable.GetOneByName(productName)
	if err == sql.ErrNoRows {
		utils.Success(ctx, "no such product", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", productInfo)
}

func AddNewProduct(ctx *gin.Context) {
	log.Logger().Info("[AddNewProduct] %s", ctx.ClientIP())

	var productInfo model.ProductInfo
	if err := ctx.Bind(&productInfo); err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	if _, err := model.ProductTable.InsertNewProductInfo(productInfo); err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)
}

func AddNewParentProduct(ctx *gin.Context) {
	log.Logger().Info("[AddNewParentProduct] %s", ctx.ClientIP())

	newParentProduct := ctx.PostForm("parent_product_name")

	if _, err := model.ProductTable.Insert(map[string]interface{}{"parent_product_name": newParentProduct, "is_on_sale": "-1"}); err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)
}

func DeleteProduct(ctx *gin.Context) {
	log.Logger().Info("[DeleteProduct] %s", ctx.ClientIP())

	productIDs := ctx.PostFormArray("pid")

	if _, err := model.ProductTable.DeleteProductInfoById(productIDs); err == sql.ErrNoRows {
		utils.Fail(ctx, "no this record", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)
}

func UpdateProductInfo(ctx *gin.Context) {
	log.Logger().Info("[UpdateProductInfo] %s", ctx.ClientIP())

	productID := ctx.PostForm("pid")

	var productInfo model.ProductInfo
	if err := ctx.Bind(&productInfo); err != nil {
		utils.Fail(ctx, "invalid params", nil)
		return
	}

	if _, err := model.ProductTable.UpdateProductInfoById(productInfo, productID); err == sql.ErrNoRows {
		utils.Fail(ctx, "no this record", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)
}

func AdminDeleteParentProduct(ctx *gin.Context) {
	log.Logger().Info("[AdminDeleteParentProduct] %s", ctx.ClientIP())

	parentProductName := ctx.Query("parent_product_name")

	if _, err := model.ProductTable.DeleteParentProduct(parentProductName); err == sql.ErrNoRows {
		utils.Fail(ctx, "no this record", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", parentProductName)
}
