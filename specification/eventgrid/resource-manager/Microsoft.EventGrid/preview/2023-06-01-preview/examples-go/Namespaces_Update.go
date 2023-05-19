package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Namespaces_Update.json
func ExampleNamespacesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamespacesClient().BeginUpdate(ctx, "examplerg", "exampleNamespaceName1", armeventgrid.NamespaceUpdateParameters{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1Updated"),
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
	// res.Namespace = armeventgrid.Namespace{
	// 	Name: to.Ptr("exampleNamespaceName1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/namespaces"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/exampleNamespaceName1"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armeventgrid.NamespaceProperties{
	// 		ProvisioningState: to.Ptr(armeventgrid.NamespaceProvisioningStateSucceeded),
	// 		TopicSpacesConfiguration: &armeventgrid.TopicSpacesConfiguration{
	// 			Hostname: to.Ptr("exampleNamespaceName1.westus-1.mqtt.eventgrid-int.azure.net"),
	// 			RouteTopicResourceID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampleTopic1"),
	// 			State: to.Ptr(armeventgrid.TopicSpacesConfigurationStateEnabled),
	// 		},
	// 	},
	// }
}
