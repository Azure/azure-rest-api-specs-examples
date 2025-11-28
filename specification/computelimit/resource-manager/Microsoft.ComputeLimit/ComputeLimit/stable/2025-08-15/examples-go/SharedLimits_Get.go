package armcomputelimit_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computelimit/armcomputelimit"
)

// Generated from example definition: 2025-08-15/SharedLimits_Get.json
func ExampleSharedLimitsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputelimit.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSharedLimitsClient().Get(ctx, "eastus", "StandardDSv3Family", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputelimit.SharedLimitsClientGetResponse{
	// 	SharedLimit: &armcomputelimit.SharedLimit{
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/providers/Microsoft.ComputeLimit/locations/eastus/sharedLimits/StandardDSv3Family"),
	// 		Name: to.Ptr("StandardDSv3Family"),
	// 		Type: to.Ptr("Microsoft.ComputeLimit/locations/sharedLimits"),
	// 		Properties: &armcomputelimit.SharedLimitProperties{
	// 			ResourceName: &armcomputelimit.LimitName{
	// 				Value: to.Ptr("StandardDSv3Family"),
	// 				LocalizedValue: to.Ptr("Standard DSv3 Family vCPUs"),
	// 			},
	// 			Limit: to.Ptr[int32](100),
	// 			Unit: to.Ptr("Count"),
	// 			ProvisioningState: to.Ptr(armcomputelimit.ResourceProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
