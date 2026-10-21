package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/ChangeDataCaptureConfigurations_Update.json
func ExampleChangeDataCaptureConfigurationsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewChangeDataCaptureConfigurationsClient().BeginUpdate(ctx, "example-resource-group", "example-cluster", "sales-database", armhorizondb.ChangeDataCaptureConfigurationForPatchUpdate{
		Properties: &armhorizondb.ChangeDataCaptureConfigurationPropertiesForPatchUpdate{
			DesiredState: to.Ptr(armhorizondb.ChangeDataCaptureDesiredStateDisabled),
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
	// res = armhorizondb.ChangeDataCaptureConfigurationsClientUpdateResponse{
	// 	ChangeDataCaptureConfiguration: armhorizondb.ChangeDataCaptureConfiguration{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/changeDataCaptureConfigurations/sales-database"),
	// 		Name: to.Ptr("sales-database"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/clusters/changeDataCaptureConfigurations"),
	// 		ETag: to.Ptr("\"00000000-0000-0000-0000-000000000031\""),
	// 		Properties: &armhorizondb.ChangeDataCaptureConfigurationProperties{
	// 			DatabaseName: to.Ptr("sales"),
	// 			UsageProfile: to.Ptr(armhorizondb.ChangeDataCaptureUsageProfileMirroring),
	// 			DestinationResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/changeDataCaptureDestinations/fabric-mirroring"),
	// 			LogicalStreamID: to.Ptr("aaaaaaaa-1111-2222-3333-bbbbbbbbbbbb"),
	// 			DesiredState: to.Ptr(armhorizondb.ChangeDataCaptureDesiredStateDisabled),
	// 			EffectiveState: to.Ptr(armhorizondb.ChangeDataCaptureEffectiveStateDisabled),
	// 			ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
	// 			Extension: &armhorizondb.ChangeDataCaptureExtensionStatus{
	// 				Name: to.Ptr("horizondb_change_data_capture"),
	// 				Version: to.Ptr("1.0"),
	// 				State: to.Ptr(armhorizondb.ChangeDataCaptureExtensionStateReady),
	// 			},
	// 			Replication: &armhorizondb.ChangeDataCaptureReplicationStatus{
	// 				State: to.Ptr(armhorizondb.ChangeDataCaptureReplicationStateNotConnected),
	// 			},
	// 			Topology: &armhorizondb.ChangeDataCaptureTopologyBindingStatus{
	// 				State: to.Ptr(armhorizondb.ChangeDataCaptureTopologyBindingStateActive),
	// 				Generation: to.Ptr[int64](1),
	// 				ActiveClusterResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster"),
	// 			},
	// 			Status: &armhorizondb.ChangeDataCaptureStatus{
	// 				Code: to.Ptr("Disabled"),
	// 				Message: to.Ptr("The database is not publishing changes to Microsoft Fabric."),
	// 			},
	// 		},
	// 	},
	// }
}
