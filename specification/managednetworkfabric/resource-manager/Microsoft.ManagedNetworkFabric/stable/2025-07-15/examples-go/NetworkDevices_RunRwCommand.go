package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkDevices_RunRwCommand.json
func ExampleNetworkDevicesClient_BeginRunRwCommand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkDevicesClient().BeginRunRwCommand(ctx, "example-rg", "example-device", armmanagednetworkfabric.DeviceRwCommand{
		Command:    to.Ptr("show running-config"),
		CommandURL: to.Ptr("https://example.blob.core.windows.net/commands/config.txt"),
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
	// res = armmanagednetworkfabric.NetworkDevicesClientRunRwCommandResponse{
	// 	NetworkDeviceRunRwCommandResponse: &armmanagednetworkfabric.NetworkDeviceRunRwCommandResponse{
	// 		Status: to.Ptr("Failed"),
	// 		Error: &armmanagednetworkfabric.ErrorDetail{
	// 			Code: to.Ptr("CommandExecutionPending"),
	// 			Message: to.Ptr("Command execution requires administrative approval"),
	// 			Target: to.Ptr("NetworkDevice"),
	// 			Details: []*armmanagednetworkfabric.ErrorDetail{
	// 				{
	// 					Code: to.Ptr("PendingAdministrativeUpdate"),
	// 					Message: to.Ptr("Command execution discovered configuration changes requiring administrative update"),
	// 					Target: to.Ptr("ConfigurationState"),
	// 				},
	// 			},
	// 			AdditionalInfo: []*armmanagednetworkfabric.ErrorAdditionalInfo{
	// 				{
	// 					Type: to.Ptr("CommandExecutionDetails"),
	// 					Info: map[string]any{
	// 						"requiredAction": "ApplyAdministrativeUpdate",
	// 						"impactedConfiguration": "running-config",
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Properties: &armmanagednetworkfabric.NetworkDeviceRwCommandResponseProperties{
	// 			OutputURL: to.Ptr("https://example.blob.core.windows.net/outputs/command-output.txt"),
	// 		},
	// 	},
	// }
}
