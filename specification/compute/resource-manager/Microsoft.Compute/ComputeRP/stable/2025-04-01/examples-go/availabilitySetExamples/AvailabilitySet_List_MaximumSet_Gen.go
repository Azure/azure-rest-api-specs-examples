package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/availabilitySetExamples/AvailabilitySet_List_MaximumSet_Gen.json
func ExampleAvailabilitySetsClient_NewListPager_availabilitySetListMaximumSetGen() {
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
		// 			Name: to.Ptr("{availabilitySetName}"),
		// 			Type: to.Ptr("Microsoft.Compute/availabilitySets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 			Location: to.Ptr("australiasoutheast"),
		// 			Tags: map[string]*string{
		// 				"key2039": to.Ptr("aaaaaaaaaaaaa"),
		// 			},
		// 			Properties: &armcompute.AvailabilitySetProperties{
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				PlatformUpdateDomainCount: to.Ptr[int32](5),
		// 				ProximityPlacementGroup: &armcompute.SubResource{
		// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 				},
		// 				Statuses: []*armcompute.InstanceViewStatus{
		// 					{
		// 						Code: to.Ptr("aaaaaaaaaaaaaaa"),
		// 						DisplayStatus: to.Ptr("aaaaaaaaaaa"),
		// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 						Message: to.Ptr("aaaaaa"),
		// 						Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T13:39:56.427Z"); return t}()),
		// 				}},
		// 				VirtualMachines: []*armcompute.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}"),
		// 				}},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Classic"),
		// 				Capacity: to.Ptr[int64](22),
		// 				Tier: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("{availabilitySetName}"),
		// 			Type: to.Ptr("Microsoft.Compute/availabilitySets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 			Location: to.Ptr("australiasoutheast"),
		// 			Tags: map[string]*string{
		// 				"key5849": to.Ptr("aaaaaaaaaaaaaaa"),
		// 			},
		// 			Properties: &armcompute.AvailabilitySetProperties{
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				PlatformUpdateDomainCount: to.Ptr[int32](5),
		// 				ProximityPlacementGroup: &armcompute.SubResource{
		// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 				},
		// 				Statuses: []*armcompute.InstanceViewStatus{
		// 					{
		// 						Code: to.Ptr("aaaaaaaaaaaaaaa"),
		// 						DisplayStatus: to.Ptr("aaaaaaaaaaa"),
		// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 						Message: to.Ptr("aaaaaa"),
		// 						Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T13:39:56.427Z"); return t}()),
		// 				}},
		// 				VirtualMachines: []*armcompute.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}"),
		// 				}},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Classic"),
		// 				Capacity: to.Ptr[int64](23),
		// 				Tier: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("{availabilitySetName}"),
		// 			Type: to.Ptr("Microsoft.Compute/availabilitySets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcompute.AvailabilitySetProperties{
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				PlatformUpdateDomainCount: to.Ptr[int32](5),
		// 				ProximityPlacementGroup: &armcompute.SubResource{
		// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 				},
		// 				Statuses: []*armcompute.InstanceViewStatus{
		// 					{
		// 						Code: to.Ptr("aaaaaaaaaaaaaaa"),
		// 						DisplayStatus: to.Ptr("aaaaaaaaaaa"),
		// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 						Message: to.Ptr("aaaaaa"),
		// 						Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T13:39:56.427Z"); return t}()),
		// 				}},
		// 				VirtualMachines: []*armcompute.SubResource{
		// 				},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Classic"),
		// 				Capacity: to.Ptr[int64](26),
		// 				Tier: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("{availabilitySetName}"),
		// 			Type: to.Ptr("Microsoft.Compute/availabilitySets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcompute.AvailabilitySetProperties{
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				PlatformUpdateDomainCount: to.Ptr[int32](5),
		// 				ProximityPlacementGroup: &armcompute.SubResource{
		// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 				},
		// 				Statuses: []*armcompute.InstanceViewStatus{
		// 					{
		// 						Code: to.Ptr("aaaaaaaaaaaaaaa"),
		// 						DisplayStatus: to.Ptr("aaaaaaaaaaa"),
		// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 						Message: to.Ptr("aaaaaa"),
		// 						Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T13:39:56.427Z"); return t}()),
		// 				}},
		// 				VirtualMachines: []*armcompute.SubResource{
		// 				},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Classic"),
		// 				Capacity: to.Ptr[int64](6),
		// 				Tier: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 			},
		// 	}},
		// }
	}
}
