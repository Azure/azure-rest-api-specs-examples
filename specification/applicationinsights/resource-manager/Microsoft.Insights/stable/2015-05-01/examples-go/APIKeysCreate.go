package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a0168458930c86636a76bcd7acfdc9c81291bfc/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/APIKeysCreate.json
func ExampleAPIKeysClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPIKeysClient().Create(ctx, "my-resource-group", "my-component", armapplicationinsights.APIKeyRequest{
		Name: to.Ptr("test2"),
		LinkedReadProperties: []*string{
			to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/api"),
			to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/agentconfig")},
		LinkedWriteProperties: []*string{
			to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/annotations")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComponentAPIKey = armapplicationinsights.ComponentAPIKey{
	// 	Name: to.Ptr("test"),
	// 	APIKey: to.Ptr("0000000000000000000000000000000000000000"),
	// 	CreatedDate: to.Ptr("Thu, 28 Sep 2017 16:58:52 GMT"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/my-resource-group/providers/Microsoft.Insights/components/my-component/apikeys/00000000-0000-0000-0000-000000000000"),
	// 	LinkedReadProperties: []*string{
	// 		to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/api"),
	// 		to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/agentconfig")},
	// 		LinkedWriteProperties: []*string{
	// 			to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/annotations")},
	// 		}
}
