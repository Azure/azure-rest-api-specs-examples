package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/availabilitySetExamples/AvailabilitySet_Update_MaximumSet_Gen.json
func ExampleAvailabilitySetsClient_Update_availabilitySetUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAvailabilitySetsClient().Update(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaa", armcompute.AvailabilitySetUpdate{
		Tags: map[string]*string{
			"key2574": to.Ptr("aaaaaaaa"),
		},
		Properties: &armcompute.AvailabilitySetProperties{
			PlatformFaultDomainCount:  to.Ptr[int32](2),
			PlatformUpdateDomainCount: to.Ptr[int32](20),
			ProximityPlacementGroup: &armcompute.SubResource{
				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
			},
			VirtualMachines: []*armcompute.SubResource{
				{
					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
				}},
		},
		SKU: &armcompute.SKU{
			Name:     to.Ptr("DSv3-Type1"),
			Capacity: to.Ptr[int64](7),
			Tier:     to.Ptr("aaa"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AvailabilitySet = armcompute.AvailabilitySet{
	// 	Name: to.Ptr("myAvailabilitySet"),
	// 	Type: to.Ptr("Microsoft.Compute/availabilitySets"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/availabilitySets/myAvailabilitySet"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"key2505": to.Ptr("aa"),
	// 		"key9626": to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 	},
	// 	Properties: &armcompute.AvailabilitySetProperties{
	// 		PlatformFaultDomainCount: to.Ptr[int32](2),
	// 		PlatformUpdateDomainCount: to.Ptr[int32](20),
	// 		ProximityPlacementGroup: &armcompute.SubResource{
	// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 		},
	// 		Statuses: []*armcompute.InstanceViewStatus{
	// 			{
	// 				Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 				DisplayStatus: to.Ptr("aaaaaa"),
	// 				Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 				Message: to.Ptr("a"),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
	// 		}},
	// 		VirtualMachines: []*armcompute.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 		}},
	// 	},
	// 	SKU: &armcompute.SKU{
	// 		Name: to.Ptr("Classic"),
	// 		Capacity: to.Ptr[int64](29),
	// 		Tier: to.Ptr("aaaaaaaaaaaaaa"),
	// 	},
	// }
}
