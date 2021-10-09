package main

import (
	"fmt"
	"strings"
)

func main() {

	str := ""
	num := 0
	fmt.Scan(&str)
	fmt.Scan(&num)
	str = strings.ToUpper(str)

	strList := strings.SplitN(str, "-", 2)

	str1 := strList[0]

	if len(strList) == 1 {
		fmt.Println(str1)
		return
	}

	str2 := strings.ReplaceAll(strList[1], "-", "")

	strResList := []string{str1}
	//asdasd-asd-asd-asd
	for i := 0; i < len(str2); i += num {
		if i+num > len(str2) {
			strResList = append(strResList, str2[i:])
			break
		}
		strResList = append(strResList, str2[i:i+num])

		if i == len(str2) {
			break
		}
	}
	strRes := strings.Join(strResList, "-")
	fmt.Println(strRes)
}

//func main() {
//	IP := ""
//	fmt.Scan(&IP)
//
//	ip := net.ParseIP(IP)
//	if ip == nil {
//		fmt.Println("Neither")
//		return
//	}
//
//	if strings.Contains(IP, ".") {
//		ipList := strings.Split(IP, ".")
//		if len(ipList) != 4 {
//			fmt.Println("Neither")
//			return
//		}
//		for _, v := range ipList {
//			if strings.HasPrefix(v, "0") {
//				fmt.Println("Neither")
//				return
//			}
//			num, err := strconv.Atoi(v)
//			if err != nil {
//				fmt.Println("Neither")
//				return
//			}
//			if num > 255 || num < 0 {
//				fmt.Println("Neither")
//				return
//			}
//		}
//		fmt.Println("IPv4")
//	} else if strings.Contains(IP, ":") {
//
//		if strings.Contains(IP,"::"){
//			fmt.Println("Neither")
//			return
//		}
//
//		ipList := strings.Split(IP, ":")
//		if len(ipList) != 8 {
//			fmt.Println("Neither")
//			return
//		}
//
//		for _, v := range ipList {
//			if strings.HasPrefix(v, "0000") {
//				fmt.Println("Neither")
//				return
//			}
//		}
//		fmt.Println("IPv6")
//	} else {
//		fmt.Println("Neither")
//	}
//}
