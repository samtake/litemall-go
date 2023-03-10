
## 说明：

### 商品表：goods、brand、category、
    是一种商品的基本信息，主要包括商品介绍，商品图片，商品所属类目，商品品牌商等；
### 商品参数（商品属性）：litemall_goods_attribute
    商品参数表其实也是商品的基本信息，但是由于是一对多关系，因此不能直接保存在商品表中（虽然采用JSON也可以但是不合理）， 因此采用独立的商品参数表，通常是商品的一些公共基本商品参数；
### 商品规格表 ：litemall_goods_specification
    商品规格表是商品进一步区分货品的标识，例如同样一款衣服，基本信息一致，基本属性一致，但是在尺寸这个属性上可以 把衣服区分成多个货品，而且造成对应的数量和价格不一致。商品规格可以看着是商品属性，但具有特殊特征。
商品规格和规格值存在以下几种关系：
- 单一规格和单一规格值，最常见的，即当前商品存在一种货品；
- 单一规格和多个规格值，较常见，即当前商品基于某个规格存在多种货品，通常价格都是相同的，当然也可能不相同；
- 多个规格和单一规格值，可以简化成第一种情况，或者采用第四种情况，通常实际情况下不常见；
- 多个规格和多个规格值，通常是两种规格或者三种规格较为常见，而且对应的价格不完全相同。
### 商品货品表：litemall_goods_product
    商品货品表则是最终实现商品库存管理、购买业务的实体对象，存在多个规格值、数量和价格。 例如，同样的衣服品牌，可能因为不能尺寸和颜色而存在最终的货品，这里每个货品的价格可以一样，也可以不一样。

总结一下，一个普通商品，实际上在数据库中，存在一个商品表项，存在（至少0个）多个商品属性表项目，存在（至少一个）多个商品规格表项， 存在（至少一个）多个货品表项。
举例如下：
- 一个商品“2018春季衣服商品编号1111111”，
- 存在两个商品参数， 
  - 属性名称“面向人群”，属性值“男士”
  - 属性名称“面料”，属性值“100%棉”
- 存在两种规格（分别五个规格值和三个规格值）共八个商品规格项，
  - 规格名称“尺寸”，规则值“S”
  - 规格名称“尺寸”，规则值“M”
  - 规格名称“尺寸”，规则值“L”
  - 规格名称“尺寸”，规则值“XL”
  - 规格名称“尺寸”，规则值“XXL”
  - 规格名称“颜色”，规格值“蓝色”
  - 规格名称“颜色”，规格值“灰色”
  - 规格名称“颜色”，规格值“黑色” 
- 存在15个货品（尺寸*颜色=15个货品)
  - 货品“S蓝”，数量 100， 价格 100
  - 货品“M蓝”，数量 100， 价格 100
  - 货品“L蓝”，数量 100， 价格 100
  - 货品“XL蓝”，数量 100， 价格 100
  - 货品“XXL蓝”，数量 100， 价格 100
  - 货品“S灰”，数量 100， 价格 100
  - 货品“M灰”，数量 100， 价格 100
  - 货品“L灰”，数量 100， 价格 100
  - 货品“XL灰”，数量 100， 价格 100
  - 货品“XXL灰”，数量 100， 价格 100
  - 货品“S黑”，数量 100， 价格 100
  - 货品“M黑”，数量 100， 价格 100
  - 货品“L黑”，数量 100， 价格 100
  - 货品“XL黑”，数量 0， 价格 100
  - 货品“XXL黑”，数量 0， 价格 100
### 以下是一些细节的讨论：
- 商品表中可能存在数量和价格属性，而货品中也存在数量和价格属性，目前设计这样：
  - 商品表的价格应该和某个货品的价格一样，通常应该是所有货品价格的最小值，或者基本款式的价格；
  - 商品表中的数量和价格应该仅用于展示，而不能用于最终的订单价格计算；
  - 商品表的数量应该设置成所有货品数量的总和；
  - 在管理后台添加商品时，如果管理员不填写商品表的数量和价格属性，则自动填写合适的值；如果填写，则使用显示。
  - 当小商城中，用户查看商品详情时，初始显示商品表的价格，而如果用户选择具体规格后，则商品
    详情里面的价格需要自动切换到该规格的价格。
- 商品规格可以存在规格图片，效果是规格名称前放置规格图片
- 货品也可以存在货品图片，效果是所有规格选定以后对应的货品有货，则在货品价格前放置货品图片
- 如果商品是两种规格，分别是M个和N个规格值，那么通常应该是M*N个货品，但是有些货品可能天然不存在。
- 目前这里要求所有货品信息都应该存在，如果实际中货品不存在，也要设置商品数量为0.

### 注意：
- 这里的设计可能与实际项目设计不一致，但是目前是可行的。 商品的中文用语“商品”和英语用语“goods”，货品的中文用语“货品”和英语用语“product”可能是不正确的。