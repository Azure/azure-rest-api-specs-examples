package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/availabilitySetExamples/AvailabilitySet_List_MinimumSet_Gen.json
func ExampleAvailabilitySetsClient_NewListPager_availabilitySetListMinimumSetGen() {
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
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/availabilitySets/myAvailabilitySet1"),
		// 			Location: to.Ptr("australiasoutheast"),
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/availabilitySets/myAvailabilitySet2"),
		// 			Location: to.Ptr("australiasoutheast"),
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/availabilitySets/myAvailabilitySet3"),
		// 			Location: to.Ptr("westcentralus"),
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/availabilitySets/myAvailabilitySet4"),
		// 			Location: to.Ptr("westcentralus"),
		// 	}},
		// }
	}
}
