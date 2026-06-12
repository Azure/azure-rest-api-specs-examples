package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkWatcherFlowLogConfigure.json
func ExampleWatchersClient_BeginSetFlowLogConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWatchersClient().BeginSetFlowLogConfiguration(ctx, "rg1", "nw1", armnetwork.FlowLogInformation{
		Identity: &armnetwork.ManagedServiceIdentity{
			Type: to.Ptr(armnetwork.ResourceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armnetwork.ManagedServiceIdentityUserAssignedIdentities{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
			},
		},
		Properties: &armnetwork.FlowLogProperties{
			Enabled:   to.Ptr(true),
			StorageID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/st1"),
		},
		TargetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.WatchersClientSetFlowLogConfigurationResponse{
	// 	FlowLogInformation: armnetwork.FlowLogInformation{
	// 		Identity: &armnetwork.ManagedServiceIdentity{
	// 			Type: to.Ptr(armnetwork.ResourceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armnetwork.ManagedServiceIdentityUserAssignedIdentities{
	// 				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": &armnetwork.ManagedServiceIdentityUserAssignedIdentities{
	// 					ClientID: to.Ptr("c16d15e1-f60a-40e4-8a05-df3d3f655c14"),
	// 					PrincipalID: to.Ptr("e3858881-e40c-43bd-9cde-88da39c05023"),
	// 				},
	// 			},
	// 		},
	// 		Properties: &armnetwork.FlowLogProperties{
	// 			Enabled: to.Ptr(true),
	// 			StorageID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/st1"),
	// 		},
	// 		TargetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1"),
	// 	},
	// }
}
