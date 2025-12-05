package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/SnapshotPolicies_ListVolumes.json
func ExampleSnapshotPoliciesClient_ListVolumes() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSnapshotPoliciesClient().ListVolumes(ctx, "myRG", "account1", "snapshotPolicyName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetapp.SnapshotPoliciesClientListVolumesResponse{
	// 	SnapshotPolicyVolumeList: &armnetapp.SnapshotPolicyVolumeList{
	// 		Value: []*armnetapp.Volume{
	// 			{
	// 				Name: to.Ptr("account1/pool1/volume1"),
	// 				Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes"),
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1"),
	// 				Location: to.Ptr("eastus"),
	// 				Properties: &armnetapp.VolumeProperties{
	// 					CreationToken: to.Ptr("my-unique-file-path"),
	// 					FileSystemID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 					ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
	// 					SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 					ThroughputMibps: to.Ptr[float32](128),
	// 					UsageThreshold: to.Ptr[int64](107374182400),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
