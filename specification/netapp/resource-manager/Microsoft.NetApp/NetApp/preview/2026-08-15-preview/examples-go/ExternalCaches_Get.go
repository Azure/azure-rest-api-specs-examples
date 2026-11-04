package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v11"
)

// Generated from example definition: 2026-08-15-preview/ExternalCaches_Get.json
func ExampleExternalCachesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExternalCachesClient().Get(ctx, "contosoNetAppRg", "contosoExternalCluster", "contosoExternalSvm", "contosoCache_AB12", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetapp.ExternalCachesClientGetResponse{
	// 	ExternalCache: armnetapp.ExternalCache{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.NetApp/externalClusters/contosoExternalCluster/externalSvms/contosoExternalSvm/externalCaches/contosoCache_AB12"),
	// 		Name: to.Ptr("contosoCache_AB12"),
	// 		Type: to.Ptr("Microsoft.NetApp/externalClusters/externalSvms/externalCaches"),
	// 		Properties: &armnetapp.ExternalCacheProperties{
	// 			ExternalCacheName: to.Ptr("contosoCache"),
	// 			OriginVolumeResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.NetApp/netAppAccounts/contosoAccount/capacityPools/contosoPool/volumes/contosoVolume"),
	// 			OriginVolumeName: to.Ptr("contosoVolume"),
	// 			OriginSVMName: to.Ptr("anfSvm"),
	// 			SizeInBytes: to.Ptr[int64](107374182400),
	// 			State: to.Ptr(armnetapp.ExternalCacheStateOnline),
	// 			ConnectionStatus: to.Ptr(armnetapp.ExternalCacheConnectionStatusConnected),
	// 			CreateCommand: to.Ptr("volume flexcache create -vserver contosoExternalSvm -volume contosoCache -origin-vserver anfSvm -origin-volume contosoVolume"),
	// 		},
	// 	},
	// }
}
