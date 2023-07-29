package idlmanager

import (
	gateway "Gateway/biz/model/Gateway"
	"errors"
	"sync"
)

var ServiceNameMap = make(map[string]gateway.ServiceInfo)

var mapMutex = &sync.Mutex{}
var ioMutex = &sync.Mutex{}

func GetService(serviceName string) *gateway.ServiceInfo {
	mapMutex.Lock() // 获取锁定
	defer mapMutex.Unlock()
	service, ok := ServiceNameMap[serviceName]
	if !ok {
		err := errors.New("service is not found！")
		panic(err)
	}
	return &service
}

func GetIdlPath(serviceName string) string {
	service := GetService(serviceName)
	return service.ServiceIdlName
}

func AddService(service gateway.ServiceInfo) {
	mapMutex.Lock() // 获取锁定
	defer mapMutex.Unlock()
	if _, ok := ServiceNameMap[service.ServiceName]; !ok {
		ServiceNameMap[service.ServiceName] = service
	} else {
		err := errors.New("service" + service.ServiceName + "has been added！")
		panic(err)
	}
}

func DeleteService(serviceName string) {
	mapMutex.Lock() // 获取锁定
	defer mapMutex.Unlock()
	delete(ServiceNameMap, serviceName)
}

func UpdateService(service gateway.ServiceInfo) {
	mapMutex.Lock() // 获取锁定
	defer mapMutex.Unlock()
	if _, ok := ServiceNameMap[service.ServiceName]; ok {
		ServiceNameMap[service.ServiceName] = service
	} else {
		AddService(service)
	}
}

func GetAllService() []*gateway.ServiceInfo {
	var services []*gateway.ServiceInfo
	mapMutex.Lock() // 获取锁定
	defer mapMutex.Unlock()
	for k := range ServiceNameMap {
		service := ServiceNameMap[k]
		services = append(services, &service)
	}
	return services
}
