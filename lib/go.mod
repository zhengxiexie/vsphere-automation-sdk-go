module github.com/vmware/vsphere-automation-sdk-go/lib

go 1.17

replace (
	github.com/vmware/vsphere-automation-sdk-go/runtime => github.com/zhengxiexie/vsphere-automation-sdk-go/runtime v0.7.3-0.20240627100112-b3371cbc535c
)

require github.com/vmware/vsphere-automation-sdk-go/runtime v0.6.0

require (
	github.com/beevik/etree v1.1.0 // indirect
	github.com/gibson042/canonicaljson-go v1.0.3 // indirect
	github.com/golang-jwt/jwt/v4 v4.3.0 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	golang.org/x/text v0.3.8 // indirect
)
