Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkusto%2Farmkusto%2Fv1.0.0/sdk/resourcemanager/kusto/armkusto/README.md) on how to add the SDK to your project and authenticate.

```go
package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoDatabaseRemovePrincipals.json
func ExampleDatabasesClient_RemovePrincipals() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armkusto.NewDatabasesClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.RemovePrincipals(ctx,
		"kustorptest",
		"kustoCluster",
		"KustoDatabase8",
		armkusto.DatabasePrincipalListRequest{
			Value: []*armkusto.DatabasePrincipal{
				{
					Name:  to.Ptr("Some User"),
					Type:  to.Ptr(armkusto.DatabasePrincipalTypeUser),
					AppID: to.Ptr(""),
					Email: to.Ptr("user@microsoft.com"),
					Fqn:   to.Ptr("aaduser=some_guid"),
					Role:  to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
				},
				{
					Name:  to.Ptr("Kusto"),
					Type:  to.Ptr(armkusto.DatabasePrincipalTypeGroup),
					AppID: to.Ptr(""),
					Email: to.Ptr("kusto@microsoft.com"),
					Fqn:   to.Ptr("aadgroup=some_guid"),
					Role:  to.Ptr(armkusto.DatabasePrincipalRoleViewer),
				},
				{
					Name:  to.Ptr("SomeApp"),
					Type:  to.Ptr(armkusto.DatabasePrincipalTypeApp),
					AppID: to.Ptr("some_guid_app_id"),
					Email: to.Ptr(""),
					Fqn:   to.Ptr("aadapp=some_guid_app_id"),
					Role:  to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
				}},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
