Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsolutions%2Farmmanagedapplications%2Fv0.1.0/sdk/resourcemanager/solutions/armmanagedapplications/README.md) on how to add the SDK to your project and authenticate.

```go
package armmanagedapplications_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications"
)

// x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/updateJitRequest.json
func ExampleJitRequestsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmanagedapplications.NewJitRequestsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<jit-request-name>",
		armmanagedapplications.JitRequestPatchable{
			Tags: map[string]*string{
				"department": to.StringPtr("Finance"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("JitRequestDefinition.ID: %s\n", *res.ID)
}
```
