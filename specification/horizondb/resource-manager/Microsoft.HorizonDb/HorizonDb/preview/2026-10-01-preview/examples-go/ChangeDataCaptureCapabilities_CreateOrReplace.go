package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/ChangeDataCaptureCapabilities_CreateOrReplace.json
func ExampleChangeDataCaptureCapabilitiesClient_BeginCreateOrReplace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewChangeDataCaptureCapabilitiesClient().BeginCreateOrReplace(ctx, "example-resource-group", "example-cluster", armhorizondb.ChangeDataCaptureCapability{
		Properties: &armhorizondb.ChangeDataCaptureCapabilityProperties{
			DesiredState:         to.Ptr(armhorizondb.ChangeDataCaptureDesiredStateEnabled),
			AutoUpgradeExtension: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhorizondb.ChangeDataCaptureCapabilitiesClientCreateOrReplaceResponse{
	// 	ChangeDataCaptureCapability: armhorizondb.ChangeDataCaptureCapability{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/changeDataCapture/default"),
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/clusters/changeDataCapture"),
	// 		ETag: to.Ptr("\"00000000-0000-0000-0000-000000000001\""),
	// 		Properties: &armhorizondb.ChangeDataCaptureCapabilityProperties{
	// 			DesiredState: to.Ptr(armhorizondb.ChangeDataCaptureDesiredStateEnabled),
	// 			EffectiveState: to.Ptr(armhorizondb.ChangeDataCaptureEffectiveStateEnabled),
	// 			ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
	// 			AutoUpgradeExtension: to.Ptr(true),
	// 			ReplicationArchitecture: to.Ptr(armhorizondb.ChangeDataCaptureReplicationArchitectureHybrid),
	// 			Extension: &armhorizondb.ChangeDataCaptureExtensionStatus{
	// 				Name: to.Ptr("horizondb_change_data_capture"),
	// 				Version: to.Ptr("1.0"),
	// 				State: to.Ptr(armhorizondb.ChangeDataCaptureExtensionStateReady),
	// 			},
	// 			Prerequisites: &armhorizondb.ChangeDataCapturePrerequisiteStatus{
	// 				EngineConfigurationState: to.Ptr(armhorizondb.ChangeDataCaptureEngineConfigurationStateReady),
	// 				ReadWriteEndpointState: to.Ptr(armhorizondb.ChangeDataCaptureReadWriteEndpointStateReady),
	// 			},
	// 			DestinationCount: to.Ptr[int32](0),
	// 			ConfigurationCount: to.Ptr[int32](0),
	// 			Status: &armhorizondb.ChangeDataCaptureStatus{
	// 				Code: to.Ptr("Ready"),
	// 				Message: to.Ptr("Change data capture is ready."),
	// 			},
	// 		},
	// 	},
	// }
}
