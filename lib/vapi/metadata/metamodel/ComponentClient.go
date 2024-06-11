// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Component
// Used by client-side stubs.

package metamodel

import (
	vapiStdErrors_ "github.com/zhengxiexie/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

// The ``Component`` interface providers methods to retrieve metamodel information of a component element.
//
//  A component defines a set of functionality that is deployed together and versioned together. For example, all the interfaces that belong to VMware Content Library are part of a single component. A component element describes a component. A component element contains one or more package elements.
//
//  The methods for package elements are provided by interface Package.
type ComponentClient interface {

	// Returns the identifiers for the component elements that are registered with the infrastructure.
	// @return The list of identifiers for the component elements that are registered with the infrastructure.
	// The return value will contain identifiers for the resource type: ``com.vmware.vapi.component``.
	List() ([]string, error)

	// Retrieves metamodel information about the component element corresponding to ``component_id``.
	//
	//  The ComponentData contains the metamodel information about the component and it's fingerprint. It contains information about all the package elements that are contained in this component element.
	//
	// @param componentIdParam Identifier of the component element.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.component``.
	// @return The ComponentData instance that corresponds to ``component_id``.
	//
	// @throws NotFound if the component element associated with ``component_id`` is not registered with the infrastructure.
	Get(componentIdParam string) (ComponentData, error)

	// Retrieves the fingerprint computed from the metamodel metadata of the component element corresponding to ``component_id``.
	//
	//  The fingerprint provides clients an efficient way to check if the metadata for a particular component element has been modified on the server. The client can do this by comparing the result of this operation with the fingerprint returned in the result of Component#get.
	//
	// @param componentIdParam Identifier of the component element.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.component``.
	// @return The fingerprint computed from the metamodel metadata of the component element.
	//
	// @throws NotFound if the component element associated with ``component_id`` is not registered with the infrastructure.
	Fingerprint(componentIdParam string) (string, error)
}

type componentClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewComponentClient(connector vapiProtocolClient_.Connector) *componentClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vapi.metadata.metamodel.component")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"list":        vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"get":         vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"fingerprint": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "fingerprint"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	cIface := componentClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *componentClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *componentClient) List() ([]string, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := componentListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(componentListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []string
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.metamodel.component", "list", inputDataValue, executionContext)
	var emptyOutput []string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ComponentListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.([]string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *componentClient) Get(componentIdParam string) (ComponentData, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := componentGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(componentGetInputType(), typeConverter)
	sv.AddStructField("ComponentId", componentIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput ComponentData
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.metamodel.component", "get", inputDataValue, executionContext)
	var emptyOutput ComponentData
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ComponentGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(ComponentData), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *componentClient) Fingerprint(componentIdParam string) (string, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := componentFingerprintRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(componentFingerprintInputType(), typeConverter)
	sv.AddStructField("ComponentId", componentIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput string
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.metamodel.component", "fingerprint", inputDataValue, executionContext)
	var emptyOutput string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ComponentFingerprintOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
