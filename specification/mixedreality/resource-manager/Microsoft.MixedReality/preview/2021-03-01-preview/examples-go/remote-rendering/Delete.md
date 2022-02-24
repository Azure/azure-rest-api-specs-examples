Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmixedreality%2Farmmixedreality%2Fv0.2.1/sdk/resourcemanager/mixedreality/armmixedreality/README.md) on how to add the SDK to your project and authenticate.

```go
package armmixedreality_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mixedreality/armmixedreality"
)

// x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/Delete.json
func ExampleRemoteRenderingAccountsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmixedreality.NewRemoteRenderingAccountsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<account-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
