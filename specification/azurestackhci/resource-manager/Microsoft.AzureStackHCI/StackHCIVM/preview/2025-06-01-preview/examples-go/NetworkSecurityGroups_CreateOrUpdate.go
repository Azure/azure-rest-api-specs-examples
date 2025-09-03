package armazurestackhcivm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhcivm"
)

// Generated from example definition: 2025-06-01-preview/NetworkSecurityGroups_CreateOrUpdate.json
func ExampleNetworkSecurityGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhcivm.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkSecurityGroupsClient().BeginCreateOrUpdate(ctx, "testrg", "testnsg", armazurestackhcivm.NetworkSecurityGroup{
		Location: to.Ptr("eastus"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armazurestackhcivm.NetworkSecurityGroupsClientCreateOrUpdateResponse{
	// 	NetworkSecurityGroup: &armazurestackhcivm.NetworkSecurityGroup{
	// 		Name: to.Ptr("testnsg"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/networkSecurityGroups"),
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.AzureStackHCI/networkSecurityGroups/testnsg"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armazurestackhcivm.NetworkSecurityGroupProperties{
	// 			ProvisioningState: to.Ptr(armazurestackhcivm.ProvisioningStateEnumSucceeded),
	// 		},
	// 	},
	// }
}
