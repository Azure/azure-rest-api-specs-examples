package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad60d7f8eba124edc6999677c55aba2184e303b0/specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2024-05-02-preview/examples/AutoUpgradeProfiles_ListByFleet.json
func ExampleAutoUpgradeProfilesClient_NewListByFleetPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAutoUpgradeProfilesClient().NewListByFleetPager("rg1", "fleet1", nil)
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
		// page.AutoUpgradeProfileListResult = armcontainerservicefleet.AutoUpgradeProfileListResult{
		// 	Value: []*armcontainerservicefleet.AutoUpgradeProfile{
		// 		{
		// 			Name: to.Ptr("autoupgradeprofile1"),
		// 			Type: to.Ptr("Microsoft.ContainerService/fleets/autoUpgradeProfiles"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/fleets/myFleet/autoUpgradeProfiles/autoupgradeprofile1"),
		// 			SystemData: &armcontainerservicefleet.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T01:10:08.395Z"); return t}()),
		// 				CreatedBy: to.Ptr("@contoso.com"),
		// 				CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T01:10:08.395Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 			},
		// 			ETag: to.Ptr("\"EtagValue\""),
		// 			Properties: &armcontainerservicefleet.AutoUpgradeProfileProperties{
		// 				Channel: to.Ptr(armcontainerservicefleet.UpgradeChannelStable),
		// 				ProvisioningState: to.Ptr(armcontainerservicefleet.AutoUpgradeProfileProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
