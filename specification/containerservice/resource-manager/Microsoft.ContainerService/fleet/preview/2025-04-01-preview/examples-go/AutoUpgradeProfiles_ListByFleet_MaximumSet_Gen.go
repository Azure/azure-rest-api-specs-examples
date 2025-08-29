package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v3"
)

// Generated from example definition: 2025-04-01-preview/AutoUpgradeProfiles_ListByFleet_MaximumSet_Gen.json
func ExampleAutoUpgradeProfilesClient_NewListByFleetPager_listsTheAutoUpgradeProfileResourcesByFleetGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAutoUpgradeProfilesClient().NewListByFleetPager("rgfleets", "fleet1", nil)
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
		// page = armcontainerservicefleet.AutoUpgradeProfilesClientListByFleetResponse{
		// 	AutoUpgradeProfileListResult: armcontainerservicefleet.AutoUpgradeProfileListResult{
		// 		Value: []*armcontainerservicefleet.AutoUpgradeProfile{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/fleets/myFleet/autoUpgradeProfiles/autoupgradeprofile1"),
		// 				Name: to.Ptr("autoupgradeprofile1"),
		// 				Type: to.Ptr("Microsoft.ContainerService/fleets/autoUpgradeProfiles"),
		// 				SystemData: &armcontainerservicefleet.SystemData{
		// 					CreatedBy: to.Ptr("@contoso.com"),
		// 					CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T01:10:08.395Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T01:10:08.395Z"); return t}()),
		// 				},
		// 				Properties: &armcontainerservicefleet.AutoUpgradeProfileProperties{
		// 					ProvisioningState: to.Ptr(armcontainerservicefleet.AutoUpgradeProfileProvisioningStateSucceeded),
		// 					Channel: to.Ptr(armcontainerservicefleet.UpgradeChannelStable),
		// 					UpdateStrategyID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgfleets/providers/Microsoft.ContainerService/fleets/fleet1/updateStrategies/strategy1"),
		// 					NodeImageSelection: &armcontainerservicefleet.AutoUpgradeNodeImageSelection{
		// 						Type: to.Ptr(armcontainerservicefleet.AutoUpgradeNodeImageSelectionTypeLatest),
		// 					},
		// 					Disabled: to.Ptr(true),
		// 				},
		// 				ETag: to.Ptr("\"EtagValue\""),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("http://nextlink.contoso.com"),
		// 	},
		// }
	}
}
