package armpurestorageblock_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purestorageblock/armpurestorageblock"
)

// Generated from example definition: 2024-11-01-preview/AvsVmVolumes_Get_MaximumSet_Gen.json
func ExampleAvsVMVolumesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurestorageblock.NewClientFactory("BC47D6CC-AA80-4374-86F8-19D94EC70666", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAvsVMVolumesClient().Get(ctx, "rgpurestorage", "storagePoolname", "cbdec-ddbb", "cbdec-ddbb", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpurestorageblock.AvsVMVolumesClientGetResponse{
	// 	AvsVMVolume: &armpurestorageblock.AvsVMVolume{
	// 		Properties: &armpurestorageblock.VolumeProperties{
	// 			StoragePoolInternalID: to.Ptr("xkcbzbtfhftbnowayannq"),
	// 			StoragePoolResourceID: to.Ptr("mrnioblkjhwry"),
	// 			VolumeInternalID: to.Ptr("ivdmuth"),
	// 			DisplayName: to.Ptr("mcdetkfvt"),
	// 			Space: &armpurestorageblock.Space{
	// 				TotalUsed: to.Ptr[int64](28),
	// 				Unique: to.Ptr[int64](4),
	// 				Snapshots: to.Ptr[int64](5),
	// 				Shared: to.Ptr[int64](9),
	// 			},
	// 			SoftDeletion: &armpurestorageblock.SoftDeletion{
	// 				Destroyed: to.Ptr(true),
	// 				EradicationTimestamp: to.Ptr("kaxjtehra"),
	// 			},
	// 			CreatedTimestamp: to.Ptr("pwhsjhrcszrhsmzc"),
	// 			ProvisionedSize: to.Ptr[int64](14),
	// 			VolumeType: to.Ptr(armpurestorageblock.VolumeTypeAVS),
	// 			Avs: &armpurestorageblock.AvsDiskDetails{
	// 				DiskID: to.Ptr("egehzvidnjgevc"),
	// 				DiskName: to.Ptr("gqpqpkxjyj"),
	// 				Folder: to.Ptr("oxdbbh"),
	// 				AvsVMInternalID: to.Ptr("xoiliugbcvkhdl"),
	// 				AvsVMResourceID: to.Ptr("jblyybyfhtikdhwx"),
	// 				AvsVMName: to.Ptr("gaeydwwisfuonbuwtnkdk"),
	// 				AvsStorageContainerResourceID: to.Ptr("svwilhkpx"),
	// 			},
	// 			ProvisioningState: to.Ptr(armpurestorageblock.ResourceProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
	// 		Name: to.Ptr("hnglbdwyocdmggwii"),
	// 		Type: to.Ptr("jprfxakifyliuzougfctzpfhozq"),
	// 		SystemData: &armpurestorageblock.SystemData{
	// 			CreatedBy: to.Ptr("ruoitchmuomrbscg"),
	// 			CreatedByType: to.Ptr(armpurestorageblock.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-04T05:29:25.341Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("thfyhokbrldzmghuylqbwpbublj"),
	// 			LastModifiedByType: to.Ptr(armpurestorageblock.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-04T05:29:25.345Z"); return t}()),
	// 		},
	// 	},
	// }
}
