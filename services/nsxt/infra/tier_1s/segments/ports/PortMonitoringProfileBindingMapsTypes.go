// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: PortMonitoringProfileBindingMaps.
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

func portMonitoringProfileBindingMapsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PortMonitoringProfileBindingMapsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func portMonitoringProfileBindingMapsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	paramsTypeMap["tier1_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_monitoring_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier1Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portMonitoringProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["port_monitoring_profile_binding_map_id"] = "portMonitoringProfileBindingMapId"
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
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
		"/policy/api/v1/infra/tier-1s/{tier1Id}/segments/{segmentId}/ports/{portId}/port-monitoring-profile-binding-maps/{portMonitoringProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func portMonitoringProfileBindingMapsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PortMonitoringProfileBindingMapsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PortMonitoringProfileBindingMapBindingType)
}

func portMonitoringProfileBindingMapsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	paramsTypeMap["tier1_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_monitoring_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier1Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portMonitoringProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["port_monitoring_profile_binding_map_id"] = "portMonitoringProfileBindingMapId"
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
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
		"/policy/api/v1/infra/tier-1s/{tier1Id}/segments/{segmentId}/ports/{portId}/port-monitoring-profile-binding-maps/{portMonitoringProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func portMonitoringProfileBindingMapsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PortMonitoringProfileBindingMapsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PortMonitoringProfileBindingMapListResultBindingType)
}

func portMonitoringProfileBindingMapsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["tier1_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["tier1Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portId"] = vapiBindings_.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["port_id"] = "portId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["include_mark_for_delete_objects"] = "include_mark_for_delete_objects"
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
		"/policy/api/v1/infra/tier-1s/{tier1Id}/segments/{segmentId}/ports/{portId}/port-monitoring-profile-binding-maps",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func portMonitoringProfileBindingMapsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["port_monitoring_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PortMonitoringProfileBindingMapBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	fieldNameMap["port_monitoring_profile_binding_map"] = "PortMonitoringProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PortMonitoringProfileBindingMapsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func portMonitoringProfileBindingMapsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["port_monitoring_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PortMonitoringProfileBindingMapBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	fieldNameMap["port_monitoring_profile_binding_map"] = "PortMonitoringProfileBindingMap"
	paramsTypeMap["port_monitoring_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PortMonitoringProfileBindingMapBindingType)
	paramsTypeMap["tier1_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_monitoring_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier1Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portMonitoringProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["port_monitoring_profile_binding_map_id"] = "portMonitoringProfileBindingMapId"
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
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
		"port_monitoring_profile_binding_map",
		"PATCH",
		"/policy/api/v1/infra/tier-1s/{tier1Id}/segments/{segmentId}/ports/{portId}/port-monitoring-profile-binding-maps/{portMonitoringProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func portMonitoringProfileBindingMapsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["port_monitoring_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PortMonitoringProfileBindingMapBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	fieldNameMap["port_monitoring_profile_binding_map"] = "PortMonitoringProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PortMonitoringProfileBindingMapsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PortMonitoringProfileBindingMapBindingType)
}

func portMonitoringProfileBindingMapsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["port_monitoring_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PortMonitoringProfileBindingMapBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	fieldNameMap["port_monitoring_profile_binding_map"] = "PortMonitoringProfileBindingMap"
	paramsTypeMap["port_monitoring_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PortMonitoringProfileBindingMapBindingType)
	paramsTypeMap["tier1_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_monitoring_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier1Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portMonitoringProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["port_monitoring_profile_binding_map_id"] = "portMonitoringProfileBindingMapId"
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
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
		"port_monitoring_profile_binding_map",
		"PUT",
		"/policy/api/v1/infra/tier-1s/{tier1Id}/segments/{segmentId}/ports/{portId}/port-monitoring-profile-binding-maps/{portMonitoringProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
