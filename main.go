package main

import (
	"github.com/crazytravel/mgwaas/mgw"
)

func init() {
	// fmt.Println("...init...")
	// mgw := mgw.MicroGateway{
	// 	Org:      "wangshuo866-eval",
	// 	Env:      mgw.EnvProd,
	// 	Key:      "91b0d6e10b52e2c44ea5f8d62d10c785d9f5c94481e899d743afc8db0f163318",
	// 	Secret:   "a900025dbb83998721f9509b060627aa68884459c24661613972d32d470d5f5c",
	// 	Username: "wangshuo866@gmail.com",
	// 	Password: "Wang@891369",
	// }
	// mgw.InitGateway()
}

func main() {
	// mgw.ParseTemplate()
	mgw.InitRouter()
}
