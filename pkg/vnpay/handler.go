package vnpay

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (v *PaymentGateWayReq) Handler(router *gin.Engine) {
	router.LoadHTMLGlob("../../web/templates/*.html")
	router.GET("/getPageOrder", getPageOrder)
	router.POST("/createOrder", v.createOrder)
	router.GET("/returnOrder", v.returnOrder)

}

func getPageOrder(c *gin.Context) {
	c.HTML(http.StatusOK, "pageOrder.html", gin.H{
		"title": "Dashboard"})
}

func (v *PaymentGateWayReq) createOrder(c *gin.Context) {
	var url string

	v.VnpOrderType = c.PostForm("ordertype")
	amount, err := strconv.Atoi(c.PostForm("amount"))
	if err != nil {
		c.Redirect(http.StatusNotFound, url)
	}
	v.VnpAmount = amount
	v.VnpOrderInfo = c.PostForm("orderDescription")
	v.VnpBankCode = c.PostForm("bankcode")
	v.VnpLocale = c.PostForm("language")
	submit := c.PostForm("submit")
	if strings.EqualFold(submit, "submit") {
		url, err = paymentGateWay(v)
		if err != nil {
			c.Redirect(http.StatusNotFound, url)
		}
	}
	log.Println(url)
	c.Redirect(http.StatusFound, url)
}

func NewVNPay(currCode, returnURL, cancelURL, version, command, tmnCode, vnpUrl string) *PaymentGateWayReq {
	return &PaymentGateWayReq{
		VNPVersion:   version,
		VnpCommand:   command,
		VnpTmnCode:   tmnCode,
		VnpCurrCode:  currCode,
		VnpReturnUrl: returnURL,
		VnpCancelURL: cancelURL,
		VnpUrl:       vnpUrl,
	}
}

func (v *PaymentGateWayReq) returnOrder(c *gin.Context) {

	respPayment := &PaymentGateWayResp{}

	//vnp_TmnCode
	respPayment.VnpTmnCode = c.Query("vnp_TmnCode")
	//vnp_Amount
	amount, err := strconv.Atoi(c.Query("vnp_Amount"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "VnPayReturn.html", gin.H{
			"title":       "Dashboard",
			"respPayment": respPayment,
		})
	}

	respPayment.VnpAmount = amount
	respPayment.VnpBankCode = c.Query("vnp_BankCode")
	respPayment.VnpBankTranNo = c.Query("vnp_BankTranNo")
	respPayment.VnpCardType = c.Query("vnp_CardType")
	respPayment.VnpCardType = c.Query("vnp_PayDate")
	respPayment.VnpOrderInfo = c.Query("vnp_OrderInfo")
	tranNo, _ := strconv.Atoi(c.Query("vnp_TransactionNo"))
	respPayment.VnpTranNo = tranNo
	respCode, _ := strconv.Atoi(c.Query("vnp_ResponseCode"))
	respPayment.VnpResponseCode = respCode

	transactionStatus, _ := strconv.Atoi(c.Query("vnp_TransactionStatus")) //option
	respPayment.VnpTransactionStatus = transactionStatus

	respPayment.VnpTxnRef = c.Query("vnp_TxnRef")
	log.Println("/n ==========vnp_TxnRef=========== ", respPayment.VnpTxnRef)
	// respPayment.VnpSecureHash = c.Query("vnp_SecureHash")
	respPayment.VnpSecureHashType = c.Query("vnp_SecureHashType")
	c.HTML(http.StatusOK, "vnpayReturn.html", gin.H{
		"title":       "Dashboard",
		"respPayment": respPayment,
	})

}
