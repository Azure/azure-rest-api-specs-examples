Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmysql%2Farmmysql%2Fv0.5.0/sdk/resourcemanager/mysql/armmysql/README.md) on how to add the SDK to your project and authenticate.

```go
package armmysql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2020-01-01/examples/ServerKeyCreateOrUpdate.json
func ExampleServerKeysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmysql.NewServerKeysClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<server-name>",
		"<key-name>",
		"<resource-group-name>",
		armmysql.ServerKey{
			Properties: &armmysql.ServerKeyProperties{
				ServerKeyType: to.Ptr(armmysql.ServerKeyTypeAzureKeyVault),
				URI:           to.Ptr("<uri>"),
			},
		},
		&armmysql.ServerKeysClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
