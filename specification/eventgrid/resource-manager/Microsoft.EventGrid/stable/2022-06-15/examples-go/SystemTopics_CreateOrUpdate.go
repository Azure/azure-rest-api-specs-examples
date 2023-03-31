package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f88928d723133dc392e3297e6d61b7f6d10501fd/specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/SystemTopics_CreateOrUpdate.json
func ExampleSystemTopicsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSystemTopicsClient().BeginCreateOrUpdate(ctx, "examplerg", "exampleSystemTopic1", armeventgrid.SystemTopic{
		Location: to.Ptr("westus2"),
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
		Properties: &armeventgrid.SystemTopicProperties{
			Source:    to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/azureeventgridrunnerrgcentraluseuap/providers/microsoft.storage/storageaccounts/pubstgrunnerb71cd29e"),
			TopicType: to.Ptr("microsoft.storage.storageaccounts"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SystemTopic = armeventgrid.SystemTopic{
	// 	Name: to.Ptr("exampleSystemTopic2"),
	// 	Type: to.Ptr("Microsoft.EventGrid/systemTopics"),
	// 	ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/systemTopics/exampleSystemTopic2"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Properties: &armeventgrid.SystemTopicProperties{
	// 		MetricResourceID: to.Ptr("183c0fb1-17ff-47b6-ac77-5a47420ab01e"),
	// 		ProvisioningState: to.Ptr(armeventgrid.ResourceProvisioningStateSucceeded),
	// 		Source: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/azureeventgridrunnerrgcentraluseuap/providers/microsoft.storage/storageaccounts/pubstgrunnerb71cd29e"),
	// 		TopicType: to.Ptr("microsoft.storage.storageaccounts"),
	// 	},
	// }
}
