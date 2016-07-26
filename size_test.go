package subhuman

import (
	"fmt"
	"testing"
)

func TestHumanSize(t *testing.T) {
	for k, v := range map[int64]string{
		32:         "32B",
		3232:       "3.00KB",
		323232:     "315.00KB",
		32323232:   "30.00MB",
		3232323232: "3.00GB"} {
		ret := HumanSize(k)
		if ret != v {
			t.Error("TestHumanSize       : Expected '", v, "' but got ", ret)
		} else {
			fmt.Printf("TestHumanSize       : Expected %10s and got %10s SUCCESS\n",
				v, ret)
		}
	}
}
