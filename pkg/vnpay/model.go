package vnpay

type PaymentGateWayReq struct {
	VNPVersion    string `json:"vnp_Version" validate:"required,alphanum,min=1,max=8"`
	VnpCommand    string `json:"vnp_Command" validate:"required,alpha,min=1,max=16"`         //Mã API sử dụng, mã cho giao dịch thanh toán là: pay
	VnpTmnCode    string `json:"vnp_TmnCode" validate:"required,alphanumeric8"`              //Mã website của merchant trên hệ thống của VNPAY. Ví dụ: 2QXUI4J4
	VnpAmount     int    `json:"vnp_Amount" validate:"required,min=1,max=12"`                //	Số tiền thanh toán. Số tiền không mang các ký tự phân tách thập phân, phần nghìn, ký tự tiền tệ. Để gửi số tiền thanh toán là 10,000 VND (mười nghìn VNĐ) thì merchant cần nhân thêm 100 lần (khử phần thập phân), sau đó gửi sang VNPAY là: 1000000
	VnpBankCode   string `json:"vnp_BankCode" validate:"alphanum,min=3,max=20"`              //	Mã Ngân hàng thanh toán. Ví dụ: NCB
	VnpCreateDate int    `json:"vnp_CreateDate"`                                             //validate:"required,number,min=14,max=14"` //Là thời gian phát sinh giao dịch định dạng yyyyMMddHHmmss(Time zone GMT+7)Ví dụ: 20170829103111
	VnpCurrCode   string `json:"vnp_CurrCode" validate:"required,alpha,len=3"`               //	Đơn vị tiền tệ sử dụng thanh toán. Hiện tại chỉ hỗ trợ VND
	VnpIpAddr     string `json:"vnp_IpAddr" validate:"required,alphanum,min=7,max=45"`       //Địa chỉ IP của khách hàng thực hiện giao dịch. Ví dụ: 13.160.92.202
	VnpLocale     string `json:"vnp_Locale" validate:"required,alpha,min=2,max=5"`           //Ngôn ngữ giao diện hiển thị. Hiện tại hỗ trợ Tiếng Việt (vn), Tiếng Anh (en)
	VnpOrderInfo  string `json:"vnp_OrderInfo" validate:"required,alphanum,min=1,max=255"`   //Thông tin mô tả nội dung thanh toán (Tiếng Việt, không dấu). Ví dụ: **Nap tien cho thue bao 0123456789. So tien 100,000 VND**
	VnpOrderType  string `json:"vnp_OrderType" validate:"alpha,min=1,max=100"`               //	Mã danh mục hàng hóa. Mỗi hàng hóa sẽ thuộc một nhóm danh mục do VNPAY quy định. Xem thêm bảng Danh mục hàng hóa
	VnpReturnUrl  string `json:"vnp_ReturnUrl" validate:"required,alphanum,min=10,max=255"`  //URL thông báo kết quả giao dịch khi Khách hàng kết thúc thanh toán. Ví dụ: https://domain.vn/VnPayReturn
	VnpTxnRef     string `json:"vnp_TxnRef"`                                                 // validate:"required,alphanum,min=1,max=100"`
	VnpSecureHash string `json:"vnp_SecureHash" validate:"required,alphanum,min=32,max=256"` //Mã kiểm tra (checksum) để đảm bảo dữ liệu của giao dịch không bị thay đổi trong quá trình chuyển từ merchant sang VNPAY. Việc tạo ra mã này phụ thuộc vào cấu hình của merchant và phiên bản api sử dụng. Phiên bản hiện tại hỗ trợ SHA256, HMACSHA512.
	VnpCancelURL  string `json:"vnp_cancel_url"`
	VnpUrl        string `json:"vnp_url"`
}

type PaymentGateWayResp struct {
	VnpTmnCode           string `json:"vnp_TmnCode" `
	VnpAmount            int    `json:"vnp_Amount" `
	VnpBankCode          string `json:"vnp_BankCode" `
	VnpBankTranNo        string `json:"vnp_BankTranNo"`
	VnpCardType          string `json:"vnp_CardType"`
	VnpPayDate           string `json:"vnp_PayDate"`
	VnpCreateDate        int    `json:"vnp_CreateDate"`
	VnpCurrCode          string `json:"vnp_CurrCode" `
	VnpIpAddr            string `json:"vnp_IpAddr" `
	VnpLocale            string `json:"vnp_Locale" `
	VnpOrderInfo         string `json:"vnp_OrderInfo" `
	VnpTranNo            int    `json:"vnp_TransactionNo"`
	VnpResponseCode      int    `json:"vnp_ResponseCode"`
	VnpTransactionStatus int    `json:"vnp_TransactionStatus"`
	VnpSecureHashType    string `json:"vnp_SecureHashType"`
	VnpTxnRef            string `json:"vnp_TxnRef"`
	VnpSecureHash        string `json:"vnp_SecureHash"`
}
