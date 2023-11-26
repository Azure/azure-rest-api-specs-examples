package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/assets-get-by-name.json
func ExampleAssetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssetsClient().Get(ctx, "contoso", "contosomedia", "ClimbingMountAdams", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Asset = armmediaservices.Asset{
	// 	Name: to.Ptr("ClimbingMountAdams"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/assets"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/assets/ClimbingMountAdams"),
	// 	Properties: &armmediaservices.AssetProperties{
	// 		Description: to.Ptr("A documentary showing the ascent of Mount Adams"),
	// 		AlternateID: to.Ptr("CLIMB00002"),
	// 		AssetID: to.Ptr("1b648c1a-2268-461d-a1da-742bde23db40"),
	// 		Container: to.Ptr("asset-1b648c1a-2268-461d-a1da-742bde23db40"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2013-02-01T00:00:00.000Z"); return t}()),
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-01T00:00:00.000Z"); return t}()),
	// 		StorageEncryptionFormat: to.Ptr(armmediaservices.AssetStorageEncryptionFormatNone),
	// 	},
	// }
}
