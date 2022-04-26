Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpurview%2Farmpurview%2Fv0.4.0/sdk/resourcemanager/purview/armpurview/README.md) on how to add the SDK to your project and authenticate.

```go
package armpurview_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/DefaultAccounts_Set.json
func ExampleDefaultAccountsClient_Set() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armpurview.NewDefaultAccountsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Set(ctx,
		armpurview.DefaultAccountPayload{
			AccountName:       to.Ptr("<account-name>"),
			ResourceGroupName: to.Ptr("<resource-group-name>"),
			Scope:             to.Ptr("<scope>"),
			ScopeTenantID:     to.Ptr("<scope-tenant-id>"),
			ScopeType:         to.Ptr(armpurview.ScopeTypeTenant),
			SubscriptionID:    to.Ptr("<subscription-id>"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
