package main

import (
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

//send post request with xml payload but with plain text content type (visma returns error if content type is XML ¯\_(ツ)_/¯
//some VISMA methods need not only filter but also add constraints (eg CustomerNo) into request body
//docs http://hjelp.vitari.no/VitariConnect/Vbus/Vitari%20Connect%20VismaBusiness.html
func NewApi(url string, guid string, clientId int) *Api {
	return &Api{url, guid, clientId}
}
// API object, consists of URL to Visma service with trailing / eg http://erp.vitari.no/vbws/
//guid and client id
type Api struct {
	url      string
	guid     string
	clientId int
}
//GetCustomers send request to Visma GetCustomers end point
//return struct with response. Check Message part for any error codes after receiving
func (r *Api) GetCustomers(f Filter) Customers{
	//build xml request body
	queryBody := buildQuery(GetCustomers, "{header}", BuildHeader(r.guid, r.clientId), "{filters}", RenderFilter(f))
	//make request
	customerBytes, err := makeVismaApiRequest(r, "Customer.svc/GetCustomers", queryBody)
	if err != nil {
		fmt.Println(err)
	}
	var customers Customers

	xml.Unmarshal(customerBytes, &customers)
	return customers
}

func (r *Api) GetCostUnits(f Filter) {
	//build xml request body
	//make request
	//parse xml response
	fmt.Println("GetCostUnits fired")
}

func (r *Api) GetLedgerTransactions(f Filter) LedgerTransactionInfo{
	//build xml request body
	queryBody := buildQuery(GetLedgerTransactions, "{header}", BuildHeader(r.guid, r.clientId), "{filters}", RenderFilter(f))
	//make request
	customerBytes, err := makeVismaApiRequest(r, "Accounting.svc/GetLedgerTransactions", queryBody)
	if err != nil {
		fmt.Println(err)
	}
	var ledgerTransactionInfo LedgerTransactionInfo

	xml.Unmarshal(customerBytes, &ledgerTransactionInfo)
	return ledgerTransactionInfo
}

//makeVismaApiRequest construct and make HTTP request to Visma instance
func makeVismaApiRequest(r *Api, endpoint string, queryBody string) ([]byte, error) {
	resp, err := http.Post(r.url+endpoint, "text/plain", strings.NewReader(queryBody))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//bodyString := string(bodyBytes)
	return bodyBytes, nil
}
