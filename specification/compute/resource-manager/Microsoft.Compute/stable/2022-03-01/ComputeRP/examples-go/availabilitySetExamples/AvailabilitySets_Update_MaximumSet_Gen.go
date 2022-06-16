package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/availabilitySetExamples/AvailabilitySets_Update_MaximumSet_Gen.json
func ExampleAvailabilitySetsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewAvailabilitySetsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"rgcompute",
		"aaaaaaaaaaaaaaaaaaa",
		armcompute.AvailabilitySetUpdate{
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
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
