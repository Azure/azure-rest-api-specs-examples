Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Foperationalinsights%2Farmoperationalinsights%2Fv0.3.1/sdk/resourcemanager/operationalinsights/armoperationalinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights"
)

// x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/LinkedStorageAccountsCreate.json
func ExampleLinkedStorageAccountsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armoperationalinsights.NewLinkedStorageAccountsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		armoperationalinsights.DataSourceTypeCustomLogs,
		armoperationalinsights.LinkedStorageAccountsResource{
			Properties: &armoperationalinsights.LinkedStorageAccountsProperties{
				StorageAccountIDs: []*string{
					to.StringPtr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.Storage/storageAccounts/testStorageA"),
					to.StringPtr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.Storage/storageAccounts/testStorageB")},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LinkedStorageAccountsClientCreateOrUpdateResult)
}
```
