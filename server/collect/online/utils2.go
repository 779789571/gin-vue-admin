package online

import "strings"

func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

//将规则内占位符生成可用规则
func ReplaceStringWithInfos(list []string, string2 string, infos string) (ruleList []string) {

	infoList := strings.Split(infos, "|")
	for _, v := range list {
		for _, vv := range infoList {
			list := strings.Replace(v, string2, "\""+vv+"\"", 1)
			ruleList = append(ruleList, list)
		}
	}
	return ruleList
}

func AddAutoTimeBefore(ruleList []string, afterTime string) (withTimeruleList []string) {
	for _, v := range ruleList {
		withTimeruleList = append(withTimeruleList, v+afterTime)
	}
	return withTimeruleList
}
