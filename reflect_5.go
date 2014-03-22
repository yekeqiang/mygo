package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b := []byte(`{
    "Title":"go programming language",
    "Author":["john","ada","alice"],
    "Publisher":"qinghua",
    "IsPublished":true,
    "Price":99
  }`)
	//先创建一个目标类型的实例对象，用于存放解码后的值
	var inter interface{}
	err := json.Unmarshal(b, &inter)
	if err != nil {
		fmt.Println("error in translating,", err.Error())
		return
	}
	//要访问解码后的数据结构，需要先判断目标结构是否为预期的数据类型
	book, ok := inter.(map[string]interface{})
	//然后通过for循环一一访问解码后的目标数据
	if ok {
		for k, v := range book {
			switch vt := v.(type) {
			case float64:
				fmt.Println(k, " is float64 ", vt)
			case string:
				fmt.Println(k, " is string ", vt)
			case []interface{}:
				fmt.Println(k, " is an array:")
				for i, iv := range vt {
					fmt.Println(i, iv)
				}
			default:
				fmt.Println("illegle type")
			}
		}
	}
}
