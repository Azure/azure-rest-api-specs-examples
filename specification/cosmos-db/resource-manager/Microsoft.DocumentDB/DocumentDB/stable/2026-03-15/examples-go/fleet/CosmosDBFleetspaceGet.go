package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/fleet/CosmosDBFleetspaceGet.json
func ExampleFleetspaceClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFleetspaceClient().Get(ctx, "rg1", "fleet1", "fleetspace1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.FleetspaceClientGetResponse{
	// 	FleetspaceResource: armcosmos.FleetspaceResource{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/rg1/providers/Microsoft.DocumentDB/fleets/fleet1/fleetspaces/fleetspace1"),
	// 		Name: to.Ptr("fleetspace1"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/fleets/fleetspaces"),
	// 		Properties: &armcosmos.FleetspaceProperties{
	// 			ProvisioningState: to.Ptr(armcosmos.StatusSucceeded),
	// 			ServiceTier: to.Ptr(armcosmos.FleetspacePropertiesServiceTierGeneralPurpose),
	// 			DataRegions: []*string{
	// 				to.Ptr("westus2"),
	// 			},
	// 			ThroughputPoolConfiguration: &armcosmos.FleetspacePropertiesThroughputPoolConfiguration{
	// 				MinThroughput: to.Ptr[int32](100000),
	// 				MaxThroughput: to.Ptr[int32](500000),
	// 			},
	// 		},
	// 	},
	// }
}
