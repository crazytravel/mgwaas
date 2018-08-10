package mgw

import (
	"fmt"
)

// Gateway is a interface for microgateway operate
type Gateway interface {
	// InitMGW is use command to init a microgateway
	InitGateway()
	// ConfigMGW is use command to config a microgateway
	ConfigGateway()
	// StartMGW is use command start a microgateway function
	StartGateway()
}

// Env is an enum
type Env string

const (
	// EnvProd is production environment
	EnvProd Env = "prod"
	// EnvTest is test environment
	EnvTest Env = "test"
)

// MicroGateway struct will collect microgateway tings
type MicroGateway struct {
	Org      string
	Env      Env
	Key      string
	Secret   string
	Username string
	Password string
}

// InitGateway is a micro gateway implement for gateway init
func (mgw *MicroGateway) InitGateway() {
	fmt.Println(mgw.Username)
	fmt.Println("init micro gateway...")
}
