package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkDevices_RefreshConfiguration.json
func ExampleNetworkDevicesClient_BeginRefreshConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkDevicesClient().BeginRefreshConfiguration(ctx, "example-rg", "example-device", nil)
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
	// res = armmanagednetworkfabric.NetworkDevicesClientRefreshConfigurationResponse{
	// 	NetworkDeviceRefreshConfigurationResponse: &armmanagednetworkfabric.NetworkDeviceRefreshConfigurationResponse{
	// 		ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-device"),
	// 		Name: to.Ptr("example-device"),
	// 		Status: to.Ptr("Succeeded"),
	// 		PercentComplete: to.Ptr[float64](71),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-17T10:47:32.146Z"); return t}()),
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-17T10:47:32.146Z"); return t}()),
	// 		Operations: []*armmanagednetworkfabric.OperationStatusResult{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-device"),
	// 				Status: to.Ptr("Succeeded"),
	// 				ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-device"),
	// 				Name: to.Ptr("example-device"),
	// 				PercentComplete: to.Ptr[float64](17),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-17T10:47:32.140Z"); return t}()),
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-17T10:47:32.140Z"); return t}()),
	// 				Operations: []*armmanagednetworkfabric.OperationStatusResult{
	// 					{
	// 						ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-device"),
	// 						ResourceID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-device"),
	// 						Name: to.Ptr("example-device"),
	// 						Status: to.Ptr("InProgress"),
	// 						PercentComplete: to.Ptr[float64](84),
	// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-17T10:47:32.137Z"); return t}()),
	// 						EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-17T10:47:32.139Z"); return t}()),
	// 						Operations: []*armmanagednetworkfabric.OperationStatusResult{
	// 						},
	// 						Error: &armmanagednetworkfabric.ErrorDetail{
	// 							Code: to.Ptr("lvuwxyaprfiwvkejnadlwj"),
	// 							Message: to.Ptr("bzudjixb"),
	// 							Target: to.Ptr("bhdxiwfrzmnqbxfhkoxbtagirvhscr"),
	// 							Details: []*armmanagednetworkfabric.ErrorDetail{
	// 							},
	// 							AdditionalInfo: []*armmanagednetworkfabric.ErrorAdditionalInfo{
	// 								{
	// 									Type: to.Ptr("dgjnscthaeayttkopxtgztoguexdm"),
	// 									Info: map[string]any{
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				Error: &armmanagednetworkfabric.ErrorDetail{
	// 					Code: to.Ptr("lvuwxyaprfiwvkejnadlwj"),
	// 					Message: to.Ptr("bzudjixb"),
	// 					Target: to.Ptr("bhdxiwfrzmnqbxfhkoxbtagirvhscr"),
	// 					Details: []*armmanagednetworkfabric.ErrorDetail{
	// 					},
	// 					AdditionalInfo: []*armmanagednetworkfabric.ErrorAdditionalInfo{
	// 						{
	// 							Type: to.Ptr("dgjnscthaeayttkopxtgztoguexdm"),
	// 							Info: map[string]any{
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Error: &armmanagednetworkfabric.ErrorDetail{
	// 			Code: to.Ptr("uueoujlesavjup"),
	// 			Message: to.Ptr("ueiogogdennovfjlollfoqg"),
	// 			Target: to.Ptr("opnpremrj"),
	// 			Details: []*armmanagednetworkfabric.ErrorDetail{
	// 			},
	// 			AdditionalInfo: []*armmanagednetworkfabric.ErrorAdditionalInfo{
	// 				{
	// 					Type: to.Ptr("dbb"),
	// 					Info: map[string]any{
	// 					},
	// 				},
	// 			},
	// 		},
	// 		ResourceID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-device"),
	// 	},
	// }
}
