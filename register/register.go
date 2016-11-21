package register

import (
	"net"

	"Wildfire/hallo"
)

// Regist Regist listener
func Regist(lis net.Listener) {
	hallo.RegistServer(lis)
}
