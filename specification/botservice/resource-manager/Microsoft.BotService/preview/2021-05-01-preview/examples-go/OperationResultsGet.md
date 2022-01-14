Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fbotservice%2Farmbotservice%2Fv0.2.0/sdk/resourcemanager/botservice/armbotservice/README.md) on how to add the SDK to your project and authenticate.

```go
package armbotservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
)

// x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/OperationResultsGet.json
func ExampleOperationResultsClient_BeginGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbotservice.NewOperationResultsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginGet(ctx,
		"<operation-result-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.OperationResultsClientGetResult)
}
```
