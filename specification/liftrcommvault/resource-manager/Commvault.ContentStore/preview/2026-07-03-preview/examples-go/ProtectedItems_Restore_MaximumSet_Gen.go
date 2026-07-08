package armcommvaultcontentstore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commvaultcontentstore/armcommvaultcontentstore"
)

// Generated from example definition: 2026-07-03-preview/ProtectedItems_Restore_MaximumSet_Gen.json
func ExampleProtectedItemsClient_Restore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommvaultcontentstore.NewClientFactory("65D4E6D7-7063-4C4B-BAC5-13C45474009E", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProtectedItemsClient().Restore(ctx, "rgcommvault", "sample-cloudAccountName", "sample-protectionGroupName", "sample-protectedItemName", armcommvaultcontentstore.RestoreProtectionItemRequest{
		InPlaceRestore: to.Ptr(true),
		RestoreType:    to.Ptr(armcommvaultcontentstore.RestoreTypeNone),
		ToTime:         to.Ptr("kgueys"),
		VMDestinationInfo: &armcommvaultcontentstore.VMDestinationInfo{
			VMInfoList: []*armcommvaultcontentstore.VMInfo{
				{
					SourceVMGUID:          to.Ptr("40000000-aaaa-4bbb-8ccc-000000000000"),
					StorageAccountID:      to.Ptr("pldvo"),
					PowerOnVMAfterRestore: to.Ptr(true),
					Name:                  to.Ptr("ctmwbnzhxqdhshl"),
					ResourceGroup:         to.Ptr("pdfpesq"),
					Region:                to.Ptr("ywwecvwsosvatgmulaxfnja"),
					NetworkID:             to.Ptr("amtetghiqnbjuiurekikacvymxjcv"),
					SubnetID:              to.Ptr("klskhhutusnzycgxaq"),
					AttachAndSwapOsDisk:   to.Ptr(true),
					TargetVMGUID:          to.Ptr("40000000-aaaa-4bbb-8ccc-000000000001"),
					Vmtags: []*armcommvaultcontentstore.VMTag{
						{
							Name:  to.Ptr("dzabcqlzgqphznrpzhrqbsszgdzjrh"),
							Value: to.Ptr("zjtpbngipinqlwajjrf"),
						},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcommvaultcontentstore.ProtectedItemsClientRestoreResponse{
	// 	RestoreProtectionItemResponse: armcommvaultcontentstore.RestoreProtectionItemResponse{
	// 		TaskID: to.Ptr[int32](12),
	// 		JobIDs: []*string{
	// 			to.Ptr("isfz"),
	// 		},
	// 	},
	// }
}
