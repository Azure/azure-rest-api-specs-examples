package armcomputelimit_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computelimit/armcomputelimit"
)

// Generated from example definition: 2026-07-01/SharedLimitCaps_SetMemberCapOverrides.json
func ExampleSharedLimitCapsClient_SetMemberCapOverrides_replaceTheFullSetOfMemberCapOverridesForASharedLimitCap() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputelimit.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSharedLimitCapsClient().SetMemberCapOverrides(ctx, "eastus", "StandardDSv3Family", armcomputelimit.SetMemberCapOverridesRequest{
		MemberCapOverrides: []*armcomputelimit.MemberCap{
			{
				SubscriptionID: to.Ptr("11111111-1111-1111-1111-111111111111"),
				Cap:            to.Ptr[int32](200),
			},
			{
				SubscriptionID: to.Ptr("22222222-2222-2222-2222-222222222222"),
				Cap:            to.Ptr[int32](150),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputelimit.SharedLimitCapsClientSetMemberCapOverridesResponse{
	// 	SetMemberCapOverridesResult: armcomputelimit.SetMemberCapOverridesResult{
	// 		MemberCapOverrides: []*armcomputelimit.MemberCap{
	// 			{
	// 				SubscriptionID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 				Cap: to.Ptr[int32](200),
	// 			},
	// 			{
	// 				SubscriptionID: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 				Cap: to.Ptr[int32](150),
	// 			},
	// 		},
	// 	},
	// }
}
