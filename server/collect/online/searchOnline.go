package online

import "sync"

func SearchRunner(interval float64, proxy string) {
	//在这里设置循环跑脚本
	var wg sync.WaitGroup

	FofaSearch(proxy, &wg)

	wg.Wait()
	//获取结果

	//处理结果

	//发送结果

}
