package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/NetworkWatcherFlowLogCreate.json
func ExampleFlowLogsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFlowLogsClient().BeginCreateOrUpdate(ctx, "rg1", "nw1", "fl", armnetwork.FlowLog{
		Location: to.Ptr("centraluseuap"),
		Properties: &armnetwork.FlowLogPropertiesFormat{
			Format: &armnetwork.FlowLogFormatParameters{
				Type:    to.Ptr(armnetwork.FlowLogFormatTypeJSON),
				Version: to.Ptr[int32](1),
			},
			Enabled:          to.Ptr(true),
			StorageID:        to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/nwtest1mgvbfmqsigdxe"),
			TargetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/desmondcentral-nsg"),
		},
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
	// res.FlowLog = armnetwork.FlowLog{
	// 	Name: to.Ptr("Microsoft.Networkdesmond-rgdesmondcentral-nsg"),
	// 	Type: to.Ptr("Microsoft.Network/networkWatchers/FlowLogs"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkWatchers/nw/FlowLogs/fl"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
	// 	Properties: &armnetwork.FlowLogPropertiesFormat{
	// 		Format: &armnetwork.FlowLogFormatParameters{
	// 			Type: to.Ptr(armnetwork.FlowLogFormatTypeJSON),
	// 			Version: to.Ptr[int32](1),
	// 		},
	// 		Enabled: to.Ptr(true),
	// 		FlowAnalyticsConfiguration: &armnetwork.TrafficAnalyticsProperties{
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RetentionPolicy: &armnetwork.RetentionPolicyParameters{
	// 			Days: to.Ptr[int32](0),
	// 			Enabled: to.Ptr(false),
	// 		},
	// 		StorageID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/nwtest1mgvbfmqsigdxe"),
	// 		TargetResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		TargetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/desmondcentral-nsg"),
	// 	},
	// }
}
