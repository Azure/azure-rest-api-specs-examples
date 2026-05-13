package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering/v2"
)

// Generated from example definition: 2025-05-01/UpdatePeeringServiceTags.json
func ExampleServicesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("subId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().Update(ctx, "rgName", "peeringServiceName", armpeering.ResourceTags{
		Tags: map[string]*string{
			"tag0": to.Ptr("value0"),
			"tag1": to.Ptr("value1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpeering.ServicesClientUpdateResponse{
	// 	Service: &armpeering.Service{
	// 		Name: to.Ptr("peeringServiceName"),
	// 		Type: to.Ptr("Microsoft.Peering/peeringServices"),
	// 		ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peeringServices/peeringServiceName"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armpeering.ServiceProperties{
	// 			LogAnalyticsWorkspaceProperties: &armpeering.LogAnalyticsWorkspaceProperties{
	// 				ConnectedAgents: []*string{
	// 					to.Ptr("Agent1"),
	// 					to.Ptr("Agent2"),
	// 				},
	// 				Key: to.Ptr("key"),
	// 				WorkspaceID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			},
	// 			PeeringServiceLocation: to.Ptr("state1"),
	// 			PeeringServiceProvider: to.Ptr("serviceProvider1"),
	// 			ProviderBackupPeeringLocation: to.Ptr("peeringLocation2"),
	// 			ProviderPrimaryPeeringLocation: to.Ptr("peeringLocation1"),
	// 			ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 			"tag0": to.Ptr("value0"),
	// 			"tag1": to.Ptr("value1"),
	// 		},
	// 	},
	// }
}
