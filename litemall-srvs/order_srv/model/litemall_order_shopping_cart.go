package model

type ShoppingCart struct {
	BaseModel
	User    int32 `gorm:"type:int;index"` //在购物车列表中我们需要查询当前用户的购物车记录
	Goods   int32 `gorm:"type:int;index"` //加索引：我们需要查询时候， 1. 会影响插入性能 2. 会占用磁盘
	Nums    int32 `gorm:"type:int"`
	Checked bool  //是否选中
}
