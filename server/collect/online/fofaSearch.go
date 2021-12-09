package online

import (
	"errors"
	"fmt"
	autocode2 "github.com/779789571/gin-vue-admin/server/model/autocode"
	"github.com/779789571/gin-vue-admin/server/service/autocode"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	defaultFofaAPIUrl = "https://fofa.so/api/v1/search/all?"
)

type Fofa struct {
	email    []byte
	key      []byte
	limit    int
	proxy    string
	ruleType int
	*http.Client
}

func (fofa *Fofa) GetRules() (err error, list []autocode2.Rules) {
	err, list = autocode.GetValidRulesByType(fofa.ruleType)

	return err, list
}

func (*Fofa) Runner(intervalTime float64, proxy string) {
	//global.GVA_LOG.Info(fmt.Sprintf("inFofa : %v", intervalTime))

}

func (*Fofa) GenerateKeywords(list []autocode2.Rules) (err error, ruleList []string) {

	if len(list) == 0 {
		return errors.New("No Fofa rules found"), nil
	}
	for _, v := range list {
		if *v.IsDynamic {
			fmt.Printf("动态的规则：%v, 生成动态规则中...\n", v.Content)
			err, dynamicRulesList := GenerateDynamicRules(v)
			if err != nil {
				print(err)
			}
			fmt.Printf("动态规则生成完成：%v ", dynamicRulesList)

		} else {
			//fmt.Printf("静态的规则：%v\n",v.Content)
			ruleList = append(ruleList, v.Content)
		}
	}
	return nil, ruleList
}

//动态规则与配置信息匹配，生成规则
func GenerateDynamicRules(rule autocode2.Rules) (err error, dynamicRulesList []string) {
	err, requiredInfos := autocode.GetRulesByRuleName(rule.RuleName)
	if err != nil {
		print(err)
	}

	var dynamicRules []string //单个配置信息生成的动态规则
	dynamicRules = append(dynamicRules, rule.Content)
	//遍历配置信息
	for _, v := range requiredInfos {
		//fmt.Printf("配置信息： %v", v)
		ruleNameList := strings.Split(v.UseRules, "|")
		if IsContain(ruleNameList, rule.RuleName) {
			//fmt.Printf("%v包含%v\n",ruleNameList, rule.RuleName)
			//替换规则内占位符 `域名`
			if strings.Contains(rule.Content, "`域名`") {
				if len(v.Domain) == 0 {
					fmt.Printf("配置信息%v有问题！没有域名信息", v.UseRules)
					continue
				}
				dynamicRules = ReplaceStringWithInfos(dynamicRules, "`域名`", v.Domain)
				//fmt.Printf("生成的动态域名规则：%v", dynamicRules)
			}

			if strings.Contains(rule.Content, "`网站标题`") {
				//替换规则内占位符 `域名`
				if strings.Contains(rule.Content, "`网站标题`") {
					if len(v.Title) == 0 {
						fmt.Printf("配置信息%v有问题！没有网站标题信息", v.UseRules)
						continue
					}
					dynamicRules = ReplaceStringWithInfos(dynamicRules, "`网站标题`", v.Title)
					//fmt.Printf("生成的动态域名规则：%v", dynamicRules)
				}
			}
			//动态则生成前一天时间
			if *v.TimeAutoCreate {
				t := time.Now()
				oldTime := t.AddDate(0, 0, -1)
				afterTime := oldTime.Format("2006-01-02")
				dynamicRules = AddAutoTimeBefore(dynamicRules, " && after=\""+afterTime+"\" ")
			}
			if strings.Contains(rule.Content, "`https证书`") {
				if strings.Contains(rule.Content, "`https证书`") {
					if len(v.Cert) == 0 {
						fmt.Printf("配置信息%v有问题！没有https证书信息", v.UseRules)
						continue
					}
					dynamicRules = ReplaceStringWithInfos(dynamicRules, "`https证书`", v.Cert)
					//fmt.Printf("生成的动态域名规则：%v", dynamicRules)
				}
			}

			if strings.Contains(rule.Content, "`Ip或ip段`") {
				if strings.Contains(rule.Content, "`Ip或ip段`") {
					if len(v.Ip) == 0 {
						fmt.Printf("配置信息%v有问题！没有Ip或ip段信息", v.UseRules)
						continue
					}
					dynamicRules = ReplaceStringWithInfos(dynamicRules, "`Ip或ip段`", v.Ip)
					//fmt.Printf("生成的动态域名规则：%v", dynamicRules)
				}
			}

			if strings.Contains(rule.Content, "`协议`") {
				if strings.Contains(rule.Content, "`协议`") {
					if len(v.Protocol) == 0 {
						fmt.Printf("配置信息%v有问题！没有Ip或ip段信息", v.UseRules)
						continue
					}
					dynamicRules = ReplaceStringWithInfos(dynamicRules, "`协议`", v.Protocol)
					//fmt.Printf("生成的动态域名规则：%v", dynamicRules)
				}
			}

			if strings.Contains(rule.Content, "`icp备案号`") {
				if strings.Contains(rule.Content, "`icp备案号`") {
					if len(v.Icp) == 0 {
						fmt.Printf("配置信息%v有问题！没有icp备案号信息", v.UseRules)
						continue
					}
					dynamicRules = ReplaceStringWithInfos(dynamicRules, "`icp备案号`", v.Icp)
					//fmt.Printf("生成的动态域名规则：%v", dynamicRules)
				}
			}

			if strings.Contains(rule.Content, "`服务器状态码`") {
				if strings.Contains(rule.Content, "`服务器状态码`") {
					if len(v.WebStatusCode) == 0 {
						fmt.Printf("配置信息%v有问题！没有服务器状态码信息", v.UseRules)
						continue
					}
					dynamicRules = ReplaceStringWithInfos(dynamicRules, "`服务器状态码`", v.WebStatusCode)
					//fmt.Printf("生成的动态域名规则：%v", dynamicRules)
				}
			}

			if strings.Contains(rule.Content, "`HTTP请求头`") {
				if strings.Contains(rule.Content, "`HTTP请求头`") {
					if len(v.Header) == 0 {
						fmt.Printf("配置信息%v有问题！没有HTTP请求头信息", v.UseRules)
						continue
					}
					dynamicRules = ReplaceStringWithInfos(dynamicRules, "`HTTP请求头`", v.Header)
					//fmt.Printf("生成的动态域名规则：%v", dynamicRules)
				}
			}

			if strings.Contains(rule.Content, "`国家`") {
				if strings.Contains(rule.Content, "`国家`") {
					if len(v.Country) == 0 {
						fmt.Printf("配置信息%v有问题！没有国家信息", v.UseRules)
						continue
					}
					dynamicRules = ReplaceStringWithInfos(dynamicRules, "`国家`", v.Country)
					//fmt.Printf("生成的动态域名规则：%v", dynamicRules)
				}
			}

			if strings.Contains(rule.Content, "`城市`") {
				if strings.Contains(rule.Content, "`城市`") {
					if len(v.City) == 0 {
						fmt.Printf("配置信息%v有问题！没有城市信息", v.UseRules)
						continue
					}
					dynamicRules = ReplaceStringWithInfos(dynamicRules, "`城市`", v.City)
					//fmt.Printf("生成的动态域名规则：%v", dynamicRules)
				}
			}

			if strings.Contains(rule.Content, "`省份`") {
				if strings.Contains(rule.Content, "`省份`") {
					if len(v.Province) == 0 {
						fmt.Printf("配置信息%v有问题！没有省份信息", v.UseRules)
						continue
					}
					dynamicRules = ReplaceStringWithInfos(dynamicRules, "`省份`", v.Province)
					//fmt.Printf("生成的动态域名规则：%v", dynamicRules)
				}
			}
		}
		if len(dynamicRules) > 0 {
			for _, v := range dynamicRules {
				dynamicRulesList = append(dynamicRulesList, v)
			}
		} else {
			return errors.New(fmt.Sprintf("未找到动态规则 %v 匹配的配置信息", rule.Content)), nil
		}
	}
	return nil, dynamicRulesList
}

func (fofa *Fofa) GetAPi() (err error) {
	err, username, key, limitTime := autocode.GetApiToken(1)
	if err != nil {
		return err
	}
	switch username.(type) {
	case string:
		if username.(string) != "" {
			if limitTime.(int) <= 100 {
				return errors.New("fofa账号次数限制已小于100，暂停使用。")
			}

		} else {
			return errors.New("fofa api 没有用户名")
		}

		//case []string: //todo 处理多账号情况

	}
	fofa.email = []byte(username.(string))
	fofa.key = []byte(key.(string))
	fofa.limit = limitTime.(int)
	return nil
}

func (*Fofa) RunTask() {}

func (*Fofa) SaveResult() {}

func FofaSearch(proxy string, wg *sync.WaitGroup) (err error) {
	wg.Add(1)
	var fofa Fofa
	fofa.proxy = proxy
	fofa.ruleType = 0
	//todo  GetAPi()未处理多账号问题
	err = fofa.GetAPi()
	if err != nil {
		fmt.Printf("fofa API 出现问题 ： %v", err)
		return err
	} else {
		fmt.Printf("%v", fofa)
	}
	err, rulelist := fofa.GetRules()
	if err != nil {
		return err
	}
	user, err := fofa.UserInfo()
	if err != nil {
		return err
	}
	fmt.Printf("%v", &user)
	err, _ = fofa.GenerateKeywords(rulelist)
	if err != nil {
		print(err)
	}

	//username,api := fofa.GetAPi()
	//if username != "" && api!= ""{
	//	fofa.username = username
	//	fofa.token = api
	//}
	//fmt.Printf("%v", list)

	wg.Done()
	return nil
}
