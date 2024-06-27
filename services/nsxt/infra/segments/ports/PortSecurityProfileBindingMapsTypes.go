// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: PortSecurityProfileBindingMaps.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package ports

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func portSecurityProfileBindingMapsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_security_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_security_profile_binding_map_id"] = "PortSecurityProfileBindingMapId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PortSecurityProfileBindingMapsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func portSecurityProfileBindingMapsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_security_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_security_profile_binding_map_id"] = "PortSecurityProfileBindingMapId"
	paramsTypeMap["port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_security_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portSecurityProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["port_security_profile_binding_map_id"] = "portSecurityProfileBindingMapId"
	pathParams["segment_id"] = "segmentId"
	pathParams["port_id"] = "portId"
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
		"DELETE",
		"/policy/api/v1/infra/segments/{segmentId}/ports/{portId}/port-security-profile-binding-maps/{portSecurityProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func portSecurityProfileBindingMapsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_security_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_security_profile_binding_map_id"] = "PortSecurityProfileBindingMapId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PortSecurityProfileBindingMapsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PortSecurityProfileBindingMapBindingType)
}

func portSecurityProfileBindingMapsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_security_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_security_profile_binding_map_id"] = "PortSecurityProfileBindingMapId"
	paramsTypeMap["port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_security_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portSecurityProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["port_security_profile_binding_map_id"] = "portSecurityProfileBindingMapId"
	pathParams["segment_id"] = "segmentId"
	pathParams["port_id"] = "portId"
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
		"/policy/api/v1/infra/segments/{segmentId}/ports/{portId}/port-security-profile-binding-maps/{portSecurityProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func portSecurityProfileBindingMapsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PortSecurityProfileBindingMapsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PortSecurityProfileBindingMapListResultBindingType)
}

func portSecurityProfileBindingMapsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portId"] = vapiBindings_.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["port_id"] = "portId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["page_size"] = "page_size"
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
		"/policy/api/v1/infra/segments/{segmentId}/ports/{portId}/port-security-profile-binding-maps",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func portSecurityProfileBindingMapsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_security_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["port_security_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PortSecurityProfileBindingMapBindingType)
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_security_profile_binding_map_id"] = "PortSecurityProfileBindingMapId"
	fieldNameMap["port_security_profile_binding_map"] = "PortSecurityProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PortSecurityProfileBindingMapsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func portSecurityProfileBindingMapsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_security_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["port_security_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PortSecurityProfileBindingMapBindingType)
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_security_profile_binding_map_id"] = "PortSecurityProfileBindingMapId"
	fieldNameMap["port_security_profile_binding_map"] = "PortSecurityProfileBindingMap"
	paramsTypeMap["port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_security_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_security_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PortSecurityProfileBindingMapBindingType)
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portSecurityProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["port_security_profile_binding_map_id"] = "portSecurityProfileBindingMapId"
	pathParams["segment_id"] = "segmentId"
	pathParams["port_id"] = "portId"
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
		"port_security_profile_binding_map",
		"PATCH",
		"/policy/api/v1/infra/segments/{segmentId}/ports/{portId}/port-security-profile-binding-maps/{portSecurityProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func portSecurityProfileBindingMapsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_security_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["port_security_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PortSecurityProfileBindingMapBindingType)
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_security_profile_binding_map_id"] = "PortSecurityProfileBindingMapId"
	fieldNameMap["port_security_profile_binding_map"] = "PortSecurityProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PortSecurityProfileBindingMapsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PortSecurityProfileBindingMapBindingType)
}

func portSecurityProfileBindingMapsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_security_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["port_security_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PortSecurityProfileBindingMapBindingType)
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_security_profile_binding_map_id"] = "PortSecurityProfileBindingMapId"
	fieldNameMap["port_security_profile_binding_map"] = "PortSecurityProfileBindingMap"
	paramsTypeMap["port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_security_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_security_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PortSecurityProfileBindingMapBindingType)
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portSecurityProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["port_security_profile_binding_map_id"] = "portSecurityProfileBindingMapId"
	pathParams["segment_id"] = "segmentId"
	pathParams["port_id"] = "portId"
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
		"port_security_profile_binding_map",
		"PUT",
		"/policy/api/v1/infra/segments/{segmentId}/ports/{portId}/port-security-profile-binding-maps/{portSecurityProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
