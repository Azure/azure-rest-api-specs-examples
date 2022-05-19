Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fedgeorderpartner%2Farmedgeorderpartner%2Fv0.5.0/sdk/resourcemanager/edgeorderpartner/armedgeorderpartner/README.md) on how to add the SDK to your project and authenticate.

```go
package armedgeorderpartner_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorderpartner/armedgeorderpartner"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorderpartner/resource-manager/Microsoft.EdgeOrderPartner/preview/2020-12-01-preview/examples/ListOperationsPartner.json
func ExampleAPISClient_NewListOperationsPartnerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armedgeorderpartner.NewAPISClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListOperationsPartnerPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
```
