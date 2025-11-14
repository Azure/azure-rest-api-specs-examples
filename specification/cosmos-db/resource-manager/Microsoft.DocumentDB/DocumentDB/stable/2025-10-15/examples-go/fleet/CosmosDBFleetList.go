package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8b53f9cfc1fdb24dbfa28e311d3be4c645169297/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/fleet/CosmosDBFleetList.json
func ExampleFleetClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFleetClient().NewListByResourceGroupPager("rg1", nil)
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
		// page.FleetListResult = armcosmos.FleetListResult{
		// 	Value: []*armcosmos.FleetResource{
		// 		{
		// 			Name: to.Ptr("fleet1"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/fleets"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/rg1/providers/Microsoft.DocumentDB/fleets/fleet1"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"Dept": to.Ptr("Finance"),
		// 				"Environment": to.Ptr("Production"),
		// 			},
		// 			Properties: &armcosmos.FleetResourceProperties{
		// 				ProvisioningState: to.Ptr(armcosmos.StatusSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("fleet2"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/fleets"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/rg2/providers/Microsoft.DocumentDB/fleets/fleet2"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 				"Dept": to.Ptr("IT"),
		// 				"Environment": to.Ptr("Development"),
		// 			},
		// 			Properties: &armcosmos.FleetResourceProperties{
		// 				ProvisioningState: to.Ptr(armcosmos.StatusSucceeded),
		// 			},
		// 	}},
		// }
	}
}
