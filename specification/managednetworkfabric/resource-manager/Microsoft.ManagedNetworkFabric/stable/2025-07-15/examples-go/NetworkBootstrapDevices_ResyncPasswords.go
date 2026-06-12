package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkBootstrapDevices_ResyncPasswords.json
func ExampleNetworkBootstrapDevicesClient_BeginResyncPasswords() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("1234ABCD-0A1B-1234-5678-123456ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkBootstrapDevicesClient().BeginResyncPasswords(ctx, "example-rg", "example-device", nil)
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
	// res = armmanagednetworkfabric.NetworkBootstrapDevicesClientResyncPasswordsResponse{
	// 	NetworkBootstrapDeviceResyncPasswordsResponse: &armmanagednetworkfabric.NetworkBootstrapDeviceResyncPasswordsResponse{
	// 		ID: to.Ptr("/Subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkBootstrapDevices/example-device"),
	// 		Name: to.Ptr("example-device"),
	// 		Status: to.Ptr("InProgress"),
	// 		PercentComplete: to.Ptr[float64](70),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-30T12:05:46.266Z"); return t}()),
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-30T12:05:46.266Z"); return t}()),
	// 		Operations: []*armmanagednetworkfabric.OperationStatusResult{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkBootstrapDevices/example-device-1"),
	// 				Status: to.Ptr("Succeeded"),
	// 				ID: to.Ptr("/Subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkBootstrapDevices/example-device-1"),
	// 				Name: to.Ptr("example-device-1"),
	// 				PercentComplete: to.Ptr[float64](88),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-30T12:04:42.876Z"); return t}()),
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-30T12:04:42.876Z"); return t}()),
	// 				Operations: []*armmanagednetworkfabric.OperationStatusResult{
	// 					{
	// 						ID: to.Ptr("/Subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkBootstrapDevices/example-device-1"),
	// 						ResourceID: to.Ptr("/Subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkBootstrapDevices/example-device-1"),
	// 						Name: to.Ptr("example-device-1"),
	// 						Status: to.Ptr("Failed"),
	// 						PercentComplete: to.Ptr[float64](69),
	// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-30T12:04:42.871Z"); return t}()),
	// 						EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-30T12:04:42.871Z"); return t}()),
	// 						Operations: []*armmanagednetworkfabric.OperationStatusResult{
	// 						},
	// 						Error: &armmanagednetworkfabric.ErrorDetail{
	// 							Code: to.Ptr("500"),
	// 							Message: to.Ptr("An unexpected error occurred"),
	// 							Target: to.Ptr("example-target"),
	// 							Details: []*armmanagednetworkfabric.ErrorDetail{
	// 							},
	// 							AdditionalInfo: []*armmanagednetworkfabric.ErrorAdditionalInfo{
	// 								{
	// 									Type: to.Ptr("example-additional-info-type"),
	// 									Info: map[string]any{
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				Error: &armmanagednetworkfabric.ErrorDetail{
	// 					Code: to.Ptr("404"),
	// 					Message: to.Ptr("Resource not found"),
	// 					Target: to.Ptr("example-target"),
	// 					Details: []*armmanagednetworkfabric.ErrorDetail{
	// 					},
	// 					AdditionalInfo: []*armmanagednetworkfabric.ErrorAdditionalInfo{
	// 						{
	// 							Type: to.Ptr("example-additional-info-type"),
	// 							Info: map[string]any{
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Error: &armmanagednetworkfabric.ErrorDetail{
	// 			Code: to.Ptr("400"),
	// 			Message: to.Ptr("Invalid request"),
	// 			Target: to.Ptr("example-target"),
	// 			Details: []*armmanagednetworkfabric.ErrorDetail{
	// 			},
	// 			AdditionalInfo: []*armmanagednetworkfabric.ErrorAdditionalInfo{
	// 				{
	// 					Type: to.Ptr("example-type"),
	// 					Info: map[string]any{
	// 					},
	// 				},
	// 			},
	// 		},
	// 		ResourceID: to.Ptr("/Subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkBootstrapDevices/example-device"),
	// 	},
	// }
}
