package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2025-08-01/FileSharesList_ProvisionedV2.json
func ExampleFileSharesClient_NewListPager_listSharesProvisionedV2() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSharesClient().NewListPager("res9290", "sto1590", nil)
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
		// 		NextLink: to.Ptr("https://sto1590endpoint/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/fileServices/default/shares?api-version=2022-09-01&$maxpagesize=2&$skipToken=/sto1590/share5103"),
		// 		Value: []*armstorage.FileShareItem{
		// 			{
		// 				Name: to.Ptr("share1644"),
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
		// 				Etag: to.Ptr("\"0x8D589847D51C7DE\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/fileServices/default/shares/share1644"),
		// 				Properties: &armstorage.FileShareProperties{
		// 					IncludedBurstIops: to.Ptr[int32](15000),
		// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-25T01:50:50.0000000Z"); return t}()),
		// 					MaxBurstCreditsForIops: to.Ptr[int64](36000000),
		// 					NextAllowedProvisionedBandwidthDowngradeTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Fri, 25 Oct 2024 01:48:09 GMT"); return t}()),
		// 					NextAllowedProvisionedIopsDowngradeTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Fri, 25 Oct 2024 01:48:09 GMT"); return t}()),
		// 					NextAllowedQuotaDowngradeTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Sat, 26 Oct 2024 01:50:50 GMT"); return t}()),
		// 					ProvisionedBandwidthMibps: to.Ptr[int32](200),
		// 					ProvisionedIops: to.Ptr[int32](5000),
		// 					ShareQuota: to.Ptr[int32](100),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("share4052"),
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
		// 				Etag: to.Ptr("\"0x8D589847DAB5AF9\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/fileServices/default/shares/share4052"),
		// 				Properties: &armstorage.FileShareProperties{
		// 					IncludedBurstIops: to.Ptr[int32](15000),
		// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-25T01:50:50.0000000Z"); return t}()),
		// 					MaxBurstCreditsForIops: to.Ptr[int64](36000000),
		// 					NextAllowedProvisionedBandwidthDowngradeTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Fri, 25 Oct 2024 01:48:09 GMT"); return t}()),
		// 					NextAllowedProvisionedIopsDowngradeTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Fri, 25 Oct 2024 01:48:09 GMT"); return t}()),
		// 					NextAllowedQuotaDowngradeTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Sat, 26 Oct 2024 01:50:50 GMT"); return t}()),
		// 					ProvisionedBandwidthMibps: to.Ptr[int32](200),
		// 					ProvisionedIops: to.Ptr[int32](5000),
		// 					ShareQuota: to.Ptr[int32](100),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
