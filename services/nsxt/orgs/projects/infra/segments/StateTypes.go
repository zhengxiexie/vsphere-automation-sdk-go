// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: State.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package segments

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for “bgpNeighborType“ of method State#get.
const State_GET_BGP_NEIGHBOR_TYPE_INTER_SR = "INTER_SR"

// Possible value for “bgpNeighborType“ of method State#get.
const State_GET_BGP_NEIGHBOR_TYPE_USER = "USER"

// Possible value for “source“ of method State#get.
const State_GET_SOURCE_REALTIME = "realtime"

// Possible value for “source“ of method State#get.
const State_GET_SOURCE_CACHED = "cached"

// Possible value for “statsType“ of method State#get.
const State_GET_STATS_TYPE_STATS = "DATAPATH_STATS"

// Possible value for “configurationState“ of method State#list.
const State_LIST_CONFIGURATION_STATE_PENDING = "pending"

// Possible value for “configurationState“ of method State#list.
const State_LIST_CONFIGURATION_STATE_IN_PROGRESS = "in_progress"

// Possible value for “configurationState“ of method State#list.
const State_LIST_CONFIGURATION_STATE_SUCCESS = "success"

// Possible value for “configurationState“ of method State#list.
const State_LIST_CONFIGURATION_STATE_FAILED = "failed"

// Possible value for “configurationState“ of method State#list.
const State_LIST_CONFIGURATION_STATE_PARTIAL_SUCCESS = "partial_success"

// Possible value for “configurationState“ of method State#list.
const State_LIST_CONFIGURATION_STATE_ORPHANED = "orphaned"

// Possible value for “configurationState“ of method State#list.
const State_LIST_CONFIGURATION_STATE_UNKNOWN = "unknown"

// Possible value for “source“ of method State#list.
const State_LIST_SOURCE_REALTIME = "realtime"

// Possible value for “source“ of method State#list.
const State_LIST_SOURCE_CACHED = "cached"

func stateGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["segments_id"] = vapiBindings_.NewStringType()
	fields["bgp_neighbor_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["edge_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["stats_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["transport_node_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segments_id"] = "SegmentsId"
	fieldNameMap["bgp_neighbor_type"] = "BgpNeighborType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["edge_path"] = "EdgePath"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["stats_type"] = "StatsType"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func StateGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.SegmentConfigurationStateBindingType)
}

func stateGetRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["segments_id"] = vapiBindings_.NewStringType()
	fields["bgp_neighbor_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["edge_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["stats_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["transport_node_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segments_id"] = "SegmentsId"
	fieldNameMap["bgp_neighbor_type"] = "BgpNeighborType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["edge_path"] = "EdgePath"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["stats_type"] = "StatsType"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["stats_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["transport_node_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["edge_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["segments_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["bgp_neighbor_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentsId"] = vapiBindings_.NewStringType()
	pathParams["segments_id"] = "segmentsId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["edge_path"] = "edge_path"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
	queryParams["stats_type"] = "stats_type"
	queryParams["sort_by"] = "sort_by"
	queryParams["source"] = "source"
	queryParams["transport_node_id"] = "transport_node_id"
	queryParams["include_mark_for_delete_objects"] = "include_mark_for_delete_objects"
	queryParams["bgp_neighbor_type"] = "bgp_neighbor_type"
	queryParams["page_size"] = "page_size"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments/{segmentsId}/state",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func stateListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["configuration_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["configuration_state"] = "ConfigurationState"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	fieldNameMap["source"] = "Source"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func StateListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.SegmentConfigurationStateListResultBindingType)
}

func stateListRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["configuration_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["configuration_state"] = "ConfigurationState"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["configuration_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	queryParams["configuration_state"] = "configuration_state"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
	queryParams["source"] = "source"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments/state",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
