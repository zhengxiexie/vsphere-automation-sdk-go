// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: EnforcementPoints.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package sites

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func enforcementPointsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EnforcementPointsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func enforcementPointsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	paramsTypeMap["enforcementpoint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementpointId"] = vapiBindings_.NewStringType()
	pathParams["enforcementpoint_id"] = "enforcementpointId"
	pathParams["site_id"] = "siteId"
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
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementpointId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func enforcementPointsFullsyncInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EnforcementPointsFullsyncOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func enforcementPointsFullsyncRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementPointId"] = vapiBindings_.NewStringType()
	pathParams["site_id"] = "siteId"
	pathParams["enforcement_point_id"] = "enforcementPointId"
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
		"action=full-sync",
		"",
		"POST",
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementPointId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func enforcementPointsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EnforcementPointsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.EnforcementPointBindingType)
}

func enforcementPointsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	paramsTypeMap["enforcementpoint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementpointId"] = vapiBindings_.NewStringType()
	pathParams["enforcementpoint_id"] = "enforcementpointId"
	pathParams["site_id"] = "siteId"
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
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementpointId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func enforcementPointsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EnforcementPointsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.EnforcementPointListResultBindingType)
}

func enforcementPointsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	pathParams["site_id"] = "siteId"
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
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func enforcementPointsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point"] = vapiBindings_.NewReferenceType(nsx_policyModel.EnforcementPointBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["enforcement_point"] = "EnforcementPoint"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EnforcementPointsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func enforcementPointsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point"] = vapiBindings_.NewReferenceType(nsx_policyModel.EnforcementPointBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["enforcement_point"] = "EnforcementPoint"
	paramsTypeMap["enforcementpoint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point"] = vapiBindings_.NewReferenceType(nsx_policyModel.EnforcementPointBindingType)
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementpointId"] = vapiBindings_.NewStringType()
	pathParams["enforcementpoint_id"] = "enforcementpointId"
	pathParams["site_id"] = "siteId"
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
		"enforcement_point",
		"PATCH",
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementpointId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func enforcementPointsReloadInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EnforcementPointsReloadOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.EnforcementPointBindingType)
}

func enforcementPointsReloadRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	paramsTypeMap["enforcementpoint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementpointId"] = vapiBindings_.NewStringType()
	pathParams["enforcementpoint_id"] = "enforcementpointId"
	pathParams["site_id"] = "siteId"
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
		"action=reload",
		"",
		"POST",
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementpointId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func enforcementPointsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point"] = vapiBindings_.NewReferenceType(nsx_policyModel.EnforcementPointBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["enforcement_point"] = "EnforcementPoint"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EnforcementPointsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.EnforcementPointBindingType)
}

func enforcementPointsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point"] = vapiBindings_.NewReferenceType(nsx_policyModel.EnforcementPointBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["enforcement_point"] = "EnforcementPoint"
	paramsTypeMap["enforcementpoint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point"] = vapiBindings_.NewReferenceType(nsx_policyModel.EnforcementPointBindingType)
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementpointId"] = vapiBindings_.NewStringType()
	pathParams["enforcementpoint_id"] = "enforcementpointId"
	pathParams["site_id"] = "siteId"
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
		"enforcement_point",
		"PUT",
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementpointId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
