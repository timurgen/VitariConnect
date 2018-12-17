package main

func main() {
	vismaApi := NewApi("http://erp.vitari.no/vbws/", "XXXX-XXXX-XXXX-XXXX-XXXX", 9999)


	//trying to get customers
	f :=  Filter{}

	AddRowToFilter(&f, CreateFilterRow("OrgUnit1", GreaterThanOrEqualTo, "70000", ""))

	vismaApi.GetCostUnits(f, 1)
}
