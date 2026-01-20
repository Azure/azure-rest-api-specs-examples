package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/availabilitySetExamples/AvailabilitySet_ListBySubscription.json
func ExampleAvailabilitySetsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailabilitySetsClient().NewListBySubscriptionPager(&armcompute.AvailabilitySetsClientListBySubscriptionOptions{Expand: to.Ptr("virtualMachines\\$ref")})
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
		// 			Properties: &armcompute.AvailabilitySetProperties{
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				PlatformUpdateDomainCount: to.Ptr[int32](5),
		// 				VirtualMachines: []*armcompute.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}"),
		// 				}},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Classic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("{availabilitySetName}"),
		// 			Type: to.Ptr("Microsoft.Compute/availabilitySets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 			Location: to.Ptr("australiasoutheast"),
		// 			Properties: &armcompute.AvailabilitySetProperties{
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				PlatformUpdateDomainCount: to.Ptr[int32](5),
		// 				VirtualMachines: []*armcompute.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}"),
		// 				}},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Classic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("{availabilitySetName}"),
		// 			Type: to.Ptr("Microsoft.Compute/availabilitySets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Tags: map[string]*string{
		// 				"{tagName}": to.Ptr("{tagValue}"),
		// 			},
		// 			Properties: &armcompute.AvailabilitySetProperties{
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				PlatformUpdateDomainCount: to.Ptr[int32](5),
		// 				VirtualMachines: []*armcompute.SubResource{
		// 				},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Classic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("{availabilitySetName}"),
		// 			Type: to.Ptr("Microsoft.Compute/availabilitySets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Tags: map[string]*string{
		// 				"{tagName}": to.Ptr("{tagValue}"),
		// 			},
		// 			Properties: &armcompute.AvailabilitySetProperties{
		// 				PlatformFaultDomainCount: to.Ptr[int32](3),
		// 				PlatformUpdateDomainCount: to.Ptr[int32](5),
		// 				VirtualMachines: []*armcompute.SubResource{
		// 				},
		// 			},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Classic"),
		// 			},
		// 	}},
		// }
	}
}
