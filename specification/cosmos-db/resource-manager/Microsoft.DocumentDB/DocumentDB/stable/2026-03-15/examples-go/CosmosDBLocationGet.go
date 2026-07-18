package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBLocationGet.json
func ExampleLocationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocationsClient().Get(ctx, "westus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.LocationsClientGetResponse{
	// 	LocationGetResult: armcosmos.LocationGetResult{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DocumentDB/locations/westus"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/locations"),
	// 		Name: to.Ptr("westus"),
	// 		Properties: &armcosmos.LocationProperties{
	// 			SupportsAvailabilityZone: to.Ptr(true),
	// 			IsResidencyRestricted: to.Ptr(true),
	// 			BackupStorageRedundancies: []*armcosmos.BackupStorageRedundancy{
	// 				to.Ptr(armcosmos.BackupStorageRedundancyLocal),
	// 				to.Ptr(armcosmos.BackupStorageRedundancyGeo),
	// 			},
	// 			IsSubscriptionRegionAccessAllowedForRegular: to.Ptr(true),
	// 			IsSubscriptionRegionAccessAllowedForAz: to.Ptr(true),
	// 			Status: to.Ptr(armcosmos.StatusOnline),
	// 		},
	// 	},
	// }
}
