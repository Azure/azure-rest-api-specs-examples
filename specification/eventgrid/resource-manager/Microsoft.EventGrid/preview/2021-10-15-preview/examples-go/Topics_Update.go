package armeventgrid_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/Topics_Update.json
func ExampleTopicsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armeventgrid.NewTopicsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<topic-name>",
		armeventgrid.TopicUpdateParameters{
			Properties: &armeventgrid.TopicUpdateParameterProperties{
				InboundIPRules: []*armeventgrid.InboundIPRule{
					{
						Action: to.Ptr(armeventgrid.IPActionTypeAllow),
						IPMask: to.Ptr("<ipmask>"),
					},
					{
						Action: to.Ptr(armeventgrid.IPActionTypeAllow),
						IPMask: to.Ptr("<ipmask>"),
					}},
				PublicNetworkAccess: to.Ptr(armeventgrid.PublicNetworkAccessEnabled),
			},
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
		},
		&armeventgrid.TopicsClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
