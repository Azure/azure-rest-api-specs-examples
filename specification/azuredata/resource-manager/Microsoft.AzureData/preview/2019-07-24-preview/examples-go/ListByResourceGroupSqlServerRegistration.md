Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fazuredata%2Farmazuredata%2Fv0.4.0/sdk/resourcemanager/azuredata/armazuredata/README.md) on how to add the SDK to your project and authenticate.

```go
package armazuredata_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azuredata/resource-manager/Microsoft.AzureData/preview/2019-07-24-preview/examples/ListByResourceGroupSqlServerRegistration.json
func ExampleSQLServerRegistrationsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armazuredata.NewSQLServerRegistrationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListByResourceGroupPager("<resource-group-name>",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
```
