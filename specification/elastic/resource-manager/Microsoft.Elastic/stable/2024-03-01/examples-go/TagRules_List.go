package armelastic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad60d7f8eba124edc6999677c55aba2184e303b0/specification/elastic/resource-manager/Microsoft.Elastic/stable/2024-03-01/examples/TagRules_List.json
func ExampleTagRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelastic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTagRulesClient().NewListPager("myResourceGroup", "myMonitor", nil)
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
		// page.MonitoringTagRulesListResponse = armelastic.MonitoringTagRulesListResponse{
		// 	Value: []*armelastic.MonitoringTagRules{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Datadog/monitors/tagRules"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Datadog/monitors/myMonitor/tagRules/default"),
		// 			Properties: &armelastic.MonitoringTagRulesProperties{
		// 				LogRules: &armelastic.LogRules{
		// 					FilteringTags: []*armelastic.FilteringTag{
		// 						{
		// 							Name: to.Ptr("Environment"),
		// 							Action: to.Ptr(armelastic.TagActionInclude),
		// 							Value: to.Ptr("Prod"),
		// 						},
		// 						{
		// 							Name: to.Ptr("Environment"),
		// 							Action: to.Ptr(armelastic.TagActionExclude),
		// 							Value: to.Ptr("Dev"),
		// 					}},
		// 					SendAADLogs: to.Ptr(false),
		// 					SendActivityLogs: to.Ptr(true),
		// 					SendSubscriptionLogs: to.Ptr(true),
		// 				},
		// 				ProvisioningState: to.Ptr(armelastic.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
