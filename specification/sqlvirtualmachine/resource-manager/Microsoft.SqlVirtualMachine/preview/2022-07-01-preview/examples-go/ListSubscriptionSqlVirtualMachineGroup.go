package armsqlvirtualmachine_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/ListSubscriptionSqlVirtualMachineGroup.json
func ExampleGroupsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsqlvirtualmachine.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGroupsClient().NewListPager(nil)
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
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/"),
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
		// 					ClusterSubnetType: to.Ptr(armsqlvirtualmachine.ClusterSubnetTypeMultiSubnet),
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
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg1/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/"),
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
		// 					ClusterSubnetType: to.Ptr(armsqlvirtualmachine.ClusterSubnetTypeMultiSubnet),
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
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg2/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/"),
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
		// 					ClusterSubnetType: to.Ptr(armsqlvirtualmachine.ClusterSubnetTypeMultiSubnet),
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
