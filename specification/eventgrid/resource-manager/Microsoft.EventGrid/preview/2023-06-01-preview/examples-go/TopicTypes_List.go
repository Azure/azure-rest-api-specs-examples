package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/TopicTypes_List.json
func ExampleTopicTypesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTopicTypesClient().NewListPager(nil)
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
		// page.TopicTypesListResult = armeventgrid.TopicTypesListResult{
		// 	Value: []*armeventgrid.TopicTypeInfo{
		// 		{
		// 			Name: to.Ptr("Microsoft.Eventhub.Namespaces"),
		// 			Type: to.Ptr("Microsoft.EventGrid/topicTypes"),
		// 			ID: to.Ptr("providers/Microsoft.EventGrid/topicTypes/Microsoft.Eventhub.Namespaces"),
		// 			Properties: &armeventgrid.TopicTypeProperties{
		// 				Description: to.Ptr("Microsoft EventHubs service events."),
		// 				DisplayName: to.Ptr("EventHubs Namespace"),
		// 				Provider: to.Ptr("Microsoft.Eventhub"),
		// 				ProvisioningState: to.Ptr(armeventgrid.TopicTypeProvisioningStateSucceeded),
		// 				ResourceRegionType: to.Ptr(armeventgrid.ResourceRegionTypeRegionalResource),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage.StorageAccounts"),
		// 			Type: to.Ptr("Microsoft.EventGrid/topicTypes"),
		// 			ID: to.Ptr("providers/Microsoft.EventGrid/topicTypes/Microsoft.Storage.StorageAccounts"),
		// 			Properties: &armeventgrid.TopicTypeProperties{
		// 				Description: to.Ptr("Microsoft Storage service events."),
		// 				DisplayName: to.Ptr("Storage Accounts"),
		// 				Provider: to.Ptr("Microsoft.Storage"),
		// 				ProvisioningState: to.Ptr(armeventgrid.TopicTypeProvisioningStateSucceeded),
		// 				ResourceRegionType: to.Ptr(armeventgrid.ResourceRegionTypeRegionalResource),
		// 			},
		// 	}},
		// }
	}
}
