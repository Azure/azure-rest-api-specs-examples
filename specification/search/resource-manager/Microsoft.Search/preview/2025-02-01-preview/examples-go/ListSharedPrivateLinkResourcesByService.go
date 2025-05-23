package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3db6867b8e524ea6d1bc7a3bbb989fe50dd2f184/specification/search/resource-manager/Microsoft.Search/preview/2025-02-01-preview/examples/ListSharedPrivateLinkResourcesByService.json
func ExampleSharedPrivateLinkResourcesClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsearch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSharedPrivateLinkResourcesClient().NewListByServicePager("rg1", "mysearchservice", &armsearch.SearchManagementRequestOptions{ClientRequestID: nil}, nil)
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
		// page.SharedPrivateLinkResourceListResult = armsearch.SharedPrivateLinkResourceListResult{
		// 	Value: []*armsearch.SharedPrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("testResource"),
		// 			Type: to.Ptr("Microsoft.Search/searchServices/sharedPrivateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Search/searchServices/mysearchservice/sharedPrivateLinkResources/testResource"),
		// 			Properties: &armsearch.SharedPrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("blob"),
		// 				PrivateLinkResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/storageAccountName"),
		// 				RequestMessage: to.Ptr("please approve"),
		// 				Status: to.Ptr(armsearch.SharedPrivateLinkResourceStatusPending),
		// 			},
		// 	}},
		// }
	}
}
