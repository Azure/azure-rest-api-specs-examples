package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-04-01/DeletedFileSharesList.json
func ExampleFileSharesClient_NewListPager_listDeletedShares() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSharesClient().NewListPager("res9290", "sto1590", &armstorage.FileSharesClientListOptions{
		Expand: to.Ptr("deleted")})
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
		// page = armstorage.FileSharesClientListResponse{
		// 	FileShareItems: armstorage.FileShareItems{
		// 		Value: []*armstorage.FileShareItem{
		// 			{
		// 				Name: to.Ptr("share1644"),
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
		// 				Etag: to.Ptr("\"0x8D589847D51C7DE\""),
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/fileServices/default/shares/share1644_1234567890"),
		// 				Properties: &armstorage.FileShareProperties{
		// 					Deleted: to.Ptr(true),
		// 					DeletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-14T08:20:47Z"); return t}()),
		// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-14T08:20:47Z"); return t}()),
		// 					RemainingRetentionDays: to.Ptr[int32](30),
		// 					ShareQuota: to.Ptr[int32](1024),
		// 					Version: to.Ptr("1234567890"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("share4052"),
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
		// 				Etag: to.Ptr("\"0x8D589847DAB5AF9\""),
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/fileServices/default/shares/share4052"),
		// 				Properties: &armstorage.FileShareProperties{
		// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-14T08:20:47Z"); return t}()),
		// 					ShareQuota: to.Ptr[int32](1024),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
