package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/Topics_CreateOrUpdateForAzureArc.json
func ExampleTopicsClient_BeginCreateOrUpdate_topicsCreateOrUpdateForAzureArc() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armeventgrid.NewTopicsClient("5b4b650e-28b9-4790-b3ab-ddbd88d727c4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "examplerg", "exampletopic1", armeventgrid.Topic{
		Location: to.Ptr("westus2"),
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
		ExtendedLocation: &armeventgrid.ExtendedLocation{
			Name: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.ExtendedLocation/CustomLocations/exampleCustomLocation"),
			Type: to.Ptr("CustomLocation"),
		},
		Kind: to.Ptr(armeventgrid.ResourceKindAzureArc),
		Properties: &armeventgrid.TopicProperties{
			InputSchema: to.Ptr(armeventgrid.InputSchemaCloudEventSchemaV10),
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
