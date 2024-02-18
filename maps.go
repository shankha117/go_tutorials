package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Message struct {
	Symbol string `json:"symbol"`

	Size string `json:"size"`

	Date string `json:"date"`
	Last string `json:"last"`
}

func main() {
	// fmt.Println("Welcome to Maps !")
	// a := []string{"A", "B", "C", "D", "E"}
	// i := 2
	// fmt.Println(a)
	// fmt.Println(a[i:], []string{"K", "M"})
	// ret := copy(a[i:], []string{"K", "M", "L"})
	// fmt.Println(ret)

	// fmt.Println(a)

	// testMap := map[string]string{"type": "trade", "symbol": "SPY", "exch": "D", "price": "411.225", "size": "100", "cvol": "26494968", "date": "1618248244112", "last": "411.225"}

	// data_map := map[string]string{"symbol": testMap["symbol"], "size": testMap["size"], "date": testMap["date"], "last": testMap["last"]}

	// fmt.Println(data_map)

	// b, err := json.Marshal(data_map)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(b, reflect.TypeOf(b))
	// m := map[string]string{}
	// if err := json.Unmarshal(b, &m); err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%q", m)

	// data_map := map[string]string{"Date": "1618344600001", "Last": "412.86", "Size": "0", "Symbol": "SPY"}
	// x := ""
	// fmt.Println(data_map, x)
	// for _, val := range data_map {
	// 	x = x + val + ","

	// }
	// x = strings.TrimSuffix(x, ",")
	// fmt.Println(x, reflect.TypeOf(x))

	// b := []byte("1,NULL,2020-08-06 07:53:09")
	// fmt.Println(b)
	// fmt.Printf("this the data %s", b)

	st := Message{
		Symbol: "ss",
		Size:   "sss",
		Date:   "ssssss",
		Last:   "sdd",
	}

	ans := convert(st)
	fmt.Println(ans)
	fmt.Printf("%s", ans)

	// fmt.Println(st, reflect.TypeOf(st))

	// v := reflect.ValueOf(st)

	// values := make([]interface{}, v.NumField())

	// for i := 0; i < v.NumField(); i++ {
	// 	values[i] = v.Field(i).Interface()

	// }

	// fmt.Println(reflect.TypeOf(values))

	// str := ""

	// for index := 0; index < len(values); index++ {
	// 	str1 := fmt.Sprintf("%v", values[index])
	// 	str += str1 + ","
	// }
	// fmt.Println(str)
	// str = strings.TrimSuffix(str, ",")
	// fmt.Println(str)

	// b := []byte(str)
	// fmt.Println(b)
	// fmt.Printf("this the data %s", b)
}

func convert(st Message) []byte {
	fmt.Println(st, reflect.TypeOf(st))

	v := reflect.ValueOf(st)

	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()

	}

	fmt.Println(reflect.TypeOf(values))

	str := ""

	for index := 0; index < len(values); index++ {
		str1 := fmt.Sprintf("%v", values[index])
		str += str1 + ","
	}
	fmt.Println(str)
	str = strings.TrimSuffix(str, ",")
	fmt.Println(str)

	b := []byte(str)
	fmt.Println(b)
	fmt.Printf("this the data %s", b)
	return b

}
