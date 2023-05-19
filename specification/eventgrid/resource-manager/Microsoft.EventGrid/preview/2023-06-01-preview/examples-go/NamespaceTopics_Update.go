package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/NamespaceTopics_Update.json
func ExampleNamespaceTopicsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamespaceTopicsClient().BeginUpdate(ctx, "examplerg", "exampleNamespaceName1", "exampleNamespaceTopicName1", armeventgrid.NamespaceTopicUpdateParameters{
		Properties: &armeventgrid.NamespaceTopicUpdateParameterProperties{
			EventRetentionInDays: to.Ptr[int32](1),
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
	// res.NamespaceTopic = armeventgrid.NamespaceTopic{
	// 	Name: to.Ptr("examplenamespacetopic2"),
	// 	Type: to.Ptr("Microsoft.EventGrid/namespaces/topics"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/exampleNamespaceName1/topics/exampleNamespaceTopicName1"),
	// 	Properties: &armeventgrid.NamespaceTopicProperties{
	// 		EventRetentionInDays: to.Ptr[int32](1),
	// 		InputSchema: to.Ptr(armeventgrid.EventInputSchemaCloudEventSchemaV10),
	// 		ProvisioningState: to.Ptr(armeventgrid.NamespaceTopicProvisioningStateSucceeded),
	// 		PublisherType: to.Ptr(armeventgrid.PublisherTypeCustom),
	// 	},
	// }
}
