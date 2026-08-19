package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-04-01/capacityReservationExamples/OpenCapacityReservationGroup_Get.json
func ExampleCapacityReservationGroupsClient_Get_getAnOpenCapacityReservationGroupWithInstanceView() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCapacityReservationGroupsClient().Get(ctx, "myResourceGroup", "openCapacityReservationGroup", &armcompute.CapacityReservationGroupsClientGetOptions{
		Expand: to.Ptr(armcompute.CapacityReservationGroupInstanceViewTypesInstanceView)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.CapacityReservationGroupsClientGetResponse{
	// 	CapacityReservationGroup: armcompute.CapacityReservationGroup{
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/CapacityReservationGroups/openCapacityReservationGroup"),
	// 		Properties: &armcompute.CapacityReservationGroupProperties{
	// 			CapacityReservations: []*armcompute.SubResourceReadOnly{
	// 				{
	// 					ID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/openCapacityReservationGroup/capacityReservations/openCapacityReservation1"),
	// 				},
	// 			},
	// 			SharingProfile: &armcompute.ResourceSharingProfile{
	// 				SubscriptionIDs: []*armcompute.SubResource{
	// 					{
	// 						ID: to.Ptr("/subscriptions/{subscription-id1}"),
	// 					},
	// 					{
	// 						ID: to.Ptr("/subscriptions/{subscription-id2}"),
	// 					},
	// 				},
	// 			},
	// 			InstanceView: &armcompute.CapacityReservationGroupInstanceView{
	// 				CapacityReservations: []*armcompute.CapacityReservationInstanceViewWithName{
	// 					{
	// 						Name: to.Ptr("openCapacityReservation1"),
	// 						UtilizationInfo: &armcompute.CapacityReservationUtilization{
	// 							CurrentCapacity: to.Ptr[int32](10),
	// 							UsedReservedCountBySubscription: map[string]*int32{
	// 								"{subscription-id1}": to.Ptr[int32](3),
	// 								"{subscription-id2}": to.Ptr[int32](2),
	// 							},
	// 						},
	// 						Statuses: []*armcompute.InstanceViewStatus{
	// 							{
	// 								Code: to.Ptr("ProvisioningState/succeeded"),
	// 								Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 								DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 				SharedSubscriptionIDs: []*armcompute.SubResourceReadOnly{
	// 					{
	// 						ID: to.Ptr("/subscriptions/{subscription-id1}"),
	// 					},
	// 					{
	// 						ID: to.Ptr("/subscriptions/{subscription-id2}"),
	// 					},
	// 				},
	// 			},
	// 			ReservationType: to.Ptr(armcompute.ReservationTypeOpen),
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"{tagName}": to.Ptr("{tagValue}"),
	// 		},
	// 		Name: to.Ptr("openCapacityReservationGroup"),
	// 		Zones: []*string{
	// 			to.Ptr("1"),
	// 			to.Ptr("2"),
	// 		},
	// 	},
	// }
}
