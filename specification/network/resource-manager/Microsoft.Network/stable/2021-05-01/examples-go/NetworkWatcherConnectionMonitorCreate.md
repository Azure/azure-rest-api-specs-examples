Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.3.1/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherConnectionMonitorCreate.json
func ExampleConnectionMonitorsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewConnectionMonitorsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<network-watcher-name>",
		"<connection-monitor-name>",
		armnetwork.ConnectionMonitor{
			Location: to.StringPtr("<location>"),
			Properties: &armnetwork.ConnectionMonitorParameters{
				Endpoints: []*armnetwork.ConnectionMonitorEndpoint{
					{
						Name:       to.StringPtr("<name>"),
						ResourceID: to.StringPtr("<resource-id>"),
					},
					{
						Name:    to.StringPtr("<name>"),
						Address: to.StringPtr("<address>"),
					}},
				TestConfigurations: []*armnetwork.ConnectionMonitorTestConfiguration{
					{
						Name: to.StringPtr("<name>"),
						TCPConfiguration: &armnetwork.ConnectionMonitorTCPConfiguration{
							Port: to.Int32Ptr(80),
						},
						TestFrequencySec: to.Int32Ptr(60),
						Protocol:         armnetwork.ConnectionMonitorTestConfigurationProtocol("Tcp").ToPtr(),
					}},
				TestGroups: []*armnetwork.ConnectionMonitorTestGroup{
					{
						Name: to.StringPtr("<name>"),
						Destinations: []*string{
							to.StringPtr("destination")},
						Sources: []*string{
							to.StringPtr("source")},
						TestConfigurations: []*string{
							to.StringPtr("tcp")},
					}},
			},
		},
		&armnetwork.ConnectionMonitorsClientBeginCreateOrUpdateOptions{Migrate: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ConnectionMonitorsClientCreateOrUpdateResult)
}
```
