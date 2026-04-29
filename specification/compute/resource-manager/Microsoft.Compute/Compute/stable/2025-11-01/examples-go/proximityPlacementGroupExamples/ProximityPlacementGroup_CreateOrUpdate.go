package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/proximityPlacementGroupExamples/ProximityPlacementGroup_CreateOrUpdate.json
func ExampleProximityPlacementGroupsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProximityPlacementGroupsClient().CreateOrUpdate(ctx, "myResourceGroup", "myProximityPlacementGroup", armcompute.ProximityPlacementGroup{
		Location: to.Ptr("westus"),
		Zones: []*string{
			to.Ptr("1"),
		},
		Properties: &armcompute.ProximityPlacementGroupProperties{
			ProximityPlacementGroupType: to.Ptr(armcompute.ProximityPlacementGroupTypeStandard),
			Intent: &armcompute.ProximityPlacementGroupPropertiesIntent{
				VMSizes: []*string{
					to.Ptr("Basic_A0"),
					to.Ptr("Basic_A2"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.ProximityPlacementGroupsClientCreateOrUpdateResponse{
	// 	ProximityPlacementGroup: &armcompute.ProximityPlacementGroup{
	// 		Name: to.Ptr("myProximityPlacementGroup"),
	// 		Properties: &armcompute.ProximityPlacementGroupProperties{
	// 			ProximityPlacementGroupType: to.Ptr(armcompute.ProximityPlacementGroupTypeStandard),
	// 			Intent: &armcompute.ProximityPlacementGroupPropertiesIntent{
	// 				VMSizes: []*string{
	// 					to.Ptr("Basic_A0"),
	// 					to.Ptr("Basic_A2"),
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		Zones: []*string{
	// 			to.Ptr("1"),
	// 		},
	// 		Type: to.Ptr("Microsoft.Compute/proximityPlacementGroups"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/proximityPlacementGroups/myProximityPlacementGroup"),
	// 	},
	// }
}
