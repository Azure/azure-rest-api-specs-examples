Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmysql%2Farmmysqlflexibleservers%2Fv0.6.0/sdk/resourcemanager/mysql/armmysqlflexibleservers/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ConfigurationsBatchUpdate.json
func ExampleConfigurationsClient_BeginBatchUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmysqlflexibleservers.NewConfigurationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginBatchUpdate(ctx,
		"<resource-group-name>",
		"<server-name>",
		armmysqlflexibleservers.ConfigurationListForBatchUpdate{
			Value: []*armmysqlflexibleservers.ConfigurationForBatchUpdate{
				{
					Name: to.Ptr("<name>"),
					Properties: &armmysqlflexibleservers.ConfigurationForBatchUpdateProperties{
						Value: to.Ptr("<value>"),
					},
				},
				{
					Name: to.Ptr("<name>"),
					Properties: &armmysqlflexibleservers.ConfigurationForBatchUpdateProperties{
						Value: to.Ptr("<value>"),
					},
				}},
		},
		&armmysqlflexibleservers.ConfigurationsClientBeginBatchUpdateOptions{ResumeToken: ""})
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
