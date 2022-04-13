package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	Port = ":8080"
)

var (
	otpKeys = make(map[string]string, 0)
)

func main() {
	handler := http.HandlerFunc(getOtpHandler)
	http.Handle("/get-otp", handler)

	verifyHandler := http.HandlerFunc(getVerifyOtpHandler)
	http.Handle("/get-verify", verifyHandler)

	fmt.Printf("server is running on: %s", Port)
	if err := http.ListenAndServe(Port, nil); err != nil {
		log.Fatal(err)
	}
}
