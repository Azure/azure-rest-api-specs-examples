Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fbilling%2Farmbilling%2Fv0.4.0/sdk/resourcemanager/billing/armbilling/README.md) on how to add the SDK to your project and authenticate.

```go
package armbilling_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/PutInstruction.json
func ExampleInstructionsClient_Put() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armbilling.NewInstructionsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Put(ctx,
		"<billing-account-name>",
		"<billing-profile-name>",
		"<instruction-name>",
		armbilling.Instruction{
			Properties: &armbilling.InstructionProperties{
				Amount:    to.Ptr[float32](5000),
				EndDate:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-30T21:26:47.997Z"); return t }()),
				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-30T21:26:47.997Z"); return t }()),
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
