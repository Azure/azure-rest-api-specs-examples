Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fredisenterprise%2Farmredisenterprise%2Fv0.2.1/sdk/resourcemanager/redisenterprise/armredisenterprise/README.md) on how to add the SDK to your project and authenticate.

```go
package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise"
)

// x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2021-08-01/examples/RedisEnterpriseDatabasesListKeys.json
func ExampleDatabasesClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armredisenterprise.NewDatabasesClient("<subscription-id>", cred, nil)
	res, err := client.ListKeys(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<database-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DatabasesClientListKeysResult)
}
```
