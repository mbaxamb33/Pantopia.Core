package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomEmail generates a random email
func RandomEmail() string {
	return RandomString(8) + "@example.com"
}

// RandomName generates a random name
func RandomName() string {
	return strings.Title(RandomString(6))
}

// RandomPhoneNumber generates a random phone number
func RandomPhoneNumber() string {
	return "+40" + RandomString(8)
}

// RandomCompanyName generates a random company name
func RandomCompanyName() string {
	return strings.Title(RandomString(10)) + " Ltd"
}

// RandomAddress generates a random address
func RandomAddress() string {
	return RandomString(10) + " Street, City"
}

// RandomGoalName generates a random goal name
func RandomGoalName() string {
	goals := []string{
		"Increase Sales", "Improve Conversion Rate", "Expand Customer Base",
		"Enhance Team Productivity", "Reduce Customer Churn", "Increase Meeting Rates",
	}
	return goals[rand.Intn(len(goals))]
}

// RandomDescription generates a random goal description
func RandomDescription() string {
	descriptions := []string{
		"Focus on acquiring new clients.", "Enhance existing customer retention.",
		"Optimize marketing campaigns.", "Improve efficiency in sales operations.",
	}
	return descriptions[rand.Intn(len(descriptions))]
}

// RandomTargetValue generates a random target value
func RandomTargetValue() string {
	return strconv.Itoa(int(RandomInt(10, 100))) // Random target between 10 and 100
}
