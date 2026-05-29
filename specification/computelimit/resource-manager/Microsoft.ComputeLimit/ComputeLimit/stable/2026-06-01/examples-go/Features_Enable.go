package armcomputelimit_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computelimit/armcomputelimit"
)

// Generated from example definition: 2026-06-01/Features_Enable.json
func ExampleFeaturesClient_BeginEnable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputelimit.NewClientFactory("74219ad7-63fc-442f-8037-4b43c627c07d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFeaturesClient().BeginEnable(ctx, "eastus", "VmCategoryQuota", &armcomputelimit.FeaturesClientBeginEnableOptions{
		Body: &armcomputelimit.FeatureEnableRequest{
			ServiceTreeID: to.Ptr("a1b2c3d4-5678-90ab-cdef-1234567890ab"),
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputelimit.FeaturesClientEnableResponse{
	// 	OperationStatusResult: armcomputelimit.OperationStatusResult{
	// 		Status: to.Ptr("Succeeded"),
	// 	},
	// }
}
