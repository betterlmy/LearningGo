package main

import "fmt"

//map类似于字典 key-value
//Map是无序的

//Map的定义
//var map1 map[key_type]value_type

func main() {
	capitalMap := make(map[string]string)

	//向map中插入元素
	capitalMap["France"] = "巴黎"
	capitalMap["China"] = "北京"
	capitalMap["India"] = "新德里"

	//mapTest1(capitalMap)
	capitalMap = deleteOneCountry(capitalMap, "India")
	printMap(capitalMap)

}

func printMap(capitalMap map[string]string) {
	for countryName, capitalName := range capitalMap {
		fmt.Printf("%s首都是%s\n", countryName, capitalName)
	}
}

// 创建一个国家首都的map
func mapTest1(capitalMap map[string]string) {
	for countryName, capitalName := range capitalMap {
		fmt.Println(countryName, "的首都是", capitalName)
	}

	//从Map中判断元素(键)是否存在
	capitalName, isExist := capitalMap["America"]
	if isExist {
		fmt.Println("存在首都是", capitalName)
	} else {
		fmt.Println("不存在")
	}

}

func deleteOneCountry(capitalMap map[string]string, countryName string) map[string]string {
	_, isExist := capitalMap[countryName]
	if isExist {
		fmt.Printf("删除了%s,它的首都是%s\n", countryName, capitalMap[countryName])
		delete(capitalMap, countryName)
	} else {
		fmt.Println("不存在该国家")
	}

	return capitalMap
}
