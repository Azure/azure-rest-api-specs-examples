Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapplicationinsights%2Farmapplicationinsights%2Fv0.2.0/sdk/resourcemanager/applicationinsights/armapplicationinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2020-03-01-preview/examples/ComponentLinkedStorageAccountsCreateAndUpdate.json
func ExampleComponentLinkedStorageAccountsClient_CreateAndUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewComponentLinkedStorageAccountsClient("<subscription-id>", cred, nil)
	res, err := client.CreateAndUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armapplicationinsights.StorageType("ServiceProfiler"),
		armapplicationinsights.ComponentLinkedStorageAccounts{
			Properties: &armapplicationinsights.LinkedStorageAccountsProperties{
				LinkedStorageAccount: to.StringPtr("<linked-storage-account>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ComponentLinkedStorageAccountsClientCreateAndUpdateResult)
}
```
