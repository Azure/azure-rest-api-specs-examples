Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv0.5.0/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ImportDatabaseWithNetworkIsolation.json
func ExampleDatabasesClient_BeginImport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsql.NewDatabasesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginImport(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<database-name>",
		armsql.ImportExistingDatabaseDefinition{
			AdministratorLogin:         to.Ptr("<administrator-login>"),
			AdministratorLoginPassword: to.Ptr("<administrator-login-password>"),
			AuthenticationType:         to.Ptr("<authentication-type>"),
			NetworkIsolation: &armsql.NetworkIsolationSettings{
				SQLServerResourceID:      to.Ptr("<sqlserver-resource-id>"),
				StorageAccountResourceID: to.Ptr("<storage-account-resource-id>"),
			},
			StorageKey:     to.Ptr("<storage-key>"),
			StorageKeyType: to.Ptr(armsql.StorageKeyTypeStorageAccessKey),
			StorageURI:     to.Ptr("<storage-uri>"),
		},
		&armsql.DatabasesClientBeginImportOptions{ResumeToken: ""})
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
