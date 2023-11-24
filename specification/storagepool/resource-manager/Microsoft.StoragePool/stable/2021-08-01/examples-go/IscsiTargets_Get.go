package armstoragepool_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagepool/armstoragepool"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/IscsiTargets_Get.json
func ExampleIscsiTargetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragepool.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIscsiTargetsClient().Get(ctx, "myResourceGroup", "myDiskPool", "myIscsiTarget", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IscsiTarget = armstoragepool.IscsiTarget{
	// 	Name: to.Ptr("myIscsiTarget"),
	// 	Type: to.Ptr("Microsoft.StoragePool/diskPools/iscsiTargets"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.StoragePool/diskPools/myDiskPool/iscsiTargets/myIscsiTarget"),
	// 	Properties: &armstoragepool.IscsiTargetProperties{
	// 		ACLMode: to.Ptr(armstoragepool.IscsiTargetACLModeStatic),
	// 		Endpoints: []*string{
	// 			to.Ptr("10.0.0.1:3260")},
	// 			Luns: []*armstoragepool.IscsiLun{
	// 				{
	// 					Name: to.Ptr("lun0"),
	// 					Lun: to.Ptr[int32](0),
	// 					ManagedDiskAzureResourceID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_1"),
	// 			}},
	// 			Port: to.Ptr[int32](3260),
	// 			ProvisioningState: to.Ptr(armstoragepool.ProvisioningStatesSucceeded),
	// 			Sessions: []*string{
	// 				to.Ptr("iqn.2005-03.org.iscsi:client")},
	// 				StaticACLs: []*armstoragepool.ACL{
	// 					{
	// 						InitiatorIqn: to.Ptr("iqn.2005-03.org.iscsi:client"),
	// 						MappedLuns: []*string{
	// 							to.Ptr("lun0")},
	// 					}},
	// 					Status: to.Ptr(armstoragepool.OperationalStatusHealthy),
	// 					TargetIqn: to.Ptr("iqn.2005-03.org.iscsi:server1"),
	// 				},
	// 				SystemData: &armstoragepool.SystemMetadata{
	// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-24T06:53:57.000Z"); return t}()),
	// 					CreatedBy: to.Ptr("alias"),
	// 					CreatedByType: to.Ptr(armstoragepool.CreatedByTypeUser),
	// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-24T06:53:57.000Z"); return t}()),
	// 					LastModifiedBy: to.Ptr("alias"),
	// 					LastModifiedByType: to.Ptr(armstoragepool.CreatedByTypeUser),
	// 				},
	// 			}
}
