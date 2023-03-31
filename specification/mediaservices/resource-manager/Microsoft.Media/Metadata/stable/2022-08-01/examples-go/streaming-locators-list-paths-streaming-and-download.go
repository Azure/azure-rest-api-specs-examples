package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-list-paths-streaming-and-download.json
func ExampleStreamingLocatorsClient_ListPaths_listPathsWhichHasStreamingPathsAndDownloadPaths() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStreamingLocatorsClient().ListPaths(ctx, "contoso", "contosomedia", "clearStreamingLocator", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ListPathsResponse = armmediaservices.ListPathsResponse{
	// 	DownloadPaths: []*string{
	// 		to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/textTrack.vtt"),
	// 		to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/video1.mp4"),
	// 		to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/video2.mp4"),
	// 		to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/video3.mp4")},
	// 		StreamingPaths: []*armmediaservices.StreamingPath{
	// 			{
	// 				EncryptionScheme: to.Ptr(armmediaservices.EncryptionSchemeNoEncryption),
	// 				Paths: []*string{
	// 					to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/videoManifest1.ism/manifest(format=m3u8-aapl)"),
	// 					to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/videoManifest2.ism/manifest(format=m3u8-aapl)"),
	// 					to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/videoManifest3.ism/manifest(format=m3u8-aapl)")},
	// 					StreamingProtocol: to.Ptr(armmediaservices.StreamingPolicyStreamingProtocolHls),
	// 				},
	// 				{
	// 					EncryptionScheme: to.Ptr(armmediaservices.EncryptionSchemeNoEncryption),
	// 					Paths: []*string{
	// 						to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/videoManifest1.ism/manifest(format=mpd-time-csf)"),
	// 						to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/videoManifest2.ism/manifest(format=mpd-time-csf)"),
	// 						to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/videoManifest3.ism/manifest(format=mpd-time-csf)")},
	// 						StreamingProtocol: to.Ptr(armmediaservices.StreamingPolicyStreamingProtocolDash),
	// 					},
	// 					{
	// 						EncryptionScheme: to.Ptr(armmediaservices.EncryptionSchemeNoEncryption),
	// 						Paths: []*string{
	// 							to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/videoManifest1.ism/manifest"),
	// 							to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/videoManifest2.ism/manifest"),
	// 							to.Ptr("/262a87b6-b538-4657-bac1-b6897924471d/videoManifest3.ism/manifest")},
	// 							StreamingProtocol: to.Ptr(armmediaservices.StreamingPolicyStreamingProtocolSmoothStreaming),
	// 					}},
	// 				}
}
