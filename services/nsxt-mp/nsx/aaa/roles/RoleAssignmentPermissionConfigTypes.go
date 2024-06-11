// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: RoleAssignmentPermissionConfig.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package roles

import (
	vapiBindings_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/zhengxiexie/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func roleAssignmentPermissionConfigUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["role"] = vapiBindings_.NewStringType()
	fields["role_assignment_permission_config"] = vapiBindings_.NewReferenceType(nsxModel.RoleAssignmentPermissionConfigBindingType)
	fieldNameMap["role"] = "Role"
	fieldNameMap["role_assignment_permission_config"] = "RoleAssignmentPermissionConfig"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func RoleAssignmentPermissionConfigUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func roleAssignmentPermissionConfigUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["role"] = vapiBindings_.NewStringType()
	fields["role_assignment_permission_config"] = vapiBindings_.NewReferenceType(nsxModel.RoleAssignmentPermissionConfigBindingType)
	fieldNameMap["role"] = "Role"
	fieldNameMap["role_assignment_permission_config"] = "RoleAssignmentPermissionConfig"
	paramsTypeMap["role"] = vapiBindings_.NewStringType()
	paramsTypeMap["role_assignment_permission_config"] = vapiBindings_.NewReferenceType(nsxModel.RoleAssignmentPermissionConfigBindingType)
	paramsTypeMap["role"] = vapiBindings_.NewStringType()
	pathParams["role"] = "role"
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
		"role_assignment_permission_config",
		"PUT",
		"/api/v1/aaa/roles/{role}/role-assignment-permission-config",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
