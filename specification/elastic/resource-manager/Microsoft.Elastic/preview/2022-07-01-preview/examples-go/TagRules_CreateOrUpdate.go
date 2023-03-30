package armelastic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/232b858812a4f946a82bc11a81241826f5554fbd/specification/elastic/resource-manager/Microsoft.Elastic/preview/2022-07-01-preview/examples/TagRules_CreateOrUpdate.json
func ExampleTagRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelastic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTagRulesClient().CreateOrUpdate(ctx, "myResourceGroup", "myMonitor", "default", &armelastic.TagRulesClientCreateOrUpdateOptions{Body: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MonitoringTagRules = armelastic.MonitoringTagRules{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Datadog/monitors/tagRules"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Datadog/monitors/myMonitor/tagRules/default"),
	// 	Properties: &armelastic.MonitoringTagRulesProperties{
	// 		LogRules: &armelastic.LogRules{
	// 			FilteringTags: []*armelastic.FilteringTag{
	// 				{
	// 					Name: to.Ptr("Environment"),
	// 					Action: to.Ptr(armelastic.TagActionInclude),
	// 					Value: to.Ptr("Prod"),
	// 				},
	// 				{
	// 					Name: to.Ptr("Environment"),
	// 					Action: to.Ptr(armelastic.TagActionExclude),
	// 					Value: to.Ptr("Dev"),
	// 			}},
	// 			SendAADLogs: to.Ptr(false),
	// 			SendActivityLogs: to.Ptr(true),
	// 			SendSubscriptionLogs: to.Ptr(true),
	// 		},
	// 		ProvisioningState: to.Ptr(armelastic.ProvisioningStateSucceeded),
	// 	},
	// }
}
