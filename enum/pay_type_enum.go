package enum

//支付方式

type PayType int

const (
	WeChatPay PayType = 1
	AliPay    PayType = 2
)

func (p PayType) String() string {
	switch p {
	case WeChatPay:
		return "微信"
	case AliPay:
		return "支付宝"
	default:
		return "unknown"
	}
}
