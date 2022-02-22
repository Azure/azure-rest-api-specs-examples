Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpurview%2Farmpurview%2Fv0.2.1/sdk/resourcemanager/purview/armpurview/README.md) on how to add the SDK to your project and authenticate.

```go
package armpurview_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview"
)

// x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/DefaultAccounts_Set.json
func ExampleDefaultAccountsClient_Set() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpurview.NewDefaultAccountsClient(cred, nil)
	res, err := client.Set(ctx,
		armpurview.DefaultAccountPayload{
			AccountName:       to.StringPtr("<account-name>"),
			ResourceGroupName: to.StringPtr("<resource-group-name>"),
			Scope:             to.StringPtr("<scope>"),
			ScopeTenantID:     to.StringPtr("<scope-tenant-id>"),
			ScopeType:         armpurview.ScopeType("Tenant").ToPtr(),
			SubscriptionID:    to.StringPtr("<subscription-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DefaultAccountsClientSetResult)
}
```
