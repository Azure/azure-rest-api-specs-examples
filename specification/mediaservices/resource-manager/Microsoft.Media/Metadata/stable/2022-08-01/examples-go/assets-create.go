package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/assets-create.json
func ExampleAssetsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssetsClient().CreateOrUpdate(ctx, "contoso", "contosomedia", "ClimbingMountLogan", armmediaservices.Asset{
		Properties: &armmediaservices.AssetProperties{
			Description:        to.Ptr("A documentary showing the ascent of Mount Logan"),
			StorageAccountName: to.Ptr("storage0"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Asset = armmediaservices.Asset{
	// 	Name: to.Ptr("ClimbingMountLogan"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/assets"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/assets/ClimbingMountLogan"),
	// 	Properties: &armmediaservices.AssetProperties{
	// 		Description: to.Ptr("A documentary showing the ascent of Mount Logan"),
	// 		AssetID: to.Ptr("09194258-36ba-4403-abb3-68780e6bc545"),
	// 		Container: to.Ptr("asset-09194258-36ba-4403-abb3-68780e6bc545"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:26.080Z"); return t}()),
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:26.107Z"); return t}()),
	// 		StorageAccountName: to.Ptr("storage0"),
	// 		StorageEncryptionFormat: to.Ptr(armmediaservices.AssetStorageEncryptionFormatNone),
	// 	},
	// }
}
