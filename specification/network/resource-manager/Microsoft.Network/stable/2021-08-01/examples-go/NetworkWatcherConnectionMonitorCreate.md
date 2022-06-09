```go
package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/NetworkWatcherConnectionMonitorCreate.json
func ExampleConnectionMonitorsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewConnectionMonitorsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"nw1",
		"cm1",
		armnetwork.ConnectionMonitor{
			Location: to.Ptr("eastus"),
			Properties: &armnetwork.ConnectionMonitorParameters{
				Endpoints: []*armnetwork.ConnectionMonitorEndpoint{
					{
						Name:       to.Ptr("source"),
						ResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/ct1"),
					},
					{
						Name:    to.Ptr("destination"),
						Address: to.Ptr("bing.com"),
					}},
				TestConfigurations: []*armnetwork.ConnectionMonitorTestConfiguration{
					{
						Name: to.Ptr("tcp"),
						TCPConfiguration: &armnetwork.ConnectionMonitorTCPConfiguration{
							Port: to.Ptr[int32](80),
						},
						TestFrequencySec: to.Ptr[int32](60),
						Protocol:         to.Ptr(armnetwork.ConnectionMonitorTestConfigurationProtocolTCP),
					}},
				TestGroups: []*armnetwork.ConnectionMonitorTestGroup{
					{
						Name: to.Ptr("tg"),
						Destinations: []*string{
							to.Ptr("destination")},
						Sources: []*string{
							to.Ptr("source")},
						TestConfigurations: []*string{
							to.Ptr("tcp")},
					}},
			},
		},
		&armnetwork.ConnectionMonitorsClientBeginCreateOrUpdateOptions{Migrate: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv1.0.0/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.
