// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Settings.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package custom_signature_versions

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func settingsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version_id"] = vapiBindings_.NewStringType()
	fieldNameMap["version_id"] = "VersionId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SettingsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.IdsCustomSignatureSettingsBindingType)
}

func settingsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["version_id"] = vapiBindings_.NewStringType()
	fieldNameMap["version_id"] = "VersionId"
	paramsTypeMap["version_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["versionId"] = vapiBindings_.NewStringType()
	pathParams["version_id"] = "versionId"
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
		"/policy/api/v1/infra/settings/firewall/security/intrusion-services/custom-signature-versions/{versionId}/settings",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func settingsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version_id"] = vapiBindings_.NewStringType()
	fields["ids_custom_signature_settings"] = vapiBindings_.NewReferenceType(nsx_policyModel.IdsCustomSignatureSettingsBindingType)
	fieldNameMap["version_id"] = "VersionId"
	fieldNameMap["ids_custom_signature_settings"] = "IdsCustomSignatureSettings"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SettingsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func settingsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["version_id"] = vapiBindings_.NewStringType()
	fields["ids_custom_signature_settings"] = vapiBindings_.NewReferenceType(nsx_policyModel.IdsCustomSignatureSettingsBindingType)
	fieldNameMap["version_id"] = "VersionId"
	fieldNameMap["ids_custom_signature_settings"] = "IdsCustomSignatureSettings"
	paramsTypeMap["ids_custom_signature_settings"] = vapiBindings_.NewReferenceType(nsx_policyModel.IdsCustomSignatureSettingsBindingType)
	paramsTypeMap["version_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["versionId"] = vapiBindings_.NewStringType()
	pathParams["version_id"] = "versionId"
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
		"ids_custom_signature_settings",
		"PATCH",
		"/policy/api/v1/infra/settings/firewall/security/intrusion-services/custom-signature-versions/{versionId}/settings",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func settingsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version_id"] = vapiBindings_.NewStringType()
	fields["ids_custom_signature_settings"] = vapiBindings_.NewReferenceType(nsx_policyModel.IdsCustomSignatureSettingsBindingType)
	fieldNameMap["version_id"] = "VersionId"
	fieldNameMap["ids_custom_signature_settings"] = "IdsCustomSignatureSettings"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SettingsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.IdsCustomSignatureSettingsBindingType)
}

func settingsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["version_id"] = vapiBindings_.NewStringType()
	fields["ids_custom_signature_settings"] = vapiBindings_.NewReferenceType(nsx_policyModel.IdsCustomSignatureSettingsBindingType)
	fieldNameMap["version_id"] = "VersionId"
	fieldNameMap["ids_custom_signature_settings"] = "IdsCustomSignatureSettings"
	paramsTypeMap["ids_custom_signature_settings"] = vapiBindings_.NewReferenceType(nsx_policyModel.IdsCustomSignatureSettingsBindingType)
	paramsTypeMap["version_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["versionId"] = vapiBindings_.NewStringType()
	pathParams["version_id"] = "versionId"
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
		"ids_custom_signature_settings",
		"PUT",
		"/policy/api/v1/infra/settings/firewall/security/intrusion-services/custom-signature-versions/{versionId}/settings",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
