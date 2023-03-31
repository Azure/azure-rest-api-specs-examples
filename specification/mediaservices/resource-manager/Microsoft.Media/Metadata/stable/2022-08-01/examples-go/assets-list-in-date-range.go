package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/assets-list-in-date-range.json
func ExampleAssetsClient_NewListPager_listAssetCreatedInADateRange() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssetsClient().NewListPager("contoso", "contosomedia", &armmediaservices.AssetsClientListOptions{Filter: to.Ptr("properties/created gt 2012-06-01 and properties/created lt 2013-07-01"),
		Top:     nil,
		Orderby: to.Ptr("properties/created"),
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
		// page.AssetCollection = armmediaservices.AssetCollection{
		// 	Value: []*armmediaservices.Asset{
		// 		{
		// 			Name: to.Ptr("ClimbingMountRainier"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/assets"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/assets/ClimbingMountRainier"),
		// 			Properties: &armmediaservices.AssetProperties{
		// 				Description: to.Ptr("A documentary showing the ascent of Mount Rainier"),
		// 				AlternateID: to.Ptr("CLIMB00001"),
		// 				AssetID: to.Ptr("f8eea45c-b814-44c2-9c42-a5174ebdee4c"),
		// 				Container: to.Ptr("asset-f8eea45c-b814-44c2-9c42-a5174ebdee4c"),
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2012-11-01T00:00:00Z"); return t}()),
		// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2012-11-01T00:00:00Z"); return t}()),
		// 				StorageEncryptionFormat: to.Ptr(armmediaservices.AssetStorageEncryptionFormatNone),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ClimbingMountAdams"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/assets"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/assets/ClimbingMountAdams"),
		// 			Properties: &armmediaservices.AssetProperties{
		// 				Description: to.Ptr("A documentary showing the ascent of Mount Adams"),
		// 				AlternateID: to.Ptr("CLIMB00002"),
		// 				AssetID: to.Ptr("1b648c1a-2268-461d-a1da-742bde23db40"),
		// 				Container: to.Ptr("asset-1b648c1a-2268-461d-a1da-742bde23db40"),
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2013-02-01T00:00:00Z"); return t}()),
		// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-01T00:00:00Z"); return t}()),
		// 				StorageEncryptionFormat: to.Ptr(armmediaservices.AssetStorageEncryptionFormatNone),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ClimbingMountSaintHelens"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/assets"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/assets/ClimbingMountSaintHelens"),
		// 			Properties: &armmediaservices.AssetProperties{
		// 				Description: to.Ptr("A documentary showing the ascent of Saint Helens"),
		// 				AlternateID: to.Ptr("CLIMB00005"),
		// 				AssetID: to.Ptr("14d58c40-ec1f-446c-b041-f5cff949bd1d"),
		// 				Container: to.Ptr("asset-14d58c40-ec1f-446c-b041-f5cff949bd1d"),
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2013-03-01T00:00:00Z"); return t}()),
		// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2013-04-01T00:00:00Z"); return t}()),
		// 				StorageEncryptionFormat: to.Ptr(armmediaservices.AssetStorageEncryptionFormatMediaStorageClientEncryption),
		// 			},
		// 	}},
		// }
	}
}
