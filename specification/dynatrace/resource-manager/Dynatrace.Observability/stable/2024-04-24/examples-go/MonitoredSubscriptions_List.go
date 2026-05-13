package armdynatrace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dynatrace/armdynatrace/v2"
)

// Generated from example definition: 2024-04-24/MonitoredSubscriptions_List.json
func ExampleMonitoredSubscriptionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdynatrace.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
		// page = armdynatrace.MonitoredSubscriptionsClientListResponse{
		// 	MonitoredSubscriptionPropertiesList: armdynatrace.MonitoredSubscriptionPropertiesList{
		// 		Value: []*armdynatrace.MonitoredSubscriptionProperties{
		// 			{
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Dynatrace.Observability/monitors/monitoredSubscriptions"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Dynatrace.Observability/monitors/myMonitor/monitoredSubscriptions/default"),
		// 				Properties: &armdynatrace.SubscriptionList{
		// 					MonitoredSubscriptionList: []*armdynatrace.MonitoredSubscription{
		// 						{
		// 							Status: to.Ptr(armdynatrace.StatusActive),
		// 							SubscriptionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000"),
		// 							TagRules: &armdynatrace.MonitoringTagRulesProperties{
		// 								LogRules: &armdynatrace.LogRules{
		// 									FilteringTags: []*armdynatrace.FilteringTag{
		// 										{
		// 											Name: to.Ptr("Environment"),
		// 											Action: to.Ptr(armdynatrace.TagActionInclude),
		// 											Value: to.Ptr("Prod"),
		// 										},
		// 										{
		// 											Name: to.Ptr("Environment"),
		// 											Action: to.Ptr(armdynatrace.TagActionExclude),
		// 											Value: to.Ptr("Dev"),
		// 										},
		// 									},
		// 									SendAADLogs: to.Ptr(armdynatrace.SendAADLogsStatusDisabled),
		// 									SendActivityLogs: to.Ptr(armdynatrace.SendActivityLogsStatusEnabled),
		// 									SendSubscriptionLogs: to.Ptr(armdynatrace.SendSubscriptionLogsStatusEnabled),
		// 								},
		// 								MetricRules: &armdynatrace.MetricRules{
		// 									FilteringTags: []*armdynatrace.FilteringTag{
		// 									},
		// 								},
		// 							},
		// 						},
		// 						{
		// 							Status: to.Ptr(armdynatrace.StatusFailed),
		// 							SubscriptionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000001"),
		// 							TagRules: &armdynatrace.MonitoringTagRulesProperties{
		// 								LogRules: &armdynatrace.LogRules{
		// 									FilteringTags: []*armdynatrace.FilteringTag{
		// 										{
		// 											Name: to.Ptr("Environment"),
		// 											Action: to.Ptr(armdynatrace.TagActionInclude),
		// 											Value: to.Ptr("Prod"),
		// 										},
		// 										{
		// 											Name: to.Ptr("Environment"),
		// 											Action: to.Ptr(armdynatrace.TagActionExclude),
		// 											Value: to.Ptr("Dev"),
		// 										},
		// 									},
		// 									SendAADLogs: to.Ptr(armdynatrace.SendAADLogsStatusDisabled),
		// 									SendActivityLogs: to.Ptr(armdynatrace.SendActivityLogsStatusEnabled),
		// 									SendSubscriptionLogs: to.Ptr(armdynatrace.SendSubscriptionLogsStatusEnabled),
		// 								},
		// 								MetricRules: &armdynatrace.MetricRules{
		// 									FilteringTags: []*armdynatrace.FilteringTag{
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
