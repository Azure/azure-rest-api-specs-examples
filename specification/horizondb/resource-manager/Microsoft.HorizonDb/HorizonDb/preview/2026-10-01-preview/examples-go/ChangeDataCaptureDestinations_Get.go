package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/ChangeDataCaptureDestinations_Get.json
func ExampleChangeDataCaptureDestinationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewChangeDataCaptureDestinationsClient().Get(ctx, "example-resource-group", "example-cluster", "fabric-mirroring", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhorizondb.ChangeDataCaptureDestinationsClientGetResponse{
	// 	ChangeDataCaptureDestination: armhorizondb.ChangeDataCaptureDestination{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/changeDataCaptureDestinations/fabric-mirroring"),
	// 		Name: to.Ptr("fabric-mirroring"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/clusters/changeDataCaptureDestinations"),
	// 		ETag: to.Ptr("\"00000000-0000-0000-0000-000000000010\""),
	// 		Properties: &armhorizondb.ChangeDataCaptureDestinationProperties{
	// 			Destination: &armhorizondb.FabricDestinationDetails{
	// 				DestinationType: to.Ptr(armhorizondb.ChangeDataCaptureDestinationTypeFabric),
	// 				Protocol: to.Ptr(armhorizondb.ChangeDataCaptureProtocolNative),
	// 				WorkspaceID: to.Ptr("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
	// 			},
	// 			DesiredState: to.Ptr(armhorizondb.ChangeDataCaptureDesiredStateEnabled),
	// 			EffectiveState: to.Ptr(armhorizondb.ChangeDataCaptureDestinationEffectiveStateEnabled),
	// 			ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
	// 			ActiveIdentityAssignmentResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/changeDataCaptureDestinations/fabric-mirroring/identityAssignments/fabric-writer"),
	// 			ConfigurationCount: to.Ptr[int32](1),
	// 			Status: &armhorizondb.ChangeDataCaptureStatus{
	// 				Code: to.Ptr("Ready"),
	// 				Message: to.Ptr("The Microsoft Fabric destination is ready."),
	// 			},
	// 		},
	// 	},
	// }
}
