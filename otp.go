package main

import (
	"crypto/sha1"
	"fmt"
	"strconv"
	"strings"
)

func getOtp(userId int64, timeStamp int64) (string, error) {
	if userId <= 0 || timeStamp <= 0 {
		return "", fmt.Errorf("invalid userid: %d, timestamp: %d", userId, timeStamp)
	}

	// compute the hash of userid and timestamp
	otpKey := getOtpUniqueKey(userId, timeStamp)
	hasher := sha1.New()
	hasher.Write([]byte(otpKey))
	hashedOtpKey := hasher.Sum(nil)

	// reformat ky
	hashedOtpKeySliced := strings.Split(fmt.Sprintf("%x", hashedOtpKey), "")
	prefixSixHashed := strings.Join(hashedOtpKeySliced[:6], "")

	// save entry into memcache
	strKey := strconv.Itoa(int(userId))
	setOtpKey(strKey, prefixSixHashed)

	// return hashed key
	return prefixSixHashed, nil
}

func isOtpVerified(otp string, userId int64) bool {
	// if exists in map, return true
	strKey := strconv.Itoa(int(userId))
	otpInCache, err := getOtpInCacheByKey(strKey)
	if err != nil { // nil hit
		return false
	}

	return otpInCache == otp
}

// this is what is used to hash into the otp
func getOtpUniqueKey(userId, timestamp int64) string {
	return fmt.Sprintf("%d@%d", userId, timestamp)
}
