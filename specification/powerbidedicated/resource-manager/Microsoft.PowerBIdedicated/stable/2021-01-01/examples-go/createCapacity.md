Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpowerbidedicated%2Farmpowerbidedicated%2Fv0.4.0/sdk/resourcemanager/powerbidedicated/armpowerbidedicated/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/createCapacity.json
func ExampleCapacitiesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armpowerbidedicated.NewCapacitiesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<dedicated-capacity-name>",
		armpowerbidedicated.DedicatedCapacity{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"testKey": to.Ptr("testValue"),
			},
			Properties: &armpowerbidedicated.DedicatedCapacityProperties{
				Administration: &armpowerbidedicated.DedicatedCapacityAdministrators{
					Members: []*string{
						to.Ptr("azsdktest@microsoft.com"),
						to.Ptr("azsdktest2@microsoft.com")},
				},
			},
			SKU: &armpowerbidedicated.CapacitySKU{
				Name: to.Ptr("<name>"),
				Tier: to.Ptr(armpowerbidedicated.CapacitySKUTierPBIEAzure),
			},
		},
		&armpowerbidedicated.CapacitiesClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
