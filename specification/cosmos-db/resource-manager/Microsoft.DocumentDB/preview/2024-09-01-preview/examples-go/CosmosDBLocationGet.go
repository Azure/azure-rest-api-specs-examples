package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1ad29756bd141a47cac770140105a706d065ae1b/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBLocationGet.json
func ExampleLocationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.LocationGetResult = armcosmos.LocationGetResult{
	// 	Name: to.Ptr("westus"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/locations"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DocumentDB/locations/westus"),
	// 	Properties: &armcosmos.LocationProperties{
	// 		BackupStorageRedundancies: []*armcosmos.BackupStorageRedundancy{
	// 			to.Ptr(armcosmos.BackupStorageRedundancyLocal),
	// 			to.Ptr(armcosmos.BackupStorageRedundancyGeo)},
	// 			IsResidencyRestricted: to.Ptr(true),
	// 			IsSubscriptionRegionAccessAllowedForAz: to.Ptr(true),
	// 			IsSubscriptionRegionAccessAllowedForRegular: to.Ptr(true),
	// 			Status: to.Ptr(armcosmos.StatusOnline),
	// 			SupportsAvailabilityZone: to.Ptr(true),
	// 		},
	// 	}
}
