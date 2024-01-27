package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6f8faf5da91b5b9af5f3512fe609e22e99383d41/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-11-15-preview/examples/CosmosDBLocationList.json
func ExampleLocationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLocationsClient().NewListPager(nil)
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
		// page.LocationListResult = armcosmos.LocationListResult{
		// 	Value: []*armcosmos.LocationGetResult{
		// 		{
		// 			Name: to.Ptr("westus"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/locations"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DocumentDB/locations/westus"),
		// 			Properties: &armcosmos.LocationProperties{
		// 				BackupStorageRedundancies: []*armcosmos.BackupStorageRedundancy{
		// 					to.Ptr(armcosmos.BackupStorageRedundancyLocal),
		// 					to.Ptr(armcosmos.BackupStorageRedundancyGeo)},
		// 					IsResidencyRestricted: to.Ptr(false),
		// 					IsSubscriptionRegionAccessAllowedForAz: to.Ptr(false),
		// 					IsSubscriptionRegionAccessAllowedForRegular: to.Ptr(true),
		// 					Status: to.Ptr(armcosmos.StatusOnline),
		// 					SupportsAvailabilityZone: to.Ptr(false),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("centralus"),
		// 				Type: to.Ptr("Microsoft.DocumentDB/locations"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DocumentDB/locations/centralus"),
		// 				Properties: &armcosmos.LocationProperties{
		// 					BackupStorageRedundancies: []*armcosmos.BackupStorageRedundancy{
		// 						to.Ptr(armcosmos.BackupStorageRedundancyZone),
		// 						to.Ptr(armcosmos.BackupStorageRedundancyGeo)},
		// 						IsResidencyRestricted: to.Ptr(false),
		// 						IsSubscriptionRegionAccessAllowedForAz: to.Ptr(true),
		// 						IsSubscriptionRegionAccessAllowedForRegular: to.Ptr(false),
		// 						Status: to.Ptr(armcosmos.StatusOnline),
		// 						SupportsAvailabilityZone: to.Ptr(true),
		// 					},
		// 			}},
		// 		}
	}
}
