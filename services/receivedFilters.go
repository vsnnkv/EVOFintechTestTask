package services

type ReceivedFilters struct {
	TransactionId    []string
	TerminalId       []string
	Status           string
	PaymentType      string
	DatePostFrom     string
	DatePostTo       string
	PaymentNarrative string
}

func NewReceivedFiltersFilters(transactionId []string, terminalId []string, status string, paymentType string,
	datePostFrom string, datePostTo string, paymentNarrative string) *ReceivedFilters {
	return &ReceivedFilters{
		TransactionId:    transactionId,
		TerminalId:       terminalId,
		Status:           status,
		PaymentType:      paymentType,
		DatePostFrom:     datePostFrom,
		DatePostTo:       datePostTo,
		PaymentNarrative: paymentNarrative,
	}
}

const (
	transactionField = "transaction_id"
	terminalField    = "terminal_id"
	statusField      = "status"
	paymentTypeField = "payment_type"
)
