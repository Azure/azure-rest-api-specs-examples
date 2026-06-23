package armcomputelimit_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computelimit/armcomputelimit"
)

// Generated from example definition: 2026-07-01/MemberCapOverrides_ListByParent.json
func ExampleMemberCapOverridesClient_NewListByParentPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputelimit.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMemberCapOverridesClient().NewListByParentPager("eastus", "StandardDSv3Family", nil)
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
		// page = armcomputelimit.MemberCapOverridesClientListByParentResponse{
		// 	MemberCapOverrideListResult: armcomputelimit.MemberCapOverrideListResult{
		// 		Value: []*armcomputelimit.MemberCapOverride{
		// 			{
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/providers/Microsoft.ComputeLimit/locations/eastus/sharedLimitCaps/StandardDSv3Family/memberCapOverrides/11111111-1111-1111-1111-111111111111"),
		// 				Name: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 				Type: to.Ptr("Microsoft.ComputeLimit/locations/sharedLimitCaps/memberCapOverrides"),
		// 				Properties: &armcomputelimit.MemberCapOverrideProperties{
		// 					Cap: to.Ptr[int32](200),
		// 					ProvisioningState: to.Ptr(armcomputelimit.ResourceProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/providers/Microsoft.ComputeLimit/locations/eastus/sharedLimitCaps/StandardDSv3Family/memberCapOverrides/22222222-2222-2222-2222-222222222222"),
		// 				Name: to.Ptr("22222222-2222-2222-2222-222222222222"),
		// 				Type: to.Ptr("Microsoft.ComputeLimit/locations/sharedLimitCaps/memberCapOverrides"),
		// 				Properties: &armcomputelimit.MemberCapOverrideProperties{
		// 					Cap: to.Ptr[int32](350),
		// 					ProvisioningState: to.Ptr(armcomputelimit.ResourceProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
