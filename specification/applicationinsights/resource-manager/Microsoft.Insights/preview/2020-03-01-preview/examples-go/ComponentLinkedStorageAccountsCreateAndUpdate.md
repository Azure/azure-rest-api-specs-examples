Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapplicationinsights%2Farmapplicationinsights%2Fv0.5.0/sdk/resourcemanager/applicationinsights/armapplicationinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2020-03-01-preview/examples/ComponentLinkedStorageAccountsCreateAndUpdate.json
func ExampleComponentLinkedStorageAccountsClient_CreateAndUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewComponentLinkedStorageAccountsClient("86dc51d3-92ed-4d7e-947a-775ea79b4918", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateAndUpdate(ctx,
		"someResourceGroupName",
		"myComponent",
		armapplicationinsights.StorageTypeServiceProfiler,
		armapplicationinsights.ComponentLinkedStorageAccounts{
			Properties: &armapplicationinsights.LinkedStorageAccountsProperties{
				LinkedStorageAccount: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4918/resourceGroups/someResourceGroupName/providers/Microsoft.Storage/storageAccounts/storageaccountname"),
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
