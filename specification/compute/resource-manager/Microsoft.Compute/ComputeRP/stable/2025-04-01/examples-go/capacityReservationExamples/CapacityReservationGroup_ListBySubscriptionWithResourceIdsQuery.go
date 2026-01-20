package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/capacityReservationExamples/CapacityReservationGroup_ListBySubscriptionWithResourceIdsQuery.json
func ExampleCapacityReservationGroupsClient_NewListBySubscriptionPager_listCapacityReservationGroupsWithResourceIdsOnlyInSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCapacityReservationGroupsClient().NewListBySubscriptionPager(&armcompute.CapacityReservationGroupsClientListBySubscriptionOptions{Expand: nil,
		ResourceIDsOnly: to.Ptr(armcompute.ResourceIDOptionsForGetCapacityReservationGroupsAll),
	})
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
		// page.CapacityReservationGroupListResult = armcompute.CapacityReservationGroupListResult{
		// 	Value: []*armcompute.CapacityReservationGroup{
		// 		{
		// 			Type: to.Ptr("Microsoft.Compute/capacityReservationGroups"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup1/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName1}"),
		// 			Location: to.Ptr("SouthCentralUSSTG"),
		// 		},
		// 		{
		// 			Type: to.Ptr("Microsoft.Compute/capacityReservationGroups"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId2}/resourceGroups/myResourceGroup2/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName2}"),
		// 			Location: to.Ptr("SouthCentralUSSTG"),
		// 	}},
		// }
	}
}
