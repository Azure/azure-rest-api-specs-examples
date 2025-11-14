package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8b53f9cfc1fdb24dbfa28e311d3be4c645169297/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/fleet/CosmosDBFleetspaceCreate.json
func ExampleFleetspaceClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFleetspaceClient().BeginCreate(ctx, "rg1", "fleet1", "fleetspace1", armcosmos.FleetspaceResource{
		Properties: &armcosmos.FleetspaceProperties{
			DataRegions: []*string{
				to.Ptr("westus2")},
			FleetspaceAPIKind: to.Ptr(armcosmos.FleetspacePropertiesFleetspaceAPIKindNoSQL),
			ServiceTier:       to.Ptr(armcosmos.FleetspacePropertiesServiceTierGeneralPurpose),
			ThroughputPoolConfiguration: &armcosmos.FleetspacePropertiesThroughputPoolConfiguration{
				MaxThroughput: to.Ptr[int32](500000),
				MinThroughput: to.Ptr[int32](100000),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FleetspaceResource = armcosmos.FleetspaceResource{
	// 	Name: to.Ptr("fleetspace1"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/fleets/fleetspaces"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/rg1/providers/Microsoft.DocumentDB/fleets/fleet1/fleetspaces/fleetspace1"),
	// 	Properties: &armcosmos.FleetspaceProperties{
	// 		DataRegions: []*string{
	// 			to.Ptr("westus2")},
	// 			FleetspaceAPIKind: to.Ptr(armcosmos.FleetspacePropertiesFleetspaceAPIKindNoSQL),
	// 			ProvisioningState: to.Ptr(armcosmos.StatusSucceeded),
	// 			ServiceTier: to.Ptr(armcosmos.FleetspacePropertiesServiceTierGeneralPurpose),
	// 			ThroughputPoolConfiguration: &armcosmos.FleetspacePropertiesThroughputPoolConfiguration{
	// 				MaxThroughput: to.Ptr[int32](500000),
	// 				MinThroughput: to.Ptr[int32](100000),
	// 			},
	// 		},
	// 	}
}
