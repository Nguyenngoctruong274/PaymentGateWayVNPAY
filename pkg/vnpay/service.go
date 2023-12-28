package vnpay

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"net/url"
	"os"
	"time"
)

func paymentGateWay(paymentReq *PaymentGateWayReq) (string, error) {
	params := url.Values{}
	params.Set("vnp_Version", paymentReq.VNPVersion)
	params.Set("vnp_Command", paymentReq.VnpCommand)
	params.Set("vnp_TmnCode", paymentReq.VnpTmnCode)
	params.Set("vnp_Amount", fmt.Sprintf("%v", paymentReq.VnpAmount*100))
	params.Set("vnp_CreateDate", time.Now().Format("20060102150405"))
	params.Set("vnp_CurrCode", paymentReq.VnpCurrCode)

	//GetIp
	hostName, err := os.Hostname()
	if err != nil {
		return fmt.Sprintln("Failed to get hostname"), err
	}
	addresses, err := net.LookupIP(hostName)
	if err != nil {
		return fmt.Sprintln("Failed to Lookup IP"), err
	}
	paymentReq.VnpIpAddr = addresses[0].String()
	params.Set("vnp_IpAddr", paymentReq.VnpIpAddr)
	//

	params.Set("vnp_Locale", paymentReq.VnpLocale)
	params.Set("vnp_OrderInfo", paymentReq.VnpOrderInfo)
	params.Set("vnp_OrderType", paymentReq.VnpOrderType)
	params.Set("vnp_ReturnUrl", returnURL)
	paymentReq.VnpTxnRef = generateOrderID()
	params.Set("vnp_TxnRef", paymentReq.VnpTxnRef)
	if paymentReq.VnpBankCode != "" {
		params.Set("vnp_BankCode", paymentReq.VnpBankCode)
	}
	signData := params.Encode()
	h := hmac.New(sha512.New, []byte(secretKey))
	h.Write([]byte(signData))
	signature := hex.EncodeToString(h.Sum(nil))

	params.Set("vnp_SecureHash", signature)
	vnpUrl := paymentReq.VnpUrl
	vnpUrl += "?" + params.Encode()

	return vnpUrl, nil
}
func generateOrderID() string {
	timestamp := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s%d", timestamp, rand.Intn(10000))
}
