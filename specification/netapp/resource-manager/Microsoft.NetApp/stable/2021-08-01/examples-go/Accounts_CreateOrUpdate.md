Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetapp%2Farmnetapp%2Fv0.1.0/sdk/resourcemanager/netapp/armnetapp/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetapp_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp"
)

// x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-08-01/examples/Accounts_CreateOrUpdate.json
func ExampleAccountsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetapp.NewAccountsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armnetapp.NetAppAccount{
			Location: to.StringPtr("<location>"),
			Properties: &armnetapp.AccountProperties{
				ActiveDirectories: []*armnetapp.ActiveDirectory{
					{
						AesEncryption:      to.BoolPtr(true),
						DNS:                to.StringPtr("<dns>"),
						Domain:             to.StringPtr("<domain>"),
						LdapOverTLS:        to.BoolPtr(false),
						LdapSigning:        to.BoolPtr(false),
						OrganizationalUnit: to.StringPtr("<organizational-unit>"),
						Password:           to.StringPtr("<password>"),
						Site:               to.StringPtr("<site>"),
						SmbServerName:      to.StringPtr("<smb-server-name>"),
						Username:           to.StringPtr("<username>"),
					}},
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
	log.Printf("NetAppAccount.ID: %s\n", *res.ID)
}
```
