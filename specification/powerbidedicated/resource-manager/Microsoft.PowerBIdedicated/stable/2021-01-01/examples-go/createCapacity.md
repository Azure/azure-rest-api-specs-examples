Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpowerbidedicated%2Farmpowerbidedicated%2Fv0.2.0/sdk/resourcemanager/powerbidedicated/armpowerbidedicated/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/createCapacity.json
func ExampleCapacitiesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpowerbidedicated.NewCapacitiesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<dedicated-capacity-name>",
		armpowerbidedicated.DedicatedCapacity{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"testKey": to.StringPtr("testValue"),
			},
			Properties: &armpowerbidedicated.DedicatedCapacityProperties{
				Administration: &armpowerbidedicated.DedicatedCapacityAdministrators{
					Members: []*string{
						to.StringPtr("azsdktest@microsoft.com"),
						to.StringPtr("azsdktest2@microsoft.com")},
				},
			},
			SKU: &armpowerbidedicated.CapacitySKU{
				Name: to.StringPtr("<name>"),
				Tier: armpowerbidedicated.CapacitySKUTier("PBIE_Azure").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.CapacitiesClientCreateResult)
}
```
