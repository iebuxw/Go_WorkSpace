package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("Hello World!")

	arr1 := [...]int{1, 2, 3}
	is_in, _ := in_array(arr1, 2)
	fmt.Println(is_in)

	slice1 := []string{"1", "2", "3"}
	imp_str := Implode(",", slice1)
	fmt.Println("imp_str=", imp_str)
}

func in_array(haystack interface{}, needle interface{}) (bool, error) {
	sVal := reflect.ValueOf(haystack)
	kind := sVal.Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < sVal.Len(); i++ {
			if sVal.Index(i).Interface() == needle {
				return true, nil
			}
		}
		return false, nil
	}
	return false, nil
}

func Implode(glue string, pieces []string) string {
	return strings.Join(pieces, glue)
}
