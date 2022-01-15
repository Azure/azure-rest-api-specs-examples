Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fazurearcdata%2Farmazurearcdata%2Fv0.2.0/sdk/resourcemanager/azurearcdata/armazurearcdata/README.md) on how to add the SDK to your project and authenticate.

```go
package armazurearcdata_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
)

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/UpdateSqlManagedInstance.json
func ExampleSQLManagedInstancesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewSQLManagedInstancesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<sql-managed-instance-name>",
		armazurearcdata.SQLManagedInstanceUpdate{
			Tags: map[string]*string{
				"mytag": to.StringPtr("myval"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SQLManagedInstancesClientUpdateResult)
}
```
