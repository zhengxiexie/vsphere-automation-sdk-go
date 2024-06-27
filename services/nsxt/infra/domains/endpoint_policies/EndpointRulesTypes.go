// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: EndpointRules.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package endpoint_policies

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func endpointRulesDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["endpoint_policy_id"] = vapiBindings_.NewStringType()
	fields["endpoint_rule_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["endpoint_policy_id"] = "EndpointPolicyId"
	fieldNameMap["endpoint_rule_id"] = "EndpointRuleId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EndpointRulesDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func endpointRulesDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["endpoint_policy_id"] = vapiBindings_.NewStringType()
	fields["endpoint_rule_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["endpoint_policy_id"] = "EndpointPolicyId"
	fieldNameMap["endpoint_rule_id"] = "EndpointRuleId"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["endpoint_rule_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["endpoint_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["endpointPolicyId"] = vapiBindings_.NewStringType()
	paramsTypeMap["endpointRuleId"] = vapiBindings_.NewStringType()
	pathParams["endpoint_policy_id"] = "endpointPolicyId"
	pathParams["endpoint_rule_id"] = "endpointRuleId"
	pathParams["domain_id"] = "domainId"
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
		"/policy/api/v1/infra/domains/{domainId}/endpoint-policies/{endpointPolicyId}/endpoint-rules/{endpointRuleId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func endpointRulesGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["endpoint_policy_id"] = vapiBindings_.NewStringType()
	fields["endpoint_rule_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["endpoint_policy_id"] = "EndpointPolicyId"
	fieldNameMap["endpoint_rule_id"] = "EndpointRuleId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EndpointRulesGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.EndpointRuleBindingType)
}

func endpointRulesGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["endpoint_policy_id"] = vapiBindings_.NewStringType()
	fields["endpoint_rule_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["endpoint_policy_id"] = "EndpointPolicyId"
	fieldNameMap["endpoint_rule_id"] = "EndpointRuleId"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["endpoint_rule_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["endpoint_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["endpointPolicyId"] = vapiBindings_.NewStringType()
	paramsTypeMap["endpointRuleId"] = vapiBindings_.NewStringType()
	pathParams["endpoint_policy_id"] = "endpointPolicyId"
	pathParams["endpoint_rule_id"] = "endpointRuleId"
	pathParams["domain_id"] = "domainId"
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
		"/policy/api/v1/infra/domains/{domainId}/endpoint-policies/{endpointPolicyId}/endpoint-rules/{endpointRuleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func endpointRulesListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["endpoint_policy_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["endpoint_policy_id"] = "EndpointPolicyId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EndpointRulesListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.EndpointRuleListResultBindingType)
}

func endpointRulesListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["endpoint_policy_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["endpoint_policy_id"] = "EndpointPolicyId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["endpoint_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["endpointPolicyId"] = vapiBindings_.NewStringType()
	pathParams["endpoint_policy_id"] = "endpointPolicyId"
	pathParams["domain_id"] = "domainId"
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
		"/policy/api/v1/infra/domains/{domainId}/endpoint-policies/{endpointPolicyId}/endpoint-rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func endpointRulesPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["endpoint_policy_id"] = vapiBindings_.NewStringType()
	fields["endpoint_rule_id"] = vapiBindings_.NewStringType()
	fields["endpoint_rule"] = vapiBindings_.NewReferenceType(nsx_policyModel.EndpointRuleBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["endpoint_policy_id"] = "EndpointPolicyId"
	fieldNameMap["endpoint_rule_id"] = "EndpointRuleId"
	fieldNameMap["endpoint_rule"] = "EndpointRule"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EndpointRulesPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func endpointRulesPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["endpoint_policy_id"] = vapiBindings_.NewStringType()
	fields["endpoint_rule_id"] = vapiBindings_.NewStringType()
	fields["endpoint_rule"] = vapiBindings_.NewReferenceType(nsx_policyModel.EndpointRuleBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["endpoint_policy_id"] = "EndpointPolicyId"
	fieldNameMap["endpoint_rule_id"] = "EndpointRuleId"
	fieldNameMap["endpoint_rule"] = "EndpointRule"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["endpoint_rule"] = vapiBindings_.NewReferenceType(nsx_policyModel.EndpointRuleBindingType)
	paramsTypeMap["endpoint_rule_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["endpoint_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["endpointPolicyId"] = vapiBindings_.NewStringType()
	paramsTypeMap["endpointRuleId"] = vapiBindings_.NewStringType()
	pathParams["endpoint_policy_id"] = "endpointPolicyId"
	pathParams["endpoint_rule_id"] = "endpointRuleId"
	pathParams["domain_id"] = "domainId"
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
		"endpoint_rule",
		"PATCH",
		"/policy/api/v1/infra/domains/{domainId}/endpoint-policies/{endpointPolicyId}/endpoint-rules/{endpointRuleId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func endpointRulesUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["endpoint_policy_id"] = vapiBindings_.NewStringType()
	fields["endpoint_rule_id"] = vapiBindings_.NewStringType()
	fields["endpoint_rule"] = vapiBindings_.NewReferenceType(nsx_policyModel.EndpointRuleBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["endpoint_policy_id"] = "EndpointPolicyId"
	fieldNameMap["endpoint_rule_id"] = "EndpointRuleId"
	fieldNameMap["endpoint_rule"] = "EndpointRule"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EndpointRulesUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.EndpointRuleBindingType)
}

func endpointRulesUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["endpoint_policy_id"] = vapiBindings_.NewStringType()
	fields["endpoint_rule_id"] = vapiBindings_.NewStringType()
	fields["endpoint_rule"] = vapiBindings_.NewReferenceType(nsx_policyModel.EndpointRuleBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["endpoint_policy_id"] = "EndpointPolicyId"
	fieldNameMap["endpoint_rule_id"] = "EndpointRuleId"
	fieldNameMap["endpoint_rule"] = "EndpointRule"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["endpoint_rule"] = vapiBindings_.NewReferenceType(nsx_policyModel.EndpointRuleBindingType)
	paramsTypeMap["endpoint_rule_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["endpoint_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["endpointPolicyId"] = vapiBindings_.NewStringType()
	paramsTypeMap["endpointRuleId"] = vapiBindings_.NewStringType()
	pathParams["endpoint_policy_id"] = "endpointPolicyId"
	pathParams["endpoint_rule_id"] = "endpointRuleId"
	pathParams["domain_id"] = "domainId"
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
		"endpoint_rule",
		"PUT",
		"/policy/api/v1/infra/domains/{domainId}/endpoint-policies/{endpointPolicyId}/endpoint-rules/{endpointRuleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
