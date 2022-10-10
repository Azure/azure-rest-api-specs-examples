package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkWatcherConnectionMonitorV2Create.json
func ExampleConnectionMonitorsClient_BeginCreateOrUpdate_createConnectionMonitorV2() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewConnectionMonitorsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "nw1", "cm1", armnetwork.ConnectionMonitor{
		Properties: &armnetwork.ConnectionMonitorParameters{
			Endpoints: []*armnetwork.ConnectionMonitorEndpoint{
				{
					Name:       to.Ptr("vm1"),
					ResourceID: to.Ptr("/subscriptions/96e68903-0a56-4819-9987-8d08ad6a1f99/resourceGroups/NwRgIrinaCentralUSEUAP/providers/Microsoft.Compute/virtualMachines/vm1"),
				},
				{
					Name: to.Ptr("CanaryWorkspaceVamshi"),
					Filter: &armnetwork.ConnectionMonitorEndpointFilter{
						Type: to.Ptr(armnetwork.ConnectionMonitorEndpointFilterTypeInclude),
						Items: []*armnetwork.ConnectionMonitorEndpointFilterItem{
							{
								Type:    to.Ptr(armnetwork.ConnectionMonitorEndpointFilterItemTypeAgentAddress),
								Address: to.Ptr("npmuser"),
							}},
					},
					ResourceID: to.Ptr("/subscriptions/96e68903-0a56-4819-9987-8d08ad6a1f99/resourceGroups/vasamudrRG/providers/Microsoft.OperationalInsights/workspaces/vasamudrWorkspace"),
				},
				{
					Name:    to.Ptr("bing"),
					Address: to.Ptr("bing.com"),
				},
				{
					Name:    to.Ptr("google"),
					Address: to.Ptr("google.com"),
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
						to.Ptr("CanaryWorkspaceVamshi")},
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
	// TODO: use response item
	_ = res
}
