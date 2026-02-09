package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkWatcherConnectionMonitorCreate.json
func ExampleConnectionMonitorsClient_BeginCreateOrUpdate_createConnectionMonitorV1() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectionMonitorsClient().BeginCreateOrUpdate(ctx, "rg1", "nw1", "cm1", armnetwork.ConnectionMonitor{
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
	}, &armnetwork.ConnectionMonitorsClientBeginCreateOrUpdateOptions{Migrate: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectionMonitorResult = armnetwork.ConnectionMonitorResult{
	// 	Name: to.Ptr("cm1"),
	// 	Type: to.Ptr("Microsoft.Network/networkWatchers/connectionMonitors"),
	// 	Etag: to.Ptr("W/\"e7497f26-5f09-4559-900b-fe98f3dedb6f\""),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkWatchers/nw1/connectionMonitors/cm1"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armnetwork.ConnectionMonitorResultProperties{
	// 		Endpoints: []*armnetwork.ConnectionMonitorEndpoint{
	// 			{
	// 				Name: to.Ptr("source"),
	// 				ResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/ct1"),
	// 			},
	// 			{
	// 				Name: to.Ptr("destination"),
	// 				Address: to.Ptr("bing.com"),
	// 		}},
	// 		TestConfigurations: []*armnetwork.ConnectionMonitorTestConfiguration{
	// 			{
	// 				Name: to.Ptr("tcp"),
	// 				TCPConfiguration: &armnetwork.ConnectionMonitorTCPConfiguration{
	// 					Port: to.Ptr[int32](80),
	// 				},
	// 				TestFrequencySec: to.Ptr[int32](60),
	// 				Protocol: to.Ptr(armnetwork.ConnectionMonitorTestConfigurationProtocolTCP),
	// 		}},
	// 		TestGroups: []*armnetwork.ConnectionMonitorTestGroup{
	// 			{
	// 				Name: to.Ptr("tg"),
	// 				Destinations: []*string{
	// 					to.Ptr("destination")},
	// 					Sources: []*string{
	// 						to.Ptr("source")},
	// 						TestConfigurations: []*string{
	// 							to.Ptr("tcp")},
	// 					}},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				},
	// 			}
}
