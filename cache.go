package main

import "fmt"

// directly set, no matter what values there could be, use latest as valid
func setOtpKey(key string, value string) {
	otpKeys[key] = value
}

// key -> userid
func getOtpInCacheByKey(key string) (string, error) {
	if value, ok := otpKeys[key]; ok {
		return value, nil
	}
	return "", fmt.Errorf("nil hit")
}
