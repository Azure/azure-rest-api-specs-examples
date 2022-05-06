Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fanalysisservices%2Farmanalysisservices%2Fv0.4.0/sdk/resourcemanager/analysisservices/armanalysisservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armanalysisservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/analysisservices/armanalysisservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2017-08-01/examples/getServer.json
func ExampleServersClient_GetDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armanalysisservices.NewServersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetDetails(ctx,
		"<resource-group-name>",
		"<server-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
