package iso8583

type MTI string

const (
	MTIAuthorizationRequest       MTI = "0100"
	MTIAuthorizationAdviceRequest MTI = "0120"
	MTIReversalRequest            MTI = "0400"
	MTINetworkManagementRequest   MTI = "0800"
	MTINetworkAdviceRequest       MTI = "0820"

	MTIAuthorizationResponse       MTI = "0110"
	MTIAuthorizationAdviceResponse MTI = "0130"
	MTIReversalResponse            MTI = "0410"
	MTINetworkManagementResponse   MTI = "0810"
	MTINetworkAdviceResponse       MTI = "0830"

	UNKNOWN MTI = "unknown"
)

func (m MTI) ResponseMTI() MTI {
	switch m {
	case MTIAuthorizationRequest:
		return MTIAuthorizationResponse
	case MTIAuthorizationAdviceRequest:
		return MTIAuthorizationAdviceResponse
	case MTIReversalRequest:
		return MTIReversalResponse
	case MTINetworkManagementRequest:
		return MTINetworkManagementResponse
	case MTINetworkAdviceRequest:
		return MTINetworkAdviceResponse
	default:
		return UNKNOWN
	}
}

type ProcessingCode string

const (
	ServicePurchase     ProcessingCode = "000000" // 00 Электронная коммерция
	ATMWithdrawal       ProcessingCode = "000001" // 01 Снятие через АТМ
	AccountFunding      ProcessingCode = "000010" // Списание при p2p
	Unique              ProcessingCode = "000011" // WAY4 транзакции / Транзит VISA/MC
	CashAdvance         ProcessingCode = "000017" // Оплата с POS-терминала
	PurchaseReturn      ProcessingCode = "000020" // Возврат (VISA/MasterCard) todo обработать в рамках 04XX
	Payment             ProcessingCode = "000028" // Пополнение при p2p / баланса карты
	CashIn              ProcessingCode = "000029" // Внесение наличных через АТМ
	BalanceInquiry      ProcessingCode = "000030" // Запрос баланса карты
	MiniStatement       ProcessingCode = "000032" // Минивыписка о последних 10 операциях
	AccountVerification ProcessingCode = "000039" // Проверка счета
	CardControl         ProcessingCode = "000091" // Управление картой
	PinChange           ProcessingCode = "000092" // Смена ПИН-кода

	NotImplementedCode ProcessingCode = "666" // Если пришел неизвестный код
)

type ResponseCode string

const (
	Approved          ResponseCode = "00"
	NoActionTaken     ResponseCode = "21"
	FormatError       ResponseCode = "30"
	NoCardRecord      ResponseCode = "56"
	SystemMalfunction ResponseCode = "96"
)
