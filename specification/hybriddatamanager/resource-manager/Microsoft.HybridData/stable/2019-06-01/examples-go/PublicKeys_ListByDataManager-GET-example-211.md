Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybriddatamanager%2Farmhybriddatamanager%2Fv0.2.0/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager/README.md) on how to add the SDK to your project and authenticate.

```go
package armhybriddatamanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// x-ms-original-file: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/PublicKeys_ListByDataManager-GET-example-211.json
func ExamplePublicKeysClient_ListByDataManager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybriddatamanager.NewPublicKeysClient("<subscription-id>", cred, nil)
	pager := client.ListByDataManager("<resource-group-name>",
		"<data-manager-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}
```
