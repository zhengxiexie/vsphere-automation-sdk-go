// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: RoutingTable.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package transit_gateways

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for ``componentType`` of method RoutingTable#list.
const RoutingTable_LIST_COMPONENT_TYPE_ROUTES = "DR_ROUTES"

// Possible value for ``routeSource`` of method RoutingTable#list.
const RoutingTable_LIST_ROUTE_SOURCE_BGP = "BGP"

// Possible value for ``routeSource`` of method RoutingTable#list.
const RoutingTable_LIST_ROUTE_SOURCE_STATIC = "STATIC"

// Possible value for ``routeSource`` of method RoutingTable#list.
const RoutingTable_LIST_ROUTE_SOURCE_CONNECTED = "CONNECTED"

// Possible value for ``routeSource`` of method RoutingTable#list.
const RoutingTable_LIST_ROUTE_SOURCE_OSPF = "OSPF"

func routingTableListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["transit_gateway_id"] = vapiBindings_.NewStringType()
	fields["component_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["edge_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["edge_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["host_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["host_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["network_prefix"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["route_source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["transit_gateway_id"] = "TransitGatewayId"
	fieldNameMap["component_type"] = "ComponentType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["edge_id"] = "EdgeId"
	fieldNameMap["edge_path"] = "EdgePath"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	fieldNameMap["host_id"] = "HostId"
	fieldNameMap["host_path"] = "HostPath"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["network_prefix"] = "NetworkPrefix"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["route_source"] = "RouteSource"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func RoutingTableListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.RoutingTableListResultBindingType)
}

func routingTableListRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["transit_gateway_id"] = vapiBindings_.NewStringType()
	fields["component_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["edge_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["edge_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["host_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["host_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["network_prefix"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["route_source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["transit_gateway_id"] = "TransitGatewayId"
	fieldNameMap["component_type"] = "ComponentType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["edge_id"] = "EdgeId"
	fieldNameMap["edge_path"] = "EdgePath"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	fieldNameMap["host_id"] = "HostId"
	fieldNameMap["host_path"] = "HostPath"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["network_prefix"] = "NetworkPrefix"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["route_source"] = "RouteSource"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["host_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["transit_gateway_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["network_prefix"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["edge_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["host_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["component_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["edge_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["route_source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["transitGatewayId"] = vapiBindings_.NewStringType()
	pathParams["transit_gateway_id"] = "transitGatewayId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	queryParams["cursor"] = "cursor"
	queryParams["host_path"] = "host_path"
	queryParams["network_prefix"] = "network_prefix"
	queryParams["edge_id"] = "edge_id"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
	queryParams["sort_by"] = "sort_by"
	queryParams["host_id"] = "host_id"
	queryParams["component_type"] = "component_type"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["edge_path"] = "edge_path"
	queryParams["route_source"] = "route_source"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/transit-gateways/{transitGatewayId}/routing-table",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
