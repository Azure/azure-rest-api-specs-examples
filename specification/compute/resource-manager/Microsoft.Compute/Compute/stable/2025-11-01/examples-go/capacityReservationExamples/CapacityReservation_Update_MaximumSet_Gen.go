package armcompute_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/capacityReservationExamples/CapacityReservation_Update_MaximumSet_Gen.json
func ExampleCapacityReservationsClient_BeginUpdate_capacityReservationUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCapacityReservationsClient().BeginUpdate(ctx, "rgcompute", "aaaaaaaaaa", "aaaaaaaaaaaaaaaaaaa", armcompute.CapacityReservationUpdate{
		Properties: &armcompute.CapacityReservationProperties{
			InstanceView: &armcompute.CapacityReservationInstanceView{
				UtilizationInfo: &armcompute.CapacityReservationUtilization{},
				Statuses: []*armcompute.InstanceViewStatus{
					{
						Code:          to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
						Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
						DisplayStatus: to.Ptr("aaaaaa"),
						Message:       to.Ptr("a"),
						Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
					},
				},
			},
		},
		SKU: &armcompute.SKU{
			Name:     to.Ptr("Standard_DS1_v2"),
			Tier:     to.Ptr("aaa"),
			Capacity: to.Ptr[int64](7),
		},
		Tags: map[string]*string{
			"key4974": to.Ptr("aaaaaaaaaaaaaaaa"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.CapacityReservationsClientUpdateResponse{
	// 	CapacityReservation: &armcompute.CapacityReservation{
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 		},
	// 		SKU: &armcompute.SKU{
	// 			Name: to.Ptr("Standard_DS1_v2"),
	// 			Capacity: to.Ptr[int64](4),
	// 			Tier: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 		},
	// 		Zones: []*string{
	// 			to.Ptr("1"),
	// 		},
	// 		Properties: &armcompute.CapacityReservationProperties{
	// 			PlatformFaultDomainCount: to.Ptr[int32](3),
	// 			ReservationID: to.Ptr("{GUID}"),
	// 			VirtualMachinesAssociated: []*armcompute.SubResourceReadOnly{
	// 				{
	// 					ID: to.Ptr("aaaa"),
	// 				},
	// 			},
	// 			ProvisioningTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-27T01:02:38.3138469+00:00"); return t}()),
	// 			ProvisioningState: to.Ptr("Creating"),
	// 			InstanceView: &armcompute.CapacityReservationInstanceView{
	// 				UtilizationInfo: &armcompute.CapacityReservationUtilization{
	// 					VirtualMachinesAllocated: []*armcompute.SubResourceReadOnly{
	// 						{
	// 							ID: to.Ptr("aaaa"),
	// 						},
	// 					},
	// 				},
	// 				Statuses: []*armcompute.InstanceViewStatus{
	// 					{
	// 						Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 						DisplayStatus: to.Ptr("aaaaaa"),
	// 						Message: to.Ptr("a"),
	// 						Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/capacityReservationGroups/myCapacityReservationGroup/capacityReservations/myCapacityReservation"),
	// 		Name: to.Ptr("myCapacityReservation"),
	// 		Type: to.Ptr("aaaaaaaaaaaaaaa"),
	// 	},
	// }
}
