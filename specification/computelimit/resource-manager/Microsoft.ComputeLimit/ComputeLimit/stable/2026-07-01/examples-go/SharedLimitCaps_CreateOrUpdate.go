package armcomputelimit_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computelimit/armcomputelimit"
)

// Generated from example definition: 2026-07-01/SharedLimitCaps_CreateOrUpdate.json
func ExampleSharedLimitCapsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputelimit.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSharedLimitCapsClient().CreateOrUpdate(ctx, "eastus", "StandardDSv3Family", armcomputelimit.SharedLimitCap{
		Properties: &armcomputelimit.SharedLimitCapProperties{
			DefaultMemberCap: to.Ptr[int32](100),
			IsBoundedCap:     to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputelimit.SharedLimitCapsClientCreateOrUpdateResponse{
	// 	SharedLimitCap: armcomputelimit.SharedLimitCap{
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/providers/Microsoft.ComputeLimit/locations/eastus/sharedLimitCaps/StandardDSv3Family"),
	// 		Name: to.Ptr("StandardDSv3Family"),
	// 		Type: to.Ptr("Microsoft.ComputeLimit/locations/sharedLimitCaps"),
	// 		Properties: &armcomputelimit.SharedLimitCapProperties{
	// 			DefaultMemberCap: to.Ptr[int32](100),
	// 			IsBoundedCap: to.Ptr(true),
	// 			ProvisioningState: to.Ptr(armcomputelimit.ResourceProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
