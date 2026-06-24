package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/capacityReservationExamples/CapacityReservationGroup_ListBySubscriptionWithResourceIdsQuery.json
func ExampleCapacityReservationGroupsClient_NewListBySubscriptionPager_listCapacityReservationGroupsWithResourceIdsOnlyInSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCapacityReservationGroupsClient().NewListBySubscriptionPager(&armcompute.CapacityReservationGroupsClientListBySubscriptionOptions{
		ResourceIDsOnly: to.Ptr(armcompute.ResourceIDOptionsForGetCapacityReservationGroupsAll)})
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
		// page = armcompute.CapacityReservationGroupsClientListBySubscriptionResponse{
		// 	CapacityReservationGroupListResult: armcompute.CapacityReservationGroupListResult{
		// 		Value: []*armcompute.CapacityReservationGroup{
		// 			{
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup1/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName1}"),
		// 				Type: to.Ptr("Microsoft.Compute/capacityReservationGroups"),
		// 				Location: to.Ptr("SouthCentralUSSTG"),
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/{subscriptionId2}/resourceGroups/myResourceGroup2/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName2}"),
		// 				Type: to.Ptr("Microsoft.Compute/capacityReservationGroups"),
		// 				Location: to.Ptr("SouthCentralUSSTG"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
