package main

import (
	"log"
	"sync"
)

type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		singleInstance = &Singleton{}
	})
	return singleInstance
}

func TestGetInstance() {
	singleInstance1 := GetInstance()
	singleInstance2 := GetInstance()
	if singleInstance1 == singleInstance2 {
		log.Println("same")
	} else {
		log.Println("different")
	}
}
