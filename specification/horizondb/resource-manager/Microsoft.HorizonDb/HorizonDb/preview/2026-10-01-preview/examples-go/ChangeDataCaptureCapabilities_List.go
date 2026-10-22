package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/ChangeDataCaptureCapabilities_List.json
func ExampleChangeDataCaptureCapabilitiesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewChangeDataCaptureCapabilitiesClient().NewListPager("example-resource-group", "example-cluster", nil)
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
		// page = armhorizondb.ChangeDataCaptureCapabilitiesClientListResponse{
		// 	ChangeDataCaptureCapabilityListResult: armhorizondb.ChangeDataCaptureCapabilityListResult{
		// 		Value: []*armhorizondb.ChangeDataCaptureCapability{
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/changeDataCapture/default"),
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters/changeDataCapture"),
		// 				ETag: to.Ptr("\"00000000-0000-0000-0000-000000000001\""),
		// 				Properties: &armhorizondb.ChangeDataCaptureCapabilityProperties{
		// 					DesiredState: to.Ptr(armhorizondb.ChangeDataCaptureDesiredStateEnabled),
		// 					EffectiveState: to.Ptr(armhorizondb.ChangeDataCaptureEffectiveStateEnabled),
		// 					ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
		// 					AutoUpgradeExtension: to.Ptr(true),
		// 					ReplicationArchitecture: to.Ptr(armhorizondb.ChangeDataCaptureReplicationArchitectureHybrid),
		// 					Extension: &armhorizondb.ChangeDataCaptureExtensionStatus{
		// 						Name: to.Ptr("horizondb_change_data_capture"),
		// 						Version: to.Ptr("1.0"),
		// 						State: to.Ptr(armhorizondb.ChangeDataCaptureExtensionStateReady),
		// 					},
		// 					Prerequisites: &armhorizondb.ChangeDataCapturePrerequisiteStatus{
		// 						EngineConfigurationState: to.Ptr(armhorizondb.ChangeDataCaptureEngineConfigurationStateReady),
		// 						ReadWriteEndpointState: to.Ptr(armhorizondb.ChangeDataCaptureReadWriteEndpointStateReady),
		// 					},
		// 					DestinationCount: to.Ptr[int32](1),
		// 					ConfigurationCount: to.Ptr[int32](1),
		// 					Status: &armhorizondb.ChangeDataCaptureStatus{
		// 						Code: to.Ptr("Ready"),
		// 						Message: to.Ptr("Change data capture is ready."),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
