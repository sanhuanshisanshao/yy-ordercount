package util

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestGetToday(t *testing.T) {
	fmt.Println(GetToday())
}

func TestSubtraction(t *testing.T) {
	fmt.Println(Subtraction("80.0", "-2.0"))
}

func TestTimeConvert(t *testing.T) {
	s := "2018-08-31 16:00"
	tim, _ := time.Parse("2006-01-02 15:04", s)
	//tim.Unix()
	fmt.Println(tim.After(time.Now().Add(8 * time.Hour).UTC()))
	fmt.Println(tim)
	fmt.Println(time.Now().Add(8 * time.Hour).UTC())
}

func TestGetExternalIP(t *testing.T) {
	fmt.Println(strings.Replace(GetExternalIP(), "\n", "", -1))
}
