package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/streaming-policies-create-clear.json
func ExampleStreamingPoliciesClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewStreamingPoliciesClient("<subscription-id>", cred, nil)
	_, err = client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<streaming-policy-name>",
		armmediaservices.StreamingPolicy{
			Properties: &armmediaservices.StreamingPolicyProperties{
				NoEncryption: &armmediaservices.NoEncryption{
					EnabledProtocols: &armmediaservices.EnabledProtocols{
						Dash:            to.BoolPtr(true),
						Download:        to.BoolPtr(true),
						Hls:             to.BoolPtr(true),
						SmoothStreaming: to.BoolPtr(true),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
