package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering/v2"
)

// Generated from example definition: 2025-05-01/CreatePeeringService.json
func ExampleServicesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("subId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().CreateOrUpdate(ctx, "rgName", "peeringServiceName", armpeering.Service{
		Location: to.Ptr("eastus"),
		Properties: &armpeering.ServiceProperties{
			PeeringServiceLocation:         to.Ptr("state1"),
			PeeringServiceProvider:         to.Ptr("serviceProvider1"),
			ProviderBackupPeeringLocation:  to.Ptr("peeringLocation2"),
			ProviderPrimaryPeeringLocation: to.Ptr("peeringLocation1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpeering.ServicesClientCreateOrUpdateResponse{
	// 	Service: &armpeering.Service{
	// 		Name: to.Ptr("peeringServiceName"),
	// 		Type: to.Ptr("Microsoft.Peering/peeringServices"),
	// 		ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peeringServices/peeringServiceName"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armpeering.ServiceProperties{
	// 			PeeringServiceLocation: to.Ptr("state1"),
	// 			PeeringServiceProvider: to.Ptr("serviceProvider1"),
	// 			ProviderBackupPeeringLocation: to.Ptr("peeringLocation2"),
	// 			ProviderPrimaryPeeringLocation: to.Ptr("peeringLocation1"),
	// 			ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
