package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/proximityPlacementGroupExamples/ProximityPlacementGroup_CreateOrUpdate.json
func ExampleProximityPlacementGroupsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewProximityPlacementGroupsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "myResourceGroup", "myProximityPlacementGroup", armcompute.ProximityPlacementGroup{
		Location: to.Ptr("westus"),
		Properties: &armcompute.ProximityPlacementGroupProperties{
			Intent: &armcompute.ProximityPlacementGroupPropertiesIntent{
				VMSizes: []*string{
					to.Ptr("Basic_A0"),
					to.Ptr("Basic_A2")},
			},
			ProximityPlacementGroupType: to.Ptr(armcompute.ProximityPlacementGroupTypeStandard),
		},
		Zones: []*string{
			to.Ptr("1")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
