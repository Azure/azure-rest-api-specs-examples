package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f6f50c6388fd5836fa142384641b8353a99874ef/specification/search/resource-manager/Microsoft.Search/preview/2024-06-01-preview/examples/CreateOrUpdateSharedPrivateLinkResource.json
func ExampleSharedPrivateLinkResourcesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsearch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSharedPrivateLinkResourcesClient().BeginCreateOrUpdate(ctx, "rg1", "mysearchservice", "testResource", armsearch.SharedPrivateLinkResource{
		Properties: &armsearch.SharedPrivateLinkResourceProperties{
			GroupID:               to.Ptr("blob"),
			PrivateLinkResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/storageAccountName"),
			RequestMessage:        to.Ptr("please approve"),
		},
	}, &armsearch.SearchManagementRequestOptions{ClientRequestID: nil}, nil)
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
	// res.SharedPrivateLinkResource = armsearch.SharedPrivateLinkResource{
	// 	Name: to.Ptr("testResource"),
	// 	Type: to.Ptr("Microsoft.Search/searchServices/sharedPrivateLinkResources"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Search/searchServices/mysearchservice/sharedPrivateLinkResources/testResource"),
	// 	Properties: &armsearch.SharedPrivateLinkResourceProperties{
	// 		GroupID: to.Ptr("blob"),
	// 		PrivateLinkResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/storageAccountName"),
	// 		RequestMessage: to.Ptr("please approve"),
	// 		Status: to.Ptr(armsearch.SharedPrivateLinkResourceStatusPending),
	// 	},
	// }
}
