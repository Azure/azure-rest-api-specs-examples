package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/assets-get-encryption-keys.json
func ExampleAssetsClient_GetEncryptionKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssetsClient().GetEncryptionKey(ctx, "contoso", "contosomedia", "ClimbingMountSaintHelens", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StorageEncryptedAssetDecryptionData = armmediaservices.StorageEncryptedAssetDecryptionData{
	// 	AssetFileEncryptionMetadata: []*armmediaservices.AssetFileEncryptionMetadata{
	// 		{
	// 			AssetFileID: to.Ptr("a9536fa7-bd5d-4f84-a137-d1290982fe95"),
	// 			AssetFileName: to.Ptr("AssetFile0"),
	// 			InitializationVector: to.Ptr("-988929866"),
	// 		},
	// 		{
	// 			AssetFileID: to.Ptr("f4060046-94ac-422d-824c-3f1d6aa3ecf2"),
	// 			AssetFileName: to.Ptr("AssetFile1"),
	// 			InitializationVector: to.Ptr("1604993689"),
	// 		},
	// 		{
	// 			AssetFileID: to.Ptr("485968d3-ddae-4b13-98e7-901201a9620b"),
	// 			AssetFileName: to.Ptr("AssetFile2"),
	// 			InitializationVector: to.Ptr("100082635"),
	// 	}},
	// 	Key: []byte("AAAAAAAAAAAAAAAAAAAAAA=="),
	// }
}
