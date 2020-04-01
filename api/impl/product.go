package impl

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"pet-paradise/log"
	"pet-paradise/model"
	"pet-paradise/utils"
	"strconv"
)

func GetProductInfoByParentName(ctx *gin.Context) {
	log.Logger().Info("[GetProductInfoByParentName] ", ctx.Request.URL)

	parentProductName := ctx.Query("parent_product_name")

	productInfo, err := model.ProductTable.SelectByParentProductName(parentProductName)
	if err == sql.ErrNoRows {
		utils.Fail(ctx, "none", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", productInfo)
}

func GetProductInfoByName(ctx *gin.Context) {
	log.Logger().Info("[GetProductInfoByName] ", ctx.Request.URL)

	productName := ctx.Query("product_name")
	productInfo, err := model.ProductTable.GetOneByName(productName)
	if err == sql.ErrNoRows {
		utils.Fail(ctx, "no such product", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", productInfo)
}

func AddNewProduct(ctx *gin.Context) {
	log.Logger().Info("[AddNewProduct] ", ctx.Request.URL)

	var productInfo model.ProductInfo
	if err := ctx.Bind(&productInfo); err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	if err := model.ProductTable.InsertNewProductInfo(productInfo); err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)
}

func DeleteProduct(ctx *gin.Context) {
	log.Logger().Info("[DeleteProduct] ", ctx.Request.URL)

	productID, err := strconv.Atoi(ctx.Query("pid"))
	if err != nil {
		utils.Fail(ctx, "invalid param: product_id", nil)
		return
	}

	if err := model.ProductTable.DeleteProductInfoById(productID); err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)
}

func UpdateProductInfo(ctx *gin.Context) {
	log.Logger().Info("[UpdateProductInfo] ", ctx.Request.URL)

	productID, err := strconv.Atoi(ctx.Query("pid"))
	if err != nil {
		utils.Fail(ctx, "invalid param: product_id", nil)
		return
	}

	var productInfo model.ProductInfo
	if err := ctx.Bind(&productInfo); err != nil {
		utils.Fail(ctx, "invalid params", nil)
		return
	}

	if err := model.ProductTable.UpdateProductInfoById(productInfo, productID); err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)
}
