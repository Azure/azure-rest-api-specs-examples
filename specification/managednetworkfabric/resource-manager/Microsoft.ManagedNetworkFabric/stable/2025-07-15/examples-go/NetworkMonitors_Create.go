package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkMonitors_Create.json
func ExampleNetworkMonitorsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("1234ABCD-0A1B-1234-5678-123456ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkMonitorsClient().BeginCreate(ctx, "example-rg", "example-monitor", armmanagednetworkfabric.NetworkMonitor{
		Properties: &armmanagednetworkfabric.NetworkMonitorProperties{
			Annotation: to.Ptr("annotation"),
			BmpConfiguration: &armmanagednetworkfabric.BmpConfigurationProperties{
				StationConfigurationState: to.Ptr(armmanagednetworkfabric.StationConfigurationStateEnabled),
				ScopeResourceID:           to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric"),
				StationName:               to.Ptr("name"),
				StationIP:                 to.Ptr("10.0.0.1"),
				StationPort:               to.Ptr[int32](62695),
				StationConnectionMode:     to.Ptr(armmanagednetworkfabric.StationConnectionModeActive),
				StationConnectionProperties: &armmanagednetworkfabric.StationConnectionProperties{
					KeepaliveIdleTime: to.Ptr[int32](49),
					ProbeInterval:     to.Ptr[int32](3558),
					ProbeCount:        to.Ptr[int32](43),
				},
				StationNetwork: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/example-l3domain/internalNetworks/example-internalnetwork"),
				MonitoredNetworks: []*string{
					to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/example-l3domain"),
				},
				ExportPolicy: to.Ptr(armmanagednetworkfabric.BmpExportPolicyPrePolicy),
				ExportPolicyConfiguration: &armmanagednetworkfabric.BmpExportPolicyProperties{
					ExportPolicies: []*armmanagednetworkfabric.BmpExportPolicy{
						to.Ptr(armmanagednetworkfabric.BmpExportPolicyPrePolicy),
					},
				},
				MonitoredAddressFamilies: []*armmanagednetworkfabric.BmpMonitoredAddressFamily{
					to.Ptr(armmanagednetworkfabric.BmpMonitoredAddressFamilyIPv4Unicast),
				},
			},
			AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
		},
		Tags: map[string]*string{
			"key": to.Ptr("value"),
		},
		Location: to.Ptr("eastus"),
	}, nil)
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
	// res = armmanagednetworkfabric.NetworkMonitorsClientCreateResponse{
	// 	NetworkMonitor: &armmanagednetworkfabric.NetworkMonitor{
	// 		Properties: &armmanagednetworkfabric.NetworkMonitorProperties{
	// 			Annotation: to.Ptr("annotation"),
	// 			BmpConfiguration: &armmanagednetworkfabric.BmpConfigurationProperties{
	// 				ScopeResourceID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric"),
	// 				StationConfigurationState: to.Ptr(armmanagednetworkfabric.StationConfigurationStateEnabled),
	// 				StationName: to.Ptr("name"),
	// 				StationIP: to.Ptr("10.0.0.1"),
	// 				StationPort: to.Ptr[int32](62695),
	// 				StationConnectionMode: to.Ptr(armmanagednetworkfabric.StationConnectionModeActive),
	// 				StationConnectionProperties: &armmanagednetworkfabric.StationConnectionProperties{
	// 					KeepaliveIdleTime: to.Ptr[int32](49),
	// 					ProbeInterval: to.Ptr[int32](3558),
	// 					ProbeCount: to.Ptr[int32](43),
	// 				},
	// 				StationNetwork: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/example-l3domain/internalNetworks/example-internalnetwork"),
	// 				MonitoredNetworks: []*string{
	// 					to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/example-l3domain"),
	// 				},
	// 				ExportPolicy: to.Ptr(armmanagednetworkfabric.BmpExportPolicyPrePolicy),
	// 				ExportPolicyConfiguration: &armmanagednetworkfabric.BmpExportPolicyProperties{
	// 					ExportPolicies: []*armmanagednetworkfabric.BmpExportPolicy{
	// 						to.Ptr(armmanagednetworkfabric.BmpExportPolicyPrePolicy),
	// 					},
	// 				},
	// 				MonitoredAddressFamilies: []*armmanagednetworkfabric.BmpMonitoredAddressFamily{
	// 					to.Ptr(armmanagednetworkfabric.BmpMonitoredAddressFamilyIPv4Unicast),
	// 				},
	// 			},
	// 			ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
	// 			AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key": to.Ptr("value"),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkMonitors/example-monitor"),
	// 		Name: to.Ptr("example-monitor"),
	// 		Type: to.Ptr("microsoft.managednetworkfabric/networkmonitors"),
	// 		SystemData: &armmanagednetworkfabric.SystemData{
	// 			CreatedBy: to.Ptr("email@address.com"),
	// 			CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("UserId"),
	// 			LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
	// 		},
	// 	},
	// }
}
