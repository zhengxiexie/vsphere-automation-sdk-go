// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: EdgeClusterHighAvailabilityProfiles
// Used by client-side stubs.

package enforcement_points

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type EdgeClusterHighAvailabilityProfilesClient interface {

	// Delete a edge cluster high availability profile.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param edgeClusterHighAvailabilityProfileIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(siteIdParam string, enforcementpointIdParam string, edgeClusterHighAvailabilityProfileIdParam string) error

	// Get a edge cluster high availability profile.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param edgeClusterHighAvailabilityProfileIdParam (required)
	// @return com.vmware.nsx_policy.model.PolicyEdgeHighAvailabilityProfile
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(siteIdParam string, enforcementpointIdParam string, edgeClusterHighAvailabilityProfileIdParam string) (nsx_policyModel.PolicyEdgeHighAvailabilityProfile, error)

	// Creates a new edge cluster high availability profile.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param edgeClusterHighAvailabilityProfileIdParam (required)
	// @param policyEdgeHighAvailabilityProfileParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(siteIdParam string, enforcementpointIdParam string, edgeClusterHighAvailabilityProfileIdParam string, policyEdgeHighAvailabilityProfileParam nsx_policyModel.PolicyEdgeHighAvailabilityProfile) error

	// Create or Update a edge cluster high availability profile.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param edgeClusterHighAvailabilityProfileIdParam (required)
	// @param policyEdgeHighAvailabilityProfileParam (required)
	// @return com.vmware.nsx_policy.model.PolicyEdgeHighAvailabilityProfile
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(siteIdParam string, enforcementpointIdParam string, edgeClusterHighAvailabilityProfileIdParam string, policyEdgeHighAvailabilityProfileParam nsx_policyModel.PolicyEdgeHighAvailabilityProfile) (nsx_policyModel.PolicyEdgeHighAvailabilityProfile, error)
}

type edgeClusterHighAvailabilityProfilesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewEdgeClusterHighAvailabilityProfilesClient(connector vapiProtocolClient_.Connector) *edgeClusterHighAvailabilityProfilesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.enforcement_points.edge_cluster_high_availability_profiles")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	eIface := edgeClusterHighAvailabilityProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &eIface
}

func (eIface *edgeClusterHighAvailabilityProfilesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := eIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (eIface *edgeClusterHighAvailabilityProfilesClient) Delete(siteIdParam string, enforcementpointIdParam string, edgeClusterHighAvailabilityProfileIdParam string) error {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := edgeClusterHighAvailabilityProfilesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(edgeClusterHighAvailabilityProfilesDeleteInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("EdgeClusterHighAvailabilityProfileId", edgeClusterHighAvailabilityProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.edge_cluster_high_availability_profiles", "delete", inputDataValue, executionContext)
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

func (eIface *edgeClusterHighAvailabilityProfilesClient) Get(siteIdParam string, enforcementpointIdParam string, edgeClusterHighAvailabilityProfileIdParam string) (nsx_policyModel.PolicyEdgeHighAvailabilityProfile, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := edgeClusterHighAvailabilityProfilesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(edgeClusterHighAvailabilityProfilesGetInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("EdgeClusterHighAvailabilityProfileId", edgeClusterHighAvailabilityProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyEdgeHighAvailabilityProfile
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.edge_cluster_high_availability_profiles", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyEdgeHighAvailabilityProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), EdgeClusterHighAvailabilityProfilesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyEdgeHighAvailabilityProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *edgeClusterHighAvailabilityProfilesClient) Patch(siteIdParam string, enforcementpointIdParam string, edgeClusterHighAvailabilityProfileIdParam string, policyEdgeHighAvailabilityProfileParam nsx_policyModel.PolicyEdgeHighAvailabilityProfile) error {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := edgeClusterHighAvailabilityProfilesPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(edgeClusterHighAvailabilityProfilesPatchInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("EdgeClusterHighAvailabilityProfileId", edgeClusterHighAvailabilityProfileIdParam)
	sv.AddStructField("PolicyEdgeHighAvailabilityProfile", policyEdgeHighAvailabilityProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.edge_cluster_high_availability_profiles", "patch", inputDataValue, executionContext)
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

func (eIface *edgeClusterHighAvailabilityProfilesClient) Update(siteIdParam string, enforcementpointIdParam string, edgeClusterHighAvailabilityProfileIdParam string, policyEdgeHighAvailabilityProfileParam nsx_policyModel.PolicyEdgeHighAvailabilityProfile) (nsx_policyModel.PolicyEdgeHighAvailabilityProfile, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := edgeClusterHighAvailabilityProfilesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(edgeClusterHighAvailabilityProfilesUpdateInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("EdgeClusterHighAvailabilityProfileId", edgeClusterHighAvailabilityProfileIdParam)
	sv.AddStructField("PolicyEdgeHighAvailabilityProfile", policyEdgeHighAvailabilityProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyEdgeHighAvailabilityProfile
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.edge_cluster_high_availability_profiles", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyEdgeHighAvailabilityProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), EdgeClusterHighAvailabilityProfilesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyEdgeHighAvailabilityProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
