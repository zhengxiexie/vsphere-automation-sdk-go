// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: MigrationUnits.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package migration

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func migrationUnitsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["migration_unit_id"] = vapiBindings_.NewStringType()
	fieldNameMap["migration_unit_id"] = "MigrationUnitId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func MigrationUnitsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.MigrationUnitBindingType)
}

func migrationUnitsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["migration_unit_id"] = vapiBindings_.NewStringType()
	fieldNameMap["migration_unit_id"] = "MigrationUnitId"
	paramsTypeMap["migration_unit_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["migrationUnitId"] = vapiBindings_.NewStringType()
	pathParams["migration_unit_id"] = "migrationUnitId"
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
		"/api/v1/migration/migration-units/{migrationUnitId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func migrationUnitsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["component_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["current_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["has_warnings"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["metadata"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["migration_unit_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["component_type"] = "ComponentType"
	fieldNameMap["current_version"] = "CurrentVersion"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["has_warnings"] = "HasWarnings"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["metadata"] = "Metadata"
	fieldNameMap["migration_unit_type"] = "MigrationUnitType"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func MigrationUnitsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.MigrationUnitListResultBindingType)
}

func migrationUnitsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["component_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["current_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["has_warnings"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["metadata"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["migration_unit_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["component_type"] = "ComponentType"
	fieldNameMap["current_version"] = "CurrentVersion"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["has_warnings"] = "HasWarnings"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["metadata"] = "Metadata"
	fieldNameMap["migration_unit_type"] = "MigrationUnitType"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["component_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["migration_unit_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["metadata"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["current_version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["has_warnings"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	queryParams["cursor"] = "cursor"
	queryParams["component_type"] = "component_type"
	queryParams["migration_unit_type"] = "migration_unit_type"
	queryParams["metadata"] = "metadata"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["group_id"] = "group_id"
	queryParams["included_fields"] = "included_fields"
	queryParams["current_version"] = "current_version"
	queryParams["has_warnings"] = "has_warnings"
	queryParams["sort_by"] = "sort_by"
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
		"/api/v1/migration/migration-units",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
