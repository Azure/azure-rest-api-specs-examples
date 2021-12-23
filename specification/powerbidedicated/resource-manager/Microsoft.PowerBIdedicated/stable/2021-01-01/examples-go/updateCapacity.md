Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpowerbidedicated%2Farmpowerbidedicated%2Fv0.1.0/sdk/resourcemanager/powerbidedicated/armpowerbidedicated/README.md) on how to add the SDK to your project and authenticate.

```go
package armpowerbidedicated_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated"
)

// x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/updateCapacity.json
func ExampleCapacitiesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpowerbidedicated.NewCapacitiesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<dedicated-capacity-name>",
		armpowerbidedicated.DedicatedCapacityUpdateParameters{
			Properties: &armpowerbidedicated.DedicatedCapacityMutableProperties{
				Administration: &armpowerbidedicated.DedicatedCapacityAdministrators{
					Members: []*string{
						to.StringPtr("azsdktest@microsoft.com"),
						to.StringPtr("azsdktest2@microsoft.com")},
				},
			},
			SKU: &armpowerbidedicated.CapacitySKU{
				Name: to.StringPtr("<name>"),
				Tier: armpowerbidedicated.CapacitySKUTierPBIEAzure.ToPtr(),
			},
			Tags: map[string]*string{
				"testKey": to.StringPtr("testValue"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DedicatedCapacity.ID: %s\n", *res.ID)
}
```
