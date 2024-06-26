package armdatadog_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c280892951a9e45c059132c05aace25a9c752d48/specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/MonitoredSubscriptions_List.json
func ExampleMonitoredSubscriptionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitoredSubscriptionsClient().NewListPager("myResourceGroup", "myMonitor", nil)
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
		// page.MonitoredSubscriptionPropertiesList = armdatadog.MonitoredSubscriptionPropertiesList{
		// 	Value: []*armdatadog.MonitoredSubscriptionProperties{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Datadog/monitors/monitoredSubscriptions"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Datadog/monitors/myMonitor/monitoredSubscriptions/default"),
		// 			Properties: &armdatadog.SubscriptionList{
		// 				MonitoredSubscriptionList: []*armdatadog.MonitoredSubscription{
		// 					{
		// 						Status: to.Ptr(armdatadog.StatusActive),
		// 						SubscriptionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000"),
		// 						TagRules: &armdatadog.MonitoringTagRulesProperties{
		// 							Automuting: to.Ptr(true),
		// 							LogRules: &armdatadog.LogRules{
		// 								FilteringTags: []*armdatadog.FilteringTag{
		// 									{
		// 										Name: to.Ptr("Environment"),
		// 										Action: to.Ptr(armdatadog.TagActionInclude),
		// 										Value: to.Ptr("Prod"),
		// 									},
		// 									{
		// 										Name: to.Ptr("Environment"),
		// 										Action: to.Ptr(armdatadog.TagActionExclude),
		// 										Value: to.Ptr("Dev"),
		// 								}},
		// 								SendAADLogs: to.Ptr(false),
		// 								SendResourceLogs: to.Ptr(true),
		// 								SendSubscriptionLogs: to.Ptr(true),
		// 							},
		// 							MetricRules: &armdatadog.MetricRules{
		// 								FilteringTags: []*armdatadog.FilteringTag{
		// 								},
		// 							},
		// 						},
		// 					},
		// 					{
		// 						Status: to.Ptr(armdatadog.StatusFailed),
		// 						SubscriptionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000001"),
		// 						TagRules: &armdatadog.MonitoringTagRulesProperties{
		// 							Automuting: to.Ptr(true),
		// 							LogRules: &armdatadog.LogRules{
		// 								FilteringTags: []*armdatadog.FilteringTag{
		// 									{
		// 										Name: to.Ptr("Environment"),
		// 										Action: to.Ptr(armdatadog.TagActionInclude),
		// 										Value: to.Ptr("Prod"),
		// 									},
		// 									{
		// 										Name: to.Ptr("Environment"),
		// 										Action: to.Ptr(armdatadog.TagActionExclude),
		// 										Value: to.Ptr("Dev"),
		// 								}},
		// 								SendAADLogs: to.Ptr(false),
		// 								SendResourceLogs: to.Ptr(true),
		// 								SendSubscriptionLogs: to.Ptr(true),
		// 							},
		// 							MetricRules: &armdatadog.MetricRules{
		// 								FilteringTags: []*armdatadog.FilteringTag{
		// 								},
		// 							},
		// 						},
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
