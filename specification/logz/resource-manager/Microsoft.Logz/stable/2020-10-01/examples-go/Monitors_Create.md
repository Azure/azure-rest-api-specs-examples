Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flogz%2Farmlogz%2Fv0.1.0/sdk/resourcemanager/logz/armlogz/README.md) on how to add the SDK to your project and authenticate.

```go
package armlogz_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logz/armlogz"
)

// x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/Monitors_Create.json
func ExampleMonitorsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogz.NewMonitorsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<monitor-name>",
		&armlogz.MonitorsBeginCreateOptions{Body: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("LogzMonitorResource.ID: %s\n", *res.ID)
}
```
