package enum

type UserAction int

const (
	userPay            UserAction = 1
	userCancel         UserAction = 2
	userRefund         UserAction = 3
	userConfirmReceipt UserAction = 4
	userApplyForReturn UserAction = 5
	userEvaluate       UserAction = 6
	userBuyAgain       UserAction = 7
	userDelete         UserAction = 8
)

func (u UserAction) String() string {
	switch u {
	case userPay:
		return "支付"
	case userCancel:
		return "取消"
	case userRefund:
		return "退款"
	case userConfirmReceipt:
		return "确认收货"
	case userApplyForReturn:
		return "申请退货"
	case userEvaluate:
		return "去评价"
	case userBuyAgain:
		return "再次购买"
	case userDelete:
		return "删除"
	default:
		return "unknown"
	}
}

/*
订单状态码标识了订单的状态，但是对于用户而言，真正关心的只是他们能够进行的操作， 也就是在小商场的小程序端用户可以进行点击的按钮操作，目前支持：
	支付，如果下单后未立即支付，则订单详情页面会出现支付按钮；
	取消，如果用户未支付，则订单详情页面会出现取消按钮；
	退款，如果用户支付后但是管理员未发货，则订单详情页面会出现退款按钮；
	确认收货，如果管理员已发货，则订单详情页面会出现确认收货按钮；
	申请退货，如果用户已经确认收货同时未超过一段时间，则订单详情页面会出现申请退货按钮；注意，这里如果是系统超时自动确认收货，则不会出现；
	去评价，如果用户确认收货以后，则订单详情页面会出现去评价按钮；
	再次购买，如果用户确认收货以后，则订单详情页面会出现再次购买按钮；
	删除，如果当前订单状态码是102、103、203、401和402时，则订单详情页面会出现删除订单按钮；注意，这里的删除操作是逻辑删除，即设置订单的删除状态deleted。
*/

/*
因此订单状态码和小商场用户操作之间存在映射关系：
	101
	用户可以支付、取消
	102
	用户可以删除
	103
	用户可以删除
	201
	用户可以退款
	203
	用户可以删除
	301
	用户可以确认收货
	401
	用户可以删除、去评价、申请售后、再次购买
	402
	用户可以删除、去评价、申请售后、再次购买
*/
