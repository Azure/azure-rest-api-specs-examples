package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/availabilitySetExamples/AvailabilitySet_Create.json
func ExampleAvailabilitySetsClient_CreateOrUpdate_createAnAvailabilitySet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.AvailabilitySet = armcompute.AvailabilitySet{
	// 	Name: to.Ptr("myAvailabilitySet"),
	// 	Type: to.Ptr("Microsoft.Compute/availabilitySets"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/availabilitySets/myAvailabilitySet"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armcompute.AvailabilitySetProperties{
	// 		PlatformFaultDomainCount: to.Ptr[int32](2),
	// 		PlatformUpdateDomainCount: to.Ptr[int32](20),
	// 	},
	// 	SKU: &armcompute.SKU{
	// 		Name: to.Ptr("Classic"),
	// 	},
	// }
}
