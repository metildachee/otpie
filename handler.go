package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func validateCommonRequest(w http.ResponseWriter, r *http.Request, path string) bool {
	if r.URL.Path != path {
		http.Error(w, "404 not found", http.StatusNotFound)
		return false
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return false
	}

	if headerContentType := r.Header.Get("Content-Type"); headerContentType != "application/json" {
		http.Error(w, "invalid parsing type", http.StatusUnsupportedMediaType)
	}
	return true
}

func getOtpHandler(w http.ResponseWriter, r *http.Request) {
	if valid := validateCommonRequest(w, r, "/get-otp"); !valid {
		return
	}

	otpRequest := &GetOtpRequest{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(otpRequest); err != nil {
		var unmarshalErr *json.UnmarshalTypeError
		if errors.As(err, &unmarshalErr) {
			http.Error(w, fmt.Sprintf("bad request, wrong type provided for field: %v, ", unmarshalErr.Field), http.StatusBadRequest)
		} else {
			http.Error(w, fmt.Sprintf("bad request: %v", err.Error()), http.StatusBadRequest)
		}
		return
	}

	eventTime := time.Now().Unix()
	otp, err := getOtp(otpRequest.UserId, eventTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := &GetOtpResponse{
		UserId: otpRequest.UserId,
		Otp:    Otp{otp},
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonResp)
}

func getVerifyOtpHandler(w http.ResponseWriter, r *http.Request) {
	if valid := validateCommonRequest(w, r, "/get-verify"); !valid {
		return
	}

	getVerifyOtpRequest := &GetVerifyOtpRequest{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(getVerifyOtpRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	verificationStatus := isOtpVerified(getVerifyOtpRequest.OneTimePassword, getVerifyOtpRequest.UserId)
	res := &GetVerifyOtpResponse{
		UserId:     getVerifyOtpRequest.UserId,
		IsVerified: verificationStatus,
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonResp)
}
