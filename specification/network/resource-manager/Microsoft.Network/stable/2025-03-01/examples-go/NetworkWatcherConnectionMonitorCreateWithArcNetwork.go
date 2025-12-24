package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkWatcherConnectionMonitorCreateWithArcNetwork.json
func ExampleConnectionMonitorsClient_BeginCreateOrUpdate_createConnectionMonitorWithArcNetwork() {
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
		Properties: &armnetwork.ConnectionMonitorParameters{
			Endpoints: []*armnetwork.ConnectionMonitorEndpoint{
				{
					Name:       to.Ptr("vm1"),
					Type:       to.Ptr(armnetwork.EndpointTypeAzureVM),
					ResourceID: to.Ptr("/subscriptions/9cece3e3-0f7d-47ca-af0e-9772773f90b7/resourceGroups/testRG/providers/Microsoft.Compute/virtualMachines/TESTVM"),
				},
				{
					Name:    to.Ptr("bing"),
					Type:    to.Ptr(armnetwork.EndpointTypeExternalAddress),
					Address: to.Ptr("bing.com"),
				},
				{
					Name:    to.Ptr("google"),
					Type:    to.Ptr(armnetwork.EndpointTypeExternalAddress),
					Address: to.Ptr("google.com"),
				},
				{
					Name: to.Ptr("ArcBasedNetwork"),
					Type: to.Ptr(armnetwork.EndpointTypeAzureArcNetwork),
					LocationDetails: &armnetwork.ConnectionMonitorEndpointLocationDetails{
						Region: to.Ptr("eastus"),
					},
					Scope: &armnetwork.ConnectionMonitorEndpointScope{
						Include: []*armnetwork.ConnectionMonitorEndpointScopeItem{
							{
								Address: to.Ptr("172.21.128.0/20"),
							}},
					},
					SubscriptionID: to.Ptr("9cece3e3-0f7d-47ca-af0e-9772773f90b7"),
				}},
			Outputs: []*armnetwork.ConnectionMonitorOutput{},
			TestConfigurations: []*armnetwork.ConnectionMonitorTestConfiguration{
				{
					Name: to.Ptr("testConfig1"),
					TCPConfiguration: &armnetwork.ConnectionMonitorTCPConfiguration{
						DisableTraceRoute: to.Ptr(false),
						Port:              to.Ptr[int32](80),
					},
					TestFrequencySec: to.Ptr[int32](60),
					Protocol:         to.Ptr(armnetwork.ConnectionMonitorTestConfigurationProtocolTCP),
				}},
			TestGroups: []*armnetwork.ConnectionMonitorTestGroup{
				{
					Name: to.Ptr("test1"),
					Destinations: []*string{
						to.Ptr("bing"),
						to.Ptr("google")},
					Disable: to.Ptr(false),
					Sources: []*string{
						to.Ptr("vm1"),
						to.Ptr("ArcBasedNetwork")},
					TestConfigurations: []*string{
						to.Ptr("testConfig1")},
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
	// 	Location: to.Ptr("centraluseuap"),
	// 	Properties: &armnetwork.ConnectionMonitorResultProperties{
	// 		Endpoints: []*armnetwork.ConnectionMonitorEndpoint{
	// 			{
	// 				Name: to.Ptr("vm1"),
	// 				ResourceID: to.Ptr("/subscriptions/9cece3e3-0f7d-47ca-af0e-9772773f90b7/resourceGroups/testRG/providers/Microsoft.Compute/virtualMachines/TESTVM"),
	// 			},
	// 			{
	// 				Name: to.Ptr("ArcBasedNetwork"),
	// 				Type: to.Ptr(armnetwork.EndpointTypeAzureArcNetwork),
	// 				LocationDetails: &armnetwork.ConnectionMonitorEndpointLocationDetails{
	// 					Region: to.Ptr("eastus"),
	// 				},
	// 				Scope: &armnetwork.ConnectionMonitorEndpointScope{
	// 					Include: []*armnetwork.ConnectionMonitorEndpointScopeItem{
	// 						{
	// 							Address: to.Ptr("172.21.128.0/20"),
	// 					}},
	// 				},
	// 				SubscriptionID: to.Ptr("9cece3e3-0f7d-47ca-af0e-9772773f90b7"),
	// 			},
	// 			{
	// 				Name: to.Ptr("bing"),
	// 				Address: to.Ptr("bing.com"),
	// 			},
	// 			{
	// 				Name: to.Ptr("google"),
	// 				Address: to.Ptr("google.com"),
	// 		}},
	// 		Outputs: []*armnetwork.ConnectionMonitorOutput{
	// 		},
	// 		TestConfigurations: []*armnetwork.ConnectionMonitorTestConfiguration{
	// 			{
	// 				Name: to.Ptr("testConfig1"),
	// 				TCPConfiguration: &armnetwork.ConnectionMonitorTCPConfiguration{
	// 					DisableTraceRoute: to.Ptr(false),
	// 					Port: to.Ptr[int32](80),
	// 				},
	// 				TestFrequencySec: to.Ptr[int32](60),
	// 				Protocol: to.Ptr(armnetwork.ConnectionMonitorTestConfigurationProtocolTCP),
	// 		}},
	// 		TestGroups: []*armnetwork.ConnectionMonitorTestGroup{
	// 			{
	// 				Name: to.Ptr("test1"),
	// 				Destinations: []*string{
	// 					to.Ptr("bing"),
	// 					to.Ptr("google")},
	// 					Disable: to.Ptr(false),
	// 					Sources: []*string{
	// 						to.Ptr("vm1"),
	// 						to.Ptr("ArcBasedNetwork")},
	// 						TestConfigurations: []*string{
	// 							to.Ptr("testConfig1")},
	// 					}},
	// 				},
	// 			}
}
