package main

// this is to store all the structures

type Otp struct {
	OneTimePassword string `json:"otp"`
}

type GetOtpRequest struct {
	UserId int64 `json:"userid"`
}

type GetOtpResponse struct {
	UserId int64 `json:"userid"`
	Otp
}

type GetVerifyOtpRequest struct {
	Otp
	UserId int64 `json:"userid"`
}

type GetVerifyOtpResponse struct {
	IsVerified bool `json:"isverified"`
	UserId int64 `json:"userid"`
}
