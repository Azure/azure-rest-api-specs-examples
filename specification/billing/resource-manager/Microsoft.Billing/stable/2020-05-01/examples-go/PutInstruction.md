Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fbilling%2Farmbilling%2Fv0.1.0/sdk/resourcemanager/billing/armbilling/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/PutInstruction.json
func ExampleInstructionsClient_Put() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewInstructionsClient(cred, nil)
	res, err := client.Put(ctx,
		"<billing-account-name>",
		"<billing-profile-name>",
		"<instruction-name>",
		armbilling.Instruction{
			Properties: &armbilling.InstructionProperties{
				Amount:    to.Float32Ptr(5000),
				EndDate:   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-30T21:26:47.997Z"); return t }()),
				StartDate: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-30T21:26:47.997Z"); return t }()),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Instruction.ID: %s\n", *res.ID)
}
```
