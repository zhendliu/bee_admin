package util

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func GetNegative (s string) float64{
	er :=fmt.Sprintf("%b",HexDec(s))

	shi :=Str2DEC(er[1:])
	a :=fmt.Sprintf("%b",shi-1)
	return float64(0-ByBitNegation(a))
}
//十六转十进制
func HexDec(h string) (n int64) {
	s := strings.Split(strings.ToUpper(h), "")
	l := len(s)
	i := 0
	d := float64(0)
	hex := map[string]string{"A": "10", "B": "11", "C": "12", "D": "13", "E": "14", "F": "15"}
	for i = 0; i < l; i++ {
		c := s[i]
		if v, ok := hex[c]; ok {
			c = v
		}
		f, err := strconv.ParseFloat(c, 10)
		if err != nil {
			log.Println("Hexadecimal to decimal error:", err.Error())
			return -1
		}
		d += f * math.Pow(16, float64(l-i-1))
	}
	return int64(d)
}

func Str2DEC(s string) (num int) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (int(s[l-i-1]) & 0xf) << uint8(i)
	}
	return
}

func ByBitNegation(s string)(num int){

	val :=""
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == "0"{
			val = val+"1"
		}else{
			val = val+"0"
		}
	}
	return Str2DEC(val)
}

