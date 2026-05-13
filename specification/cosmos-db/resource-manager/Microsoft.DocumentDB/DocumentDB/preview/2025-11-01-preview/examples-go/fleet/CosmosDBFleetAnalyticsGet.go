package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2025-11-01-preview/fleet/CosmosDBFleetAnalyticsGet.json
func ExampleFleetAnalyticsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFleetAnalyticsClient().Get(ctx, "rg1", "fleet1", "storageAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.FleetAnalyticsClientGetResponse{
	// 	FleetAnalyticsResource: &armcosmos.FleetAnalyticsResource{
	// 		Name: to.Ptr("storageAccount"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/fleets/fleetAnalytics"),
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/rg1/providers/Microsoft.DocumentDB/fleets/fleet1/fleetAnalytics/storageAccount"),
	// 		Properties: &armcosmos.FleetAnalyticsProperties{
	// 			ProvisioningState: to.Ptr(armcosmos.StatusSucceeded),
	// 			StorageLocationType: to.Ptr(armcosmos.FleetAnalyticsPropertiesStorageLocationTypeStorageAccount),
	// 			StorageLocationURI: to.Ptr("/subscriptions/d1eb41bc-1b7f-4404-bd2a-868c222852d/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/testStorageAccount1"),
	// 		},
	// 	},
	// }
}
