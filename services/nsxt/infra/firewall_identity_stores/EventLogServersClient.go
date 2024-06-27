// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: EventLogServers
// Used by client-side stubs.

package firewall_identity_stores

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type EventLogServersClient interface {

	// Delete a Event Log server for Firewall Identity store
	//
	//  Use the following Policy API -
	//  DELETE /infra/identity-firewall-stores/<identity-firewall-store-id>/event-log-servers/<event-log-server-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param firewallIdentityStoreIdParam Firewall Identity store identifier (required)
	// @param eventLogServerIdParam Event Log server identifier (required)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(firewallIdentityStoreIdParam string, eventLogServerIdParam string, enforcementPointPathParam *string) error

	// Get a specific Event Log server for a given Firewall Identity store
	//
	//  Use the following Policy API -
	//  GET /infra/identity-firewall-stores/<identity-firewall-store-id>/event-log-servers/<event-log-server-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param firewallIdentityStoreIdParam Firewall Identity store identifier (required)
	// @param eventLogServerIdParam Event Log server identifier (required)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	// @return com.vmware.nsx_policy.model.DirectoryEventLogServer
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(firewallIdentityStoreIdParam string, eventLogServerIdParam string, enforcementPointPathParam *string) (nsx_policyModel.DirectoryEventLogServer, error)

	// More than one Event Log server can be created and only one event log server is used to synchronize directory objects. If more than one Event Log server is configured, NSX will try all the servers until it is able to successfully connect to one.
	//
	//  Use the following Policy API -
	//  PATCH /infra/identity-firewall-stores/<identity-firewall-store-id>/event-log-servers/<event-log-server-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param firewallIdentityStoreIdParam Firewall Identity store identifier (required)
	// @param eventLogServerIdParam Event Log server identifier (required)
	// @param directoryEventLogServerParam (required)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(firewallIdentityStoreIdParam string, eventLogServerIdParam string, directoryEventLogServerParam nsx_policyModel.DirectoryEventLogServer, enforcementPointPathParam *string) error

	// Update a event log server for Firewall Identity store
	//
	//  Use the following Policy API -
	//  PUT /infra/identity-firewall-stores/<identity-firewall-store-id>/event-log-servers/<event-log-server-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param firewallIdentityStoreIdParam Firewall Identity store identifier (required)
	// @param eventLogServerIdParam Event Log Server identifier (required)
	// @param directoryEventLogServerParam (required)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	// @return com.vmware.nsx_policy.model.DirectoryEventLogServer
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(firewallIdentityStoreIdParam string, eventLogServerIdParam string, directoryEventLogServerParam nsx_policyModel.DirectoryEventLogServer, enforcementPointPathParam *string) (nsx_policyModel.DirectoryEventLogServer, error)
}

type eventLogServersClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewEventLogServersClient(connector vapiProtocolClient_.Connector) *eventLogServersClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.firewall_identity_stores.event_log_servers")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	eIface := eventLogServersClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &eIface
}

func (eIface *eventLogServersClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := eIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (eIface *eventLogServersClient) Delete(firewallIdentityStoreIdParam string, eventLogServerIdParam string, enforcementPointPathParam *string) error {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := eventLogServersDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(eventLogServersDeleteInputType(), typeConverter)
	sv.AddStructField("FirewallIdentityStoreId", firewallIdentityStoreIdParam)
	sv.AddStructField("EventLogServerId", eventLogServerIdParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.firewall_identity_stores.event_log_servers", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (eIface *eventLogServersClient) Get(firewallIdentityStoreIdParam string, eventLogServerIdParam string, enforcementPointPathParam *string) (nsx_policyModel.DirectoryEventLogServer, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := eventLogServersGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(eventLogServersGetInputType(), typeConverter)
	sv.AddStructField("FirewallIdentityStoreId", firewallIdentityStoreIdParam)
	sv.AddStructField("EventLogServerId", eventLogServerIdParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.DirectoryEventLogServer
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.firewall_identity_stores.event_log_servers", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.DirectoryEventLogServer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), EventLogServersGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.DirectoryEventLogServer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *eventLogServersClient) Patch(firewallIdentityStoreIdParam string, eventLogServerIdParam string, directoryEventLogServerParam nsx_policyModel.DirectoryEventLogServer, enforcementPointPathParam *string) error {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := eventLogServersPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(eventLogServersPatchInputType(), typeConverter)
	sv.AddStructField("FirewallIdentityStoreId", firewallIdentityStoreIdParam)
	sv.AddStructField("EventLogServerId", eventLogServerIdParam)
	sv.AddStructField("DirectoryEventLogServer", directoryEventLogServerParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.firewall_identity_stores.event_log_servers", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (eIface *eventLogServersClient) Update(firewallIdentityStoreIdParam string, eventLogServerIdParam string, directoryEventLogServerParam nsx_policyModel.DirectoryEventLogServer, enforcementPointPathParam *string) (nsx_policyModel.DirectoryEventLogServer, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := eventLogServersUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(eventLogServersUpdateInputType(), typeConverter)
	sv.AddStructField("FirewallIdentityStoreId", firewallIdentityStoreIdParam)
	sv.AddStructField("EventLogServerId", eventLogServerIdParam)
	sv.AddStructField("DirectoryEventLogServer", directoryEventLogServerParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.DirectoryEventLogServer
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.firewall_identity_stores.event_log_servers", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.DirectoryEventLogServer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), EventLogServersUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.DirectoryEventLogServer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
