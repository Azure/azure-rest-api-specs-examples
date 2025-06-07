package armpurestorageblock_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purestorageblock/armpurestorageblock"
)

// Generated from example definition: 2024-11-01-preview/AvsVms_Update_MaximumSet_Gen.json
func ExampleAvsVMsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurestorageblock.NewClientFactory("BC47D6CC-AA80-4374-86F8-19D94EC70666", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAvsVMsClient().BeginUpdate(ctx, "rgpurestorage", "storagePoolname", "cbdec-ddbb", armpurestorageblock.AvsVMUpdate{
		Properties: &armpurestorageblock.AvsVMUpdateProperties{
			SoftDeletion: &armpurestorageblock.SoftDeletion{
				Destroyed: to.Ptr(true),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpurestorageblock.AvsVMsClientUpdateResponse{
	// 	AvsVM: &armpurestorageblock.AvsVM{
	// 		Properties: &armpurestorageblock.AvsVMProperties{
	// 			StoragePoolInternalID: to.Ptr("zpchlbmnmgympdfqeuchp"),
	// 			StoragePoolResourceID: to.Ptr("refrmhvh"),
	// 			DisplayName: to.Ptr("zcpusnyyozsxptnvia"),
	// 			CreatedTimestamp: to.Ptr("eeczpavimq"),
	// 			SoftDeletion: &armpurestorageblock.SoftDeletion{
	// 				Destroyed: to.Ptr(true),
	// 				EradicationTimestamp: to.Ptr("kaxjtehra"),
	// 			},
	// 			VolumeContainerType: to.Ptr(armpurestorageblock.VolumeContainerTypeAVS),
	// 			Avs: &armpurestorageblock.AvsVMDetails{
	// 				VMID: to.Ptr("ljsdq"),
	// 				VMName: to.Ptr("hskuenhnxpscuqikeohkyjfebgzapx"),
	// 				VMType: to.Ptr(armpurestorageblock.VMTypeVVol),
	// 				AvsVMInternalID: to.Ptr("wnamcozqs"),
	// 			},
	// 			Space: &armpurestorageblock.Space{
	// 				TotalUsed: to.Ptr[int64](28),
	// 				Unique: to.Ptr[int64](4),
	// 				Snapshots: to.Ptr[int64](5),
	// 				Shared: to.Ptr[int64](9),
	// 			},
	// 			ProvisioningState: to.Ptr(armpurestorageblock.ResourceProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
	// 		Name: to.Ptr("ifvdmtawvkzz"),
	// 		Type: to.Ptr("qsfuizcaje"),
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
