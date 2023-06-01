package armsqlvirtualmachine_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e79a5aa63c0551c1b5af1d2853cceb495283d334/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/ListByResourceGroupSqlVirtualMachineGroup.json
func ExampleGroupsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsqlvirtualmachine.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGroupsClient().NewListByResourceGroupPager("testrg", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.GroupListResult = armsqlvirtualmachine.GroupListResult{
		// 	Value: []*armsqlvirtualmachine.Group{
		// 		{
		// 			Name: to.Ptr("testvmgroup"),
		// 			Type: to.Ptr("Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/testvmgroup"),
		// 			Location: to.Ptr("northeurope"),
		// 			Tags: map[string]*string{
		// 				"mytag": to.Ptr("myval"),
		// 			},
		// 			Properties: &armsqlvirtualmachine.GroupProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SQLImageOffer: to.Ptr("SQL2017-WS2016"),
		// 				SQLImageSKU: to.Ptr(armsqlvirtualmachine.SQLVMGroupImageSKUEnterprise),
		// 				WsfcDomainProfile: &armsqlvirtualmachine.WsfcDomainProfile{
		// 					ClusterBootstrapAccount: to.Ptr("testrpadmin"),
		// 					ClusterOperatorAccount: to.Ptr("testrp@testdomain.com"),
		// 					DomainFqdn: to.Ptr("testdomain.com"),
		// 					OuPath: to.Ptr("OU=WSCluster,DC=testdomain,DC=com"),
		// 					SQLServiceAccount: to.Ptr("sqlservice@testdomain.com"),
		// 					StorageAccountURL: to.Ptr("https://storgact.blob.core.windows.net/"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testvmgroup1"),
		// 			Type: to.Ptr("Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/testvmgroup1"),
		// 			Location: to.Ptr("northeurope"),
		// 			Tags: map[string]*string{
		// 				"mytag": to.Ptr("myval"),
		// 			},
		// 			Properties: &armsqlvirtualmachine.GroupProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SQLImageOffer: to.Ptr("SQL2016-WS2016"),
		// 				SQLImageSKU: to.Ptr(armsqlvirtualmachine.SQLVMGroupImageSKUEnterprise),
		// 				WsfcDomainProfile: &armsqlvirtualmachine.WsfcDomainProfile{
		// 					ClusterBootstrapAccount: to.Ptr("testrpadmin"),
		// 					ClusterOperatorAccount: to.Ptr("testrp@testdomain.com"),
		// 					DomainFqdn: to.Ptr("testdomain.com"),
		// 					OuPath: to.Ptr("OU=WSCluster,DC=testdomain,DC=com"),
		// 					SQLServiceAccount: to.Ptr("sqlservice@testdomain.com"),
		// 					StorageAccountURL: to.Ptr("https://storgact.blob.core.windows.net/"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testvmgroup2"),
		// 			Type: to.Ptr("Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/testvmgroup2"),
		// 			Location: to.Ptr("northeurope"),
		// 			Tags: map[string]*string{
		// 				"mytag": to.Ptr("myval"),
		// 			},
		// 			Properties: &armsqlvirtualmachine.GroupProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SQLImageOffer: to.Ptr("SQL2016-WS2016"),
		// 				SQLImageSKU: to.Ptr(armsqlvirtualmachine.SQLVMGroupImageSKUEnterprise),
		// 				WsfcDomainProfile: &armsqlvirtualmachine.WsfcDomainProfile{
		// 					ClusterBootstrapAccount: to.Ptr("testrpadmin"),
		// 					ClusterOperatorAccount: to.Ptr("testrp@testdomain.com"),
		// 					DomainFqdn: to.Ptr("testdomain.com"),
		// 					OuPath: to.Ptr("OU=WSCluster,DC=testdomain,DC=com"),
		// 					SQLServiceAccount: to.Ptr("sqlservice@testdomain.com"),
		// 					StorageAccountURL: to.Ptr("https://storgact.blob.core.windows.net/"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
