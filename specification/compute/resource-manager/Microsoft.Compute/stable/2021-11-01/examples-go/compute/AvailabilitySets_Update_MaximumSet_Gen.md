Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.7.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/AvailabilitySets_Update_MaximumSet_Gen.json
func ExampleAvailabilitySetsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewAvailabilitySetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<availability-set-name>",
		armcompute.AvailabilitySetUpdate{
			Tags: map[string]*string{
				"key2574": to.Ptr("aaaaaaaa"),
			},
			Properties: &armcompute.AvailabilitySetProperties{
				PlatformFaultDomainCount:  to.Ptr[int32](2),
				PlatformUpdateDomainCount: to.Ptr[int32](20),
				ProximityPlacementGroup: &armcompute.SubResource{
					ID: to.Ptr("<id>"),
				},
				VirtualMachines: []*armcompute.SubResource{
					{
						ID: to.Ptr("<id>"),
					}},
			},
			SKU: &armcompute.SKU{
				Name:     to.Ptr("<name>"),
				Capacity: to.Ptr[int64](7),
				Tier:     to.Ptr("<tier>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
