package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
)

func compare(item1 interface{}, item2 interface{}) int {
	type1 := reflect.TypeOf(item1)
	type2 := reflect.TypeOf(item2)

	if type1.Kind() == reflect.Slice && type2.Kind() == reflect.Slice {
		// Both arrays
		arr1 := item1.([]interface{})
		arr2 := item2.([]interface{})

		minLength := len(arr1)
		if len(arr1) > len(arr2) {
			minLength = len(arr2)
		}
		for i := 0; i < minLength; i++ {
			compareValue := compare(arr1[i], arr2[i])
			if compareValue != 0 {
				return compareValue
			}
		}
		if len(arr1) > len(arr2) {
			return -1
		} else if len(arr1) < len(arr2) {
			return 1
		}

	} else if type1.Kind() != reflect.Slice && type2.Kind() == reflect.Slice {
		// 2 array
		return compare([]interface{}{item1.(float64)}, item2)
	} else if type1.Kind() == reflect.Slice && type2.Kind() != reflect.Slice {
		// 1 array
		return compare(item1, []interface{}{item2.(float64)})
	} else {
		// None array
		val1 := item1.(float64)
		val2 := item2.(float64)
		if val1 > val2 {
			return -1
		} else if val1 < val2 {
			return 1
		} else {
			return 0
		}
	}
	return 0
}

func bs[T interface{}](items []T, needle T, lo int, hi int) int {
	mid := lo + (hi-lo)/2

	res := compare(items[mid], needle)
	if hi <= lo || res == 0 {
		return hi
	}
	if res == -1 {
		return bs(items, needle, lo, mid)
	} else {
		return bs(items, needle, mid+1, hi)
	}
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	idx := 1

	packets := make([]interface{}, 0)
	for scanner.Scan() {
		line1 := scanner.Text()

		if line1 == "" {
			continue
		}
		parsedLine1 := make([]interface{}, 0)

		json.Unmarshal([]byte(line1), &parsedLine1)
		packets = append(packets, parsedLine1)
		idx++
	}

	sort.Slice(packets, func(i int, j int) bool {
		result := compare(packets[i], packets[j])
		if result > 0 {
			return true
		} else {
			return false
		}
	})

	firstDecoder := make([]interface{}, 0)
	firstDecoder = append(firstDecoder, []interface{}{float64(2)})
	firstDecoderIndex := bs([]interface{}(packets), interface{}(firstDecoder), 0, len(packets)-1)

	secondDecoder := make([]interface{}, 0)
	secondDecoder = append(secondDecoder, []interface{}{float64(6)})
	secondDecoderIndex := bs([]interface{}(packets), interface{}(secondDecoder), 0, len(packets)-1)

	fmt.Println((firstDecoderIndex + 1) * (secondDecoderIndex + 2))
}
