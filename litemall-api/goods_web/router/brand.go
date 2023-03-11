package router

import (
	"github.com/gin-gonic/gin"
	"litemall-api/goods_web/api/brand"
	"litemall-api/goods_web/middlewares"
)

// 1. 商品的api接口开发完成
// 2. 图片的坑
func InitBrandRouter(Router *gin.RouterGroup) {
	BrandRouter := Router.Group("brands").Use(middlewares.Trace())
	{
		BrandRouter.GET("", brand.BrandList)          // 品牌列表页
		BrandRouter.DELETE("/:id", brand.DeleteBrand) // 删除品牌
		BrandRouter.POST("", brand.NewBrand)          //新建品牌
		BrandRouter.PUT("/:id", brand.UpdateBrand)    //修改品牌信息
	}

	CategoryBrandRouter := Router.Group("categorybrands")
	{
		CategoryBrandRouter.GET("", brand.CategoryBrandList)          // 类别品牌列表页
		CategoryBrandRouter.DELETE("/:id", brand.DeleteCategoryBrand) // 删除类别品牌
		CategoryBrandRouter.POST("", brand.NewCategoryBrand)          //新建类别品牌
		CategoryBrandRouter.PUT("/:id", brand.UpdateCategoryBrand)    //修改类别品牌
		CategoryBrandRouter.GET("/:id", brand.GetCategoryBrandList)   //获取分类的品牌
	}
}
