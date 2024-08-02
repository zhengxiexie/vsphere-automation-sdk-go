// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: EdgeClusterHighAvailabilityProfiles.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package enforcement_points

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func edgeClusterHighAvailabilityProfilesDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["edge_cluster_high_availability_profile_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["edge_cluster_high_availability_profile_id"] = "EdgeClusterHighAvailabilityProfileId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EdgeClusterHighAvailabilityProfilesDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func edgeClusterHighAvailabilityProfilesDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["edge_cluster_high_availability_profile_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["edge_cluster_high_availability_profile_id"] = "EdgeClusterHighAvailabilityProfileId"
	paramsTypeMap["enforcementpoint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["edge_cluster_high_availability_profile_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementpointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["edgeClusterHighAvailabilityProfileId"] = vapiBindings_.NewStringType()
	pathParams["enforcementpoint_id"] = "enforcementpointId"
	pathParams["edge_cluster_high_availability_profile_id"] = "edgeClusterHighAvailabilityProfileId"
	pathParams["site_id"] = "siteId"
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
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementpointId}/edge-cluster-high-availability-profiles/{edgeClusterHighAvailabilityProfileId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func edgeClusterHighAvailabilityProfilesGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["edge_cluster_high_availability_profile_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["edge_cluster_high_availability_profile_id"] = "EdgeClusterHighAvailabilityProfileId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EdgeClusterHighAvailabilityProfilesGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PolicyEdgeHighAvailabilityProfileBindingType)
}

func edgeClusterHighAvailabilityProfilesGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["edge_cluster_high_availability_profile_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["edge_cluster_high_availability_profile_id"] = "EdgeClusterHighAvailabilityProfileId"
	paramsTypeMap["enforcementpoint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["edge_cluster_high_availability_profile_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementpointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["edgeClusterHighAvailabilityProfileId"] = vapiBindings_.NewStringType()
	pathParams["enforcementpoint_id"] = "enforcementpointId"
	pathParams["edge_cluster_high_availability_profile_id"] = "edgeClusterHighAvailabilityProfileId"
	pathParams["site_id"] = "siteId"
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
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementpointId}/edge-cluster-high-availability-profiles/{edgeClusterHighAvailabilityProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func edgeClusterHighAvailabilityProfilesPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["edge_cluster_high_availability_profile_id"] = vapiBindings_.NewStringType()
	fields["policy_edge_high_availability_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyEdgeHighAvailabilityProfileBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["edge_cluster_high_availability_profile_id"] = "EdgeClusterHighAvailabilityProfileId"
	fieldNameMap["policy_edge_high_availability_profile"] = "PolicyEdgeHighAvailabilityProfile"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EdgeClusterHighAvailabilityProfilesPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func edgeClusterHighAvailabilityProfilesPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["edge_cluster_high_availability_profile_id"] = vapiBindings_.NewStringType()
	fields["policy_edge_high_availability_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyEdgeHighAvailabilityProfileBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["edge_cluster_high_availability_profile_id"] = "EdgeClusterHighAvailabilityProfileId"
	fieldNameMap["policy_edge_high_availability_profile"] = "PolicyEdgeHighAvailabilityProfile"
	paramsTypeMap["enforcementpoint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["policy_edge_high_availability_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyEdgeHighAvailabilityProfileBindingType)
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["edge_cluster_high_availability_profile_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementpointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["edgeClusterHighAvailabilityProfileId"] = vapiBindings_.NewStringType()
	pathParams["enforcementpoint_id"] = "enforcementpointId"
	pathParams["edge_cluster_high_availability_profile_id"] = "edgeClusterHighAvailabilityProfileId"
	pathParams["site_id"] = "siteId"
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
		"policy_edge_high_availability_profile",
		"PATCH",
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementpointId}/edge-cluster-high-availability-profiles/{edgeClusterHighAvailabilityProfileId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func edgeClusterHighAvailabilityProfilesUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["edge_cluster_high_availability_profile_id"] = vapiBindings_.NewStringType()
	fields["policy_edge_high_availability_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyEdgeHighAvailabilityProfileBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["edge_cluster_high_availability_profile_id"] = "EdgeClusterHighAvailabilityProfileId"
	fieldNameMap["policy_edge_high_availability_profile"] = "PolicyEdgeHighAvailabilityProfile"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EdgeClusterHighAvailabilityProfilesUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PolicyEdgeHighAvailabilityProfileBindingType)
}

func edgeClusterHighAvailabilityProfilesUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["edge_cluster_high_availability_profile_id"] = vapiBindings_.NewStringType()
	fields["policy_edge_high_availability_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyEdgeHighAvailabilityProfileBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["edge_cluster_high_availability_profile_id"] = "EdgeClusterHighAvailabilityProfileId"
	fieldNameMap["policy_edge_high_availability_profile"] = "PolicyEdgeHighAvailabilityProfile"
	paramsTypeMap["enforcementpoint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["policy_edge_high_availability_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyEdgeHighAvailabilityProfileBindingType)
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["edge_cluster_high_availability_profile_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementpointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["edgeClusterHighAvailabilityProfileId"] = vapiBindings_.NewStringType()
	pathParams["enforcementpoint_id"] = "enforcementpointId"
	pathParams["edge_cluster_high_availability_profile_id"] = "edgeClusterHighAvailabilityProfileId"
	pathParams["site_id"] = "siteId"
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
		"policy_edge_high_availability_profile",
		"PUT",
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementpointId}/edge-cluster-high-availability-profiles/{edgeClusterHighAvailabilityProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
