package vnpay

const (
	stk              = "9704198526191432198" // NCB
	nameCT           = "NGUYEN VAN A"
	date             = "07/15"
	otp              = "123456"
	bank             = "NCB"
	merchantID       = "N1I4LANL"                          // Merchant ID của bạn
	merchantPassword = "Gnourt274"                         // Merchant Password của bạn
	secretKey        = "GTVLAWHMAVNVMIKZYDZJPKOAPUCZNHTW"  // Secure Secret của bạn
	returnURL        = "http://localhost:8081/returnOrder" // URL để chuyển hướng sau khi thanh toán thành công
)

var vnpString = []string{"vnp_Version", "vnp_Command", "vnp_TmnCode", "vnp_Amount",
	"vnp_CreateDate", "vnp_CurrCode", "vnp_IpAddr", "vnp_Locale", "vnp_OrderInfo",
	"vnp_OrderType", "vnp_ReturnUrl", "vnp_TxnRef", "vnp_BankCode", "vnp_SecureHash"}
