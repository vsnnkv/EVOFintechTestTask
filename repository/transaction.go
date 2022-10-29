package repository

import (
	"time"
)

type Transaction struct {
	TransactionId      uint     `csv:"TransactionId" gorm:"primaryKey"`
	RequestId          uint     `csv:"RequestId"`
	TerminalId         uint     `csv:"TerminalId"`
	PartnerObjectId    uint     `csv:"PartnerObjectId"`
	AmountTotal        int      `csv:"AmountTotal"`
	AmountOriginal     int      `csv:"AmountOriginal"`
	CommissionPS       float32  `csv:"CommissionPS" gorm:"type: float4"`
	CommissionClient   float32  `csv:"CommissionClient" gorm:"type: float4"`
	CommissionProvider float32  `csv:"CommissionProvider" gorm:"type: float4"`
	DateInput          DateTime `csv:"DateInput" gorm:"type: jsonb"`
	DatePost           DateTime `csv:"DatePost" gorm:"type: jsonb"`
	Status             string   `csv:"Status"`
	PaymentType        string   `csv:"PaymentType"`
	PaymentNumber      string   `csv:"PaymentNumber"`
	ServiceId          uint     `csv:"ServiceId"`
	Service            string   `csv:"Service"`
	PayeeId            uint     `csv:"PayeeId"`
	PayeeName          string   `csv:"PayeeName"`
	PayeeBankMfo       string   `csv:"PayeeBankMfo"`
	PayeeBankAccount   string   `csv:"PayeeBankAccount"`
	PaymentNarrative   string   `csv:"PaymentNarrative"`
}

type DateTime struct {
	time.Time
}

func (date *DateTime) MarshalCSV() (string, error) {
	return date.Time.Format("2006-01-02 15:04:05"), nil
}

func (date *DateTime) String() string {
	return date.String()
}

func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("2006-01-02 15:04:05", csv)
	return err
}
