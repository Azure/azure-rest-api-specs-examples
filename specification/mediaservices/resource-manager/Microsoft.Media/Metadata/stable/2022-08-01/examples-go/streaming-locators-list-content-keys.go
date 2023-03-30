package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-list-content-keys.json
func ExampleStreamingLocatorsClient_ListContentKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStreamingLocatorsClient().ListContentKeys(ctx, "contoso", "contosomedia", "secureStreamingLocator", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ListContentKeysResponse = armmediaservices.ListContentKeysResponse{
	// 	ContentKeys: []*armmediaservices.StreamingLocatorContentKey{
	// 		{
	// 			Type: to.Ptr(armmediaservices.StreamingLocatorContentKeyTypeEnvelopeEncryption),
	// 			ID: to.Ptr("9259eb06-eeee-4f77-987f-48f4ea5c649f"),
	// 			LabelReferenceInStreamingPolicy: to.Ptr("aesDefaultKey"),
	// 			PolicyName: to.Ptr("PolicyWithMultipleOptions"),
	// 			Tracks: []*armmediaservices.TrackSelection{
	// 			},
	// 			Value: to.Ptr("QpiqeQROdN5xamnfUF2Wdw=="),
	// 		},
	// 		{
	// 			Type: to.Ptr(armmediaservices.StreamingLocatorContentKeyTypeCommonEncryptionCenc),
	// 			ID: to.Ptr("06bfeff1-2bb6-4f58-af27-a2767f058bca"),
	// 			LabelReferenceInStreamingPolicy: to.Ptr("cencDefaultKey"),
	// 			PolicyName: to.Ptr("PolicyWithMultipleOptions"),
	// 			Tracks: []*armmediaservices.TrackSelection{
	// 			},
	// 			Value: to.Ptr("ZjgWhNnqnqcov/h+wrYusw=="),
	// 		},
	// 		{
	// 			Type: to.Ptr(armmediaservices.StreamingLocatorContentKeyTypeCommonEncryptionCbcs),
	// 			ID: to.Ptr("799e78a0-ed6f-4179-9222-ed4ec4223cec"),
	// 			LabelReferenceInStreamingPolicy: to.Ptr("cbcsDefaultKey"),
	// 			PolicyName: to.Ptr("PolicyWithMultipleOptions"),
	// 			Tracks: []*armmediaservices.TrackSelection{
	// 			},
	// 			Value: to.Ptr("FjZ3n3yRcVxRFftdYFbe9g=="),
	// 	}},
	// }
}
