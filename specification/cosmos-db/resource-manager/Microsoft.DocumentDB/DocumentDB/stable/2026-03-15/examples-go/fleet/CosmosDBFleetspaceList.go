package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/fleet/CosmosDBFleetspaceList.json
func ExampleFleetspaceClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFleetspaceClient().NewListPager("rg1", "fleet1", nil)
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
		// page = armcosmos.FleetspaceClientListResponse{
		// 	FleetspaceListResult: armcosmos.FleetspaceListResult{
		// 		Value: []*armcosmos.FleetspaceResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/rg1/providers/Microsoft.DocumentDB/fleets/fleet1/fleetspaces/fleetspace1"),
		// 				Name: to.Ptr("fleetspace1"),
		// 				Type: to.Ptr("Microsoft.DocumentDB/fleets/fleetspaces"),
		// 				Properties: &armcosmos.FleetspaceProperties{
		// 					ProvisioningState: to.Ptr(armcosmos.StatusSucceeded),
		// 					ServiceTier: to.Ptr(armcosmos.FleetspacePropertiesServiceTierBusinessCritical),
		// 					DataRegions: []*string{
		// 						to.Ptr("westus2"),
		// 					},
		// 					ThroughputPoolConfiguration: &armcosmos.FleetspacePropertiesThroughputPoolConfiguration{
		// 						MinThroughput: to.Ptr[int32](100000),
		// 						MaxThroughput: to.Ptr[int32](500000),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/rg1/providers/Microsoft.DocumentDB/fleets/fleet1/fleetspaces/fleetspace2"),
		// 				Name: to.Ptr("fleetspace2"),
		// 				Type: to.Ptr("Microsoft.DocumentDB/fleets/fleetspaces"),
		// 				Properties: &armcosmos.FleetspaceProperties{
		// 					ProvisioningState: to.Ptr(armcosmos.StatusSucceeded),
		// 					ServiceTier: to.Ptr(armcosmos.FleetspacePropertiesServiceTierGeneralPurpose),
		// 					DataRegions: []*string{
		// 						to.Ptr("eastus"),
		// 					},
		// 					ThroughputPoolConfiguration: &armcosmos.FleetspacePropertiesThroughputPoolConfiguration{
		// 						MinThroughput: to.Ptr[int32](200000),
		// 						MaxThroughput: to.Ptr[int32](600000),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
