package visma

import (
	"fmt"
)

//send post request with xml payload but with plain text content type (visma returns error if content type is XML ¯\_(ツ)_/¯
//some VISMA methods need not only filter but also add constraints (eg CustomerNo) into request body
//docs http://hjelp.vitari.no/VitariConnect/Vbus/Vitari%20Connect%20VismaBusiness.html
func NewApi(url string, guid string, clientId int) * Api {
	return &Api{url,guid, clientId}
}

type Api struct {
	url string
	guid string
	clientId int
}

func (r *Api) GetCustomers(f Filter)  {
	//build xml request body
	//make request
	//parse xml response
	fmt.Println("GetCustomers fired")
}

func (r *Api) GetCostUnits(f Filter)  {
	//build xml request body
	//make request
	//parse xml response
	fmt.Println("GetCostUnits fired")
}

func (r *Api) GetLedgerTransactions(f Filter)  {
	//build xml request body
	//make request
	//parse xml response
	fmt.Println("GetLedgerTransactions fired")
}