package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/Volumes_List.json
func ExampleVolumesClient_NewListPager_volumesList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVolumesClient().NewListPager("myRG", "account1", "pool1", nil)
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
		// page = armnetapp.VolumesClientListResponse{
		// 	VolumeList: armnetapp.VolumeList{
		// 		Value: []*armnetapp.Volume{
		// 			{
		// 				Name: to.Ptr("account1/pool1/volume1"),
		// 				Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armnetapp.VolumeProperties{
		// 					CreationToken: to.Ptr("my-unique-file-path"),
		// 					FileSystemID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					NetworkFeatures: to.Ptr(armnetapp.NetworkFeaturesStandard),
		// 					NetworkSiblingSetID: to.Ptr("0f434a03-ce0b-4935-81af-d98652ffb1c4"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
		// 					StorageToNetworkProximity: to.Ptr(armnetapp.VolumeStorageToNetworkProximityT2),
		// 					SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
		// 					ThroughputMibps: to.Ptr[float32](128),
		// 					UsageThreshold: to.Ptr[int64](107374182400),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
