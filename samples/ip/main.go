package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

func last(s string, b byte) int {
	i := len(s)
	for i--; i >= 0; i-- {
		if s[i] == b {
			break
		}
	}
	return i
}

func splitHostZone(s string) (host, zone string) {
	// The IPv6 scoped addressing zone identifier starts after the
	// last percent sign.
	if i := last(s, '%'); i > 0 {
		host, zone = s[:i], s[i+1:]
		return
	}
	host = s
	return
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-")

	// get server ip
	adr := r.Context().Value(http.LocalAddrContextKey)
	serverAddr, serverPort, err := net.SplitHostPort(
		fmt.Sprintf("%v", adr))
	if err != nil {
		fmt.Println(err)
		return
	}
	serverIP, serverZone := splitHostZone(serverAddr)
	fmt.Println("server ip:", serverIP, "port:", serverPort)
	fmt.Println("zone:", serverZone)
	fmt.Println(net.JoinHostPort(serverIP, serverPort))

	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-")

	// get client ip
	clientAddr, clientPort, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	clientIP, clientZone := splitHostZone(clientAddr)
	fmt.Println("client ip:", clientIP, "port:", clientPort)
	fmt.Println("zone:", clientZone)
	fmt.Println(net.JoinHostPort(clientIP, clientPort))
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", handleMain).Methods(http.MethodGet)

	err := http.ListenAndServe(`:9090`, r)
	if err != nil {
		fmt.Println(err)
	}
}