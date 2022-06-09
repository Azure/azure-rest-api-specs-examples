```go
package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/ServerCreate.json
func ExampleServersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewServersClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"sqlcrudtest-7398",
		"sqlcrudtest-4645",
		armsql.Server{
			Location: to.Ptr("Japan East"),
			Properties: &armsql.ServerProperties{
				AdministratorLogin:         to.Ptr("dummylogin"),
				AdministratorLoginPassword: to.Ptr("PLACEHOLDER"),
				Administrators: &armsql.ServerExternalAdministrator{
					AzureADOnlyAuthentication: to.Ptr(true),
					Login:                     to.Ptr("bob@contoso.com"),
					PrincipalType:             to.Ptr(armsql.PrincipalTypeUser),
					Sid:                       to.Ptr("00000011-1111-2222-2222-123456789111"),
					TenantID:                  to.Ptr("00000011-1111-2222-2222-123456789111"),
				},
				PublicNetworkAccess:           to.Ptr(armsql.ServerNetworkAccessFlagEnabled),
				RestrictOutboundNetworkAccess: to.Ptr(armsql.ServerNetworkAccessFlagEnabled),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv1.0.0/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.
