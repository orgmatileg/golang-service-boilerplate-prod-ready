package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetRandomString() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}

func GetRandomSixNumber() string {
	s := ""
	for i := 0; i < 6; i++ {
		rand.Seed(time.Now().UnixNano())
		min := 0
		max := 9
		s += strconv.Itoa(rand.Intn(max-min+1) + min)
	}
	if strings.HasPrefix(s, "0") {
		rand.Seed(time.Now().UnixNano())
		min := 1
		max := 9
		s = strings.Replace(s, "0", strconv.Itoa(rand.Intn(max-min+1)+min), 1)
	}
	return s
}

// Random Combinations for E-Catalog URL
var randomCombinationECatalog = [36]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func GetRandomCombinationsSubdomainECatalog() string {
	s := ""
	for i := 0; i < 5; i++ {
		rand.Seed(time.Now().UnixNano())
		min := 0
		max := 35
		indexRandomCombinations := rand.Intn(max-min+1) + min
		s = s + randomCombinationECatalog[indexRandomCombinations]
	}
	return s
}

func GetRandomPhoneNumber() string {
	s := ""
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		min := 0
		max := 9
		s += strconv.Itoa(rand.Intn(max-min+1) + min)
	}
	return s
}

func GetRandomReferralStore() string {
	s := ""
	for i := 0; i < 8; i++ {
		rand.Seed(time.Now().UnixNano())
		min := 0
		max := 35
		indexRandomCombinations := rand.Intn(max-min+1) + min
		s = s + randomCombinationECatalog[indexRandomCombinations]
	}
	return s
}
