package main

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
	UserId int64 `json:"userid"`
	Otp
}

type GetVerifyOtpResponse struct {
	IsVerified bool  `json:"isverified"`
	UserId     int64 `json:"userid"`
}
