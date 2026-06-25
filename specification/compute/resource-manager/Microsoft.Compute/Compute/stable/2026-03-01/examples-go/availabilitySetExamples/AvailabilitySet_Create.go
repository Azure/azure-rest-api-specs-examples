package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/availabilitySetExamples/AvailabilitySet_Create.json
func ExampleAvailabilitySetsClient_CreateOrUpdate_createAnAvailabilitySet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAvailabilitySetsClient().CreateOrUpdate(ctx, "myResourceGroup", "myAvailabilitySet", armcompute.AvailabilitySet{
		Location: to.Ptr("westus"),
		Properties: &armcompute.AvailabilitySetProperties{
			PlatformFaultDomainCount:  to.Ptr[int32](2),
			PlatformUpdateDomainCount: to.Ptr[int32](20),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.AvailabilitySetsClientCreateOrUpdateResponse{
	// 	AvailabilitySet: armcompute.AvailabilitySet{
	// 		SKU: &armcompute.SKU{
	// 			Name: to.Ptr("Classic"),
	// 		},
	// 		Name: to.Ptr("myAvailabilitySet"),
	// 		Properties: &armcompute.AvailabilitySetProperties{
	// 			PlatformFaultDomainCount: to.Ptr[int32](2),
	// 			PlatformUpdateDomainCount: to.Ptr[int32](20),
	// 			VirtualMachineScaleSetMigrationInfo: &armcompute.VirtualMachineScaleSetMigrationInfo{
	// 				DefaultVirtualMachineScaleSetInfo: &armcompute.DefaultVirtualMachineScaleSetInfo{
	// 					DefaultVirtualMachineScaleSet: &armcompute.SubResource{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}"),
	// 					},
	// 					ConstrainedMaximumCapacity: to.Ptr(false),
	// 				},
	// 				MigrateToVirtualMachineScaleSet: &armcompute.SubResource{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}"),
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		Type: to.Ptr("Microsoft.Compute/availabilitySets"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/availabilitySets/myAvailabilitySet"),
	// 	},
	// }
}
