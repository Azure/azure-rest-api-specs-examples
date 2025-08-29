package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v3"
)

// Generated from example definition: 2025-04-01-preview/AutoUpgradeProfiles_CreateOrUpdate.json
func ExampleAutoUpgradeProfilesClient_BeginCreateOrUpdate_createAnAutoUpgradeProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAutoUpgradeProfilesClient().BeginCreateOrUpdate(ctx, "rg1", "fleet1", "autoupgradeprofile1", armcontainerservicefleet.AutoUpgradeProfile{
		Properties: &armcontainerservicefleet.AutoUpgradeProfileProperties{
			TargetKubernetesVersion: to.Ptr(""),
			LongTermSupport:         to.Ptr(false),
			Channel:                 to.Ptr(armcontainerservicefleet.UpgradeChannelStable),
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
	// res = armcontainerservicefleet.AutoUpgradeProfilesClientCreateOrUpdateResponse{
	// 	AutoUpgradeProfile: &armcontainerservicefleet.AutoUpgradeProfile{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/fleets/myFleet/autoUpgradeProfiles/autoupgradeprofile1"),
	// 		Name: to.Ptr("autoupgradeprofile1"),
	// 		Type: to.Ptr("Microsoft.ContainerService/fleets/autoUpgradeProfiles"),
	// 		SystemData: &armcontainerservicefleet.SystemData{
	// 			CreatedBy: to.Ptr("@contoso.com"),
	// 			CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T01:10:08.395Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T01:10:08.395Z"); return t}()),
	// 		},
	// 		Properties: &armcontainerservicefleet.AutoUpgradeProfileProperties{
	// 			ProvisioningState: to.Ptr(armcontainerservicefleet.AutoUpgradeProfileProvisioningStateSucceeded),
	// 			TargetKubernetesVersion: to.Ptr(""),
	// 			LongTermSupport: to.Ptr(false),
	// 			Channel: to.Ptr(armcontainerservicefleet.UpgradeChannelStable),
	// 		},
	// 		ETag: to.Ptr("\"EtagValue\""),
	// 	},
	// }
}
