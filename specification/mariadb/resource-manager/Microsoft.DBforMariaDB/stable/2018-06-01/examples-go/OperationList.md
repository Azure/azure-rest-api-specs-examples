Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmariadb%2Farmmariadb%2Fv0.1.0/sdk/resourcemanager/mariadb/armmariadb/README.md) on how to add the SDK to your project and authenticate.

```go
package armmariadb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
)

// x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/OperationList.json
func ExampleOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmariadb.NewOperationsClient(cred, nil)
	_, err = client.List(ctx,
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
