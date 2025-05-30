package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v2"
)

// Generated from example definition: 2025-03-01/AutoUpgradeProfiles_CreateOrUpdate_MaximumSet_Gen.json
func ExampleAutoUpgradeProfilesClient_BeginCreateOrUpdate_createAnAutoUpgradeProfileGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAutoUpgradeProfilesClient().BeginCreateOrUpdate(ctx, "rgfleets", "fleet1", "autoupgradeprofile1", armcontainerservicefleet.AutoUpgradeProfile{
		Properties: &armcontainerservicefleet.AutoUpgradeProfileProperties{
			Channel:          to.Ptr(armcontainerservicefleet.UpgradeChannelStable),
			UpdateStrategyID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgfleets/providers/Microsoft.ContainerService/fleets/fleet1/updateStrategies/strategy1"),
			NodeImageSelection: &armcontainerservicefleet.AutoUpgradeNodeImageSelection{
				Type: to.Ptr(armcontainerservicefleet.AutoUpgradeNodeImageSelectionTypeLatest),
			},
			Disabled: to.Ptr(true),
			AutoUpgradeProfileStatus: &armcontainerservicefleet.AutoUpgradeProfileStatus{
				LastTriggerError: &armcontainerservicefleet.ErrorDetail{},
			},
		},
	}, &armcontainerservicefleet.AutoUpgradeProfilesClientBeginCreateOrUpdateOptions{
		IfMatch:     to.Ptr("teikqmg"),
		IfNoneMatch: to.Ptr("ghfmmyrekxincsxklbldnvhqd")})
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
	// res = armcontainerservicefleet.AutoUpgradeProfilesClientCreateOrUpdateResponse{
	// 	AutoUpgradeProfile: &armcontainerservicefleet.AutoUpgradeProfile{
	// 		Properties: &armcontainerservicefleet.AutoUpgradeProfileProperties{
	// 			Channel: to.Ptr(armcontainerservicefleet.UpgradeChannelStable),
	// 			ProvisioningState: to.Ptr(armcontainerservicefleet.AutoUpgradeProfileProvisioningStateSucceeded),
	// 			UpdateStrategyID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgfleets/providers/Microsoft.ContainerService/fleets/fleet1/updateStrategies/strategy1"),
	// 			NodeImageSelection: &armcontainerservicefleet.AutoUpgradeNodeImageSelection{
	// 				Type: to.Ptr(armcontainerservicefleet.AutoUpgradeNodeImageSelectionTypeLatest),
	// 			},
	// 			Disabled: to.Ptr(true),
	// 			AutoUpgradeProfileStatus: &armcontainerservicefleet.AutoUpgradeProfileStatus{
	// 				LastTriggeredAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-27T02:49:14.736Z"); return t}()),
	// 				LastTriggerStatus: to.Ptr(armcontainerservicefleet.AutoUpgradeLastTriggerStatusSucceeded),
	// 				LastTriggerError: &armcontainerservicefleet.ErrorDetail{
	// 					Code: to.Ptr("mohtaimtrqkhivtsopwiuveg"),
	// 					Message: to.Ptr("t"),
	// 					Target: to.Ptr("axgswjixhofej"),
	// 					Details: []*armcontainerservicefleet.ErrorDetail{
	// 					},
	// 					AdditionalInfo: []*armcontainerservicefleet.ErrorAdditionalInfo{
	// 						{
	// 							Type: to.Ptr("oorrhg"),
	// 							Info: &armcontainerservicefleet.ErrorAdditionalInfoInfo{
	// 							},
	// 						},
	// 					},
	// 				},
	// 				LastTriggerUpgradeVersions: []*string{
	// 					to.Ptr("twqzpwqov"),
	// 				},
	// 			},
	// 		},
	// 		ETag: to.Ptr("\"EtagValue\""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/fleets/myFleet/autoUpgradeProfiles/autoupgradeprofile1"),
	// 		Name: to.Ptr("autoupgradeprofile1"),
	// 		Type: to.Ptr("Microsoft.ContainerService/fleets/autoUpgradeProfiles"),
	// 		SystemData: &armcontainerservicefleet.SystemData{
	// 			CreatedBy: to.Ptr("someUser"),
	// 			CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("someOtherUser"),
	// 			LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
	// 		},
	// 	},
	// }
}
