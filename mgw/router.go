package mgw

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

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
	fakeUser = User{
		ID:       "1",
		Username: "WangShuo",
		Password: "123",
		Phone:    "186",
		Email:    "ws@125.com",
	}
	fakeProxy = EdgeProxy{
		ID:        "11",
		ProxyName: "Test Proxy",
		TargetURL: "http://localhost:3000",
	}
)

// InitRouter function is set up routers
func InitRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods(httpGet)
	router.HandleFunc("/start-gateway", startGateway).Methods(httpPost)
	router.HandleFunc("/stop-gateway", stopGateway).Methods(httpPost)
	router.HandleFunc("/users", usersHandler).Methods(httpGet)
	router.HandleFunc("/edge-proxies", edgeProxiesHandler).Methods(httpGet)
	http.Handle("/", router)
	logs := http.ListenAndServe(":8080", router)
	log.Fatal(logs)
}

// HomeHandler is a handler for root path
func homeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	io.WriteString(writer, "Microgateway As a Service API")
}

func startGateway(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set(jhk, jhv)
	json.NewEncoder(writer).Encode(StatusOK(""))
}

func stopGateway(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set(jhk, jhv)
	json.NewEncoder(writer).Encode(StatusOK(""))
}

// UsersHandler is a user path
func usersHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set(jhk, jhv)
	json.NewEncoder(writer).Encode(fakeUser)
}

// EdgeProxiesHandler is a proxy path
func edgeProxiesHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set(jhk, jhv)
	json.NewEncoder(writer).Encode(fakeProxy)
}
