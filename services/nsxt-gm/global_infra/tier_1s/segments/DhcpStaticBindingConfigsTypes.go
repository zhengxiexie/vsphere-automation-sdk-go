// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: DhcpStaticBindingConfigs.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package segments

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_global_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
	"reflect"
)

func dhcpStaticBindingConfigsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["binding_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["binding_id"] = "BindingId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DhcpStaticBindingConfigsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func dhcpStaticBindingConfigsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["binding_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["binding_id"] = "BindingId"
	paramsTypeMap["tier1_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["binding_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier1Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["bindingId"] = vapiBindings_.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["binding_id"] = "bindingId"
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
		"/global-manager/api/v1/global-infra/tier-1s/{tier1Id}/segments/{segmentId}/dhcp-static-binding-configs/{bindingId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func dhcpStaticBindingConfigsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["binding_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["binding_id"] = "BindingId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DhcpStaticBindingConfigsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_global_policyModel.DhcpStaticBindingConfigBindingType)})
}

func dhcpStaticBindingConfigsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["binding_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["binding_id"] = "BindingId"
	paramsTypeMap["tier1_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["binding_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier1Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["bindingId"] = vapiBindings_.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["binding_id"] = "bindingId"
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
		"/global-manager/api/v1/global-infra/tier-1s/{tier1Id}/segments/{segmentId}/dhcp-static-binding-configs/{bindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func dhcpStaticBindingConfigsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DhcpStaticBindingConfigsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_global_policyModel.DhcpStaticBindingConfigListResultBindingType)
}

func dhcpStaticBindingConfigsListRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
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
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["tier1Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
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
		"/global-manager/api/v1/global-infra/tier-1s/{tier1Id}/segments/{segmentId}/dhcp-static-binding-configs",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func dhcpStaticBindingConfigsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["binding_id"] = vapiBindings_.NewStringType()
	fields["dhcp_static_binding_config"] = vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_global_policyModel.DhcpStaticBindingConfigBindingType)})
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["dhcp_static_binding_config"] = "DhcpStaticBindingConfig"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DhcpStaticBindingConfigsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func dhcpStaticBindingConfigsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["binding_id"] = vapiBindings_.NewStringType()
	fields["dhcp_static_binding_config"] = vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_global_policyModel.DhcpStaticBindingConfigBindingType)})
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["dhcp_static_binding_config"] = "DhcpStaticBindingConfig"
	paramsTypeMap["tier1_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["dhcp_static_binding_config"] = vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_global_policyModel.DhcpStaticBindingConfigBindingType)})
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["binding_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier1Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["bindingId"] = vapiBindings_.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["binding_id"] = "bindingId"
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
		"dhcp_static_binding_config",
		"PATCH",
		"/global-manager/api/v1/global-infra/tier-1s/{tier1Id}/segments/{segmentId}/dhcp-static-binding-configs/{bindingId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func dhcpStaticBindingConfigsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["binding_id"] = vapiBindings_.NewStringType()
	fields["dhcp_static_binding_config"] = vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_global_policyModel.DhcpStaticBindingConfigBindingType)})
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["dhcp_static_binding_config"] = "DhcpStaticBindingConfig"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DhcpStaticBindingConfigsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_global_policyModel.DhcpStaticBindingConfigBindingType)})
}

func dhcpStaticBindingConfigsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["binding_id"] = vapiBindings_.NewStringType()
	fields["dhcp_static_binding_config"] = vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_global_policyModel.DhcpStaticBindingConfigBindingType)})
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["dhcp_static_binding_config"] = "DhcpStaticBindingConfig"
	paramsTypeMap["tier1_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["dhcp_static_binding_config"] = vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_global_policyModel.DhcpStaticBindingConfigBindingType)})
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["binding_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier1Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["bindingId"] = vapiBindings_.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["binding_id"] = "bindingId"
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
		"dhcp_static_binding_config",
		"PUT",
		"/global-manager/api/v1/global-infra/tier-1s/{tier1Id}/segments/{segmentId}/dhcp-static-binding-configs/{bindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
