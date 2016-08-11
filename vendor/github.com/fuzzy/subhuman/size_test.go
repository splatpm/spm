package subhuman

import (
	"fmt"
	"testing"
)

func TestHumanSize(t *testing.T) {
	for k, v := range map[int64]string{
		32:         "32B",
		3232:       "3.16KB",
		323232:     "315.66KB",
		32323232:   "30.83MB",
		3232323232: "3.01GB"} {
		ret := HumanSize(k)
		if ret != v {
			t.Error("TestHumanSize       : Expected '", v, "' but got ", ret)
		} else {
			fmt.Printf("TestHumanSize       : Expected %10s and got %10s SUCCESS\n",
				v, ret)
		}
	}
}
