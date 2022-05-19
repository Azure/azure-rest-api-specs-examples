Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmysql%2Farmmysql%2Fv1.0.0/sdk/resourcemanager/mysql/armmysql/README.md) on how to add the SDK to your project and authenticate.

```go
package armmysql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/DatabaseCreate.json
func ExampleDatabasesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmysql.NewDatabasesClient("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"TestGroup",
		"testserver",
		"db1",
		armmysql.Database{
			Properties: &armmysql.DatabaseProperties{
				Charset:   to.Ptr("utf8"),
				Collation: to.Ptr("utf8_general_ci"),
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
