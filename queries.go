package main

import (
	"fmt"
	"strings"
)

const (
	//this query needs only header and filters
	GetCustomers string = `<?xml version="1.0" encoding="UTF-8"?>
<Customerinfo>
{header}
 <Status>
   <MessageId/>
   <Message/>
   <MessageDetail/>
 </Status>
{filters}
 <Customer>
   <AssociateNo/>
   <CustomerNo/>
   <InvoiceCustomerNo/>
   <SendToAssociateNo/>
   <Name/>
   <ShortName/>
   <Mobile/>
   <Phone/>
   <Fax/>
   <EmailAddress/>
   <WebPage/>
   <CompanyNo/>
   <CountryCode/>
   <LanguageCode/>
   <BankAccountNo/>
   <PaymentTerms/>
   <AddressLine1/>
   <AddressLine2/>
   <AddressLine3/>
   <AddressLine4/>
   <PostCode/>
   <PostalArea/>
   <VisitPostCode/>
   <VisitPostalArea/>
   <OrgUnit1>2</OrgUnit1>
   <OrgUnit2/>
   <OrgUnit3/>
   <OrgUnit4/>
   <OrgUnit5/>
   <OrgUnit6/>
   <OrgUnit7/>
   <OrgUnit8/>
   <OrgUnit9/>
   <OrgUnit10/>
   <OrgUnit11/>
   <OrgUnit12/>
   <Group1/>
   <Group2/>
   <Group3/>
   <Group4/>
   <Group5/>
   <Group6/>
   <Group7/>
   <Group8/>
   <Group9/>
   <Group10/>
   <Group11/>
   <Group12/>
   <CustomerPriceGroup1/>
   <CustomerPriceGroup2/>
   <CustomerPriceGroup3/>
   <Information1/>
   <Information2/>
   <Information3/>
   <Information4/>
   <Information5/>
   <Information6/>
   <Information7/>
   <Information8/>
 </Customer>
</Customerinfo>`

	GetCostUnits string = `<?xml version="1.0" encoding="utf-8" standalone="yes"?>
<CostUnitinfo>
{header}
 <Status>
   <MessageId></MessageId>
   <Message></Message>
   <MessageDetail></MessageDetail>
 </Status>
{filters}
 <CostUnit>
   <CostUnitNumber>{costUnitNumber}</CostUnitNumber>
   <OrgUnit1></OrgUnit1>
   <OrgUnit2></OrgUnit2>
   <OrgUnit3></OrgUnit3>
   <OrgUnit4></OrgUnit4>
   <OrgUnit5></OrgUnit5>
   <OrgUnit6></OrgUnit6>
   <OrgUnit7></OrgUnit7>
   <OrgUnit8></OrgUnit8>
   <OrgUnit9></OrgUnit9>
   <OrgUnit10></OrgUnit10>
   <OrgUnit11></OrgUnit11>
   <OrgUnit12></OrgUnit12>
   <Name></Name>
   <AddressLine1></AddressLine1>
   <AddressLine2></AddressLine2>
   <AddressLine3></AddressLine3>
   <AddressLine4></AddressLine4>
   <PostCode></PostCode>
   <PostalArea></PostalArea>
   <CustomerNo></CustomerNo>
   <Status></Status>
   <PlannedStartDate></PlannedStartDate>
   <PlannedEndDate></PlannedEndDate>
   <ActualStartDate></ActualStartDate>
   <ActualEndDate></ActualEndDate>
 </CostUnit>
</CostUnitinfo>`

	GetLedgerTransactions string = `<?xml version="1.0" encoding="utf-8" standalone="yes"?>
<LedgerTransactionInfo>
{header}
 <Status>
   <MessageId></MessageId>
   <Message></Message>
   <MessageDetail></MessageDetail>
 </Status>
{filters}
 <LedgerTransactions>
   <LedgerTransaction>
     <VoucherJournalNo></VoucherJournalNo>
     <AuditNo></AuditNo>
     <Year></Year>
     <Period></Period>
     <AccountNo></AccountNo>
     <OrgUnit1></OrgUnit1>
     <PostedAmountDomestic></PostedAmountDomestic>
   </LedgerTransaction>
 </LedgerTransactions>
</LedgerTransactionInfo>`
)

//args must be in form of placeholder_name, variable = "{header}", header, "{filters}", filters
func buildQuery(query string, args ...string) string{
	r := strings.NewReplacer(args...)
	res := fmt.Sprintf(r.Replace(query))
	return res
}
