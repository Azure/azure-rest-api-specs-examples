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

// x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2021-11-01-preview/examples/PatchLink.json
func ExampleLinkerClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicelinker.NewLinkerClient(cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-uri>",
		"<linker-name>",
		armservicelinker.LinkerPatch{
			Properties: &armservicelinker.LinkerProperties{
				AuthInfo: &armservicelinker.ServicePrincipalSecretAuthInfo{
					AuthType:    armservicelinker.AuthType("servicePrincipalSecret").ToPtr(),
					ClientID:    to.StringPtr("<client-id>"),
					PrincipalID: to.StringPtr("<principal-id>"),
					Secret:      to.StringPtr("<secret>"),
				},
				TargetID: to.StringPtr("<target-id>"),
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
	log.Printf("Response result: %#v\n", res.LinkerClientUpdateResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicelinker%2Farmservicelinker%2Fv0.2.1/sdk/resourcemanager/servicelinker/armservicelinker/README.md) on how to add the SDK to your project and authenticate.
