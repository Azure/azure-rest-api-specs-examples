package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/Topics_Update.json
func ExampleTopicsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTopicsClient().BeginUpdate(ctx, "examplerg", "exampletopic1", armeventgrid.TopicUpdateParameters{
		Properties: &armeventgrid.TopicUpdateParameterProperties{
			InboundIPRules: []*armeventgrid.InboundIPRule{
				{
					Action: to.Ptr(armeventgrid.IPActionTypeAllow),
					IPMask: to.Ptr("12.18.30.15"),
				},
				{
					Action: to.Ptr(armeventgrid.IPActionTypeAllow),
					IPMask: to.Ptr("12.18.176.1"),
				}},
			PublicNetworkAccess: to.Ptr(armeventgrid.PublicNetworkAccessEnabled),
		},
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
