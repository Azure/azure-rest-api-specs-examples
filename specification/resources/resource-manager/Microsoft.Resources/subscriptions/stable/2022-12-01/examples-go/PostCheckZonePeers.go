package armsubscriptions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions/v2"
)

// Generated from example definition: 2022-12-01/PostCheckZonePeers.json
func ExampleClient_CheckZonePeers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsubscriptions.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().CheckZonePeers(ctx, "8d65815f-a5b6-402f-9298-045155da7d74", armsubscriptions.CheckZonePeersRequest{
		Location: to.Ptr("eastus"),
		SubscriptionIDs: []*string{
			to.Ptr("subscriptions/11111111-1111-1111-1111-111111111111"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsubscriptions.ClientCheckZonePeersResponse{
	// 	CheckZonePeersResult: armsubscriptions.CheckZonePeersResult{
	// 		AvailabilityZonePeers: []*armsubscriptions.AvailabilityZonePeers{
	// 			{
	// 				AvailabilityZone: to.Ptr("1"),
	// 				Peers: []*armsubscriptions.Peers{
	// 					{
	// 						AvailabilityZone: to.Ptr("3"),
	// 						SubscriptionID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				AvailabilityZone: to.Ptr("2"),
	// 				Peers: []*armsubscriptions.Peers{
	// 					{
	// 						AvailabilityZone: to.Ptr("2"),
	// 						SubscriptionID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				AvailabilityZone: to.Ptr("3"),
	// 				Peers: []*armsubscriptions.Peers{
	// 					{
	// 						AvailabilityZone: to.Ptr("1"),
	// 						SubscriptionID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("eastus2"),
	// 		SubscriptionID: to.Ptr("8d65815f-a5b6-402f-9298-045155da7d74"),
	// 	},
	// }
}
