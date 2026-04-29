package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/availabilitySetExamples/AvailabilitySet_List_MaximumSet_Gen.json
func ExampleAvailabilitySetsClient_NewListPager_availabilitySetListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
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
		// page = armcompute.AvailabilitySetsClientListResponse{
		// 	AvailabilitySetListResult: armcompute.AvailabilitySetListResult{
		// 		Value: []*armcompute.AvailabilitySet{
		// 			{
		// 				Name: to.Ptr("{availabilitySetName}"),
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 				Type: to.Ptr("Microsoft.Compute/availabilitySets"),
		// 				Location: to.Ptr("australiasoutheast"),
		// 				Properties: &armcompute.AvailabilitySetProperties{
		// 					PlatformUpdateDomainCount: to.Ptr[int32](5),
		// 					PlatformFaultDomainCount: to.Ptr[int32](3),
		// 					VirtualMachines: []*armcompute.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}"),
		// 						},
		// 					},
		// 					ProximityPlacementGroup: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					Statuses: []*armcompute.InstanceViewStatus{
		// 						{
		// 							Code: to.Ptr("aaaaaaaaaaaaaaa"),
		// 							Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 							DisplayStatus: to.Ptr("aaaaaaaaaaa"),
		// 							Message: to.Ptr("aaaaaa"),
		// 							Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T13:39:56.427Z"); return t}()),
		// 						},
		// 					},
		// 				},
		// 				SKU: &armcompute.SKU{
		// 					Name: to.Ptr("Classic"),
		// 					Tier: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 					Capacity: to.Ptr[int64](22),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key2039": to.Ptr("aaaaaaaaaaaaa"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("{availabilitySetName}"),
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 				Type: to.Ptr("Microsoft.Compute/availabilitySets"),
		// 				Location: to.Ptr("australiasoutheast"),
		// 				Properties: &armcompute.AvailabilitySetProperties{
		// 					PlatformUpdateDomainCount: to.Ptr[int32](5),
		// 					PlatformFaultDomainCount: to.Ptr[int32](3),
		// 					VirtualMachines: []*armcompute.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}"),
		// 						},
		// 					},
		// 					ProximityPlacementGroup: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					Statuses: []*armcompute.InstanceViewStatus{
		// 						{
		// 							Code: to.Ptr("aaaaaaaaaaaaaaa"),
		// 							Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 							DisplayStatus: to.Ptr("aaaaaaaaaaa"),
		// 							Message: to.Ptr("aaaaaa"),
		// 							Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T13:39:56.427Z"); return t}()),
		// 						},
		// 					},
		// 				},
		// 				SKU: &armcompute.SKU{
		// 					Name: to.Ptr("Classic"),
		// 					Tier: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					Capacity: to.Ptr[int64](23),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key5849": to.Ptr("aaaaaaaaaaaaaaa"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("{availabilitySetName}"),
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 				Type: to.Ptr("Microsoft.Compute/availabilitySets"),
		// 				Location: to.Ptr("westcentralus"),
		// 				Tags: map[string]*string{
		// 				},
		// 				Properties: &armcompute.AvailabilitySetProperties{
		// 					PlatformUpdateDomainCount: to.Ptr[int32](5),
		// 					PlatformFaultDomainCount: to.Ptr[int32](3),
		// 					VirtualMachines: []*armcompute.SubResource{
		// 					},
		// 					ProximityPlacementGroup: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					Statuses: []*armcompute.InstanceViewStatus{
		// 						{
		// 							Code: to.Ptr("aaaaaaaaaaaaaaa"),
		// 							Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 							DisplayStatus: to.Ptr("aaaaaaaaaaa"),
		// 							Message: to.Ptr("aaaaaa"),
		// 							Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T13:39:56.427Z"); return t}()),
		// 						},
		// 					},
		// 				},
		// 				SKU: &armcompute.SKU{
		// 					Name: to.Ptr("Classic"),
		// 					Tier: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 					Capacity: to.Ptr[int64](26),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("{availabilitySetName}"),
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 				Type: to.Ptr("Microsoft.Compute/availabilitySets"),
		// 				Location: to.Ptr("westcentralus"),
		// 				Tags: map[string]*string{
		// 				},
		// 				Properties: &armcompute.AvailabilitySetProperties{
		// 					PlatformUpdateDomainCount: to.Ptr[int32](5),
		// 					PlatformFaultDomainCount: to.Ptr[int32](3),
		// 					VirtualMachines: []*armcompute.SubResource{
		// 					},
		// 					ProximityPlacementGroup: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					Statuses: []*armcompute.InstanceViewStatus{
		// 						{
		// 							Code: to.Ptr("aaaaaaaaaaaaaaa"),
		// 							Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 							DisplayStatus: to.Ptr("aaaaaaaaaaa"),
		// 							Message: to.Ptr("aaaaaa"),
		// 							Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T13:39:56.427Z"); return t}()),
		// 						},
		// 					},
		// 				},
		// 				SKU: &armcompute.SKU{
		// 					Name: to.Ptr("Classic"),
		// 					Tier: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 					Capacity: to.Ptr[int64](6),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("a://example.com/aaaaa"),
		// 	},
		// }
	}
}
