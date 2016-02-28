package healpay

type LineItem struct {
	ProductRefNum string
	SKU string
	ProductName string
	Description string
	UnitPrice string
	Qty string
	Taxable bool
}

type RecurringBilling struct {
	Schedule string
	Next string
	Expire string
	NumLeft int
	Amount float64
	Enabled bool
}

type Address struct {
	FirstName string
	LastName string
	Company string
	Street string
	Street2 string
	City string
	State string
	Zip string
	Country string
	Phone string
	Fax string
	Email string
}

type CheckData struct {
	CheckNumber string
	Routing string
	Account string
	AccountType string
	DriversLicense string
	DriversLicensePlate string
	RecordType string
	AuxOnUS string
	EpcCode string
	FrontImage string
	BackImage string
}

type CreditCardData struct {
	CardType string
	CardNumber string
	CardExpiration string
	CardCode string
	AvsStreet string
	AvsZip string
	CardPresent bool
	MagStripe string
	DUKPT string
	Signature string
	TermType string
	MagSupport string
	XID string
	CAVV string
	ECI int
	InternalCardAuth bool
	Pares string
}

type TransactionDetail struct {
	Invoice string
	PONum string
	OrderID string
	Clerk string
	Terminal string
	Table string
	Description string
	Comments string
	AllowPartialAuth bool
	Amount float64
	Currency string
	Tax float64
	Tip float64
	NonTax bool
	Shipping float64
	Discount float64
	Subtotal float64
}

type TransactionRequestObject struct {
	Command string
	IgnoreDuplicate bool
	AuthCode string
	RefNum string
	AccountHolder string
	Details *TransactionDetail
	CreditCardData *CreditCardData
	CheckData *CheckData
	ClientIP string
	CustomerID string
	BillingAddress *Address
	ShippingAddress *Address
	CustReceipt bool
	Software string
	CustReceiptName string
	RecurringBilling *RecurringBilling
	LineItems *[]LineItem
}
