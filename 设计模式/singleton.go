/**
 * 单例模式是一种创建型设计模式， 让你能够保证一个类只有一个实例， 并提供一个访问该实例的全局节点。
 * https://refactoringguru.cn/design-patterns/singleton
 */
package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct{}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
