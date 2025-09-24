package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/260ed6a52537921f53a18ffaf4020e3b4d510367/specification/storage/resource-manager/Microsoft.Storage/stable/2025-01-01/examples/FileSharesList_ProvisionedV2.json
func ExampleFileSharesClient_NewListPager_listSharesProvisionedV2() {
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
		// 				IncludedBurstIops: to.Ptr[int32](15000),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-25T01:50:50.000Z"); return t}()),
		// 				MaxBurstCreditsForIops: to.Ptr[int64](36000000),
		// 				NextAllowedProvisionedBandwidthDowngradeTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2024-10-25T01:48:09.000Z"); return t}()),
		// 				NextAllowedProvisionedIopsDowngradeTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2024-10-25T01:48:09.000Z"); return t}()),
		// 				NextAllowedQuotaDowngradeTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2024-10-26T01:50:50.000Z"); return t}()),
		// 				ProvisionedBandwidthMibps: to.Ptr[int32](200),
		// 				ProvisionedIops: to.Ptr[int32](5000),
		// 				ShareQuota: to.Ptr[int32](100),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("share4052"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9290/providers/Microsoft.Storage/storageAccounts/sto1590/fileServices/default/shares/share4052"),
		// 			Etag: to.Ptr("\"0x8D589847DAB5AF9\""),
		// 			Properties: &armstorage.FileShareProperties{
		// 				IncludedBurstIops: to.Ptr[int32](15000),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-25T01:50:50.000Z"); return t}()),
		// 				MaxBurstCreditsForIops: to.Ptr[int64](36000000),
		// 				NextAllowedProvisionedBandwidthDowngradeTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2024-10-25T01:48:09.000Z"); return t}()),
		// 				NextAllowedProvisionedIopsDowngradeTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2024-10-25T01:48:09.000Z"); return t}()),
		// 				NextAllowedQuotaDowngradeTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2024-10-26T01:50:50.000Z"); return t}()),
		// 				ProvisionedBandwidthMibps: to.Ptr[int32](200),
		// 				ProvisionedIops: to.Ptr[int32](5000),
		// 				ShareQuota: to.Ptr[int32](100),
		// 			},
		// 	}},
		// }
	}
}
