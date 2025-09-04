package armazurestackhcivm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhcivm"
)

// Generated from example definition: 2025-06-01-preview/NetworkSecurityGroups_ListByResourceGroup.json
func ExampleNetworkSecurityGroupsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhcivm.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkSecurityGroupsClient().NewListByResourceGroupPager("testrg", nil)
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
		// page = armazurestackhcivm.NetworkSecurityGroupsClientListByResourceGroupResponse{
		// 	NetworkSecurityGroupListResult: armazurestackhcivm.NetworkSecurityGroupListResult{
		// 		Value: []*armazurestackhcivm.NetworkSecurityGroup{
		// 			{
		// 				Name: to.Ptr("nsg1"),
		// 				Type: to.Ptr("Microsoft.AzureStackHCI/networkSecurityGroups"),
		// 				ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.AzureStackHCI/networkSecurityGroups/nsg1"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armazurestackhcivm.NetworkSecurityGroupProperties{
		// 					ProvisioningState: to.Ptr(armazurestackhcivm.ProvisioningStateEnumSucceeded),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("nsg3"),
		// 				Type: to.Ptr("Microsoft.AzureStackHCI/networkSecurityGroups"),
		// 				ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.AzureStackHCI/networkSecurityGroups/nsg3"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armazurestackhcivm.NetworkSecurityGroupProperties{
		// 					ProvisioningState: to.Ptr(armazurestackhcivm.ProvisioningStateEnumSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
