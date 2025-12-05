package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/Volumes_ExtraLargeVolumes_Get.json
func ExampleVolumesClient_Get_volumesExtralargeVolumeGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("D633CC2E-722B-4AE1-B636-BBD9E4C60ED9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVolumesClient().Get(ctx, "myRG", "account1", "pool1", "volume1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetapp.VolumesClientGetResponse{
	// 	Volume: &armnetapp.Volume{
	// 		ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1"),
	// 		Name: to.Ptr("account1/pool1/volume1"),
	// 		Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armnetapp.VolumeProperties{
	// 			FileSystemID: to.Ptr("9760acf5-4638-11e7-9bdb-020073ca7778"),
	// 			CoolAccess: to.Ptr(true),
	// 			CreationToken: to.Ptr("my-unique-file-path"),
	// 			IsLargeVolume: to.Ptr(true),
	// 			LargeVolumeType: to.Ptr(armnetapp.LargeVolumeType("ExtraLargeVolume7Dot2PiB")),
	// 			UsageThreshold: to.Ptr[int64](109951162777600),
	// 			ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			ActualThroughputMibps: to.Ptr[float32](128),
	// 			SubnetID: to.Ptr("/subscriptions/9760acf5-4638-11e7-9bdb-020073ca7778/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 			NetworkFeatures: to.Ptr(armnetapp.NetworkFeaturesBasic),
	// 			EffectiveNetworkFeatures: to.Ptr(armnetapp.NetworkFeaturesStandard),
	// 			NetworkSiblingSetID: to.Ptr("0f434a03-ce0b-4935-81af-d98652ffb1c4"),
	// 			StorageToNetworkProximity: to.Ptr(armnetapp.VolumeStorageToNetworkProximityT2),
	// 			CoolAccessTieringPolicy: to.Ptr(armnetapp.CoolAccessTieringPolicySnapshotOnly),
	// 		},
	// 	},
	// }
}
