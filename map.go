package main

import "fmt"

func main() {
	//Map 是一种无序的键值对的集合。1、统一类型；2、键值对；3、无序
	var countryCapitalMap map[string]string /*创建集合，默认 map 是 nil */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	fmt.Println(capital)
	fmt.Println(ok)
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	// 直接创建
	countryCapitalMap2 := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	//添加元素
	countryCapitalMap2["aa"] = "bbbbbb"
	countryCapitalMap2["aa"] = "bbbbbb222"

	//删除元素
	delete(countryCapitalMap2, "France")
	fmt.Println(countryCapitalMap2)
}
