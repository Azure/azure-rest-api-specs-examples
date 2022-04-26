Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpostgresql%2Farmpostgresql%2Fv0.5.0/sdk/resourcemanager/postgresql/armpostgresql/README.md) on how to add the SDK to your project and authenticate.

```go
package armpostgresql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2018-06-01/examples/PrivateEndpointConnectionUpdateTags.json
func ExamplePrivateEndpointConnectionsClient_BeginUpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armpostgresql.NewPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdateTags(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<private-endpoint-connection-name>",
		armpostgresql.TagsObject{
			Tags: map[string]*string{
				"key1": to.Ptr("val1"),
				"key2": to.Ptr("val2"),
				"key3": to.Ptr("val3"),
			},
		},
		&armpostgresql.PrivateEndpointConnectionsClientBeginUpdateTagsOptions{ResumeToken: ""})
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
