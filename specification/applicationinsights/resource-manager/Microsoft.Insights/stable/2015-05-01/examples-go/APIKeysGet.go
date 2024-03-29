package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/APIKeysGet.json
func ExampleAPIKeysClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPIKeysClient().Get(ctx, "my-resource-group", "my-component", "bb820f1b-3110-4a8b-ba2c-8c1129d7eb6a", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComponentAPIKey = armapplicationinsights.ComponentAPIKey{
	// 	Name: to.Ptr("test2"),
	// 	CreatedDate: to.Ptr("Thu, 28 Sep 2017 16:59:18 GMT"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/my-resource-group/providers/Microsoft.Insights/components/my-component/apikeys/bb820f1b-3110-4a8b-ba2c-8c1129d7eb6a"),
	// 	LinkedReadProperties: []*string{
	// 		to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/api"),
	// 		to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/draft"),
	// 		to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/extendqueries"),
	// 		to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/search"),
	// 		to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/aggregate"),
	// 		to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/agentconfig")},
	// 		LinkedWriteProperties: []*string{
	// 		},
	// 	}
}
