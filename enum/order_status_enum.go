package enum

type OrderStatus int

/*
	订单分成几种基本的状态：
	101
	状态码101，此时订单生成，记录订单编号、收货地址信息、订单商品信息和订单相关费用信息；
	201
	状态码201，此时用户微信支付付款，系统记录微信支付订单号、支付时间、支付状态；
	301
	状态码301，此时商场已经发货，系统记录快递公司、快递单号、快递发送时间。 当快递公司反馈用户签收后，系统记录快递到达时间。
	401
	状态码401，当用户收到货以后点击确认收货，系统记录确认时间。
	以上是一个订单成功完成的基本流程，但实际中还存在其他情况。
	102
	状态码102，用户下单后未付款之前，点击取消按钮，系统记录结束时间
	103
	状态码103，用户下单后半小时未付款则系统自动取消，系统记录结束时间
	202
	状态码202，用户付款以后未发货前，点击退款按钮，系统进行设置退款状态，等待管理员退款操作
	203
	状态码203，管理员在管理后台看到用户的退款申请，点击退款按钮进行退款操作。
	402
	状态码402，用户已签收却不点击确认收货，超期7天以后，则系统自动确认收货。 用户不能再点击确认收货按钮，但是可以评价订单商品。
*/
const (
	OrderCreat                OrderStatus = 101
	OrderUserCancel           OrderStatus = 102
	OrderOvertimeSystemCancel OrderStatus = 103
	OrderPay                  OrderStatus = 201
	OrderRefundUnShipped      OrderStatus = 202
	OrderRefundUnAdmin        OrderStatus = 203
	OrderExpressStart         OrderStatus = 301
	OrderExpressEndUser       OrderStatus = 401
	OrderExpressEndSystem     OrderStatus = 402
)

func (o OrderStatus) String() string {
	switch o {
	case OrderCreat:
		return "订单生成"
	case OrderUserCancel:
		return "用户取消订单"
	case OrderOvertimeSystemCancel:
		return "超时取消订单"
	case OrderPay:
		return "支付成功"
	case OrderRefundUnShipped:
		return "未发货取消订单退款"
	case OrderRefundUnAdmin:
		return "已发货，管理员审核退款"
	case OrderExpressStart:
		return "已发货"
	case OrderExpressEndUser:
		return "用户确认收货"
	case OrderExpressEndSystem:
		return "系统确认收货"
	default:
		return "unknown"
	}
}
