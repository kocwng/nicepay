package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CartItem struct {
	ImgURL      string `json:"img_url"`
	GoodsName   string `json:"goods_name"`
	GoodsDetail string `json:"goods_detail"`
	GoodsAmt    string `json:"goods_amt"`
}

type RegistrationRequest struct {
	TimeStamp    string     `json:"timeStamp"`
	IMid         string     `json:"iMid"`
	PayMethod    string     `json:"payMethod"`
	Currency     string     `json:"currency"`
	Amt          string     `json:"amt"`
	ReferenceNo  string     `json:"referenceNo"`
	GoodsNm      string     `json:"goodsNm"`
	BillingNm    string     `json:"billingNm"`
	BillingPhone string     `json:"billingPhone"`
	BillingEmail string     `json:"billingEmail"`
	BillingAddr  string     `json:"billingAddr"`
	BillingCity  string     `json:"billingCity"`
	BillingState string     `json:"billingState"`
	BillingPostCd string    `json:"billingPostCd"`
	BillingCountry string  `json:"billingCountry"`
	CartData     []CartItem `json:"cartData"`
}

type RegistrationResponse struct {
	ResultCd       string `json:"resultCd"`
	ResultMsg      string `json:"resultMsg"`
	TXid           string `json:"tXid"`
	ReferenceNo    string `json:"referenceNo"`
	PayMethod      string `json:"payMethod"`
	Amt            string `json:"amt"`
	TransDt        string `json:"transDt"`
	TransTm        string `json:"transTm"`
	BankCd         string `json:"bankCd"`
	VacctNo        string `json:"vacctNo"`
	VacctValidDt   string `json:"vacctValidDt"`
	VacctValidTm   string `json:"vacctValidTm"`
}

type PaymentRequest struct {
	TimeStamp   string `json:"timeStamp"`
	TXid        string `json:"tXid"`
	IMid        string `json:"iMid"`
	MerchantToken string `json:"merchantToken"`
}

type PaymentResponse struct {
	ResultCd string `json:"resultCd"`
	ResultMsg string `json:"resultMsg"`
	TXid     string `json:"tXid"`
	ReferenceNo string `json:"referenceNo"`
	PayMethod string `json:"payMethod"`
	Amt string `json:"amt"`
	Currency string `json:"currency"`
	GoodsNm string `json:"goodsNm"`
	BillingNm string `json:"billingNm"`
	ReqDt string `json:"reqDt"`
	ReqTm string `json:"reqTm"`
	Status string `json:"status"`
	VacctValidDt string `json:"vacctValidDt"`
	VacctValidTm string `json:"vacctValidTm"`
	VacctNo string `json:"vacctNo"`
	BankCd string `json:"bankCd"`
}

type InquiryRequest struct {
	TimeStamp    string `json:"timeStamp"`
	IMid         string `json:"iMid"`
	TXid         string `json:"tXid"`
	Amt          string `json:"amt"`
	ReferenceNo  string `json:"referenceNo"`
	MerchantToken string `json:"merchantToken"`
}

type InquiryResponse struct {
	TXid           string `json:"tXid"`
	IMid           string `json:"iMid"`
	Currency       string `json:"currency"`
	Amt            string `json:"amt"`
	InstmntMon     string `json:"instmntMon"`
	InstmntType    string `json:"instmntType"`
	ReferenceNo    string `json:"referenceNo"`
	GoodsNm        string `json:"goodsNm"`
	PayMethod      string `json:"payMethod"`
	BillingNm      string `json:"billingNm"`
	ReqDt          string `json:"reqDt"`
	ReqTm          string `json:"reqTm"`
	Status         string `json:"status"`
	ResultCd       string `json:"resultCd"`
	ResultMsg      string `json:"resultMsg"`
	CardNo         string `json:"cardNo"`
	PreauthToken   string `json:"preauthToken"`
	AcquBankCd     string `json:"acquBankCd"`
	IssuBankCd     string `json:"issuBankCd"`
	VacctValidDt   string `json:"vacctValidDt"`
	VacctValidTm   string `json:"vacctValidTm"`
	VacctNo        string `json:"vacctNo"`
	BankCd         string `json:"bankCd"`
	PayNo          string `json:"payNo"`
	MitraCd        string `json:"mitraCd"`
	ReceiptCode    string `json:"receiptCode"`
	CancelAmt      string `json:"cancelAmt"`
	TransDt        string `json:"transDt"`
	TransTm        string `json:"transTm"`
	RecurringToken string `json:"recurringToken"`
	CcTransType    string `json:"ccTransType"`
	PayValidDt     string `json:"payValidDt"`
	PayValidTm     string `json:"payValidTm"`
	MRefNo         string `json:"mRefNo"`
	AcquStatus     string `json:"acquStatus"`
	CardExpYymm    string `json:"cardExpYymm"`
	AcquBankNm     string `json:"acquBankNm"`
	IssuBankNm     string `json:"issuBankNm"`
	DepositDt      string `json:"depositDt"`
	DepositTm      string `json:"depositTm"`
	PaymentExpDt   string `json:"paymentExpDt"`
	PaymentExpTm   string `json:"paymentExpTm"`
	PaymentTrxSn   string `json:"paymentTrxSn"`
	CancelTrxSn    string `json:"cancelTrxSn"`
	UserId         string `json:"userId"`
	ShopId         string `json:"shopId"`
}

func handleRegistration(c echo.Context) error {
	var req RegistrationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	resp := RegistrationResponse{
		ResultCd:     "0000",
		ResultMsg:    "SUCCESS",
		TXid:         "IONPAYTEST02202009231352136562",
		ReferenceNo:  req.ReferenceNo,
		PayMethod:    req.PayMethod,
		Amt:          req.Amt,
		TransDt:      req.TimeStamp[:8],
		TransTm:      req.TimeStamp[8:],
		BankCd:       "BMRI",
		VacctNo:      "70014000091352136562",
		VacctValidDt: "20200924",
		VacctValidTm: "121517",
	}

	return c.JSON(http.StatusOK, resp)
}

func handlePayment(c echo.Context) error {
	var req PaymentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	resp := PaymentResponse{
		ResultCd:    "0000",
		ResultMsg:   "SUCCESS",
		TXid:        req.TXid,
		ReferenceNo: "MerchantRefrenceNo001",
		PayMethod:   "02",
		Amt:         "529500",
		// TransDt:     "20200926",
		// TransTm:     "195526",
		BankCd:      "BMRI",
		VacctNo:     "70014000091955267274",
		VacctValidDt: "20200927",
		VacctValidTm: "121517",
	}

	return c.JSON(http.StatusOK, resp)
}

func handleInquiry(c echo.Context) error {
	var req InquiryRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	resp := InquiryResponse{
		TXid:         req.TXid,
		IMid:         req.IMid,
		Currency:     "IDR",
		Amt:          "1000",
		InstmntMon:   "1",
		InstmntType:  "2",
		ReferenceNo:  req.ReferenceNo,
		GoodsNm:      "Test Transaction Nicepay",
		PayMethod:    "02",
		BillingNm:    "John Doe",
		ReqDt:        "20200923",
		ReqTm:        "135213",
		Status:       "3",
		ResultCd:     "0000",
		ResultMsg:    "unpaid",
		VacctValidDt: "20200924",
		VacctValidTm: "121517",
		VacctNo:      "70014000091352136562",
		BankCd:       "BMRI",
	}

	return c.JSON(http.StatusOK, resp)
}

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize Echo framework
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define API routes
	e.POST("/nicepay/direct/v2/registration", handleRegistration)
	e.POST("/nicepay/direct/v2/payment", handlePayment)
	e.POST("/nicepay/direct/v2/inquiry", handleInquiry)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
