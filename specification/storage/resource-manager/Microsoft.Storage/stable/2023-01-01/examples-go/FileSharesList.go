package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0baf811c3c76c87b3c127d098519bd97141222dd/specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/FileSharesList.json
func ExampleFileSharesClient_NewListPager_listShares() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSharesClient().NewListPager("res9290", "sto1590", &armstorage.FileSharesClientListOptions{Maxpagesize: nil,
		Filter: nil,
		Expand: nil,
	})
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
		// page.FileShareItems = armstorage.FileShareItems{
		// 	Value: []*armstorage.FileShareItem{
		// 		{
		// 			Name: to.Ptr("share1644"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/fileServices/default/shares/share1644"),
		// 			Etag: to.Ptr("\"0x8D589847D51C7DE\""),
		// 			Properties: &armstorage.FileShareProperties{
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-14T08:20:47.000Z"); return t}()),
		// 				ShareQuota: to.Ptr[int32](1024),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("share4052"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/fileServices/default/shares/share4052"),
		// 			Etag: to.Ptr("\"0x8D589847DAB5AF9\""),
		// 			Properties: &armstorage.FileShareProperties{
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-14T08:20:47.000Z"); return t}()),
		// 				ShareQuota: to.Ptr[int32](1024),
		// 			},
		// 	}},
		// }
	}
}
