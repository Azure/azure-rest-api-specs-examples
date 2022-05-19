Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpowerbidedicated%2Farmpowerbidedicated%2Fv1.0.0/sdk/resourcemanager/powerbidedicated/armpowerbidedicated/README.md) on how to add the SDK to your project and authenticate.

```go
package armpowerbidedicated_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/createAutoScaleVCore.json
func ExampleAutoScaleVCoresClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpowerbidedicated.NewAutoScaleVCoresClient("613192d7-503f-477a-9cfe-4efc3ee2bd60", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"TestRG",
		"testvcore",
		armpowerbidedicated.AutoScaleVCore{
			Location: to.Ptr("West US"),
			Tags: map[string]*string{
				"testKey": to.Ptr("testValue"),
			},
			Properties: &armpowerbidedicated.AutoScaleVCoreProperties{
				CapacityLimit:    to.Ptr[int32](10),
				CapacityObjectID: to.Ptr("a28f00bd-5330-4572-88f1-fa883e074785"),
			},
			SKU: &armpowerbidedicated.AutoScaleVCoreSKU{
				Name:     to.Ptr("AutoScale"),
				Capacity: to.Ptr[int32](0),
				Tier:     to.Ptr(armpowerbidedicated.VCoreSKUTierAutoScale),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
