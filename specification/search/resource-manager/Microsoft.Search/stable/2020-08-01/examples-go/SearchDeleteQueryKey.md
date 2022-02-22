Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsearch%2Farmsearch%2Fv0.3.1/sdk/resourcemanager/search/armsearch/README.md) on how to add the SDK to your project and authenticate.

```go
package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/SearchDeleteQueryKey.json
func ExampleQueryKeysClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsearch.NewQueryKeysClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<search-service-name>",
		"<key>",
		&armsearch.SearchManagementRequestOptions{ClientRequestID: nil})
	if err != nil {
		log.Fatal(err)
	}
}
```
