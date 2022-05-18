Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpostgresql%2Farmpostgresqlflexibleservers%2Fv1.0.0/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/README.md) on how to add the SDK to your project and authenticate.

```go
package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/examples/CheckNameAvailability.json
func ExampleCheckNameAvailabilityClient_Execute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpostgresqlflexibleservers.NewCheckNameAvailabilityClient("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Execute(ctx,
		armpostgresqlflexibleservers.NameAvailabilityRequest{
			Name: to.Ptr("name1"),
			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
