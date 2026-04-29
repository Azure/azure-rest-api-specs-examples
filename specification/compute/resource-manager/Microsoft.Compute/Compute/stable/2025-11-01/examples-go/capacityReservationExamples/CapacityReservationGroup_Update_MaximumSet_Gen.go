package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/capacityReservationExamples/CapacityReservationGroup_Update_MaximumSet_Gen.json
func ExampleCapacityReservationGroupsClient_Update_capacityReservationGroupUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCapacityReservationGroupsClient().Update(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaaaaa", armcompute.CapacityReservationGroupUpdate{
		Properties: &armcompute.CapacityReservationGroupProperties{
			InstanceView: &armcompute.CapacityReservationGroupInstanceView{},
		},
		Tags: map[string]*string{
			"key5355": to.Ptr("aaa"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.CapacityReservationGroupsClientUpdateResponse{
	// 	CapacityReservationGroup: &armcompute.CapacityReservationGroup{
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 		},
	// 		Zones: []*string{
	// 			to.Ptr("1"),
	// 			to.Ptr("2"),
	// 		},
	// 		Properties: &armcompute.CapacityReservationGroupProperties{
	// 			CapacityReservations: []*armcompute.SubResourceReadOnly{
	// 				{
	// 					ID: to.Ptr("aaaa"),
	// 				},
	// 			},
	// 			VirtualMachinesAssociated: []*armcompute.SubResourceReadOnly{
	// 				{
	// 					ID: to.Ptr("aaaa"),
	// 				},
	// 			},
	// 			InstanceView: &armcompute.CapacityReservationGroupInstanceView{
	// 				CapacityReservations: []*armcompute.CapacityReservationInstanceViewWithName{
	// 					{
	// 						Name: to.Ptr("aaaaaaaaaaaaaaaa"),
	// 						UtilizationInfo: &armcompute.CapacityReservationUtilization{
	// 							VirtualMachinesAllocated: []*armcompute.SubResourceReadOnly{
	// 								{
	// 									ID: to.Ptr("aaaa"),
	// 								},
	// 							},
	// 						},
	// 						Statuses: []*armcompute.InstanceViewStatus{
	// 							{
	// 								Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 								Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 								DisplayStatus: to.Ptr("aaaaaa"),
	// 								Message: to.Ptr("a"),
	// 								Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/CapacityReservationGroups/myCapacityReservationGroup"),
	// 		Name: to.Ptr("myCapacityReservationGroup"),
	// 		Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 	},
	// }
}
