Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpostgresqlhsc%2Farmpostgresqlhsc%2Fv0.1.0/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc/README.md) on how to add the SDK to your project and authenticate.

```go
package armpostgresqlhsc_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc"
)

// x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/preview/2020-10-05-privatepreview/examples/ServerGet.json
func ExampleServersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpostgresqlhsc.NewServersClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<server-group-name>",
		"<server-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ServerGroupServer.ID: %s\n", *res.ID)
}
```
