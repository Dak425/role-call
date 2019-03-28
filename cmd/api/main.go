package main

import (
	"github.com/Dak425/role-call/pkg/api/http"
)

func main() {
	// Initialize any dependencies here (Database connections, interface implementations, etc)

	// Start server
	http.StartServer()
}
