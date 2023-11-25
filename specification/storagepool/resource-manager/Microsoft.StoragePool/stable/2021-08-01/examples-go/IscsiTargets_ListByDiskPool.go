package armstoragepool_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagepool/armstoragepool"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/IscsiTargets_ListByDiskPool.json
func ExampleIscsiTargetsClient_NewListByDiskPoolPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragepool.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIscsiTargetsClient().NewListByDiskPoolPager("myResourceGroup", "myDiskPool", nil)
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
		// page.IscsiTargetList = armstoragepool.IscsiTargetList{
		// 	Value: []*armstoragepool.IscsiTarget{
		// 		{
		// 			Name: to.Ptr("myIscsiTarget"),
		// 			Type: to.Ptr("Microsoft.StoragePool/diskPools/iscsiTargets"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.StoragePool/diskPools/myDiskPool/iscsiTargets/myIscsiTarget"),
		// 			ManagedBy: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.AVS/privateClouds/myPrivateCloud/clusters/Cluster-1/datastores/datastore1"),
		// 			ManagedByExtended: []*string{
		// 				to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.AVS/privateClouds/myPrivateCloud/clusters/Cluster-1/datastores/datastore1"),
		// 				to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.AVS/privateClouds/myPrivateCloud/clusters/Cluster-1/datastores/datastore2")},
		// 				Properties: &armstoragepool.IscsiTargetProperties{
		// 					ACLMode: to.Ptr(armstoragepool.IscsiTargetACLModeStatic),
		// 					Endpoints: []*string{
		// 						to.Ptr("10.0.0.1:3260")},
		// 						Luns: []*armstoragepool.IscsiLun{
		// 							{
		// 								Name: to.Ptr("lun0"),
		// 								Lun: to.Ptr[int32](3),
		// 								ManagedDiskAzureResourceID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_1"),
		// 						}},
		// 						Port: to.Ptr[int32](3260),
		// 						ProvisioningState: to.Ptr(armstoragepool.ProvisioningStatesSucceeded),
		// 						Sessions: []*string{
		// 							to.Ptr("iqn.2005-03.org.iscsi:client")},
		// 							StaticACLs: []*armstoragepool.ACL{
		// 								{
		// 									InitiatorIqn: to.Ptr("iqn.2005-03.org.iscsi:client"),
		// 									MappedLuns: []*string{
		// 										to.Ptr("lun0")},
		// 								}},
		// 								Status: to.Ptr(armstoragepool.OperationalStatusHealthy),
		// 								TargetIqn: to.Ptr("iqn.2005-03.org.iscsi:server1"),
		// 							},
		// 							SystemData: &armstoragepool.SystemMetadata{
		// 								CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-24T06:53:57.000Z"); return t}()),
		// 								CreatedBy: to.Ptr("alias"),
		// 								CreatedByType: to.Ptr(armstoragepool.CreatedByTypeUser),
		// 								LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-24T06:53:57.000Z"); return t}()),
		// 								LastModifiedBy: to.Ptr("alias"),
		// 								LastModifiedByType: to.Ptr(armstoragepool.CreatedByTypeUser),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("myIscsiTarget2"),
		// 							Type: to.Ptr("Microsoft.StoragePool/diskPools/iscsiTargets"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.StoragePool/diskPools/myDiskPool/iscsiTargets/myIscsiTarget"),
		// 							ManagedBy: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.AVS/privateClouds/myPrivateCloud/clusters/Cluster-1/datastores/datastore1"),
		// 							ManagedByExtended: []*string{
		// 								to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.AVS/privateClouds/myPrivateCloud/clusters/Cluster-1/datastores/datastore1"),
		// 								to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.AVS/privateClouds/myPrivateCloud/clusters/Cluster-1/datastores/datastore2")},
		// 								Properties: &armstoragepool.IscsiTargetProperties{
		// 									ACLMode: to.Ptr(armstoragepool.IscsiTargetACLModeDynamic),
		// 									Endpoints: []*string{
		// 										to.Ptr("10.0.0.1:3261")},
		// 										Luns: []*armstoragepool.IscsiLun{
		// 											{
		// 												Name: to.Ptr("lun0"),
		// 												Lun: to.Ptr[int32](2),
		// 												ManagedDiskAzureResourceID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_2"),
		// 										}},
		// 										Port: to.Ptr[int32](3261),
		// 										ProvisioningState: to.Ptr(armstoragepool.ProvisioningStatesSucceeded),
		// 										Sessions: []*string{
		// 											to.Ptr("iqn.2005-03.org.iscsi:client")},
		// 											StaticACLs: []*armstoragepool.ACL{
		// 											},
		// 											Status: to.Ptr(armstoragepool.OperationalStatusHealthy),
		// 											TargetIqn: to.Ptr("iqn.2005-03.org.iscsi:server2"),
		// 										},
		// 										SystemData: &armstoragepool.SystemMetadata{
		// 											CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-24T06:53:57.000Z"); return t}()),
		// 											CreatedBy: to.Ptr("alias"),
		// 											CreatedByType: to.Ptr(armstoragepool.CreatedByTypeUser),
		// 											LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-24T06:53:57.000Z"); return t}()),
		// 											LastModifiedBy: to.Ptr("alias"),
		// 											LastModifiedByType: to.Ptr(armstoragepool.CreatedByTypeUser),
		// 										},
		// 								}},
		// 							}
	}
}
