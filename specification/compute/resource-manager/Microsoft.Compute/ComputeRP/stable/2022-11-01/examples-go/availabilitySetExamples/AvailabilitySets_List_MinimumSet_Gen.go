package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/19f98c9f526f8db961f172276dd6d6882a86ed86/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/availabilitySetExamples/AvailabilitySets_List_MinimumSet_Gen.json
func ExampleAvailabilitySetsClient_NewListPager_availabilitySetsListMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailabilitySetsClient().NewListPager("rgcompute", nil)
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
		// page.AvailabilitySetListResult = armcompute.AvailabilitySetListResult{
		// 	Value: []*armcompute.AvailabilitySet{
		// 		{
		// 			Location: to.Ptr("australiasoutheast"),
		// 		},
		// 		{
		// 			Location: to.Ptr("australiasoutheast"),
		// 		},
		// 		{
		// 			Location: to.Ptr("westcentralus"),
		// 		},
		// 		{
		// 			Location: to.Ptr("westcentralus"),
		// 	}},
		// }
	}
}
