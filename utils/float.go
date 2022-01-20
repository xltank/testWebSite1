package utils

import (
	"bytes"
	"fmt"
	"math"
)

func ToFixed(f float64, p int) float64 {
	r := math.Pow10(p)
	return math.Round(f*r) / r
}

func ToFixedStr(f float64, p int) string {
	r := math.Pow10(p)
	return fmt.Sprintf("%.2f", math.Round(f*r)/r)
}

func StringConcat(str1, str2 string) string {
	var buf bytes.Buffer
	buf.WriteString(str1)
	buf.WriteString(str2)
	str := buf.String()
	return str
}
