Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv0.3.1/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

```go
package armsql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ImportDatabaseWithNetworkIsolation.json
func ExampleDatabasesClient_BeginImport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewDatabasesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginImport(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<database-name>",
		armsql.ImportExistingDatabaseDefinition{
			AdministratorLogin:         to.StringPtr("<administrator-login>"),
			AdministratorLoginPassword: to.StringPtr("<administrator-login-password>"),
			AuthenticationType:         to.StringPtr("<authentication-type>"),
			NetworkIsolation: &armsql.NetworkIsolationSettings{
				SQLServerResourceID:      to.StringPtr("<sqlserver-resource-id>"),
				StorageAccountResourceID: to.StringPtr("<storage-account-resource-id>"),
			},
			StorageKey:     to.StringPtr("<storage-key>"),
			StorageKeyType: armsql.StorageKeyType("StorageAccessKey").ToPtr(),
			StorageURI:     to.StringPtr("<storage-uri>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DatabasesClientImportResult)
}
```
