package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8691fbfca8fcdc5a241a0b501c32fd4a76bb0cd/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/Namespaces_ListByResourceGroup.json
func ExampleNamespacesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNamespacesClient().NewListByResourceGroupPager("examplerg", &armeventgrid.NamespacesClientListByResourceGroupOptions{Filter: nil,
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
		// page.NamespacesListResult = armeventgrid.NamespacesListResult{
		// 	Value: []*armeventgrid.Namespace{
		// 		{
		// 			Name: to.Ptr("exampleNamespaceName1"),
		// 			Type: to.Ptr("Microsoft.EventGrid/namespaces"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/exampleNamespaceName1"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 				"key3": to.Ptr("value3"),
		// 			},
		// 			Properties: &armeventgrid.NamespaceProperties{
		// 				ProvisioningState: to.Ptr(armeventgrid.NamespaceProvisioningStateSucceeded),
		// 				TopicSpacesConfiguration: &armeventgrid.TopicSpacesConfiguration{
		// 					Hostname: to.Ptr("exampleNamespaceName1.westus-1.mqtt.eventgrid-int.azure.net"),
		// 					RouteTopicResourceID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampleTopic1"),
		// 					State: to.Ptr(armeventgrid.TopicSpacesConfigurationStateEnabled),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
