// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Continue.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package action

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func continueCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vcenter_id"] = vapiBindings_.NewStringType()
	fields["workflow_request_parameters"] = vapiBindings_.NewReferenceType(nsxModel.WorkflowRequestParametersBindingType)
	fieldNameMap["vcenter_id"] = "VcenterId"
	fieldNameMap["workflow_request_parameters"] = "WorkflowRequestParameters"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ContinueCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func continueCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["vcenter_id"] = vapiBindings_.NewStringType()
	fields["workflow_request_parameters"] = vapiBindings_.NewReferenceType(nsxModel.WorkflowRequestParametersBindingType)
	fieldNameMap["vcenter_id"] = "VcenterId"
	fieldNameMap["workflow_request_parameters"] = "WorkflowRequestParameters"
	paramsTypeMap["vcenter_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["workflow_request_parameters"] = vapiBindings_.NewReferenceType(nsxModel.WorkflowRequestParametersBindingType)
	paramsTypeMap["vcenterId"] = vapiBindings_.NewStringType()
	pathParams["vcenter_id"] = "vcenterId"
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
		"workflow_request_parameters",
		"POST",
		"/api/v1/migration/overlay-adoption/vcenters/{vcenterId}/workflow/action/continue",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
