Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv0.3.1/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

```go
package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseRecommendedColumnSensitivityLabelDisable.json
func ExampleManagedDatabaseSensitivityLabelsClient_DisableRecommendation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewManagedDatabaseSensitivityLabelsClient("<subscription-id>", cred, nil)
	_, err = client.DisableRecommendation(ctx,
		"<resource-group-name>",
		"<managed-instance-name>",
		"<database-name>",
		"<schema-name>",
		"<table-name>",
		"<column-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
