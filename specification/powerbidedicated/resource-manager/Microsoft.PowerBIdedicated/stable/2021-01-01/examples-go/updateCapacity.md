```go
package armpowerbidedicated_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/updateCapacity.json
func ExampleCapacitiesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpowerbidedicated.NewCapacitiesClient("613192d7-503f-477a-9cfe-4efc3ee2bd60", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"TestRG",
		"azsdktest",
		armpowerbidedicated.DedicatedCapacityUpdateParameters{
			Properties: &armpowerbidedicated.DedicatedCapacityMutableProperties{
				Administration: &armpowerbidedicated.DedicatedCapacityAdministrators{
					Members: []*string{
						to.Ptr("azsdktest@microsoft.com"),
						to.Ptr("azsdktest2@microsoft.com")},
				},
			},
			SKU: &armpowerbidedicated.CapacitySKU{
				Name: to.Ptr("A1"),
				Tier: to.Ptr(armpowerbidedicated.CapacitySKUTierPBIEAzure),
			},
			Tags: map[string]*string{
				"testKey": to.Ptr("testValue"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpowerbidedicated%2Farmpowerbidedicated%2Fv1.0.0/sdk/resourcemanager/powerbidedicated/armpowerbidedicated/README.md) on how to add the SDK to your project and authenticate.
