package collect

import (
	"github.com/779789571/gin-vue-admin/server/collect/online"
	"testing"
)

func TestSearchRunner(t *testing.T) {
	//var wg sync.WaitGroup
	var interbval float64 = 2
	var proxy = "http://127.0.0.1:7890"
	print("123123")
	online.SearchRunner(interbval, proxy)

}
