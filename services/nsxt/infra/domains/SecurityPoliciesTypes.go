// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SecurityPolicies.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package domains

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for “operation“ of method SecurityPolicies#revise.
const SecurityPolicies_REVISE_OPERATION_TOP = "insert_top"

// Possible value for “operation“ of method SecurityPolicies#revise.
const SecurityPolicies_REVISE_OPERATION_BOTTOM = "insert_bottom"

// Possible value for “operation“ of method SecurityPolicies#revise.
const SecurityPolicies_REVISE_OPERATION_AFTER = "insert_after"

// Possible value for “operation“ of method SecurityPolicies#revise.
const SecurityPolicies_REVISE_OPERATION_BEFORE = "insert_before"

func securityPoliciesDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SecurityPoliciesDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func securityPoliciesDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["security_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["securityPolicyId"] = vapiBindings_.NewStringType()
	pathParams["security_policy_id"] = "securityPolicyId"
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
		"/policy/api/v1/infra/domains/{domainId}/security-policies/{securityPolicyId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func securityPoliciesGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SecurityPoliciesGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyBindingType)
}

func securityPoliciesGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["security_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["securityPolicyId"] = vapiBindings_.NewStringType()
	pathParams["security_policy_id"] = "securityPolicyId"
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
		"/policy/api/v1/infra/domains/{domainId}/security-policies/{securityPolicyId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func securityPoliciesListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["include_rule_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["include_rule_count"] = "IncludeRuleCount"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SecurityPoliciesListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyListResultBindingType)
}

func securityPoliciesListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["include_rule_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["include_rule_count"] = "IncludeRuleCount"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_rule_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	pathParams["domain_id"] = "domainId"
	queryParams["cursor"] = "cursor"
	queryParams["include_rule_count"] = "include_rule_count"
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
		"/policy/api/v1/infra/domains/{domainId}/security-policies",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func securityPoliciesPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fields["security_policy"] = vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["security_policy"] = "SecurityPolicy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SecurityPoliciesPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func securityPoliciesPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fields["security_policy"] = vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["security_policy"] = "SecurityPolicy"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["security_policy"] = vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyBindingType)
	paramsTypeMap["security_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["securityPolicyId"] = vapiBindings_.NewStringType()
	pathParams["security_policy_id"] = "securityPolicyId"
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
		"security_policy",
		"PATCH",
		"/policy/api/v1/infra/domains/{domainId}/security-policies/{securityPolicyId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func securityPoliciesReviseInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fields["security_policy"] = vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyBindingType)
	fields["anchor_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["security_policy"] = "SecurityPolicy"
	fieldNameMap["anchor_path"] = "AnchorPath"
	fieldNameMap["operation"] = "Operation"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SecurityPoliciesReviseOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyBindingType)
}

func securityPoliciesReviseRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fields["security_policy"] = vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyBindingType)
	fields["anchor_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["security_policy"] = "SecurityPolicy"
	fieldNameMap["anchor_path"] = "AnchorPath"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["security_policy"] = vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyBindingType)
	paramsTypeMap["security_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["anchor_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["securityPolicyId"] = vapiBindings_.NewStringType()
	pathParams["security_policy_id"] = "securityPolicyId"
	pathParams["domain_id"] = "domainId"
	queryParams["anchor_path"] = "anchor_path"
	queryParams["operation"] = "operation"
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
		"action=revise",
		"security_policy",
		"POST",
		"/policy/api/v1/infra/domains/{domainId}/security-policies/{securityPolicyId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func securityPoliciesUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fields["security_policy"] = vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["security_policy"] = "SecurityPolicy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SecurityPoliciesUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyBindingType)
}

func securityPoliciesUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fields["security_policy"] = vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["security_policy"] = "SecurityPolicy"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["security_policy"] = vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyBindingType)
	paramsTypeMap["security_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["securityPolicyId"] = vapiBindings_.NewStringType()
	pathParams["security_policy_id"] = "securityPolicyId"
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
		"security_policy",
		"PUT",
		"/policy/api/v1/infra/domains/{domainId}/security-policies/{securityPolicyId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
