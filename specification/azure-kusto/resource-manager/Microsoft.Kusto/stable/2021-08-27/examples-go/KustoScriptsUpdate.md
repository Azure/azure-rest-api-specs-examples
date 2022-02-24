Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkusto%2Farmkusto%2Fv0.2.1/sdk/resourcemanager/kusto/armkusto/README.md) on how to add the SDK to your project and authenticate.

```go
package armkusto_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoScriptsUpdate.json
func ExampleScriptsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkusto.NewScriptsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<database-name>",
		"<script-name>",
		armkusto.Script{
			Properties: &armkusto.ScriptProperties{
				ContinueOnErrors:  to.BoolPtr(true),
				ForceUpdateTag:    to.StringPtr("<force-update-tag>"),
				ScriptURL:         to.StringPtr("<script-url>"),
				ScriptURLSasToken: to.StringPtr("<script-urlsas-token>"),
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
	log.Printf("Response result: %#v\n", res.ScriptsClientUpdateResult)
}
```
