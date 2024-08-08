package main

import "fmt"

//map类型名称： map[key_data_type]value_data_type
//Map 最重要的一点是通过 key 来快速检索数据
func main() {
	//Map 是一种无序的键值对的集合。1、统一类型；2、键值对；3、无序
	var countryCapitalMap map[string]string /*创建集合，默认 map 是 nil */

	//为什么var之后还要make，map是引用类型，在声明dataMap后并未初始化它
	//所以它的值是nil, 不指向任何内存地址。需要通过make方法分配确定的内存地址
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
	capital, ok := countryCapitalMap["American"]
	fmt.Println(capital)
	fmt.Println(ok)
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	// 直接创建（推荐，不过自动扩容会有性能问题）
	countryCapitalMap2 := map[string]string{
		"France": "Paris",
		"Italy": "Rome",
		"Japan": "Tokyo",
		"India": "New delhi",
	}
	//添加元素
	countryCapitalMap2["aa"] = "bbbbbb"
	countryCapitalMap2["aa"] = "bbbbbb222"

	//删除元素
	delete(countryCapitalMap2, "France")
	fmt.Println(countryCapitalMap2)

	//map是引用类型，所以改变countryCapitalMap3就是改变countryCapitalMap2
	countryCapitalMap3 := countryCapitalMap2
	countryCapitalMap3["aa"] = "bbbbbb2"
	fmt.Println(countryCapitalMap2)
	fmt.Println(countryCapitalMap3)

}
