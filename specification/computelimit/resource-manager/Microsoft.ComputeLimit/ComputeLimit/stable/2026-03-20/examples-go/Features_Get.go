package armcomputelimit_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computelimit/armcomputelimit"
)

// Generated from example definition: 2026-03-20/Features_Get.json
func ExampleFeaturesClient_Get_getFeature() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputelimit.NewClientFactory("74219ad7-63fc-442f-8037-4b43c627c07d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFeaturesClient().Get(ctx, "eastus", "VmCategoryQuota", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputelimit.FeaturesClientGetResponse{
	// 	Feature: &armcomputelimit.Feature{
	// 		ID: to.Ptr("/subscriptions/74219ad7-63fc-442f-8037-4b43c627c07d/providers/Microsoft.ComputeLimit/locations/eastus/features/VmCategoryQuota"),
	// 		Name: to.Ptr("VmCategoryQuota"),
	// 		Type: to.Ptr("Microsoft.ComputeLimit/locations/features"),
	// 		Properties: &armcomputelimit.FeatureProperties{
	// 			ProvisioningState: to.Ptr(armcomputelimit.ResourceProvisioningStateSucceeded),
	// 			State: to.Ptr(armcomputelimit.FeatureStateEnabled),
	// 		},
	// 	},
	// }
}
