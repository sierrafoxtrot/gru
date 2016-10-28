package x

import (
	"flag"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

var debug = flag.Bool("debug", false, "Whether to print debug info")

func Debug(log interface{}) {
	if *debug {
		fmt.Println(log)
	}
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Truncate(f float64) float64 {
	fmt.Println(f * 100)
	return float64(int(f*100)) / 100
}

type Claims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}
