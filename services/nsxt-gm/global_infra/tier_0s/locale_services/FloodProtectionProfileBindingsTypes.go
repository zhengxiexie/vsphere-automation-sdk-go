// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: FloodProtectionProfileBindings.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package locale_services

import (
	vapiBindings_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/protocol"
	nsx_global_policyModel "github.com/zhengxiexie/vsphere-automation-sdk-go/services/nsxt-gm/model"
	"reflect"
)

func floodProtectionProfileBindingsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fields["locale_services_id"] = vapiBindings_.NewStringType()
	fields["flood_protection_profile_binding_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_services_id"] = "LocaleServicesId"
	fieldNameMap["flood_protection_profile_binding_id"] = "FloodProtectionProfileBindingId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func FloodProtectionProfileBindingsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func floodProtectionProfileBindingsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fields["locale_services_id"] = vapiBindings_.NewStringType()
	fields["flood_protection_profile_binding_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_services_id"] = "LocaleServicesId"
	fieldNameMap["flood_protection_profile_binding_id"] = "FloodProtectionProfileBindingId"
	paramsTypeMap["tier0_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["flood_protection_profile_binding_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["locale_services_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier0Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["localeServicesId"] = vapiBindings_.NewStringType()
	paramsTypeMap["floodProtectionProfileBindingId"] = vapiBindings_.NewStringType()
	pathParams["locale_services_id"] = "localeServicesId"
	pathParams["tier0_id"] = "tier0Id"
	pathParams["flood_protection_profile_binding_id"] = "floodProtectionProfileBindingId"
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
		"/global-manager/api/v1/global-infra/tier-0s/{tier0Id}/locale-services/{localeServicesId}/flood-protection-profile-bindings/{floodProtectionProfileBindingId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func floodProtectionProfileBindingsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fields["locale_services_id"] = vapiBindings_.NewStringType()
	fields["flood_protection_profile_binding_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_services_id"] = "LocaleServicesId"
	fieldNameMap["flood_protection_profile_binding_id"] = "FloodProtectionProfileBindingId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func FloodProtectionProfileBindingsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_global_policyModel.FloodProtectionProfileBindingMapBindingType)
}

func floodProtectionProfileBindingsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fields["locale_services_id"] = vapiBindings_.NewStringType()
	fields["flood_protection_profile_binding_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_services_id"] = "LocaleServicesId"
	fieldNameMap["flood_protection_profile_binding_id"] = "FloodProtectionProfileBindingId"
	paramsTypeMap["tier0_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["flood_protection_profile_binding_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["locale_services_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier0Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["localeServicesId"] = vapiBindings_.NewStringType()
	paramsTypeMap["floodProtectionProfileBindingId"] = vapiBindings_.NewStringType()
	pathParams["locale_services_id"] = "localeServicesId"
	pathParams["tier0_id"] = "tier0Id"
	pathParams["flood_protection_profile_binding_id"] = "floodProtectionProfileBindingId"
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
		"/global-manager/api/v1/global-infra/tier-0s/{tier0Id}/locale-services/{localeServicesId}/flood-protection-profile-bindings/{floodProtectionProfileBindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func floodProtectionProfileBindingsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fields["locale_services_id"] = vapiBindings_.NewStringType()
	fields["flood_protection_profile_binding_id"] = vapiBindings_.NewStringType()
	fields["flood_protection_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.FloodProtectionProfileBindingMapBindingType)
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_services_id"] = "LocaleServicesId"
	fieldNameMap["flood_protection_profile_binding_id"] = "FloodProtectionProfileBindingId"
	fieldNameMap["flood_protection_profile_binding_map"] = "FloodProtectionProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func FloodProtectionProfileBindingsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func floodProtectionProfileBindingsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fields["locale_services_id"] = vapiBindings_.NewStringType()
	fields["flood_protection_profile_binding_id"] = vapiBindings_.NewStringType()
	fields["flood_protection_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.FloodProtectionProfileBindingMapBindingType)
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_services_id"] = "LocaleServicesId"
	fieldNameMap["flood_protection_profile_binding_id"] = "FloodProtectionProfileBindingId"
	fieldNameMap["flood_protection_profile_binding_map"] = "FloodProtectionProfileBindingMap"
	paramsTypeMap["tier0_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["flood_protection_profile_binding_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["flood_protection_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.FloodProtectionProfileBindingMapBindingType)
	paramsTypeMap["locale_services_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier0Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["localeServicesId"] = vapiBindings_.NewStringType()
	paramsTypeMap["floodProtectionProfileBindingId"] = vapiBindings_.NewStringType()
	pathParams["locale_services_id"] = "localeServicesId"
	pathParams["tier0_id"] = "tier0Id"
	pathParams["flood_protection_profile_binding_id"] = "floodProtectionProfileBindingId"
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
		"flood_protection_profile_binding_map",
		"PATCH",
		"/global-manager/api/v1/global-infra/tier-0s/{tier0Id}/locale-services/{localeServicesId}/flood-protection-profile-bindings/{floodProtectionProfileBindingId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func floodProtectionProfileBindingsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fields["locale_services_id"] = vapiBindings_.NewStringType()
	fields["flood_protection_profile_binding_id"] = vapiBindings_.NewStringType()
	fields["flood_protection_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.FloodProtectionProfileBindingMapBindingType)
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_services_id"] = "LocaleServicesId"
	fieldNameMap["flood_protection_profile_binding_id"] = "FloodProtectionProfileBindingId"
	fieldNameMap["flood_protection_profile_binding_map"] = "FloodProtectionProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func FloodProtectionProfileBindingsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_global_policyModel.FloodProtectionProfileBindingMapBindingType)
}

func floodProtectionProfileBindingsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fields["locale_services_id"] = vapiBindings_.NewStringType()
	fields["flood_protection_profile_binding_id"] = vapiBindings_.NewStringType()
	fields["flood_protection_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.FloodProtectionProfileBindingMapBindingType)
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_services_id"] = "LocaleServicesId"
	fieldNameMap["flood_protection_profile_binding_id"] = "FloodProtectionProfileBindingId"
	fieldNameMap["flood_protection_profile_binding_map"] = "FloodProtectionProfileBindingMap"
	paramsTypeMap["tier0_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["flood_protection_profile_binding_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["flood_protection_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.FloodProtectionProfileBindingMapBindingType)
	paramsTypeMap["locale_services_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier0Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["localeServicesId"] = vapiBindings_.NewStringType()
	paramsTypeMap["floodProtectionProfileBindingId"] = vapiBindings_.NewStringType()
	pathParams["locale_services_id"] = "localeServicesId"
	pathParams["tier0_id"] = "tier0Id"
	pathParams["flood_protection_profile_binding_id"] = "floodProtectionProfileBindingId"
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
		"flood_protection_profile_binding_map",
		"PUT",
		"/global-manager/api/v1/global-infra/tier-0s/{tier0Id}/locale-services/{localeServicesId}/flood-protection-profile-bindings/{floodProtectionProfileBindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
