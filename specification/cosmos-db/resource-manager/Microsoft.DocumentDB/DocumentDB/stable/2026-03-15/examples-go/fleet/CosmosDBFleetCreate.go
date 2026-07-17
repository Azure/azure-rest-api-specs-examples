package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/fleet/CosmosDBFleetCreate.json
func ExampleFleetClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFleetClient().Create(ctx, "rg1", "fleet1", armcosmos.FleetResource{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"Dept":        to.Ptr("Finance"),
			"Environment": to.Ptr("Production"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.FleetClientCreateResponse{
	// 	FleetResource: armcosmos.FleetResource{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/rg1/providers/Microsoft.DocumentDB/fleets/fleet1"),
	// 		Name: to.Ptr("fleet1"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/fleets"),
	// 		Location: to.Ptr("West US"),
	// 		Tags: map[string]*string{
	// 			"Dept": to.Ptr("Finance"),
	// 			"Environment": to.Ptr("Production"),
	// 		},
	// 		Properties: &armcosmos.FleetResourceProperties{
	// 			ProvisioningState: to.Ptr(armcosmos.StatusSucceeded),
	// 		},
	// 	},
	// }
}
