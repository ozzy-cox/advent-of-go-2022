package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

// func printRecursively(arr []interface{}, level int) {
// 	for _, v := range arr {
// 		rt := reflect.TypeOf(v)
// 		if rt.Kind() == reflect.Slice {
// 			print("\n")
// 			val, ok := v.([]interface{})
// 			if ok {
// 				printRecursively(val, level+1)
// 			}
// 		} else {
// 			fmt.Print(v)
// 		}
// 	}
// }

func compare(item1 interface{}, item2 interface{}) int {
	type1 := reflect.TypeOf(item1)
	type2 := reflect.TypeOf(item2)

	if type1.Kind() == reflect.Slice && type2.Kind() == reflect.Slice {
		// Both arrays
		arr1 := []interface{}(item1.([]interface{}))
		arr2 := []interface{}(item2.([]interface{}))
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

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	sum := 0
	idx := 1
	for scanner.Scan() {
		line1 := scanner.Text()

		if line1 == "" {
			continue
		}
		scanner.Scan()
		line2 := scanner.Text()
		parsedLine1 := make([]interface{}, 0)
		parsedLine2 := make([]interface{}, 0)

		json.Unmarshal([]byte(line1), &parsedLine1)
		json.Unmarshal([]byte(line2), &parsedLine2)

		if compare(parsedLine1, parsedLine2) == 1 {
			sum += idx
		}
		idx++
	}
	fmt.Println(sum)
}
