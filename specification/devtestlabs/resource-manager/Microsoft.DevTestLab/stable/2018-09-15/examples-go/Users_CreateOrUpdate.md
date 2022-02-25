Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdevtestlabs%2Farmdevtestlabs%2Fv0.2.1/sdk/resourcemanager/devtestlabs/armdevtestlabs/README.md) on how to add the SDK to your project and authenticate.

```go
package armdevtestlabs_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Users_CreateOrUpdate.json
func ExampleUsersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdevtestlabs.NewUsersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<lab-name>",
		"<name>",
		armdevtestlabs.User{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"tagName1": to.StringPtr("tagValue1"),
			},
			Properties: &armdevtestlabs.UserProperties{
				Identity: &armdevtestlabs.UserIdentity{
					AppID:         to.StringPtr("<app-id>"),
					ObjectID:      to.StringPtr("<object-id>"),
					PrincipalID:   to.StringPtr("<principal-id>"),
					PrincipalName: to.StringPtr("<principal-name>"),
					TenantID:      to.StringPtr("<tenant-id>"),
				},
				SecretStore: &armdevtestlabs.UserSecretStore{
					KeyVaultID:  to.StringPtr("<key-vault-id>"),
					KeyVaultURI: to.StringPtr("<key-vault-uri>"),
				},
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
	log.Printf("Response result: %#v\n", res.UsersClientCreateOrUpdateResult)
}
```
