package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/streaming-policies-create-clear.json
func ExampleStreamingPoliciesClient_Create_createsAStreamingPolicyWithClearStreaming() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewStreamingPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Create(ctx, "contoso", "contosomedia", "UserCreatedClearStreamingPolicy", armmediaservices.StreamingPolicy{
		Properties: &armmediaservices.StreamingPolicyProperties{
			NoEncryption: &armmediaservices.NoEncryption{
				EnabledProtocols: &armmediaservices.EnabledProtocols{
					Dash:            to.Ptr(true),
					Download:        to.Ptr(true),
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
