// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Statistics
// Used by client-side stubs.

package sessions

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type StatisticsClient interface {

	// Resets the statistics of the given VPN session. Since source of data is enforcement point, data is reset there.
	//
	// @param tier0IdParam (required)
	// @param serviceIdParam (required)
	// @param sessionIdParam (required)
	// @param actionParam Action on statistics (required)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(tier0IdParam string, serviceIdParam string, sessionIdParam string, actionParam string, enforcementPointPathParam *string) error

	// - no enforcement point path specified: statistics are evaluated on each enforcement point. - an enforcement point path is specified: statistics are evaluated only on the given enforcement point. - source=realtime: statistics are fetched realtime from the enforcement point. - source=cached: cached statistics are not supported for IPSec VPN session statistics. Statistics will always be fetched realtime even if source=cached.
	//
	// @param tier0IdParam (required)
	// @param serviceIdParam (required)
	// @param sessionIdParam (required)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.nsx_policy.model.AggregateIPSecVpnSessionStatistics
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(tier0IdParam string, serviceIdParam string, sessionIdParam string, enforcementPointPathParam *string, sourceParam *string) (nsx_policyModel.AggregateIPSecVpnSessionStatistics, error)
}

type statisticsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewStatisticsClient(connector vapiProtocolClient_.Connector) *statisticsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.tier_0s.ipsec_vpn_services.sessions.statistics")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := statisticsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *statisticsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *statisticsClient) Create(tier0IdParam string, serviceIdParam string, sessionIdParam string, actionParam string, enforcementPointPathParam *string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := statisticsCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(statisticsCreateInputType(), typeConverter)
	sv.AddStructField("Tier0Id", tier0IdParam)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("SessionId", sessionIdParam)
	sv.AddStructField("Action", actionParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.tier_0s.ipsec_vpn_services.sessions.statistics", "create", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *statisticsClient) Get(tier0IdParam string, serviceIdParam string, sessionIdParam string, enforcementPointPathParam *string, sourceParam *string) (nsx_policyModel.AggregateIPSecVpnSessionStatistics, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := statisticsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(statisticsGetInputType(), typeConverter)
	sv.AddStructField("Tier0Id", tier0IdParam)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("SessionId", sessionIdParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.AggregateIPSecVpnSessionStatistics
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.tier_0s.ipsec_vpn_services.sessions.statistics", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.AggregateIPSecVpnSessionStatistics
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), StatisticsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.AggregateIPSecVpnSessionStatistics), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
