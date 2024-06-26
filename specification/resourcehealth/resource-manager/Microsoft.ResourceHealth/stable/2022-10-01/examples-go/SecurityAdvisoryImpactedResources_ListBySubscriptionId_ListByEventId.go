package armresourcehealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b74978708bb95475562412d4654c00fbcedd9f89/specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/SecurityAdvisoryImpactedResources_ListBySubscriptionId_ListByEventId.json
func ExampleSecurityAdvisoryImpactedResourcesClient_NewListBySubscriptionIDAndEventIDPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcehealth.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSecurityAdvisoryImpactedResourcesClient().NewListBySubscriptionIDAndEventIDPager("BC_1-FXZ", &armresourcehealth.SecurityAdvisoryImpactedResourcesClientListBySubscriptionIDAndEventIDOptions{Filter: nil})
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
		// page.EventImpactedResourceListResult = armresourcehealth.EventImpactedResourceListResult{
		// 	Value: []*armresourcehealth.EventImpactedResource{
		// 		{
		// 			Name: to.Ptr("jkl-901-hgy-445"),
		// 			Type: to.Ptr("Microsoft.ResourceHealth/events/impactedResources"),
		// 			ID: to.Ptr("/subscriptions/{subscripitionId}/providers/Microsoft.ResourceHealth/events/BC_1-FXZ/impactedResources/jkl-901-hgy-445"),
		// 			Properties: &armresourcehealth.EventImpactedResourceProperties{
		// 				Info: []*armresourcehealth.KeyValueItem{
		// 					{
		// 						Key: to.Ptr("key-A"),
		// 						Value: to.Ptr("sample-1"),
		// 					},
		// 					{
		// 						Key: to.Ptr("key-B"),
		// 						Value: to.Ptr("sample-2"),
		// 					},
		// 					{
		// 						Key: to.Ptr("key-C"),
		// 						Value: to.Ptr("sample-3"),
		// 					},
		// 					{
		// 						Key: to.Ptr("key-D"),
		// 						Value: to.Ptr("sample-4"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("wer-345-tyu-789"),
		// 			Type: to.Ptr("Microsoft.ResourceHealth/events/impactedResources"),
		// 			ID: to.Ptr("/subscriptions/{subscripitionId}/providers/Microsoft.ResourceHealth/events/BC_1-FXZ/impactedResources/wer-345-tyu-789"),
		// 			Properties: &armresourcehealth.EventImpactedResourceProperties{
		// 				Info: []*armresourcehealth.KeyValueItem{
		// 					{
		// 						Key: to.Ptr("key-E"),
		// 						Value: to.Ptr("sample-5"),
		// 					},
		// 					{
		// 						Key: to.Ptr("key-F"),
		// 						Value: to.Ptr("sample-6"),
		// 					},
		// 					{
		// 						Key: to.Ptr("key-G"),
		// 						Value: to.Ptr("sample-7"),
		// 					},
		// 					{
		// 						Key: to.Ptr("key-H"),
		// 						Value: to.Ptr("sample-8"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
