package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-policies-list.json
func ExampleStreamingPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStreamingPoliciesClient().NewListPager("contoso", "contosomedia", &armmediaservices.StreamingPoliciesClientListOptions{Filter: nil,
		Top:     nil,
		Orderby: nil,
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
		// page.StreamingPolicyCollection = armmediaservices.StreamingPolicyCollection{
		// 	Value: []*armmediaservices.StreamingPolicy{
		// 		{
		// 			Name: to.Ptr("clearStreamingPolicy"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/streamingPolicies"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/streamingPolicies/clearStreamingPolicy"),
		// 			Properties: &armmediaservices.StreamingPolicyProperties{
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:30.676Z"); return t}()),
		// 				NoEncryption: &armmediaservices.NoEncryption{
		// 					EnabledProtocols: &armmediaservices.EnabledProtocols{
		// 						Dash: to.Ptr(true),
		// 						Download: to.Ptr(true),
		// 						Hls: to.Ptr(true),
		// 						SmoothStreaming: to.Ptr(true),
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("secureStreamingPolicy"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/streamingPolicies"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/streamingPolicies/secureStreamingPolicy"),
		// 			Properties: &armmediaservices.StreamingPolicyProperties{
		// 				CommonEncryptionCbcs: &armmediaservices.CommonEncryptionCbcs{
		// 					ClearTracks: []*armmediaservices.TrackSelection{
		// 					},
		// 					ContentKeys: &armmediaservices.StreamingPolicyContentKeys{
		// 						DefaultKey: &armmediaservices.DefaultKey{
		// 							Label: to.Ptr("cbcsDefaultKey"),
		// 						},
		// 						KeyToTrackMappings: []*armmediaservices.StreamingPolicyContentKey{
		// 						},
		// 					},
		// 					Drm: &armmediaservices.CbcsDrmConfiguration{
		// 						FairPlay: &armmediaservices.StreamingPolicyFairPlayConfiguration{
		// 							AllowPersistentLicense: to.Ptr(true),
		// 							CustomLicenseAcquisitionURLTemplate: to.Ptr("https://contoso.com/{AssetAlternativeId}/fairplay/{ContentKeyId}"),
		// 						},
		// 					},
		// 					EnabledProtocols: &armmediaservices.EnabledProtocols{
		// 						Dash: to.Ptr(false),
		// 						Download: to.Ptr(false),
		// 						Hls: to.Ptr(true),
		// 						SmoothStreaming: to.Ptr(false),
		// 					},
		// 				},
		// 				CommonEncryptionCenc: &armmediaservices.CommonEncryptionCenc{
		// 					ClearTracks: []*armmediaservices.TrackSelection{
		// 						{
		// 							TrackSelections: []*armmediaservices.TrackPropertyCondition{
		// 								{
		// 									Operation: to.Ptr(armmediaservices.TrackPropertyCompareOperationUnknown),
		// 									Property: to.Ptr(armmediaservices.TrackPropertyTypeFourCC),
		// 									Value: to.Ptr("hev1"),
		// 							}},
		// 					}},
		// 					ContentKeys: &armmediaservices.StreamingPolicyContentKeys{
		// 						DefaultKey: &armmediaservices.DefaultKey{
		// 							Label: to.Ptr("cencDefaultKey"),
		// 						},
		// 						KeyToTrackMappings: []*armmediaservices.StreamingPolicyContentKey{
		// 						},
		// 					},
		// 					Drm: &armmediaservices.CencDrmConfiguration{
		// 						PlayReady: &armmediaservices.StreamingPolicyPlayReadyConfiguration{
		// 							CustomLicenseAcquisitionURLTemplate: to.Ptr("https://contoso.com/{AssetAlternativeId}/playready/{ContentKeyId}"),
		// 							PlayReadyCustomAttributes: to.Ptr("PlayReady CustomAttributes"),
		// 						},
		// 						Widevine: &armmediaservices.StreamingPolicyWidevineConfiguration{
		// 							CustomLicenseAcquisitionURLTemplate: to.Ptr("https://contoso.com/{AssetAlternativeId}/widevine/{ContentKeyId"),
		// 						},
		// 					},
		// 					EnabledProtocols: &armmediaservices.EnabledProtocols{
		// 						Dash: to.Ptr(true),
		// 						Download: to.Ptr(false),
		// 						Hls: to.Ptr(false),
		// 						SmoothStreaming: to.Ptr(true),
		// 					},
		// 				},
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:30.678Z"); return t}()),
		// 				DefaultContentKeyPolicyName: to.Ptr("PolicyWithMultipleOptions"),
		// 				EnvelopeEncryption: &armmediaservices.EnvelopeEncryption{
		// 					ClearTracks: []*armmediaservices.TrackSelection{
		// 					},
		// 					ContentKeys: &armmediaservices.StreamingPolicyContentKeys{
		// 						DefaultKey: &armmediaservices.DefaultKey{
		// 							Label: to.Ptr("aesDefaultKey"),
		// 						},
		// 						KeyToTrackMappings: []*armmediaservices.StreamingPolicyContentKey{
		// 						},
		// 					},
		// 					CustomKeyAcquisitionURLTemplate: to.Ptr("https://contoso.com/{AssetAlternativeId}/envelope/{ContentKeyId}"),
		// 					EnabledProtocols: &armmediaservices.EnabledProtocols{
		// 						Dash: to.Ptr(true),
		// 						Download: to.Ptr(false),
		// 						Hls: to.Ptr(true),
		// 						SmoothStreaming: to.Ptr(true),
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("secureStreamingPolicyWithCommonEncryptionCbcsOnly"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/streamingPolicies"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/streamingPolicies/secureStreamingPolicyWithCommonEncryptionCbcsOnly"),
		// 			Properties: &armmediaservices.StreamingPolicyProperties{
		// 				CommonEncryptionCbcs: &armmediaservices.CommonEncryptionCbcs{
		// 					ClearTracks: []*armmediaservices.TrackSelection{
		// 					},
		// 					ContentKeys: &armmediaservices.StreamingPolicyContentKeys{
		// 						DefaultKey: &armmediaservices.DefaultKey{
		// 							Label: to.Ptr("cbcsDefaultKey"),
		// 						},
		// 						KeyToTrackMappings: []*armmediaservices.StreamingPolicyContentKey{
		// 						},
		// 					},
		// 					Drm: &armmediaservices.CbcsDrmConfiguration{
		// 						FairPlay: &armmediaservices.StreamingPolicyFairPlayConfiguration{
		// 							AllowPersistentLicense: to.Ptr(true),
		// 							CustomLicenseAcquisitionURLTemplate: to.Ptr("https://contoso.com/{AssetAlternativeId}/fairplay/{ContentKeyId}"),
		// 						},
		// 					},
		// 					EnabledProtocols: &armmediaservices.EnabledProtocols{
		// 						Dash: to.Ptr(false),
		// 						Download: to.Ptr(false),
		// 						Hls: to.Ptr(true),
		// 						SmoothStreaming: to.Ptr(false),
		// 					},
		// 				},
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:30.678Z"); return t}()),
		// 				DefaultContentKeyPolicyName: to.Ptr("PolicyWithMultipleOptions"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("secureStreamingPolicyWithCommonEncryptionCencOnly"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/streamingPolicies"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/streamingPolicies/secureStreamingPolicyWithCommonEncryptionCencOnly"),
		// 			Properties: &armmediaservices.StreamingPolicyProperties{
		// 				CommonEncryptionCenc: &armmediaservices.CommonEncryptionCenc{
		// 					ClearTracks: []*armmediaservices.TrackSelection{
		// 						{
		// 							TrackSelections: []*armmediaservices.TrackPropertyCondition{
		// 								{
		// 									Operation: to.Ptr(armmediaservices.TrackPropertyCompareOperationUnknown),
		// 									Property: to.Ptr(armmediaservices.TrackPropertyTypeFourCC),
		// 									Value: to.Ptr("hev1"),
		// 							}},
		// 					}},
		// 					ContentKeys: &armmediaservices.StreamingPolicyContentKeys{
		// 						DefaultKey: &armmediaservices.DefaultKey{
		// 							Label: to.Ptr("cencDefaultKey"),
		// 						},
		// 						KeyToTrackMappings: []*armmediaservices.StreamingPolicyContentKey{
		// 						},
		// 					},
		// 					Drm: &armmediaservices.CencDrmConfiguration{
		// 						PlayReady: &armmediaservices.StreamingPolicyPlayReadyConfiguration{
		// 							CustomLicenseAcquisitionURLTemplate: to.Ptr("https://contoso.com/{AssetAlternativeId}/playready/{ContentKeyId}"),
		// 							PlayReadyCustomAttributes: to.Ptr("PlayReady CustomAttributes"),
		// 						},
		// 						Widevine: &armmediaservices.StreamingPolicyWidevineConfiguration{
		// 							CustomLicenseAcquisitionURLTemplate: to.Ptr("https://contoso.com/{AssetAlternativeId}/widevine/{ContentKeyId"),
		// 						},
		// 					},
		// 					EnabledProtocols: &armmediaservices.EnabledProtocols{
		// 						Dash: to.Ptr(true),
		// 						Download: to.Ptr(false),
		// 						Hls: to.Ptr(false),
		// 						SmoothStreaming: to.Ptr(true),
		// 					},
		// 				},
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:30.678Z"); return t}()),
		// 				DefaultContentKeyPolicyName: to.Ptr("PolicyWithMultipleOptions"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("secureStreamingPolicyWithEnvelopeEncryptionOnly"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/streamingPolicies"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/streamingPolicies/secureStreamingPolicyWithEnvelopeEncryptionOnly"),
		// 			Properties: &armmediaservices.StreamingPolicyProperties{
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:30.678Z"); return t}()),
		// 				DefaultContentKeyPolicyName: to.Ptr("PolicyWithClearKeyOptionAndTokenRestriction"),
		// 				EnvelopeEncryption: &armmediaservices.EnvelopeEncryption{
		// 					ClearTracks: []*armmediaservices.TrackSelection{
		// 					},
		// 					ContentKeys: &armmediaservices.StreamingPolicyContentKeys{
		// 						DefaultKey: &armmediaservices.DefaultKey{
		// 							Label: to.Ptr("aesDefaultKey"),
		// 						},
		// 						KeyToTrackMappings: []*armmediaservices.StreamingPolicyContentKey{
		// 						},
		// 					},
		// 					CustomKeyAcquisitionURLTemplate: to.Ptr("https://contoso.com/{AssetAlternativeId}/envelope/{ContentKeyId}"),
		// 					EnabledProtocols: &armmediaservices.EnabledProtocols{
		// 						Dash: to.Ptr(true),
		// 						Download: to.Ptr(false),
		// 						Hls: to.Ptr(true),
		// 						SmoothStreaming: to.Ptr(true),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
