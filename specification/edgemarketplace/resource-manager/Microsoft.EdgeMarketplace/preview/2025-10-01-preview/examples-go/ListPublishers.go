package armedgemarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgemarketplace/armedgemarketplace"
)

// Generated from example definition: 2025-10-01-preview/ListPublishers.json
func ExamplePublishersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgemarketplace.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPublishersClient().NewListPager("subscriptions/4bed37fd-19a1-4d31-8b44-40267555bec5/resourceGroups/edgemarketplace-rg/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/edgemarketplace-demo", nil)
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
		// page = armedgemarketplace.PublishersClientListResponse{
		// 	PublisherListResult: armedgemarketplace.PublisherListResult{
		// 		Value: []*armedgemarketplace.Publisher{
		// 			{
		// 				ID: to.Ptr("/subscriptions/4bed37fd-19a1-4d31-8b44-40267555bec5/resourceGroups/edgemarketplace-rg/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/edgemarketplace-demo/providers/Microsoft.EdgeMarketplace/publishers/canonical"),
		// 				Name: to.Ptr("canonical"),
		// 				Type: to.Ptr("Microsoft.EdgeMarketplace/publishers"),
		// 				Properties: &armedgemarketplace.PublisherProperties{
		// 					ProvisioningState: to.Ptr(armedgemarketplace.ResourceProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/4bed37fd-19a1-4d31-8b44-40267555bec5/resourceGroups/edgemarketplace-rg/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/edgemarketplace-demo/providers/Microsoft.EdgeMarketplace/publishers/ntegralinc1586961136942"),
		// 				Name: to.Ptr("ntegralinc1586961136942"),
		// 				Type: to.Ptr("Microsoft.EdgeMarketplace/publishers"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
