Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fprivatedns%2Farmprivatedns%2Fv0.3.1/sdk/resourcemanager/privatedns/armprivatedns/README.md) on how to add the SDK to your project and authenticate.

```go
package armprivatedns_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
)

// x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/PrivateZonePatch.json
func ExamplePrivateZonesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armprivatedns.NewPrivateZonesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<private-zone-name>",
		armprivatedns.PrivateZone{
			Tags: map[string]*string{
				"key2": to.StringPtr("value2"),
			},
		},
		&armprivatedns.PrivateZonesClientBeginUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateZonesClientUpdateResult)
}
```
