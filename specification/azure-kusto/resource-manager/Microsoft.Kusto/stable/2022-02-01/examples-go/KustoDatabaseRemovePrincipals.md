Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkusto%2Farmkusto%2Fv0.4.0/sdk/resourcemanager/kusto/armkusto/README.md) on how to add the SDK to your project and authenticate.

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
		return
	}
	ctx := context.Background()
	client, err := armkusto.NewDatabasesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.RemovePrincipals(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<database-name>",
		armkusto.DatabasePrincipalListRequest{
			Value: []*armkusto.DatabasePrincipal{
				{
					Name:  to.Ptr("<name>"),
					Type:  to.Ptr(armkusto.DatabasePrincipalTypeUser),
					AppID: to.Ptr("<app-id>"),
					Email: to.Ptr("<email>"),
					Fqn:   to.Ptr("<fqn>"),
					Role:  to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
				},
				{
					Name:  to.Ptr("<name>"),
					Type:  to.Ptr(armkusto.DatabasePrincipalTypeGroup),
					AppID: to.Ptr("<app-id>"),
					Email: to.Ptr("<email>"),
					Fqn:   to.Ptr("<fqn>"),
					Role:  to.Ptr(armkusto.DatabasePrincipalRoleViewer),
				},
				{
					Name:  to.Ptr("<name>"),
					Type:  to.Ptr(armkusto.DatabasePrincipalTypeApp),
					AppID: to.Ptr("<app-id>"),
					Email: to.Ptr("<email>"),
					Fqn:   to.Ptr("<fqn>"),
					Role:  to.Ptr(armkusto.DatabasePrincipalRoleAdmin),
				}},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
