package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherConnectionMonitorCreate.json
func ExampleConnectionMonitorsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewConnectionMonitorsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<network-watcher-name>",
		"<connection-monitor-name>",
		armnetwork.ConnectionMonitor{
			Location: to.Ptr("<location>"),
			Properties: &armnetwork.ConnectionMonitorParameters{
				Endpoints: []*armnetwork.ConnectionMonitorEndpoint{
					{
						Name:       to.Ptr("<name>"),
						ResourceID: to.Ptr("<resource-id>"),
					},
					{
						Name:    to.Ptr("<name>"),
						Address: to.Ptr("<address>"),
					}},
				TestConfigurations: []*armnetwork.ConnectionMonitorTestConfiguration{
					{
						Name: to.Ptr("<name>"),
						TCPConfiguration: &armnetwork.ConnectionMonitorTCPConfiguration{
							Port: to.Ptr[int32](80),
						},
						TestFrequencySec: to.Ptr[int32](60),
						Protocol:         to.Ptr(armnetwork.ConnectionMonitorTestConfigurationProtocolTCP),
					}},
				TestGroups: []*armnetwork.ConnectionMonitorTestGroup{
					{
						Name: to.Ptr("<name>"),
						Destinations: []*string{
							to.Ptr("destination")},
						Sources: []*string{
							to.Ptr("source")},
						TestConfigurations: []*string{
							to.Ptr("tcp")},
					}},
			},
		},
		&armnetwork.ConnectionMonitorsClientBeginCreateOrUpdateOptions{Migrate: nil,
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
