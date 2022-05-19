Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicelinker%2Farmservicelinker%2Fv1.0.0/sdk/resourcemanager/servicelinker/armservicelinker/README.md) on how to add the SDK to your project and authenticate.

```go
package armservicelinker_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/stable/2022-05-01/examples/PutLink.json
func ExampleLinkerClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicelinker.NewLinkerClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app",
		"linkName",
		armservicelinker.LinkerResource{
			Properties: &armservicelinker.LinkerProperties{
				AuthInfo: &armservicelinker.SecretAuthInfo{
					AuthType: to.Ptr(armservicelinker.AuthTypeSecret),
					Name:     to.Ptr("name"),
					SecretInfo: &armservicelinker.ValueSecretInfo{
						SecretType: to.Ptr(armservicelinker.SecretTypeRawValue),
						Value:      to.Ptr("secret"),
					},
				},
				TargetService: &armservicelinker.AzureResource{
					Type: to.Ptr(armservicelinker.TargetServiceTypeAzureResource),
					ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DBforPostgreSQL/servers/test-pg/databases/test-db"),
				},
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
