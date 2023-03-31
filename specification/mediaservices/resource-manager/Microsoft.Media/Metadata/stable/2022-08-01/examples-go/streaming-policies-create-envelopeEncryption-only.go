package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-policies-create-envelopeEncryption-only.json
func ExampleStreamingPoliciesClient_Create_createsAStreamingPolicyWithEnvelopeEncryptionOnly() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewStreamingPoliciesClient().Create(ctx, "contoso", "contosomedia", "UserCreatedSecureStreamingPolicyWithEnvelopeEncryptionOnly", armmediaservices.StreamingPolicy{
		Properties: &armmediaservices.StreamingPolicyProperties{
			DefaultContentKeyPolicyName: to.Ptr("PolicyWithClearKeyOptionAndTokenRestriction"),
			EnvelopeEncryption: &armmediaservices.EnvelopeEncryption{
				ContentKeys: &armmediaservices.StreamingPolicyContentKeys{
					DefaultKey: &armmediaservices.DefaultKey{
						Label: to.Ptr("aesDefaultKey"),
					},
				},
				CustomKeyAcquisitionURLTemplate: to.Ptr("https://contoso.com/{AssetAlternativeId}/envelope/{ContentKeyId}"),
				EnabledProtocols: &armmediaservices.EnabledProtocols{
					Dash:            to.Ptr(true),
					Download:        to.Ptr(false),
					Hls:             to.Ptr(true),
					SmoothStreaming: to.Ptr(true),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
