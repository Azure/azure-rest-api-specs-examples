Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpowerbidedicated%2Farmpowerbidedicated%2Fv0.2.0/sdk/resourcemanager/powerbidedicated/armpowerbidedicated/README.md) on how to add the SDK to your project and authenticate.

```go
package armpowerbidedicated_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated"
)

// x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/createAutoScaleVCore.json
func ExampleAutoScaleVCoresClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpowerbidedicated.NewAutoScaleVCoresClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<vcore-name>",
		armpowerbidedicated.AutoScaleVCore{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"testKey": to.StringPtr("testValue"),
			},
			Properties: &armpowerbidedicated.AutoScaleVCoreProperties{
				CapacityLimit:    to.Int32Ptr(10),
				CapacityObjectID: to.StringPtr("<capacity-object-id>"),
			},
			SKU: &armpowerbidedicated.AutoScaleVCoreSKU{
				Name:     to.StringPtr("<name>"),
				Capacity: to.Int32Ptr(0),
				Tier:     armpowerbidedicated.VCoreSKUTier("AutoScale").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AutoScaleVCoresClientCreateResult)
}
```
