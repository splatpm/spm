package subhuman

import (
	"fmt"
	"testing"
)

func TestHumanTimeColon(t *testing.T) {
	for k, v := range map[int64]string{
		32:   "00:00:32",
		64:   "00:01:04",
		3667: "01:01:07",
	} {
		ret := HumanTimeColon(k)
		if ret != v {
			t.Error("TestHumanTimeColon  : Expected '", v, "' but got ", ret)
		} else {
			fmt.Printf("TestHumanTimeColon  : Expected %10s and got %10s SUCCESS\n",
				v, ret)
		}
	}
}

func TestHumanTimeConcise(t *testing.T) {
	for k, v := range map[int64]string{
		32:   "00h00m32s",
		64:   "00h01m04s",
		3667: "01h01m07s",
	} {
		ret := HumanTimeConcise(k)
		if ret != v {
			t.Error("TestHumanTimeConcise: Expected '", v, "' but got ", ret)
		} else {
			fmt.Printf("TestHumanTimeConcise: Expected %10s and got %10s SUCCESS\n",
				v, ret)
		}
	}
}
