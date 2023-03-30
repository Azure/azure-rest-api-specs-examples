package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-list-paths-streaming-only.json
func ExampleStreamingLocatorsClient_ListPaths_listPathsWhichHasStreamingPathsOnly() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStreamingLocatorsClient().ListPaths(ctx, "contoso", "contosomedia", "secureStreamingLocator", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ListPathsResponse = armmediaservices.ListPathsResponse{
	// 	DownloadPaths: []*string{
	// 	},
	// 	StreamingPaths: []*armmediaservices.StreamingPath{
	// 		{
	// 			EncryptionScheme: to.Ptr(armmediaservices.EncryptionSchemeEnvelopeEncryption),
	// 			Paths: []*string{
	// 				to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest1.ism/manifest(format=m3u8-aapl,encryption=cbc)"),
	// 				to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest2.ism/manifest(format=m3u8-aapl,encryption=cbc)"),
	// 				to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest3.ism/manifest(format=m3u8-aapl,encryption=cbc)")},
	// 				StreamingProtocol: to.Ptr(armmediaservices.StreamingPolicyStreamingProtocolHls),
	// 			},
	// 			{
	// 				EncryptionScheme: to.Ptr(armmediaservices.EncryptionSchemeEnvelopeEncryption),
	// 				Paths: []*string{
	// 					to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest1.ism/manifest(format=mpd-time-csf,encryption=cbc)"),
	// 					to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest2.ism/manifest(format=mpd-time-csf,encryption=cbc)"),
	// 					to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest3.ism/manifest(format=mpd-time-csf,encryption=cbc)")},
	// 					StreamingProtocol: to.Ptr(armmediaservices.StreamingPolicyStreamingProtocolDash),
	// 				},
	// 				{
	// 					EncryptionScheme: to.Ptr(armmediaservices.EncryptionSchemeEnvelopeEncryption),
	// 					Paths: []*string{
	// 						to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest1.ism/manifest(encryption=cbc)"),
	// 						to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest2.ism/manifest(encryption=cbc)"),
	// 						to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest3.ism/manifest(encryption=cbc)")},
	// 						StreamingProtocol: to.Ptr(armmediaservices.StreamingPolicyStreamingProtocolSmoothStreaming),
	// 					},
	// 					{
	// 						EncryptionScheme: to.Ptr(armmediaservices.EncryptionSchemeCommonEncryptionCenc),
	// 						Paths: []*string{
	// 							to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest1.ism/manifest(format=mpd-time-csf,encryption=cenc)"),
	// 							to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest2.ism/manifest(format=mpd-time-csf,encryption=cenc)"),
	// 							to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest3.ism/manifest(format=mpd-time-csf,encryption=cenc)")},
	// 							StreamingProtocol: to.Ptr(armmediaservices.StreamingPolicyStreamingProtocolDash),
	// 						},
	// 						{
	// 							EncryptionScheme: to.Ptr(armmediaservices.EncryptionSchemeCommonEncryptionCenc),
	// 							Paths: []*string{
	// 								to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest1.ism/manifest(encryption=cenc)"),
	// 								to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest2.ism/manifest(encryption=cenc)"),
	// 								to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest3.ism/manifest(encryption=cenc)")},
	// 								StreamingProtocol: to.Ptr(armmediaservices.StreamingPolicyStreamingProtocolSmoothStreaming),
	// 							},
	// 							{
	// 								EncryptionScheme: to.Ptr(armmediaservices.EncryptionSchemeCommonEncryptionCbcs),
	// 								Paths: []*string{
	// 									to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest1.ism/manifest(format=m3u8-aapl,encryption=cbcs-aapl)"),
	// 									to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest2.ism/manifest(format=m3u8-aapl,encryption=cbcs-aapl)"),
	// 									to.Ptr("/c3cd62e3-d117-4619-bcbd-99f96edd8dbe/videoManifest3.ism/manifest(format=m3u8-aapl,encryption=cbcs-aapl)")},
	// 									StreamingProtocol: to.Ptr(armmediaservices.StreamingPolicyStreamingProtocolHls),
	// 							}},
	// 						}
}
