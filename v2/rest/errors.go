package rest

import (
	"fmt"
)

// APIError return the api error
type APIError struct {
	Status  uint32
	Message string
}

// Error return the error message
func (e APIError) Error() string {
	var s string
	switch e.Status {
	// System errors
	case 10000:
		s = "Url not found"
	case 10001:
		s = "System error"
	case 10002:
		s = "Malformed request"
	case 10003:
		s = "System error"
	case 10005:
		s = "Timeout waiting for response"
	case 10007:
		s = "System maintenance"
	case 10008:
		s = "Server is busy. Retry later"
	case 10009:
		s = "You sent requests too frequently. Retry later with decreased requests"

		// Authentication errors
	case 20001:
		s = "Authentication failed api authorization"
	case 20002:
		s = "Invalid api key"
	case 20003:
		s = "Api key not found"
	case 20004:
		s = "Invalid api nonce"
	case 20005:
		s = "Invalid api signature"
	case 20011:
		s = "MFA failed"
	case 20014:
		s = "SMS verification failed"
	case 20018:
		s = "Please login. (This happens when you request API without /v1/.)"
	case 20023:
		s = "Missing OTP code"
	case 20024:
		s = "Missing SMS code"
	case 20025:
		s = "Missing OTP and SMS code"
	case 20026:
		s = "MFA is temporarily locked because too many failures. Please retry after 60 seconds"

	// Required parameter errors
	case 30001:
		s = "Missing order quantity"
	case 30006:
		s = "Missing order id"
	case 30007:
		s = "Missing order id array"
	case 30009:
		s = "Missing asset"
	case 30012:
		s = "Missing order price"
	case 30013:
		s = "Missing side"
	case 30015:
		s = "Missing order type"
	case 30016:
		s = "Missing asset"
	case 30019:
		s = "Missing uuid"
	case 30039:
		s = "Missing withdraw amount"
	case 30101:
		s = "Missing trigger price"
	case 30103:
		s = "Missing withdrawal type"
	case 30104:
		s = "Missing withdrawal name"
	case 30105:
		s = "Missing VASP"
	case 30106:
		s = "Missing beneficiary type"
	case 30107:
		s = "Missing beneficiary last name"
	case 30108:
		s = "Missing beneficiary first name"
	case 30109:
		s = "Missing beneficiary last kana"
	case 30110:
		s = "Missing beneficiary first kana"
	case 30111:
		s = "Missing beneficiary company name"
	case 30112:
		s = "Missing beneficiary company kana"
	case 30113:
		s = "Missing beneficiary company type"
	case 30114:
		s = "Missing beneficiary company type position"
	case 30115:
		s = "Missing uploaded documents"
	case 30116:
		s = "Missing withdrawal purpose"
	case 30117:
		s = "Missing beneficiary country"
	case 30118:
		s = "Missing beneficiary zip code"
	case 30119:
		s = "Missing beneficiary prefecture"
	case 30120:
		s = "Missing beneficiary city"
	case 30121:
		s = "Missing beneficiary address"
	case 30122:
		s = "Missing beneficiary building"
	case 30123:
		s = "Missing extraction request category"

	// Invalid parameter errors
	case 40001:
		s = "Invalid order quantity"
	case 40006:
		s = "Invalid count"
	case 40007:
		s = "Invalid end param"
	case 40008:
		s = "Invalid end_id"
	case 40009:
		s = "Invalid from_id"
	case 40013:
		s = "Invalid order id"
	case 40014:
		s = "Invalid order id array"
	case 40015:
		s = "Too many orders are specified"
	case 40017:
		s = "Invalid asset"
	case 40020:
		s = "Invalid order price"
	case 40021:
		s = "Invalid order side"
	case 40022:
		s = "Invalid trading start time"
	case 40024:
		s = "Invalid order type"
	case 40025:
		s = "Invalid asset"
	case 40028:
		s = "Invalid uuid"
	case 40048:
		s = "Invalid withdraw amount"
	case 40112:
		s = "Invalid trigger price"
	case 40113:
		s = "Invalid post_only"
	case 40114:
		s = "post_only can not be specified with such order type"
	case 40116:
		s = "Invalid withdrawal type"
	case 40117:
		s = "Invalid withdrawal name"
	case 40118:
		s = "Invalid VASP"
	case 40119:
		s = "Invalid beneficiary type"
	case 40120:
		s = "Invalid beneficiary last name"
	case 40121:
		s = "Invalid beneficiary first name"
	case 40122:
		s = "Invalid beneficiary last kana"
	case 40123:
		s = "Invalid beneficiary first kana"
	case 40124:
		s = "Invalid beneficiary company name"
	case 40125:
		s = "Invalid beneficiary company kana"
	case 40126:
		s = "Invalid beneficiary company type"
	case 40127:
		s = "Invalid beneficiary company type position"

	// Data errors
	case 50003:
		s = "Account is restricted"
	case 50004:
		s = "Account is provisional"
	case 50005:
		s = "Account is blocked"
	case 50006:
		s = "Account is blocked"
	case 50008:
		s = "Identity verification is not finished"
	case 50009:
		s = "Order not found"
	case 50010:
		s = "Order can not be canceled"
	case 50011:
		s = "Api not found"
	case 50026:
		s = "Order has already been canceled"
	case 50027:
		s = "Order has already been executed"
	case 50033:
		s = "Withdrawals to this address require additional entries"
	case 50034:
		s = "VASP not found"
	case 50035:
		s = "Company information is not registerd"
	case 50037:
		s = "We are temporarily restricting withdrawals while we verify your last deposit. Please try again in a few minutes"
	case 50038:
		s = "Cannot withdraw to chosen VASP service"

	// value errors
	case 60001:
		s = "Insufficient amount"
	case 60002:
		s = "Market buy order quantity has exceeded the upper limit"
	case 60003:
		s = "Order quantity has exceeded the limit"
	case 60004:
		s = "Order quantity has exceeded the lower threshold"
	case 60005:
		s = "Order price has exceeded the upper limit"
	case 60006:
		s = "Order price has exceeded the lower limit"
	case 60011:
		s = "Too many Simultaneous orders, current limit is 30"
	case 60016:
		s = "Trigger price has exceeded the upper limit"
	case 60017:
		s = "Withdrawal amount has exceeded the upper limit"

	// Stop update request system status
	case 70001:
		s = "System error"
	case 70002:
		s = "System error"
	case 70003:
		s = "System error"
	case 70004:
		s = "Order is restricted during suspension of transactions"
	case 70005:
		s = "Buy order has been temporarily restricted"
	case 70006:
		s = "Sell order has been temporarily restricted"
	case 70009:
		s = "Market order has been temporarily restricted. Please use limit order instead"
	case 70010:
		s = "Minimum Order Quantity is increased temporarily"
	case 70011:
		s = "System is busy. Please try again"
	case 70012:
		s = "System error"
	case 70013:
		s = "Order and cancel has been temporarily restricted"
	case 70014:
		s = "Withdraw and cancel request has been temporarily restricted"
	case 70015:
		s = "Lending and cancel request has been temporarily restricted"
	case 70016:
		s = "Lending and cancel request has been restricted"
	case 70017:
		s = "Orders on pair have been suspended"
	case 70018:
		s = "Order and cancel on pair have been suspended"
	case 70019:
		s = "Order cancel request is in process"
	case 70020:
		s = "Market order has been temporarily restricted"
	case 70021:
		s = "Limit order price is over the threshold"
	case 70022:
		s = "Stop limit order has been temporarily restricted"
	case 70023:
		s = "Stop order has been temporarily restricted"

	}

	return fmt.Sprintf("APIError: code=%d, error=%s, message=%s", e.Status, s, e.Message)
}
