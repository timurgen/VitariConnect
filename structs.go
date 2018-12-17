package main

import "encoding/xml"

//common XML status block in all responses
type Status struct {
	MessageID int `xml:"MessageId"`
	Message string `xml:"Message"`
	MessageDetail string `xml:"MessageDetail"`
}

//Customer data

type Customers struct {
	Customerinfo xml.Name `xml:"Customerinfo"`
	Status Status `xml:"Status"`
	Customers []Customer `xml:"Customer"`

}

type Customer struct {
	AssociateNo int `xml:"AssociateNo"`
	CustomerNo int `xml:"CustomerNo"`
	InvoiceCustomerNo int `xml:"InvoiceCustomerNo"`
	SendToAssociateNo int `xml:"SendToAssociateNo"`
	Name string `xml:"Name"`
	ShortName string `xml:"ShortName"`
	Mobile string `xml:"Mobile"`
	Phone string
	Fax string
	EmailAddress string
	WebPage string
	CompanyNo string
	CountryCode string
	LanguageCode string
	BankAccountNo string
	PaymentTerms int
	AddressLine1 string
	AddressLine2 string
	AddressLine3 string
	AddressLine4 string
	PostCode string
	PostalArea string
	VisitPostCode string
	VisitPostalArea string
	OrgUnit1 string
	OrgUnit2 string
	OrgUnit3 string
	OrgUnit4 string
	OrgUnit5 string
	OrgUnit6 string
	OrgUnit7 string
	OrgUnit8 string
	OrgUnit9 string
	OrgUnit10 string
	OrgUnit11 string
	OrgUnit12 string
	Group1 int
	Group2 int
	Group3 int
	Group4 int
	Group5 int
	Group6 int
	Group7 int
	Group8 int
	Group9 int
	Group10 int
	Group11 int
	Group12 int
	CustomerPriceGroup1 int
	CustomerPriceGroup2 int
	CustomerPriceGroup3 int
	Information1 string
	Information2 string
	Information3 string
	Information4 string
	Information5 string
	Information6 string
	Information7 string
	Information8 string
}

///// Accounting data /////

//LedgerTransactions
type LedgerTransactionInfo struct {
	LedgerTransactionInfo string `xml:"LedgerTransactionInfo"`
	Status Status `xml:"Status"`
	LedgerTransactions LedgerTransactions `xml:"LedgerTransactions"`
}

type LedgerTransactions struct {
	LedgerTransaction []LedgerTransaction `xml:"LedgerTransaction"`
}

type LedgerTransaction struct {
	VoucherJournalNo int
	AuditNo int
	Year int
	Period int
	AccountNo int
	OrgUnit1 int
	PostedAmountDomestic float32
}

//CostUnits

type CostUnitinfo struct {
	CostUnitinfo string `xml:"CostUnitinfo"`
	Status Status `xml:"Status"`
	CostUnit []CostUnit `xml:"CostUnit"`
}

type CostUnit struct {
	CostUnitNumber int
	OrgUnit1 int
	OrgUnit2 int
	OrgUnit3 int
	OrgUnit4 int
	OrgUnit5 int
	OrgUnit6 int
	OrgUnit7 int
	OrgUnit8 int
	OrgUnit9 int
	OrgUnit10 int
	OrgUnit11 int
	OrgUnit12 int
	Name string
	AddressLine1 string
	AddressLine2 string
	AddressLine3 string
	AddressLine4 string
	PostCode string
	PostalArea string
	CustomerNo int
	Status int
	PlannedStartDate string
	PlannedEndDate string
	ActualStartDate string
	ActualEndDate string
}
