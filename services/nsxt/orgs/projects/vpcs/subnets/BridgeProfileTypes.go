// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: BridgeProfile.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package subnets

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func bridgeProfileDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["subnet_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func BridgeProfileDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func bridgeProfileDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["subnet_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnet_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnetId"] = vapiBindings_.NewStringType()
	pathParams["subnet_id"] = "subnetId"
	pathParams["vpc_id"] = "vpcId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/subnets/{subnetId}/bridge-profile",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func bridgeProfileGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["subnet_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func BridgeProfileGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.VpcSubnetBridgeProfileBindingType)
}

func bridgeProfileGetRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["subnet_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnet_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnetId"] = vapiBindings_.NewStringType()
	pathParams["subnet_id"] = "subnetId"
	pathParams["vpc_id"] = "vpcId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/subnets/{subnetId}/bridge-profile",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func bridgeProfilePatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["subnet_id"] = vapiBindings_.NewStringType()
	fields["vpc_subnet_bridge_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.VpcSubnetBridgeProfileBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	fieldNameMap["vpc_subnet_bridge_profile"] = "VpcSubnetBridgeProfile"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func BridgeProfilePatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func bridgeProfilePatchRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["subnet_id"] = vapiBindings_.NewStringType()
	fields["vpc_subnet_bridge_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.VpcSubnetBridgeProfileBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	fieldNameMap["vpc_subnet_bridge_profile"] = "VpcSubnetBridgeProfile"
	paramsTypeMap["vpc_subnet_bridge_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.VpcSubnetBridgeProfileBindingType)
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnet_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnetId"] = vapiBindings_.NewStringType()
	pathParams["subnet_id"] = "subnetId"
	pathParams["vpc_id"] = "vpcId"
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
		"vpc_subnet_bridge_profile",
		"PATCH",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/subnets/{subnetId}/bridge-profile",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func bridgeProfileUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["subnet_id"] = vapiBindings_.NewStringType()
	fields["vpc_subnet_bridge_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.VpcSubnetBridgeProfileBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	fieldNameMap["vpc_subnet_bridge_profile"] = "VpcSubnetBridgeProfile"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func BridgeProfileUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.VpcSubnetBridgeProfileBindingType)
}

func bridgeProfileUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["subnet_id"] = vapiBindings_.NewStringType()
	fields["vpc_subnet_bridge_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.VpcSubnetBridgeProfileBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	fieldNameMap["vpc_subnet_bridge_profile"] = "VpcSubnetBridgeProfile"
	paramsTypeMap["vpc_subnet_bridge_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.VpcSubnetBridgeProfileBindingType)
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnet_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnetId"] = vapiBindings_.NewStringType()
	pathParams["subnet_id"] = "subnetId"
	pathParams["vpc_id"] = "vpcId"
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
		"vpc_subnet_bridge_profile",
		"PUT",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/subnets/{subnetId}/bridge-profile",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
