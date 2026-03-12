package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: 2025-07-15-preview/NamespaceTopics_ListByNamespace.json
func ExampleNamespaceTopicsClient_NewListByNamespacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("8f6b6269-84f2-4d09-9e31-1127efcd1e40", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNamespaceTopicsClient().NewListByNamespacePager("examplerg", "examplenamespace2", nil)
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
		// page = armeventgrid.NamespaceTopicsClientListByNamespaceResponse{
		// 	NamespaceTopicsListResult: armeventgrid.NamespaceTopicsListResult{
		// 		Value: []*armeventgrid.NamespaceTopic{
		// 			{
		// 				Name: to.Ptr("examplenamespacetopic2"),
		// 				Type: to.Ptr("Microsoft.EventGrid/namespaces/topics"),
		// 				ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/examplenamespace2/topics/examplenamespacetopic2"),
		// 				Properties: &armeventgrid.NamespaceTopicProperties{
		// 					EventRetentionInDays: to.Ptr[int32](1),
		// 					InputSchema: to.Ptr(armeventgrid.EventInputSchemaCloudEventSchemaV10),
		// 					ProvisioningState: to.Ptr(armeventgrid.NamespaceTopicProvisioningStateSucceeded),
		// 					PublisherType: to.Ptr(armeventgrid.PublisherTypeCustom),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("examplenamespacetopic4"),
		// 				Type: to.Ptr("Microsoft.EventGrid/namespaces/topics"),
		// 				ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/examplenamespace2/topics/examplenamespacetopic4"),
		// 				Properties: &armeventgrid.NamespaceTopicProperties{
		// 					EventRetentionInDays: to.Ptr[int32](1),
		// 					InputSchema: to.Ptr(armeventgrid.EventInputSchemaCloudEventSchemaV10),
		// 					ProvisioningState: to.Ptr(armeventgrid.NamespaceTopicProvisioningStateSucceeded),
		// 					PublisherType: to.Ptr(armeventgrid.PublisherTypeCustom),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
