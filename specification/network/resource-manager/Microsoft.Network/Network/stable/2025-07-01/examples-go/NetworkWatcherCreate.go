package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkWatcherCreate.json
func ExampleWatchersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWatchersClient().CreateOrUpdate(ctx, "rg1", "nw1", armnetwork.Watcher{
		Location:   to.Ptr("eastus"),
		Properties: &armnetwork.WatcherPropertiesFormat{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.WatchersClientCreateOrUpdateResponse{
	// 	Watcher: armnetwork.Watcher{
	// 		Name: to.Ptr("nw1"),
	// 		Type: to.Ptr("Microsoft.Network/networkWatchers"),
	// 		Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkWatchers/nw1"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armnetwork.WatcherPropertiesFormat{
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
