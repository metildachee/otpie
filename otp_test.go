package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetOtp(t *testing.T) {
	var (
		timestamp = time.Now().Unix()
		userId    = int64(438635847365943)
	)

	otp, err := getOtp(userId, timestamp)
	fmt.Printf("otp: %s, err: %v\n", otp, err)
	assert.NoErrorf(t, err, "got err when getting otp")
	assert.NotEqualValues(t, "", otp, "otp is empty")
	assert.Len(t, otp, 6, "otp is more than 6 characters")
}
