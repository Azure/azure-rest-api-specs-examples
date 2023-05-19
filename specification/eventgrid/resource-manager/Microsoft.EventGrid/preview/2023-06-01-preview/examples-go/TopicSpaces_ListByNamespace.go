package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/TopicSpaces_ListByNamespace.json
func ExampleTopicSpacesClient_NewListByNamespacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTopicSpacesClient().NewListByNamespacePager("examplerg", "namespace123", &armeventgrid.TopicSpacesClientListByNamespaceOptions{Filter: nil,
		Top: nil,
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
		// page.TopicSpacesListResult = armeventgrid.TopicSpacesListResult{
		// 	Value: []*armeventgrid.TopicSpace{
		// 		{
		// 			Name: to.Ptr("exampleTopicSpaceName1"),
		// 			Type: to.Ptr("Microsoft.EventGrid/namespaces/topicSpaces"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/exampleNamespaceName1/topicSpaces/exampleTopicSpaceName1"),
		// 			Properties: &armeventgrid.TopicSpaceProperties{
		// 				ProvisioningState: to.Ptr(armeventgrid.TopicSpaceProvisioningStateSucceeded),
		// 				TopicTemplates: []*string{
		// 					to.Ptr("filter1"),
		// 					to.Ptr("filter2")},
		// 				},
		// 		}},
		// 	}
	}
}
