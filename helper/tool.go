package helper

import (
	"reflect"
	"time"
)

//UserTicker 自定义定时器，使用channel实现控制定时器的关闭
func UserTicker(dur int, callback interface{}, args ...interface{}) chan<- bool {
	ticker := time.NewTicker(time.Duration(dur) * time.Millisecond)
	fun := reflect.ValueOf(callback)
	if fun.Kind() != reflect.Func {
		panic("not a function")
	}
	vargs := make([]reflect.Value, len(args))
	for i, arg := range args {
		vargs[i] = reflect.ValueOf(arg)
	}

	stopChan := make(chan bool, 2) //这里多给一个方便函数本身调取chan
	go func(ticker *time.Ticker, fun reflect.Value, vargs []reflect.Value, stopChan <-chan bool) {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fun.Call(vargs)
			case stop := <-stopChan:
				if stop {
					return
				}
			}
		}
	}(ticker, fun, vargs, stopChan)

	return stopChan
}
