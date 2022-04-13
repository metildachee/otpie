package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func getOtpHandler(w http.ResponseWriter, r *http.Request) {
	// do some checking to make sure the request is valid
	// --> check if GET request, check your routing
	otpRequest := &GetOtpRequest{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(otpRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
