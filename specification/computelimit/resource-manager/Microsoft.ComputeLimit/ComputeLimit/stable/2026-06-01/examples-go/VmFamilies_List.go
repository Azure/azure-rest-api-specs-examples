package armcomputelimit_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computelimit/armcomputelimit"
)

// Generated from example definition: 2026-06-01/VmFamilies_List.json
func ExampleVMFamiliesClient_NewListBySubscriptionLocationResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputelimit.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVMFamiliesClient().NewListBySubscriptionLocationResourcePager("eastus", nil)
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
		// page = armcomputelimit.VMFamiliesClientListBySubscriptionLocationResourceResponse{
		// 	VMFamilyListResult: armcomputelimit.VMFamilyListResult{
		// 		Value: []*armcomputelimit.VMFamily{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ComputeLimit/locations/eastus/vmFamilies/standardDSv2Family"),
		// 				Name: to.Ptr("standardDSv2Family"),
		// 				Type: to.Ptr("Microsoft.ComputeLimit/locations/vmFamilies"),
		// 				Properties: &armcomputelimit.VMFamilyProperties{
		// 					Category: to.Ptr("generalPurposeCategory"),
		// 					ProvisioningState: to.Ptr(armcomputelimit.ResourceProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ComputeLimit/locations/eastus/vmFamilies/standardDv2Family"),
		// 				Name: to.Ptr("standardDv2Family"),
		// 				Type: to.Ptr("Microsoft.ComputeLimit/locations/vmFamilies"),
		// 				Properties: &armcomputelimit.VMFamilyProperties{
		// 					Category: to.Ptr("generalPurposeCategory"),
		// 					ProvisioningState: to.Ptr(armcomputelimit.ResourceProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
