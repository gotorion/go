package singleton

import (
	"sync"
)

type singleton struct {}
var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	// Do底层使用atomic保证func只执行一次，defer关键字保证unlock
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}