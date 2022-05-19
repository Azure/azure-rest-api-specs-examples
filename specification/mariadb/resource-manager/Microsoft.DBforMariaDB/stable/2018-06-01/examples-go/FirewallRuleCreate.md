Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmariadb%2Farmmariadb%2Fv1.0.0/sdk/resourcemanager/mariadb/armmariadb/README.md) on how to add the SDK to your project and authenticate.

```go
package armmariadb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/FirewallRuleCreate.json
func ExampleFirewallRulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmariadb.NewFirewallRulesClient("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"TestGroup",
		"testserver",
		"rule1",
		armmariadb.FirewallRule{
			Properties: &armmariadb.FirewallRuleProperties{
				EndIPAddress:   to.Ptr("255.255.255.255"),
				StartIPAddress: to.Ptr("0.0.0.0"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
