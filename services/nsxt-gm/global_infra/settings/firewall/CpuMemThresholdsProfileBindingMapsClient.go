// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: CpuMemThresholdsProfileBindingMaps
// Used by client-side stubs.

package firewall

import (
	vapiStdErrors_ "github.com/zhengxiexie/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_global_policyModel "github.com/zhengxiexie/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type CpuMemThresholdsProfileBindingMapsClient interface {

	// API will delete Firewall CPU Memory Thresholds Profile Binding.
	//
	// @param cpuMemThresholdsProfileBindingMapIdParam Firewall CPU Memory Thresholds Profile Binding Map ID (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(cpuMemThresholdsProfileBindingMapIdParam string) error

	// API will list all Firewall CPU Memory Thresholds Profile Binding Maps.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_global_policy.model.PolicyFirewallCPUMemThresholdsProfileBindingMapListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_global_policyModel.PolicyFirewallCPUMemThresholdsProfileBindingMapListResult, error)
}

type cpuMemThresholdsProfileBindingMapsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewCpuMemThresholdsProfileBindingMapsClient(connector vapiProtocolClient_.Connector) *cpuMemThresholdsProfileBindingMapsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_global_policy.global_infra.settings.firewall.cpu_mem_thresholds_profile_binding_maps")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	cIface := cpuMemThresholdsProfileBindingMapsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *cpuMemThresholdsProfileBindingMapsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *cpuMemThresholdsProfileBindingMapsClient) Delete(cpuMemThresholdsProfileBindingMapIdParam string) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := cpuMemThresholdsProfileBindingMapsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(cpuMemThresholdsProfileBindingMapsDeleteInputType(), typeConverter)
	sv.AddStructField("CpuMemThresholdsProfileBindingMapId", cpuMemThresholdsProfileBindingMapIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.settings.firewall.cpu_mem_thresholds_profile_binding_maps", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (cIface *cpuMemThresholdsProfileBindingMapsClient) List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_global_policyModel.PolicyFirewallCPUMemThresholdsProfileBindingMapListResult, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := cpuMemThresholdsProfileBindingMapsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(cpuMemThresholdsProfileBindingMapsListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.PolicyFirewallCPUMemThresholdsProfileBindingMapListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.settings.firewall.cpu_mem_thresholds_profile_binding_maps", "list", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.PolicyFirewallCPUMemThresholdsProfileBindingMapListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CpuMemThresholdsProfileBindingMapsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.PolicyFirewallCPUMemThresholdsProfileBindingMapListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
