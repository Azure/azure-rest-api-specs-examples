Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpostgresql%2Farmpostgresql%2Fv0.3.0/sdk/resourcemanager/postgresql/armpostgresql/README.md) on how to add the SDK to your project and authenticate.

```go
package armpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
)

// x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2018-06-01/examples/PrivateEndpointConnectionGet.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpostgresql.NewPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<private-endpoint-connection-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateEndpointConnectionsClientGetResult)
}
```
