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

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/ServerCreate.json
func ExampleServersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewServersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<server-name>",
		armsql.Server{
			Location: to.StringPtr("<location>"),
			Properties: &armsql.ServerProperties{
				AdministratorLogin:         to.StringPtr("<administrator-login>"),
				AdministratorLoginPassword: to.StringPtr("<administrator-login-password>"),
				Administrators: &armsql.ServerExternalAdministrator{
					AzureADOnlyAuthentication: to.BoolPtr(true),
					Login:                     to.StringPtr("<login>"),
					PrincipalType:             armsql.PrincipalType("User").ToPtr(),
					Sid:                       to.StringPtr("<sid>"),
					TenantID:                  to.StringPtr("<tenant-id>"),
				},
				PublicNetworkAccess:           armsql.ServerNetworkAccessFlag("Enabled").ToPtr(),
				RestrictOutboundNetworkAccess: armsql.ServerNetworkAccessFlag("Enabled").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServersClientCreateOrUpdateResult)
}
```
