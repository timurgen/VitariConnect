package main

func main() {
	vismaApi := NewApi("http://erp.vitari.no/vbws/", "XXXX-XXXX-XXXX-XXXX-XXXX", 9999)


	//trying to get customers
	f :=  Filter{}

	AddRowToFilter(&f, CreateFilterRow("Year", EqualTo, "2017", ""))
	AddRowToFilter(&f, CreateFilterRow("Period", EqualTo, "1", "AND"))

	vismaApi.GetLedgerTransactions(f)
}
