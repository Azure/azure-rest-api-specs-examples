package armcomputelimit_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computelimit/armcomputelimit"
)

// Generated from example definition: 2026-07-01/MemberCapOverrides_CreateOrUpdate.json
func ExampleMemberCapOverridesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputelimit.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMemberCapOverridesClient().CreateOrUpdate(ctx, "eastus", "StandardDSv3Family", "11111111-1111-1111-1111-111111111111", armcomputelimit.MemberCapOverride{
		Properties: &armcomputelimit.MemberCapOverrideProperties{
			Cap: to.Ptr[int32](250),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputelimit.MemberCapOverridesClientCreateOrUpdateResponse{
	// 	MemberCapOverride: armcomputelimit.MemberCapOverride{
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/providers/Microsoft.ComputeLimit/locations/eastus/sharedLimitCaps/StandardDSv3Family/memberCapOverrides/11111111-1111-1111-1111-111111111111"),
	// 		Name: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		Type: to.Ptr("Microsoft.ComputeLimit/locations/sharedLimitCaps/memberCapOverrides"),
	// 		Properties: &armcomputelimit.MemberCapOverrideProperties{
	// 			Cap: to.Ptr[int32](250),
	// 			ProvisioningState: to.Ptr(armcomputelimit.ResourceProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
