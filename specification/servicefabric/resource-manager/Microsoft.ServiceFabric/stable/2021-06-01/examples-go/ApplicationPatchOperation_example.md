Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicefabric%2Farmservicefabric%2Fv0.3.1/sdk/resourcemanager/servicefabric/armservicefabric/README.md) on how to add the SDK to your project and authenticate.

```go
package armservicefabric_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric"
)

// x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationPatchOperation_example.json
func ExampleApplicationsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabric.NewApplicationsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<application-name>",
		armservicefabric.ApplicationResourceUpdate{
			Location: to.StringPtr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armservicefabric.ApplicationResourceUpdateProperties{
				Metrics: []*armservicefabric.ApplicationMetricDescription{
					{
						Name:                     to.StringPtr("<name>"),
						MaximumCapacity:          to.Int64Ptr(3),
						ReservationCapacity:      to.Int64Ptr(1),
						TotalApplicationCapacity: to.Int64Ptr(5),
					}},
				RemoveApplicationCapacity: to.BoolPtr(false),
				TypeVersion:               to.StringPtr("<type-version>"),
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
