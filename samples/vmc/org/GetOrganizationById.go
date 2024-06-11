// Copyright © 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package org

import (
	"fmt"

	"github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/zhengxiexie/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_org/api/organization/core"
	"github.com/zhengxiexie/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_org/model"
)

func GetOrganizationById(orgId string, connector client.Connector) model.Organization {

	// API call to GetOrganisationByID
	orgsClient := core.NewOrgsClient(connector)
	organization, err := orgsClient.GetOrganizationById(orgId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Name: ", organization.Name)
	fmt.Println("Id: ", organization.Id)

	return organization
}
