Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmysql%2Farmmysqlflexibleservers%2Fv0.3.0/sdk/resourcemanager/mysql/armmysqlflexibleservers/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ConfigurationsBatchUpdate.json
func ExampleConfigurationsClient_BeginBatchUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmysqlflexibleservers.NewConfigurationsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginBatchUpdate(ctx,
		"<resource-group-name>",
		"<server-name>",
		armmysqlflexibleservers.ConfigurationListForBatchUpdate{
			Value: []*armmysqlflexibleservers.ConfigurationForBatchUpdate{
				{
					Name: to.StringPtr("<name>"),
					Properties: &armmysqlflexibleservers.ConfigurationForBatchUpdateProperties{
						Value: to.StringPtr("<value>"),
					},
				},
				{
					Name: to.StringPtr("<name>"),
					Properties: &armmysqlflexibleservers.ConfigurationForBatchUpdateProperties{
						Value: to.StringPtr("<value>"),
					},
				}},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ConfigurationsClientBatchUpdateResult)
}
```
