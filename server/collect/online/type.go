package online

import (
	"sync"
)

type SearchResult struct {
	lock        sync.Mutex
	Results     []string
	FofaResults []FofaResults
}

func (searchResult *SearchResult) Add(f interface{}) {
	if data, ok := f.(FofaResults); ok {
		searchResult.lock.Unlock()
		searchResult.FofaResults = append(searchResult.FofaResults, data)
		searchResult.lock.Lock()

	}

}
