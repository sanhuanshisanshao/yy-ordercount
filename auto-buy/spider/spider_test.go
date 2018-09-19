package spider

import (
	"fmt"
	"testing"
)

func TestGetRemainFreeTimes(t *testing.T) {
	str := conf
	times, err := GetRemainFreeTimes(str)
	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println(times)
}
