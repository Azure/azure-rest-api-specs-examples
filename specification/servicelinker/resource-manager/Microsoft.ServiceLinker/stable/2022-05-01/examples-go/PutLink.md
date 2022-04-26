Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicelinker%2Farmservicelinker%2Fv0.4.0/sdk/resourcemanager/servicelinker/armservicelinker/README.md) on how to add the SDK to your project and authenticate.

```go
package armservicelinker_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/stable/2022-05-01/examples/PutLink.json
func ExampleLinkerClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armservicelinker.NewLinkerClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-uri>",
		"<linker-name>",
		armservicelinker.LinkerResource{
			Properties: &armservicelinker.LinkerProperties{
				AuthInfo: &armservicelinker.SecretAuthInfo{
					AuthType: to.Ptr(armservicelinker.AuthTypeSecret),
					Name:     to.Ptr("<name>"),
					SecretInfo: &armservicelinker.ValueSecretInfo{
						SecretType: to.Ptr(armservicelinker.SecretTypeRawValue),
						Value:      to.Ptr("<value>"),
					},
				},
				TargetService: &armservicelinker.AzureResource{
					Type: to.Ptr(armservicelinker.TypeAzureResource),
					ID:   to.Ptr("<id>"),
				},
			},
		},
		&armservicelinker.LinkerClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
