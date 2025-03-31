package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/250861bb6a886b75255edfa0aa5ee2dd0d6e7a11/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/capacityReservationExamples/CapacityReservationGroup_Update_MaximumSet_Gen.json
func ExampleCapacityReservationGroupsClient_Update_capacityReservationGroupUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCapacityReservationGroupsClient().Update(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaaaaa", armcompute.CapacityReservationGroupUpdate{
		Tags: map[string]*string{
			"key5355": to.Ptr("aaa"),
		},
		Properties: &armcompute.CapacityReservationGroupProperties{
			InstanceView: &armcompute.CapacityReservationGroupInstanceView{},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CapacityReservationGroup = armcompute.CapacityReservationGroup{
	// 	Name: to.Ptr("myCapacityReservationGroup"),
	// 	Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/CapacityReservationGroups/myCapacityReservationGroup"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcompute.CapacityReservationGroupProperties{
	// 		CapacityReservations: []*armcompute.SubResourceReadOnly{
	// 			{
	// 				ID: to.Ptr("aaaa"),
	// 		}},
	// 		InstanceView: &armcompute.CapacityReservationGroupInstanceView{
	// 			CapacityReservations: []*armcompute.CapacityReservationInstanceViewWithName{
	// 				{
	// 					Statuses: []*armcompute.InstanceViewStatus{
	// 						{
	// 							Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 							DisplayStatus: to.Ptr("aaaaaa"),
	// 							Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 							Message: to.Ptr("a"),
	// 							Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
	// 					}},
	// 					UtilizationInfo: &armcompute.CapacityReservationUtilization{
	// 						VirtualMachinesAllocated: []*armcompute.SubResourceReadOnly{
	// 							{
	// 								ID: to.Ptr("aaaa"),
	// 						}},
	// 					},
	// 					Name: to.Ptr("aaaaaaaaaaaaaaaa"),
	// 			}},
	// 		},
	// 		VirtualMachinesAssociated: []*armcompute.SubResourceReadOnly{
	// 			{
	// 				ID: to.Ptr("aaaa"),
	// 		}},
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1"),
	// 		to.Ptr("2")},
	// 	}
}
