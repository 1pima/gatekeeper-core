package iso8583

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"gatekeeper-core/internal/config"

	mooviso "github.com/moov-io/iso8583"
)

type Message struct {
	MTI                                 MTI            `iso8583:"0"`
	PrimaryAccountNumber                string         `iso8583:"2"`
	ProcessingCode                      ProcessingCode `iso8583:"3"`
	Amount                              int64          `iso8583:"4"`
	SettlementAmount                    int64          `iso8583:"5"`
	BillingAmount                       string         `iso8583:"6"`
	TransmissionDate                    string         `iso8583:"7"`
	BillingFeeAmount                    string         `iso8583:"8"`
	SettlementConversionRate            string         `iso8583:"9"`
	CardholderBillingConversionRate     string         `iso8583:"10"`
	STAN                                string         `iso8583:"11"`
	LocalTransactionTime                string         `iso8583:"12"`
	LocalTransactionDate                string         `iso8583:"13"`
	ExpirationDate                      string         `iso8583:"14"`
	SettlementDate                      string         `iso8583:"15"`
	CurrencyConversionDate              string         `iso8583:"16"`
	CaptureDate                         string         `iso8583:"17"`
	MCC                                 int64          `iso8583:"18"`
	AcquiringInstitutionCountryCode     string         `iso8583:"19"`
	PANExtendedCountryCode              string         `iso8583:"20"`
	ForwardingInstitutionCountryCode    string         `iso8583:"21"`
	POSEntryMode                        string         `iso8583:"22"`
	CardSequenceNumber                  string         `iso8583:"23"`
	FunctionCode                        string         `iso8583:"24"`
	POSConditionCode                    string         `iso8583:"25"`
	POSPINCaptureCode                   string         `iso8583:"26"`
	AuthIdentificationResponseLength    string         `iso8583:"27"`
	TransactionFeeAmount                string         `iso8583:"28"`
	SettlementFeeAmount                 string         `iso8583:"29"`
	TransactionProcessingFeeAmount      string         `iso8583:"30"`
	SettlementProcessingFeeAmount       string         `iso8583:"31"`
	AcquiringInstitutionIDCode          string         `iso8583:"32"`
	ForwardingInstitutionIDCode         string         `iso8583:"33"`
	ExtendedPAN                         string         `iso8583:"34"`
	Track2Data                          string         `iso8583:"35"`
	Track3Data                          string         `iso8583:"36"`
	RetrievalReferenceNumber            string         `iso8583:"37"`
	AuthorizationIdentificationResponse string         `iso8583:"38"`
	ResponseCode                        ResponseCode   `iso8583:"39"`
	ServiceRestrictionCode              string         `iso8583:"40"`
	CardAcceptorTerminalID              string         `iso8583:"41"`
	CardAcceptorIDCode                  string         `iso8583:"42"`
	CardAcceptorNameLocation            string         `iso8583:"43"`
	AdditionalData                      string         `iso8583:"44"`
	Track1Data                          string         `iso8583:"45"`
	AdditionalDataISO                   string         `iso8583:"46"`
	AdditionalDataNational              string         `iso8583:"47"`
	AdditionalDataPrivate               string         `iso8583:"48"`
	TransactionCurrencyCode             int64          `iso8583:"49"`
	SettlementCurrencyCode              int64          `iso8583:"50"`
	CardholderBillingCurrencyCode       string         `iso8583:"51"`
	PINData                             string         `iso8583:"52"`
	SecurityRelatedControlInfo          string         `iso8583:"53"`
	AdditionalAmounts                   string         `iso8583:"54"`
	ICCDataEMV                          string         `iso8583:"55"`
	ReservedISO                         string         `iso8583:"56"`
	ReservedNational1                   string         `iso8583:"57"`
	ReservedNational2                   string         `iso8583:"58"`
	ReservedNational3                   string         `iso8583:"59"`
	ReservedNational4                   string         `iso8583:"60"`
	ReservedPrivate1                    string         `iso8583:"61"`
	ReservedPrivate2                    string         `iso8583:"62"`
	ReservedPrivate3                    string         `iso8583:"63"`
	MAC                                 string         `iso8583:"64"`
	NetworkManagementInfo               string         `iso8583:"70"`
	OriginalDataElements                string         `iso8583:"90"`
}

func (m *Message) CardHash() string {
	key := []byte(config.Settings.HMACKey)

	h := hmac.New(sha256.New, key)
	h.Write([]byte(m.PrimaryAccountNumber))

	return hex.EncodeToString(h.Sum(nil))
}

func (m *Message) Unmarshal(data []byte) error {
	msg := mooviso.NewMessage(customSpec)

	// 1. Парсим байты в абстрактный ISO-объект
	if err := msg.Unpack(data); err != nil {
		return fmt.Errorf("unpacking raw bytes: %w", err)
	}

	// 2. Мапим ISO-объект в переданную бизнес-схему
	if err := msg.Unmarshal(m); err != nil {
		return fmt.Errorf("unmarshaling to schema: %w", err)
	}

	return nil
}

// Marshal берет твою структуру и превращает ее в сырые байты для сети
func (m *Message) Marshal() ([]byte, error) {
	msg := mooviso.NewMessage(customSpec)

	// 1. Мапим бизнес-схему в абстрактный ISO-объект
	if err := msg.Marshal(m); err != nil {
		return nil, fmt.Errorf("marshaling from schema: %w", err)
	}

	// 2. Пакуем в байты с применением всех правил из customSpec
	packed, err := msg.Pack()
	if err != nil {
		return nil, fmt.Errorf("packing to raw bytes: %w", err)
	}

	return packed, nil
}
