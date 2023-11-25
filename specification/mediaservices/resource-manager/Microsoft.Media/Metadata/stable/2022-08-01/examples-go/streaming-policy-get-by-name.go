package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-policy-get-by-name.json
func ExampleStreamingPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStreamingPoliciesClient().Get(ctx, "contoso", "contosomedia", "clearStreamingPolicy", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StreamingPolicy = armmediaservices.StreamingPolicy{
	// 	Name: to.Ptr("clearStreamingPolicy"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/streamingPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/streamingPolicies/clearStreamingPolicy"),
	// 	Properties: &armmediaservices.StreamingPolicyProperties{
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:30.850Z"); return t}()),
	// 		NoEncryption: &armmediaservices.NoEncryption{
	// 			EnabledProtocols: &armmediaservices.EnabledProtocols{
	// 				Dash: to.Ptr(true),
	// 				Download: to.Ptr(true),
	// 				Hls: to.Ptr(true),
	// 				SmoothStreaming: to.Ptr(true),
	// 			},
	// 		},
	// 	},
	// }
}
