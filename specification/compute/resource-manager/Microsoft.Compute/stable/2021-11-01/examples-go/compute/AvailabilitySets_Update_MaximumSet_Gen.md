Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.4.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/AvailabilitySets_Update_MaximumSet_Gen.json
func ExampleAvailabilitySetsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewAvailabilitySetsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<availability-set-name>",
		armcompute.AvailabilitySetUpdate{
			Tags: map[string]*string{
				"key2574": to.StringPtr("aaaaaaaa"),
			},
			Properties: &armcompute.AvailabilitySetProperties{
				PlatformFaultDomainCount:  to.Int32Ptr(2),
				PlatformUpdateDomainCount: to.Int32Ptr(20),
				ProximityPlacementGroup: &armcompute.SubResource{
					ID: to.StringPtr("<id>"),
				},
				VirtualMachines: []*armcompute.SubResource{
					{
						ID: to.StringPtr("<id>"),
					}},
			},
			SKU: &armcompute.SKU{
				Name:     to.StringPtr("<name>"),
				Capacity: to.Int64Ptr(7),
				Tier:     to.StringPtr("<tier>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AvailabilitySetsClientUpdateResult)
}
```
