package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkDevices_RunRoCommand.json
func ExampleNetworkDevicesClient_BeginRunRoCommand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("1234ABCD-0A1B-1234-5678-123456ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkDevicesClient().BeginRunRoCommand(ctx, "example-rg", "example-device", armmanagednetworkfabric.DeviceRoCommand{
		Command: to.Ptr("show version"),
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
	// res = armmanagednetworkfabric.NetworkDevicesClientRunRoCommandResponse{
	// 	CommonPostActionResponseForDeviceROCommandsOperationStatusResult: &armmanagednetworkfabric.CommonPostActionResponseForDeviceROCommandsOperationStatusResult{
	// 		Error: &armmanagednetworkfabric.ErrorDetail{
	// 			Code: to.Ptr("zuwgsrizc"),
	// 			Message: to.Ptr("v"),
	// 			Target: to.Ptr("elwgfga"),
	// 			Details: []*armmanagednetworkfabric.ErrorDetail{
	// 			},
	// 			AdditionalInfo: []*armmanagednetworkfabric.ErrorAdditionalInfo{
	// 				{
	// 					Type: to.Ptr("modrphiaxpggwxenqcitxvbeevku"),
	// 					Info: map[string]any{
	// 					},
	// 				},
	// 			},
	// 		},
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-25T08:06:00.912Z"); return t}()),
	// 		ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/providers/Microsoft.ManagedNetworkFabric/locations/UKSOUTH/operationStatuses/b7c3d8f2-56e4-4a9b-8c7d-123456789ABC*1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF"),
	// 		Name: to.Ptr("b7c3d8f2-56e4-4a9b-8c7d-123456789ABC*1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF"),
	// 		Properties: &armmanagednetworkfabric.CommonPostActionResponseForDeviceROCommands{
	// 			ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 			OutputURL: to.Ptr("https://testsast1.blob.core.windows.net/container1"),
	// 			DeviceConfigurationPreview: to.Ptr("{\n  \"architecture\": \"arm64\",\n  \"bootupTimestamp\": 1745279999.1234567,\n  \"configMacAddress\": \"11:22:33:44:55:66\",\n  \"hardwareRevision\": \"RevB\",\n  \"hwMacAddress\": \"66:55:44:33:22:11\",\n  \"imageFormatVersion\": \"2.0\",\n  \"imageOptimization\": \"High\",\n  \"internalBuildId\": \"a123b456-c789-0123-d456-7890e123f456\",\n  \"internalVersion\": \"5.12.1FX-NX-12345678.5121FXNX\",\n  \"isIntlVersion\": true,\n  \"kernelVersion\": \"6.0.0-1.cm3\",\n  \"memFree\": 209715200,\n  \"memTotal\": 262144000,\n  \"mfgName\": \"Cisco\",\n  \"modelName\": \"vEOS\",\n  \"serialNumber\": \"1234567890ABCDEF1234567890ABCDEF\",\n  \"systemMacAddress\": \"aa:bb:cc:dd:ee:ff\",\n  \"uptime\": 123456.78901234567,\n  \"version\": \"5.12.1FX-NX-12345678.5121FXNX (release build)\"\n}"),
	// 		},
	// 		ResourceID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-device"),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-25T08:06:00.912Z"); return t}()),
	// 		Status: to.Ptr("Succeeded"),
	// 	},
	// }
}
