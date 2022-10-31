package models

type Transaction struct {
	TransactionId      uint    `csv:"TransactionId" gorm:"primaryKey" json:"transaction_id"`
	RequestId          uint    `csv:"RequestId" json:"request_id"`
	TerminalId         uint    `csv:"TerminalId" json:"terminal_id"`
	PartnerObjectId    uint    `csv:"PartnerObjectId" json:"partner_object_id"`
	AmountTotal        int     `csv:"AmountTotal" json:"amount_total"`
	AmountOriginal     int     `csv:"AmountOriginal" json:"amount_original"`
	CommissionPS       float32 `csv:"CommissionPS" gorm:"type:decimal(10,2)" json:"commission_ps"`
	CommissionClient   float32 `csv:"CommissionClient" gorm:"type:decimal(10,2)" json:"commission_client"`
	CommissionProvider float32 `csv:"CommissionProvider" gorm:"type:decimal(10,2)" json:"commission_provider"`
	DateInput          string  `csv:"DateInput" json:"date_input"`
	DatePost           string  `csv:"DatePost" json:"date_post"`
	Status             string  `csv:"Status" json:"status"`
	PaymentType        string  `csv:"PaymentType" json:"payment_type"`
	PaymentNumber      string  `csv:"PaymentNumber" json:"payment_number"`
	ServiceId          uint    `csv:"ServiceId" json:"service_id"`
	Service            string  `csv:"Service" json:"service"`
	PayeeId            uint    `csv:"PayeeId" json:"payee_id"`
	PayeeName          string  `csv:"PayeeName" json:"payee_name"`
	PayeeBankMfo       string  `csv:"PayeeBankMfo" json:"payee_bank_mfo"`
	PayeeBankAccount   string  `csv:"PayeeBankAccount" json:"payee_bank_account"`
	PaymentNarrative   string  `csv:"PaymentNarrative" json:"payment_narrative"`
}
