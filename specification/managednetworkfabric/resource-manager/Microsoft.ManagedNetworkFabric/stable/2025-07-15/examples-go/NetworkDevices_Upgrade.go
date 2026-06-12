package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkDevices_Upgrade.json
func ExampleNetworkDevicesClient_BeginUpgrade() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("1234ABCD-0A1B-1234-5678-123456ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkDevicesClient().BeginUpgrade(ctx, "rgmanagednetworkfabric", "example-device", armmanagednetworkfabric.NetworkDeviceUpgradeRequest{
		Version:           to.Ptr("1.0.0"),
		RwDeviceConfigURL: to.Ptr("https://microsoft.com/a"),
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
	// res = armmanagednetworkfabric.NetworkDevicesClientUpgradeResponse{
	// 	NetworkDeviceUpgradeResponse: &armmanagednetworkfabric.NetworkDeviceUpgradeResponse{
	// 		Status: to.Ptr("Failed"),
	// 		PercentComplete: to.Ptr[float64](30),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-28T04:16:32.382Z"); return t}()),
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-28T04:16:32.382Z"); return t}()),
	// 		Error: &armmanagednetworkfabric.ErrorDetail{
	// 			Code: to.Ptr("InsufficientResources"),
	// 			Message: to.Ptr("The system has insufficient resources to complete the operation."),
	// 		},
	// 	},
	// }
}
