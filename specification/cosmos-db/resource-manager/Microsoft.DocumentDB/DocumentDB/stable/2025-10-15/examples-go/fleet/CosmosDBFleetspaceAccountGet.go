package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8b53f9cfc1fdb24dbfa28e311d3be4c645169297/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/fleet/CosmosDBFleetspaceAccountGet.json
func ExampleFleetspaceAccountClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFleetspaceAccountClient().Get(ctx, "rg1", "fleet1", "fleetspace1", "db1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FleetspaceAccountResource = armcosmos.FleetspaceAccountResource{
	// 	Name: to.Ptr("db1"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/fleets/fleetspaces/fleetspaceAccounts"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/rg1/providers/Microsoft.DocumentDB/fleets/fleet1/fleetspaces/fleetspace1/fleetspaceAccounts/db1"),
	// 	Properties: &armcosmos.FleetspaceAccountProperties{
	// 		GlobalDatabaseAccountProperties: &armcosmos.FleetspaceAccountPropertiesGlobalDatabaseAccountProperties{
	// 			ArmLocation: to.Ptr("West US"),
	// 			ResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/providers/Microsoft.DocumentDB/resourceGroup/rg1/databaseAccounts/db1"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcosmos.StatusSucceeded),
	// 	},
	// }
}
