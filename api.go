package main

import (
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

//send post request with xml payload but with plain text content type (visma returns error if content type is XML ¯\_(ツ)_/¯
//some VISMA methods need not only filter but also add constraints (eg CustomerNo) into request body
//docs http://hjelp.vitari.no/VitariConnect/Vbus/Vitari%20Connect%20VismaBusiness.html
func NewApi(url string, guid string, clientId string) *Api {
	return &Api{url, guid, clientId}
}

// API object, consists of URL to Visma service with trailing / eg http://erp.vitari.no/vbws/
//guid and client id
type Api struct {
	url      string
	guid     string
	clientId string
}

//GetCustomers send request to Visma GetCustomers end point
//return struct with response. Check Message part for any error codes after receiving. any MessageID > 0 are error codes
func (r *Api) GetCustomers(f Filter) Customers {
	queryBody := buildQuery(GetCustomers, "{header}", BuildHeader(r.guid, r.clientId), "{filters}",
		RenderFilter(f))
	responseBytes, err := makeVismaApiRequest(r, "Customer.svc/GetCustomers", queryBody)
	if err != nil {
		fmt.Println(err)
	}

	var customers Customers
	xml.Unmarshal(responseBytes, &customers)

	return customers
}

//GetCostUnits returns CostUnits based on filter and costUnitNumber parameter
//return struct with response. Check Message part for any error codes after receiving. any MessageID > 0 are error codes
func (r *Api) GetCostUnits(f Filter, costUnitNumber int) CostUnitinfo {
	queryBody := buildQuery(GetCostUnits, "{header}", BuildHeader(r.guid, r.clientId),
		"{filters}", RenderFilter(f), "{costUnitNumber}", strconv.Itoa(costUnitNumber))
	responseBytes, err := makeVismaApiRequest(r, "Accounting.svc/GetCostUnits", queryBody)
	if err != nil {
		fmt.Println(err)
	}

	var costUnitInfo CostUnitinfo
	xml.Unmarshal(responseBytes, &costUnitInfo)

	return costUnitInfo
}

//GetLedgerTransactions returns GetLedgerTransactions based on filter
//return struct with response. Check Message part for any error codes after receiving. any MessageID > 0 are error codes
func (r *Api) GetLedgerTransactions(f Filter) LedgerTransactionInfo {
	queryBody := buildQuery(GetLedgerTransactions, "{header}", BuildHeader(r.guid, r.clientId), "{filters}",
		RenderFilter(f))
	responseBytes, err := makeVismaApiRequest(r, "Accounting.svc/GetLedgerTransactions", queryBody)
	if err != nil {
		fmt.Println(err)
	}

	var ledgerTransactionInfo LedgerTransactionInfo
	xml.Unmarshal(responseBytes, &ledgerTransactionInfo)

	return ledgerTransactionInfo
}

//makeVismaApiRequest internal function construct and make HTTP request to Visma instance
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
