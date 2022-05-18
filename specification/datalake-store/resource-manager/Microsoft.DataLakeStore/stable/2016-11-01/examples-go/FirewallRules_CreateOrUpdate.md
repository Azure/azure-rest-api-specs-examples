Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatalake-store%2Farmdatalakestore%2Fv1.0.0/sdk/resourcemanager/datalake-store/armdatalakestore/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatalakestore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/FirewallRules_CreateOrUpdate.json
func ExampleFirewallRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatalakestore.NewFirewallRulesClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"contosorg",
		"contosoadla",
		"test_rule",
		armdatalakestore.CreateOrUpdateFirewallRuleParameters{
			Properties: &armdatalakestore.CreateOrUpdateFirewallRuleProperties{
				EndIPAddress:   to.Ptr("2.2.2.2"),
				StartIPAddress: to.Ptr("1.1.1.1"),
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
