package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
)

var api *Api

//simple REST service to fetch data from Vitari Connect API and return as JSON
func main() {
	vismaUrl := os.Getenv("VISMA_URL")
	if vismaUrl == "" {
		//panic("URL for Vitari Connect instance not found")
	}

	vismaGuid := os.Getenv("VISMA_GUID")
	vismaClientId := os.Getenv("VISMA_CLIENTID")
	api = NewApi(vismaUrl, vismaGuid, vismaClientId)

	wsPort := os.Getenv("PORT")
	if wsPort == "" {
		wsPort = "8080"
	}

	router := mux.NewRouter()
	//Get data
	router.HandleFunc("/datasets/Customer/entities", _GetCustomers).Methods("GET")
	router.HandleFunc("/datasets/CostUnit/entities", _GetCostUnits).Methods("GET")
	router.HandleFunc("/datasets/LedgerTransaction/entities", _GetLedgerTransactions).Methods("GET")
	//Update data
	router.HandleFunc("/GetNextAvailableCostUnitAssignAndUpdate", _GetNextAvailableCostUnitAssignAndUpdate).Methods("POST")

	log.Printf("Starting service on port %s", wsPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", wsPort), router))

}

func _GetLedgerTransactions(w http.ResponseWriter, r *http.Request) {
	log.Printf("Servinq request %s from %s", r.RequestURI, r.Host)
	currentTime := time.Now()
	currentYear := currentTime.Year()
	currentMonth := int(currentTime.Month())
	since := r.URL.Query().Get("since")
	var year int
	var month int
	if since == "" {
		year = currentTime.Year()
		month = currentMonth
	} else {
		layout := "2006-01-02T15:04:05.000-0700"
		//layout:= time.RFC3339
		t, err := time.Parse(layout, since)
		if err != nil {
			log.Printf("Couldn't parse date time from since %s", since)
			year = currentTime.Year()
			month = int(currentTime.Month())
		} else {
			year = t.Year()
			month = int(t.Month())
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("["))
	first := true
	for ; year <= currentYear+2; year++ {
		runtime.GC()
		for ; month <= 12; month++ {
			if year == currentYear+2 && month > 1 {
				break
			}
			log.Printf("Fetching transactions for %d/%d", year, month)
			var f Filter
			AddRowToFilter(&f, CreateFilterRow("Year", EqualTo, strconv.Itoa(year), ""))
			AddRowToFilter(&f, CreateFilterRow("Period", EqualTo, strconv.Itoa(month), "AND"))
			transactions := api.GetLedgerTransactions(f)
			if transactions.Status.MessageID != 0 {
				log.Printf("Couldn't fetch transactions: %s", transactions.Status.Message)
				log.Println("Trying next period")
			} else {
				for _, transaction := range transactions.LedgerTransactions.LedgerTransaction {
					if first {
						first = false
					} else {
						w.Write([]byte(","))
					}
					jsonData, err := json.Marshal(transaction)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					w.Write(jsonData)

				}
			}
		}
		month = 1
	}
	w.Write([]byte("]"))

}

func _GetCustomers(w http.ResponseWriter, r *http.Request) {
	log.Printf("Servinq request %s from %s", r.RequestURI, r.Host)
	var f Filter
	var first = true
	AddRowToFilter(&f, CreateFilterRow("CustomerNo", GreaterThanOrEqualTo, "0", ""))
	customers := api.GetCustomers(f)
	if customers.Status.MessageID != 0 {
		log.Printf("Couldn't fetch customers: %s", customers.Status.Message)
		http.Error(w, customers.Status.Message, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("["))
	for _, customer := range customers.Customers {
		if first {
			first = false
		} else {
			w.Write([]byte(","))
		}
		jsonData, err := json.Marshal(customer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)

	}
	w.Write([]byte("]"))

}
func _GetCostUnits(w http.ResponseWriter, r *http.Request) {
	log.Printf("Servinq request %s from %s", r.RequestURI, r.Host)

	var f Filter
	var first = true
	var orgUnit = r.URL.Query().Get("orgUnit")
	costUnitNumber, _ := strconv.Atoi(r.URL.Query().Get("costUnitNumber"))

	AddRowToFilter(&f, CreateFilterRow("OrgUnit1", GreaterThanOrEqualTo, orgUnit, ""))
	costUnits := api.GetCostUnits(f, costUnitNumber)

	if costUnits.Status.MessageID != 0 {
		log.Printf("Couldn't fetch CostUnits: %s", costUnits.Status.Message)
		http.Error(w, costUnits.Status.Message, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("["))
	for _, customer := range costUnits.CostUnit {
		if first {
			first = false
		} else {
			w.Write([]byte(","))
		}
		jsonData, err := json.Marshal(customer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)

	}
	w.Write([]byte("]"))

}

type SharepointDofiToVisma struct {
	CommonProjNumber        string `json:"Common_ProjNumber"`
	PO_DOFINum              string `json:"PO_DOFINum"`
	PO_ProjectName          string `json:"PO_ProjectName"`
	IdInternal              string `json:"_id"`
	DH_Is_Updated           bool   `json:"DH_Is_Updated"`
	Sys_Is_Request_To_Visma bool   `json:"Sys_Is_Request_To_Visma"`
	PO_ProjectGUID          string `json:"PO_ProjectGUID"`
	ID                      int    `json:"ID"`
	Status                  string
	Already_Exists          bool `json:"Already_Exists"`
}

//Brukes som HTTP transformasjon, tar en eller flere prosjekter fra Sesam,
//henter tilgjengelige cost units fra Visma, knytter til, push tilbake til visma, og returnerer tilbake til sesam med
//tilordnet prosjekt nr
func _GetNextAvailableCostUnitAssignAndUpdate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Servinq request %s from %s", r.RequestURI, r.Host)
	decoder := json.NewDecoder(r.Body)
	var inputData []SharepointDofiToVisma
	err := decoder.Decode(&inputData)
	if err != nil {
		panic(err)
	}

	var f Filter
	var orgUnit = r.URL.Query().Get("orgUnit")
	costUnitNumber, _ := strconv.Atoi(r.URL.Query().Get("costUnitNumber"))

	AddRowToFilter(&f, CreateFilterRow("OrgUnit1", GreaterThanOrEqualTo, orgUnit, ""))
	AddRowToFilter(&f, CreateFilterRow("Name", EqualTo, "NN", "AND"))
	AddRowToFilter(&f, CreateFilterRow("Name", EqualTo, "", "OR"))
	costUnits := api.GetCostUnits(f, costUnitNumber)
	if costUnits.Status.MessageID != 0 {
		log.Printf("Couldn'inputData fetch CostUnits: %s", costUnits.Status.Message)
		http.Error(w, costUnits.Status.Message, http.StatusInternalServerError)
		return
	}

	if len(costUnits.CostUnit) < len(inputData) {
		log.Printf("Det kommet %d projects og Visma har kun %d ledige!", len(inputData), len(costUnits.CostUnit))
		w.WriteHeader(500)
		return
	}

	var first = true
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("["))

	for key, value := range inputData {
		if len(value.CommonProjNumber) > 0 {
			log.Printf("Only entities without assigned project are expected here. Got entity with projectnr %s",
				value.CommonProjNumber)
			continue
		}
		//log.Printf("%v",value)
		var nextOrgUnit1Number = costUnits.CostUnit[key].OrgUnit1
		var ProjName = value.PO_ProjectName

		if len(ProjName) == 0 {
			ProjName = "Project without name"
		}

		value.CommonProjNumber = strconv.Itoa(nextOrgUnit1Number)

		if first {
			first = false
		} else {
			w.Write([]byte(","))
		}

		//update Visma costUnit
		costUnits := api.PutCostUnit(ProjName, costUnitNumber, nextOrgUnit1Number)
		if costUnits.Status.MessageID != 0 {
			log.Printf("Couldn'inputData update cost unit: %s", costUnits.Status.Message)
			value.Status = costUnits.Status.Message
		} else {
			log.Printf("Cost unit with orgNumber %d updated %s", nextOrgUnit1Number, costUnits.Status.Message)
			value.Sys_Is_Request_To_Visma = true
		}

		jsonData, err := json.Marshal(value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)

	}
	w.Write([]byte("]"))
}
