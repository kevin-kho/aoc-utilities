package common

import (
	"bytes"
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadInput(filePath string) ([]byte, error) {
	var data []byte

	data, err := os.ReadFile(filePath)
	if err != nil {
		return data, err
	}

	return data, nil

}

func ParseIntArray(data []byte) ([]int, error) {
	var res []int

	for entry := range strings.SplitSeq(string(data), ",") {
		entry = strings.TrimSpace(entry)

		i, err := strconv.Atoi(entry)
		if err != nil {
			return res, err
		}

		res = append(res, i)

	}

	return res, nil

}

func TrimNewLineSuffix(byteArr []byte) []byte {
	return bytes.TrimSuffix(byteArr, []byte{10})
}

func IntPow(x int, pow int) int {
	return int(math.Pow(float64(x), float64(pow)))
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
