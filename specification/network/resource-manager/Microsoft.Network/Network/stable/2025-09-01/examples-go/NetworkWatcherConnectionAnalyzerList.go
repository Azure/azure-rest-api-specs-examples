package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/NetworkWatcherConnectionAnalyzerList.json
func ExampleWatchersClient_NewConnectionAnalyzersListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("7f4a1d92-3b6e-4c8f-9a25-e1b8c3d7f024", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWatchersClient().NewConnectionAnalyzersListPager("connectionAnalyzerRG", "nw1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armnetwork.WatchersClientConnectionAnalyzersListResponse{
		// 	ConnectionAnalyzerListResult: armnetwork.ConnectionAnalyzerListResult{
		// 		Value: []*armnetwork.ConnectionAnalyzer{
		// 			{
		// 				Name: to.Ptr("ca1"),
		// 				ID: to.Ptr("/subscriptions/7f4a1d92-3b6e-4c8f-9a25-e1b8c3d7f024/resourceGroups/connectionAnalyzerRG/providers/Microsoft.Network/networkWatchers/nw1/connectionAnalyzers/ca1"),
		// 				Type: to.Ptr("Microsoft.Network/networkWatchers/connectionAnalyzers"),
		// 				Etag: to.Ptr("W/\"e7497f26-5f09-4559-900b-fe98f3dedb6f\""),
		// 				SystemData: &armnetwork.SystemData{
		// 					CreatedBy: to.Ptr("user1@contoso.com"),
		// 					CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2025, time.September, 1, 0, 0, 0, 0, time.UTC)),
		// 					LastModifiedBy: to.Ptr("user1@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2025, time.September, 1, 0, 0, 0, 0, time.UTC)),
		// 				},
		// 				Properties: &armnetwork.ConnectionAnalyzerProperties{
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					Status: to.Ptr(armnetwork.ConnectionAnalyzerStatusRunning),
		// 					DiagnosticOperations: []*armnetwork.DiagnosticOperation{
		// 						to.Ptr(armnetwork.DiagnosticOperationConnectivityCheck),
		// 					},
		// 					Source: &armnetwork.ConnectionAnalyzerEndpoint{
		// 						Type: to.Ptr(armnetwork.ConnectionAnalyzerEndpointTypeVM),
		// 						ResourceID: to.Ptr("/subscriptions/7f4a1d92-3b6e-4c8f-9a25-e1b8c3d7f024/resourceGroups/connectionAnalyzerRG/providers/Microsoft.Compute/virtualMachines/ct1"),
		// 					},
		// 					Destination: &armnetwork.ConnectionAnalyzerEndpoint{
		// 						Address: to.Ptr("www.bing.com"),
		// 						Type: to.Ptr(armnetwork.ConnectionAnalyzerEndpointTypeExternalAddress),
		// 					},
		// 					ProtocolSettings: &armnetwork.ProtocolSettings{
		// 						Protocol: to.Ptr(armnetwork.ProtocolTCP),
		// 					},
		// 					DiagnosticOperationsSettings: &armnetwork.DiagnosticOperationsSettings{
		// 						ConnectivityCheckSettings: &armnetwork.ConnectivityCheckSettings{
		// 							GeneratePath: to.Ptr(true),
		// 							PreferredIPVersion: to.Ptr(armnetwork.PreferredIPVersionIPv4),
		// 						},
		// 					},
		// 					ExpiryInDays: to.Ptr[int32](30),
		// 					OutputSettings: &armnetwork.OutputSettings{
		// 						StorageAccountSettings: &armnetwork.StorageAccountSettings{
		// 							StorageAccountID: to.Ptr("/subscriptions/7f4a1d92-3b6e-4c8f-9a25-e1b8c3d7f024/resourceGroups/connectionAnalyzerRG/providers/Microsoft.Storage/storageAccounts/sa1"),
		// 							Path: to.Ptr("connectionanalyzer/results2"),
		// 						},
		// 					},
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 			},
		// 			{
		// 				Name: to.Ptr("ca2"),
		// 				ID: to.Ptr("/subscriptions/7f4a1d92-3b6e-4c8f-9a25-e1b8c3d7f024/resourceGroups/connectionAnalyzerRG/providers/Microsoft.Network/networkWatchers/nw1/connectionAnalyzers/ca2"),
		// 				Type: to.Ptr("Microsoft.Network/networkWatchers/connectionAnalyzers"),
		// 				Etag: to.Ptr("W/\"e7497f26-5f09-4559-900b-fe98f3dedb6l\""),
		// 				SystemData: &armnetwork.SystemData{
		// 					CreatedBy: to.Ptr("user1@contoso.com"),
		// 					CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2025, time.September, 1, 0, 0, 0, 0, time.UTC)),
		// 					LastModifiedBy: to.Ptr("user1@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2025, time.September, 1, 0, 0, 0, 0, time.UTC)),
		// 				},
		// 				Properties: &armnetwork.ConnectionAnalyzerProperties{
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					Status: to.Ptr(armnetwork.ConnectionAnalyzerStatusRunning),
		// 					DiagnosticOperations: []*armnetwork.DiagnosticOperation{
		// 						to.Ptr(armnetwork.DiagnosticOperationConnectivityCheck),
		// 					},
		// 					Source: &armnetwork.ConnectionAnalyzerEndpoint{
		// 						Type: to.Ptr(armnetwork.ConnectionAnalyzerEndpointTypeVM),
		// 						ResourceID: to.Ptr("/subscriptions/7f4a1d92-3b6e-4c8f-9a25-e1b8c3d7f024/resourceGroups/connectionAnalyzerRG/providers/Microsoft.Compute/virtualMachines/ct2"),
		// 					},
		// 					Destination: &armnetwork.ConnectionAnalyzerEndpoint{
		// 						Address: to.Ptr("www.bing.com"),
		// 						Type: to.Ptr(armnetwork.ConnectionAnalyzerEndpointTypeExternalAddress),
		// 					},
		// 					ProtocolSettings: &armnetwork.ProtocolSettings{
		// 						Protocol: to.Ptr(armnetwork.ProtocolTCP),
		// 					},
		// 					DiagnosticOperationsSettings: &armnetwork.DiagnosticOperationsSettings{
		// 						ConnectivityCheckSettings: &armnetwork.ConnectivityCheckSettings{
		// 							GeneratePath: to.Ptr(true),
		// 							PreferredIPVersion: to.Ptr(armnetwork.PreferredIPVersionIPv4),
		// 						},
		// 					},
		// 					ExpiryInDays: to.Ptr[int32](30),
		// 					OutputSettings: &armnetwork.OutputSettings{
		// 						StorageAccountSettings: &armnetwork.StorageAccountSettings{
		// 							StorageAccountID: to.Ptr("/subscriptions/7f4a1d92-3b6e-4c8f-9a25-e1b8c3d7f024/resourceGroups/connectionAnalyzerRG/providers/Microsoft.Storage/storageAccounts/sa1"),
		// 							Path: to.Ptr("connectionanalyzer/results"),
		// 						},
		// 					},
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
