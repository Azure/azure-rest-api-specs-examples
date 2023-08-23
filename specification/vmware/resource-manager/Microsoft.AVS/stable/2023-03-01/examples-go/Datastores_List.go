package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a5e7ff51c8af3781e7f6dd3b82a3fc922e2f6f93/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/Datastores_List.json
func ExampleDatastoresClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.DatastoreList = armavs.DatastoreList{
		// 	Value: []*armavs.Datastore{
		// 		{
		// 			Name: to.Ptr("datastore1"),
		// 			Type: to.Ptr("Microsoft.AVS/privateClouds/clusters/datastores"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/datastores/datastore1"),
		// 			Properties: &armavs.DatastoreProperties{
		// 				NetAppVolume: &armavs.NetAppVolume{
		// 					ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/ResourceGroup1/providers/Microsoft.NetApp/netAppAccounts/NetAppAccount1/capacityPools/CapacityPool1/volumes/NFSVol1"),
		// 				},
		// 				ProvisioningState: to.Ptr(armavs.DatastoreProvisioningStateSucceeded),
		// 				Status: to.Ptr(armavs.DatastoreStatusAccessible),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("datastore2"),
		// 			Type: to.Ptr("Microsoft.AVS/privateClouds/clusters/datastores"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/datastores/datastore2"),
		// 			Properties: &armavs.DatastoreProperties{
		// 				DiskPoolVolume: &armavs.DiskPoolVolume{
		// 					Path: to.Ptr("/vmfs/devices/disks/naa.6001405f75f6bdf7f6f49db8b4b21723"),
		// 					LunName: to.Ptr("lun0"),
		// 					MountOption: to.Ptr(armavs.MountOptionEnumMOUNT),
		// 					TargetID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/ResourceGroup1/providers/Microsoft.StoragePool/diskPools/DiskPool1/targets/Target1"),
		// 				},
		// 				ProvisioningState: to.Ptr(armavs.DatastoreProvisioningStateSucceeded),
		// 				Status: to.Ptr(armavs.DatastoreStatusAccessible),
		// 			},
		// 	}},
		// }
	}
}
