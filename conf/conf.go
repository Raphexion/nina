package conf

import (
	"nina/backend"
	"sync"
)

var (
	lock     = &sync.Mutex{}
	instance *conf
)

type conf struct {
	back backend.Backend
}

func SetBackend(back backend.Backend) {
	instance := getInstance()
	instance.back = back
}

func GetBackend() backend.Backend {
	instance := getInstance()
	return instance.back
}

func getInstance() *conf {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &conf{}
	}

	return instance
}
