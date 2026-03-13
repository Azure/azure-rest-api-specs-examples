package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: 2025-07-15-preview/Topics_CreateOrUpdateForAzureArc.json
func ExampleTopicsClient_BeginCreateOrUpdate_topicsCreateOrUpdateForAzureArc() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("8f6b6269-84f2-4d09-9e31-1127efcd1e40", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTopicsClient().BeginCreateOrUpdate(ctx, "examplerg", "exampletopic1", armeventgrid.Topic{
		ExtendedLocation: &armeventgrid.ExtendedLocation{
			Name: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourcegroups/examplerg/providers/Microsoft.ExtendedLocation/CustomLocations/exampleCustomLocation"),
			Type: to.Ptr("CustomLocation"),
		},
		Kind:     to.Ptr(armeventgrid.ResourceKindAzureArc),
		Location: to.Ptr("westus2"),
		Properties: &armeventgrid.TopicProperties{
			InputSchema: to.Ptr(armeventgrid.InputSchemaCloudEventSchemaV10),
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
