Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.2.1/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

```go
package armsynapse_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolLanguageExtensionsAdd.json
func ExampleKustoPoolsClient_BeginAddLanguageExtensions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginAddLanguageExtensions(ctx,
		"<workspace-name>",
		"<kusto-pool-name>",
		"<resource-group-name>",
		armsynapse.LanguageExtensionsList{
			Value: []*armsynapse.LanguageExtension{
				{
					LanguageExtensionName: armsynapse.LanguageExtensionName("PYTHON").ToPtr(),
				},
				{
					LanguageExtensionName: armsynapse.LanguageExtensionName("R").ToPtr(),
				}},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```
