// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Statistics.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package virtual_servers

import (
	vapiBindings_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/zhengxiexie/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``source`` of method Statistics#get.
const Statistics_GET_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method Statistics#get.
const Statistics_GET_SOURCE_CACHED = "cached"

// Possible value for ``source`` of method Statistics#list.
const Statistics_LIST_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method Statistics#list.
const Statistics_LIST_SOURCE_CACHED = "cached"

func statisticsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["virtual_server_id"] = vapiBindings_.NewStringType()
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["virtual_server_id"] = "VirtualServerId"
	fieldNameMap["source"] = "Source"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func StatisticsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.LbVirtualServerStatisticsBindingType)
}

func statisticsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["virtual_server_id"] = vapiBindings_.NewStringType()
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["virtual_server_id"] = "VirtualServerId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["service_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["virtual_server_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["virtualServerId"] = vapiBindings_.NewStringType()
	pathParams["virtual_server_id"] = "virtualServerId"
	pathParams["service_id"] = "serviceId"
	queryParams["source"] = "source"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/loadbalancer/services/nsxt-mp/{serviceId}/virtual-servers/{virtualServerId}/statistics",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func statisticsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["source"] = "Source"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func StatisticsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.LbVirtualServerStatisticsListResultBindingType)
}

func statisticsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["service_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["serviceId"] = vapiBindings_.NewStringType()
	pathParams["service_id"] = "serviceId"
	queryParams["source"] = "source"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/loadbalancer/services/nsxt-mp/{serviceId}/virtual-servers/statistics",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
