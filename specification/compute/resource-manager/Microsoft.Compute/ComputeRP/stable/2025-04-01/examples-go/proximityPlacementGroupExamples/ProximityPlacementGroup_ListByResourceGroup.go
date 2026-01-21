package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/proximityPlacementGroupExamples/ProximityPlacementGroup_ListByResourceGroup.json
func ExampleProximityPlacementGroupsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProximityPlacementGroupsClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page.ProximityPlacementGroupListResult = armcompute.ProximityPlacementGroupListResult{
		// 	Value: []*armcompute.ProximityPlacementGroup{
		// 		{
		// 			Name: to.Ptr("myProximityPlacementGroup"),
		// 			Type: to.Ptr("Microsoft.Compute/proximityPlacementGroups"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/proximityPlacementGroups/myProximityPlacementGroup"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armcompute.ProximityPlacementGroupProperties{
		// 				AvailabilitySets: []*armcompute.SubResourceWithColocationStatus{
		// 					{
		// 						ID: to.Ptr("string"),
		// 				}},
		// 				Intent: &armcompute.ProximityPlacementGroupPropertiesIntent{
		// 					VMSizes: []*string{
		// 						to.Ptr("Basic_A0"),
		// 						to.Ptr("Basic_A2")},
		// 					},
		// 					ProximityPlacementGroupType: to.Ptr(armcompute.ProximityPlacementGroupTypeStandard),
		// 					VirtualMachineScaleSets: []*armcompute.SubResourceWithColocationStatus{
		// 						{
		// 							ID: to.Ptr("string"),
		// 					}},
		// 					VirtualMachines: []*armcompute.SubResourceWithColocationStatus{
		// 						{
		// 							ID: to.Ptr("string"),
		// 					}},
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("1")},
		// 			}},
		// 		}
	}
}
