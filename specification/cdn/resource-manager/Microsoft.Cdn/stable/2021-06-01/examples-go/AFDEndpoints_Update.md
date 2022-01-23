Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcdn%2Farmcdn%2Fv0.3.0/sdk/resourcemanager/cdn/armcdn/README.md) on how to add the SDK to your project and authenticate.

```go
package armcdn_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDEndpoints_Update.json
func ExampleAFDEndpointsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewAFDEndpointsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<endpoint-name>",
		armcdn.AFDEndpointUpdateParameters{
			Properties: &armcdn.AFDEndpointPropertiesUpdateParameters{
				EnabledState: armcdn.EnabledState("Enabled").ToPtr(),
			},
			Tags: map[string]*string{},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AFDEndpointsClientUpdateResult)
}
```
