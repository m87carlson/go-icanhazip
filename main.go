package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	port := "8080"
	fmt.Printf("running http server on port %s...\n", port)
	http.HandleFunc("/", iCanHazIp)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

type Client struct {
	Ip string `json:"ip"`
}

func iCanHazIp(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	c := Client{}

	if r.Header.Get("X-Forwarded-For") != "" {
		c.Ip = r.Header.Get("X-Forwarded-For")
		fmt.Println("no x-forwarded-for, next...")
	} else if r.Header.Get("X-Real-IP") != "" {
		fmt.Println("no x-real-ip, next...")
		c.Ip = r.Header.Get("X-Real-IP")
	} else {
		fmt.Println("falling back to remote addr")
		c.Ip = strings.Split(r.RemoteAddr, ":")[0]
	}

	j, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
