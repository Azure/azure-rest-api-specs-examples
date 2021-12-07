Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdataboxedge%2Farmdataboxedge%2Fv0.1.0/sdk/resourcemanager/databoxedge/armdataboxedge/README.md) on how to add the SDK to your project and authenticate.

```go
package armdataboxedge_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/TriggerSupportPackage.json
func ExampleSupportPackagesClient_BeginTriggerSupportPackage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewSupportPackagesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginTriggerSupportPackage(ctx,
		"<device-name>",
		"<resource-group-name>",
		armdataboxedge.TriggerSupportPackageRequest{
			Properties: &armdataboxedge.SupportPackageRequestProperties{
				Include:          to.StringPtr("<include>"),
				MaximumTimeStamp: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-12-18T02:19:51.4270267Z"); return t }()),
				MinimumTimeStamp: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-12-18T02:18:51.4270267Z"); return t }()),
			},
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
