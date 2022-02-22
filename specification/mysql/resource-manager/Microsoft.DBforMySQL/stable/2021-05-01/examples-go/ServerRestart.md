Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmysql%2Farmmysqlflexibleservers%2Fv0.3.1/sdk/resourcemanager/mysql/armmysqlflexibleservers/README.md) on how to add the SDK to your project and authenticate.

```go
package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers"
)

// x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ServerRestart.json
func ExampleServersClient_BeginRestart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmysqlflexibleservers.NewServersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginRestart(ctx,
		"<resource-group-name>",
		"<server-name>",
		armmysqlflexibleservers.ServerRestartParameter{
			MaxFailoverSeconds:  to.Int32Ptr(60),
			RestartWithFailover: armmysqlflexibleservers.EnableStatusEnum("Enabled").ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```
