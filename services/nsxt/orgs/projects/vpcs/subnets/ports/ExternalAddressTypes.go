// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ExternalAddress.
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

func externalAddressDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["subnet_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	fieldNameMap["port_id"] = "PortId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ExternalAddressDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func externalAddressDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["port_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	fieldNameMap["port_id"] = "PortId"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnet_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnetId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portId"] = vapiBindings_.NewStringType()
	pathParams["subnet_id"] = "subnetId"
	pathParams["vpc_id"] = "vpcId"
	pathParams["port_id"] = "portId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/subnets/{subnetId}/ports/{portId}/external-address",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func externalAddressGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["subnet_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	fieldNameMap["port_id"] = "PortId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ExternalAddressGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.ExternalAddressBindingBindingType)
}

func externalAddressGetRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["port_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	fieldNameMap["port_id"] = "PortId"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnet_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnetId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portId"] = vapiBindings_.NewStringType()
	pathParams["subnet_id"] = "subnetId"
	pathParams["vpc_id"] = "vpcId"
	pathParams["port_id"] = "portId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/subnets/{subnetId}/ports/{portId}/external-address",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func externalAddressPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["subnet_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["external_address_binding"] = vapiBindings_.NewReferenceType(nsx_policyModel.ExternalAddressBindingBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["external_address_binding"] = "ExternalAddressBinding"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ExternalAddressPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func externalAddressPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["external_address_binding"] = vapiBindings_.NewReferenceType(nsx_policyModel.ExternalAddressBindingBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["external_address_binding"] = "ExternalAddressBinding"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnet_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["external_address_binding"] = vapiBindings_.NewReferenceType(nsx_policyModel.ExternalAddressBindingBindingType)
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnetId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portId"] = vapiBindings_.NewStringType()
	pathParams["subnet_id"] = "subnetId"
	pathParams["vpc_id"] = "vpcId"
	pathParams["port_id"] = "portId"
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
		"external_address_binding",
		"PATCH",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/subnets/{subnetId}/ports/{portId}/external-address",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func externalAddressUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["subnet_id"] = vapiBindings_.NewStringType()
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["external_address_binding"] = vapiBindings_.NewReferenceType(nsx_policyModel.ExternalAddressBindingBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["external_address_binding"] = "ExternalAddressBinding"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ExternalAddressUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.ExternalAddressBindingBindingType)
}

func externalAddressUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["port_id"] = vapiBindings_.NewStringType()
	fields["external_address_binding"] = vapiBindings_.NewReferenceType(nsx_policyModel.ExternalAddressBindingBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["external_address_binding"] = "ExternalAddressBinding"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnet_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["external_address_binding"] = vapiBindings_.NewReferenceType(nsx_policyModel.ExternalAddressBindingBindingType)
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["subnetId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portId"] = vapiBindings_.NewStringType()
	pathParams["subnet_id"] = "subnetId"
	pathParams["vpc_id"] = "vpcId"
	pathParams["port_id"] = "portId"
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
		"external_address_binding",
		"PUT",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/subnets/{subnetId}/ports/{portId}/external-address",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
