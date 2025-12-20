package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/Datastores_List.json
func ExampleDatastoresClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatastoresClient().NewListPager("group1", "cloud1", "cluster1", nil)
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
		// page = armavs.DatastoresClientListResponse{
		// 	DatastoreList: armavs.DatastoreList{
		// 		Value: []*armavs.Datastore{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/datastores/datastore1"),
		// 				Name: to.Ptr("datastore1"),
		// 				Properties: &armavs.DatastoreProperties{
		// 					NetAppVolume: &armavs.NetAppVolume{
		// 						ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/ResourceGroup1/providers/Microsoft.NetApp/netAppAccounts/NetAppAccount1/capacityPools/CapacityPool1/volumes/NFSVol1"),
		// 					},
		// 					ProvisioningState: to.Ptr(armavs.DatastoreProvisioningStateSucceeded),
		// 					Status: to.Ptr(armavs.DatastoreStatusAccessible),
		// 				},
		// 				Type: to.Ptr("Microsoft.AVS/privateClouds/clusters/datastores"),
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/datastores/datastore2"),
		// 				Name: to.Ptr("datastore2"),
		// 				Properties: &armavs.DatastoreProperties{
		// 					DiskPoolVolume: &armavs.DiskPoolVolume{
		// 						TargetID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/ResourceGroup1/providers/Microsoft.StoragePool/diskPools/DiskPool1/targets/Target1"),
		// 						LunName: to.Ptr("lun0"),
		// 						MountOption: to.Ptr(armavs.MountOptionEnumMOUNT),
		// 						Path: to.Ptr("/vmfs/devices/disks/naa.6001405f75f6bdf7f6f49db8b4b21723"),
		// 					},
		// 					ProvisioningState: to.Ptr(armavs.DatastoreProvisioningStateSucceeded),
		// 					Status: to.Ptr(armavs.DatastoreStatusAccessible),
		// 				},
		// 				Type: to.Ptr("Microsoft.AVS/privateClouds/clusters/datastores"),
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/datastores/datastore2"),
		// 				Name: to.Ptr("datastore2"),
		// 				Properties: &armavs.DatastoreProperties{
		// 					ElasticSanVolume: &armavs.ElasticSanVolume{
		// 						TargetID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/ResourceGroup1/providers/Microsoft.ElasticSan/elasticSans/ElasticSan1/volumeGroups/VolumeGroup1/volumes/Volume1"),
		// 					},
		// 					ProvisioningState: to.Ptr(armavs.DatastoreProvisioningStateSucceeded),
		// 					Status: to.Ptr(armavs.DatastoreStatusAccessible),
		// 				},
		// 				Type: to.Ptr("Microsoft.AVS/privateClouds/clusters/datastores"),
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/datastores/datastore3"),
		// 				Name: to.Ptr("datastore3"),
		// 				Properties: &armavs.DatastoreProperties{
		// 					PureStorageVolume: &armavs.PureStorageVolume{
		// 						StoragePoolID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/ResourceGroup1/providers/PureStorage.Block/storagePools/storagePool1"),
		// 						SizeGb: to.Ptr[int32](64),
		// 					},
		// 					ProvisioningState: to.Ptr(armavs.DatastoreProvisioningStateSucceeded),
		// 					Status: to.Ptr(armavs.DatastoreStatusAccessible),
		// 				},
		// 				Type: to.Ptr("Microsoft.AVS/privateClouds/clusters/datastores"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
