package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v3"
)

// Generated from example definition: 2025-06-01/Profiles_Get.json
func ExampleProfilesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProfilesClient().Get(ctx, "RG", "profile1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcdn.ProfilesClientGetResponse{
	// 	Profile: armcdn.Profile{
	// 		Name: to.Ptr("profile1"),
	// 		Type: to.Ptr("Microsoft.Cdn/profiles"),
	// 		ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1"),
	// 		Kind: to.Ptr("frontdoor"),
	// 		Location: to.Ptr("global"),
	// 		Properties: &armcdn.ProfileProperties{
	// 			FrontDoorID: to.Ptr("3b4682da-b3e2-47a1-96ca-08ab3cb7294e"),
	// 			LogScrubbing: &armcdn.ProfileLogScrubbing{
	// 				ScrubbingRules: []*armcdn.ProfileScrubbingRules{
	// 				},
	// 				State: to.Ptr(armcdn.ProfileScrubbingStateEnabled),
	// 			},
	// 			OriginResponseTimeoutSeconds: to.Ptr[int32](30),
	// 			ProvisioningState: to.Ptr(armcdn.ProfileProvisioningStateSucceeded),
	// 			ResourceState: to.Ptr(armcdn.ProfileResourceStateActive),
	// 		},
	// 		SKU: &armcdn.SKU{
	// 			Name: to.Ptr(armcdn.SKUNamePremiumAzureFrontDoor),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
