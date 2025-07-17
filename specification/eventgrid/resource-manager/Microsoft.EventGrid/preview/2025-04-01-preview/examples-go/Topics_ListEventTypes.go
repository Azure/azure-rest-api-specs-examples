package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/Topics_ListEventTypes.json
func ExampleTopicsClient_NewListEventTypesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTopicsClient().NewListEventTypesPager("examplerg", "Microsoft.Storage", "storageAccounts", "ExampleStorageAccount", nil)
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
		// page.EventTypesListResult = armeventgrid.EventTypesListResult{
		// 	Value: []*armeventgrid.EventType{
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage.BlobCreated"),
		// 			Type: to.Ptr("Microsoft.EventGrid/topicTypes/eventTypes"),
		// 			ID: to.Ptr("providers/Microsoft.EventGrid/topicTypes/Microsoft.Storage.StorageAccounts/eventTypes/Microsoft.Storage.BlobCreated"),
		// 			Properties: &armeventgrid.EventTypeProperties{
		// 				Description: to.Ptr("Raised when a blob is created."),
		// 				DisplayName: to.Ptr("Blob Created"),
		// 				SchemaURL: to.Ptr("tbd"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage.BlobDeleted"),
		// 			Type: to.Ptr("Microsoft.EventGrid/topicTypes/eventTypes"),
		// 			ID: to.Ptr("providers/Microsoft.EventGrid/topicTypes/Microsoft.Storage.StorageAccounts/eventTypes/Microsoft.Storage.BlobDeleted"),
		// 			Properties: &armeventgrid.EventTypeProperties{
		// 				Description: to.Ptr("Raised when a blob is deleted."),
		// 				DisplayName: to.Ptr("Blob Deleted"),
		// 				SchemaURL: to.Ptr("tbd"),
		// 			},
		// 	}},
		// }
	}
}
