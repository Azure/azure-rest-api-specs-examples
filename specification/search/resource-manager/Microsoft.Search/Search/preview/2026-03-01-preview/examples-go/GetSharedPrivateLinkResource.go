package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch/v2"
)

// Generated from example definition: 2026-03-01-preview/GetSharedPrivateLinkResource.json
func ExampleSharedPrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsearch.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSharedPrivateLinkResourcesClient().Get(ctx, "rg1", "mysearchservice", "testResource", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsearch.SharedPrivateLinkResourcesClientGetResponse{
	// 	SharedPrivateLinkResource: &armsearch.SharedPrivateLinkResource{
	// 		Name: to.Ptr("testResource"),
	// 		Type: to.Ptr("Microsoft.Search/searchServices/sharedPrivateLinkResources"),
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Search/searchServices/mysearchservice/sharedPrivateLinkResources/testResource"),
	// 		Properties: &armsearch.SharedPrivateLinkResourceProperties{
	// 			RequestMessage: to.Ptr("please approve"),
	// 			GroupID: to.Ptr("blob"),
	// 			PrivateLinkResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/storageAccountName"),
	// 			Status: to.Ptr(armsearch.SharedPrivateLinkResourceStatusPending),
	// 		},
	// 	},
	// }
}
