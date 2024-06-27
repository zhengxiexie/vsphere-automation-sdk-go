// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: NatRules.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package nat

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func natRulesDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["tgw_id"] = vapiBindings_.NewStringType()
	fields["nat_id"] = vapiBindings_.NewStringType()
	fields["nat_rule_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["tgw_id"] = "TgwId"
	fieldNameMap["nat_id"] = "NatId"
	fieldNameMap["nat_rule_id"] = "NatRuleId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NatRulesDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func natRulesDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["tgw_id"] = vapiBindings_.NewStringType()
	fields["nat_id"] = vapiBindings_.NewStringType()
	fields["nat_rule_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["tgw_id"] = "TgwId"
	fieldNameMap["nat_id"] = "NatId"
	fieldNameMap["nat_rule_id"] = "NatRuleId"
	paramsTypeMap["tgw_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["nat_rule_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["nat_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["tgwId"] = vapiBindings_.NewStringType()
	paramsTypeMap["natId"] = vapiBindings_.NewStringType()
	paramsTypeMap["natRuleId"] = vapiBindings_.NewStringType()
	pathParams["nat_rule_id"] = "natRuleId"
	pathParams["nat_id"] = "natId"
	pathParams["tgw_id"] = "tgwId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/transit-gateways/{tgwId}/nat/{natId}/nat-rules/{natRuleId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func natRulesGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["tgw_id"] = vapiBindings_.NewStringType()
	fields["nat_id"] = vapiBindings_.NewStringType()
	fields["nat_rule_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["tgw_id"] = "TgwId"
	fieldNameMap["nat_id"] = "NatId"
	fieldNameMap["nat_rule_id"] = "NatRuleId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NatRulesGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.TransitGatewayNatRuleBindingType)
}

func natRulesGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["tgw_id"] = vapiBindings_.NewStringType()
	fields["nat_id"] = vapiBindings_.NewStringType()
	fields["nat_rule_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["tgw_id"] = "TgwId"
	fieldNameMap["nat_id"] = "NatId"
	fieldNameMap["nat_rule_id"] = "NatRuleId"
	paramsTypeMap["tgw_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["nat_rule_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["nat_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["tgwId"] = vapiBindings_.NewStringType()
	paramsTypeMap["natId"] = vapiBindings_.NewStringType()
	paramsTypeMap["natRuleId"] = vapiBindings_.NewStringType()
	pathParams["nat_rule_id"] = "natRuleId"
	pathParams["nat_id"] = "natId"
	pathParams["tgw_id"] = "tgwId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/transit-gateways/{tgwId}/nat/{natId}/nat-rules/{natRuleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func natRulesListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["tgw_id"] = vapiBindings_.NewStringType()
	fields["nat_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["tgw_id"] = "TgwId"
	fieldNameMap["nat_id"] = "NatId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NatRulesListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.TransitGatewayNatRuleListResultBindingType)
}

func natRulesListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["tgw_id"] = vapiBindings_.NewStringType()
	fields["nat_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["tgw_id"] = "TgwId"
	fieldNameMap["nat_id"] = "NatId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["tgw_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["nat_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["tgwId"] = vapiBindings_.NewStringType()
	paramsTypeMap["natId"] = vapiBindings_.NewStringType()
	pathParams["nat_id"] = "natId"
	pathParams["tgw_id"] = "tgwId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/transit-gateways/{tgwId}/nat/{natId}/nat-rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func natRulesPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["tgw_id"] = vapiBindings_.NewStringType()
	fields["nat_id"] = vapiBindings_.NewStringType()
	fields["nat_rule_id"] = vapiBindings_.NewStringType()
	fields["transit_gateway_nat_rule"] = vapiBindings_.NewReferenceType(nsx_policyModel.TransitGatewayNatRuleBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["tgw_id"] = "TgwId"
	fieldNameMap["nat_id"] = "NatId"
	fieldNameMap["nat_rule_id"] = "NatRuleId"
	fieldNameMap["transit_gateway_nat_rule"] = "TransitGatewayNatRule"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NatRulesPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func natRulesPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["tgw_id"] = vapiBindings_.NewStringType()
	fields["nat_id"] = vapiBindings_.NewStringType()
	fields["nat_rule_id"] = vapiBindings_.NewStringType()
	fields["transit_gateway_nat_rule"] = vapiBindings_.NewReferenceType(nsx_policyModel.TransitGatewayNatRuleBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["tgw_id"] = "TgwId"
	fieldNameMap["nat_id"] = "NatId"
	fieldNameMap["nat_rule_id"] = "NatRuleId"
	fieldNameMap["transit_gateway_nat_rule"] = "TransitGatewayNatRule"
	paramsTypeMap["tgw_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["transit_gateway_nat_rule"] = vapiBindings_.NewReferenceType(nsx_policyModel.TransitGatewayNatRuleBindingType)
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["nat_rule_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["nat_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["tgwId"] = vapiBindings_.NewStringType()
	paramsTypeMap["natId"] = vapiBindings_.NewStringType()
	paramsTypeMap["natRuleId"] = vapiBindings_.NewStringType()
	pathParams["nat_rule_id"] = "natRuleId"
	pathParams["nat_id"] = "natId"
	pathParams["tgw_id"] = "tgwId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
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
		"transit_gateway_nat_rule",
		"PATCH",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/transit-gateways/{tgwId}/nat/{natId}/nat-rules/{natRuleId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func natRulesUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["tgw_id"] = vapiBindings_.NewStringType()
	fields["nat_id"] = vapiBindings_.NewStringType()
	fields["nat_rule_id"] = vapiBindings_.NewStringType()
	fields["transit_gateway_nat_rule"] = vapiBindings_.NewReferenceType(nsx_policyModel.TransitGatewayNatRuleBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["tgw_id"] = "TgwId"
	fieldNameMap["nat_id"] = "NatId"
	fieldNameMap["nat_rule_id"] = "NatRuleId"
	fieldNameMap["transit_gateway_nat_rule"] = "TransitGatewayNatRule"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NatRulesUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.TransitGatewayNatRuleBindingType)
}

func natRulesUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["tgw_id"] = vapiBindings_.NewStringType()
	fields["nat_id"] = vapiBindings_.NewStringType()
	fields["nat_rule_id"] = vapiBindings_.NewStringType()
	fields["transit_gateway_nat_rule"] = vapiBindings_.NewReferenceType(nsx_policyModel.TransitGatewayNatRuleBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["tgw_id"] = "TgwId"
	fieldNameMap["nat_id"] = "NatId"
	fieldNameMap["nat_rule_id"] = "NatRuleId"
	fieldNameMap["transit_gateway_nat_rule"] = "TransitGatewayNatRule"
	paramsTypeMap["tgw_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["transit_gateway_nat_rule"] = vapiBindings_.NewReferenceType(nsx_policyModel.TransitGatewayNatRuleBindingType)
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["nat_rule_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["nat_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["tgwId"] = vapiBindings_.NewStringType()
	paramsTypeMap["natId"] = vapiBindings_.NewStringType()
	paramsTypeMap["natRuleId"] = vapiBindings_.NewStringType()
	pathParams["nat_rule_id"] = "natRuleId"
	pathParams["nat_id"] = "natId"
	pathParams["tgw_id"] = "tgwId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
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
		"transit_gateway_nat_rule",
		"PUT",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/transit-gateways/{tgwId}/nat/{natId}/nat-rules/{natRuleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
