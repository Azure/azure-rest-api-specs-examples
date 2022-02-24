Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmixedreality%2Farmmixedreality%2Fv0.2.1/sdk/resourcemanager/mixedreality/armmixedreality/README.md) on how to add the SDK to your project and authenticate.

```go
package armmixedreality_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mixedreality/armmixedreality"
)

// x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/Patch.json
func ExampleRemoteRenderingAccountsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmixedreality.NewRemoteRenderingAccountsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		armmixedreality.RemoteRenderingAccount{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"hero":    to.StringPtr("romeo"),
				"heroine": to.StringPtr("juliet"),
			},
			Identity: &armmixedreality.Identity{
				Type: to.StringPtr("<type>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RemoteRenderingAccountsClientUpdateResult)
}
```
