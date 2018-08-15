package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/crazytravel/mgwaas/utils"

	"github.com/crazytravel/mgwaas/mgw"
	"github.com/gorilla/mux"
)

const (
	jhk, jhv   = "Content-Type", "application/json"
	httpGet    = "GET"
	httpPost   = "POST"
	httpPut    = "PUT"
	httpPatch  = "PATCH"
	httpDELETE = "DELETE"
)

var (
	fakeUser = mgw.User{
		ID:       "1",
		Username: "WangShuo",
		Password: "123",
		Phone:    "186",
		Email:    "ws@125.com",
	}
	fakeProxy = mgw.EdgeProxy{
		ID:        "11",
		ProxyName: "Test Proxy",
		TargetURL: "http://localhost:3000",
	}
)

// InitRouter function is set up routers
func InitRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods(httpGet)
	router.HandleFunc("/users", usersHandler).Methods(httpGet)
	router.HandleFunc("/start-gateway", startGatewayHandler).Methods(httpPost)
	router.HandleFunc("/stop-gateway", stopGatewayHandler).Methods(httpPost)
	router.HandleFunc("/load-template", loadTemplateHandler).Methods(httpGet)
	http.Handle("/", router)
	logs := http.ListenAndServe(":8080", router)
	log.Fatal(logs)
}

func loadTemplateHandler(wr http.ResponseWriter, req *http.Request) {
	wr.WriteHeader(http.StatusOK)
	mgw.ParseTemplate(wr)
}

func homeHandler(wr http.ResponseWriter, req *http.Request) {
	wr.WriteHeader(http.StatusOK)
	io.WriteString(wr, "Microgateway As a Service API")
}

func startGatewayHandler(wr http.ResponseWriter, req *http.Request) {
	wr.WriteHeader(http.StatusOK)
	wr.Header().Set(jhk, jhv)
	json.NewEncoder(wr).Encode(utils.StatusOK(""))
}

func stopGatewayHandler(wr http.ResponseWriter, req *http.Request) {
	wr.WriteHeader(http.StatusOK)
	wr.Header().Set(jhk, jhv)
	json.NewEncoder(wr).Encode(utils.StatusOK(""))
}

func usersHandler(wr http.ResponseWriter, req *http.Request) {
	wr.WriteHeader(http.StatusOK)
	wr.Header().Set(jhk, jhv)
	json.NewEncoder(wr).Encode(fakeUser)
}
