package armdynatrace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dynatrace/armdynatrace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3751f321654db00858e70649291af5c81e94611e/specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/TagRules_Update_MaximumSet_Gen.json
func ExampleTagRulesClient_Update_tagRulesUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdynatrace.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTagRulesClient().Update(ctx, "myResourceGroup", "myMonitor", "default", armdynatrace.TagRuleUpdate{
		LogRules: &armdynatrace.LogRules{
			FilteringTags: []*armdynatrace.FilteringTag{
				{
					Name:   to.Ptr("Environment"),
					Action: to.Ptr(armdynatrace.TagActionInclude),
					Value:  to.Ptr("Prod"),
				},
				{
					Name:   to.Ptr("Environment"),
					Action: to.Ptr(armdynatrace.TagActionExclude),
					Value:  to.Ptr("Dev"),
				}},
			SendAADLogs:          to.Ptr(armdynatrace.SendAADLogsStatusEnabled),
			SendActivityLogs:     to.Ptr(armdynatrace.SendActivityLogsStatusEnabled),
			SendSubscriptionLogs: to.Ptr(armdynatrace.SendSubscriptionLogsStatusEnabled),
		},
		MetricRules: &armdynatrace.MetricRules{
			FilteringTags: []*armdynatrace.FilteringTag{
				{
					Name:   to.Ptr("Environment"),
					Action: to.Ptr(armdynatrace.TagActionInclude),
					Value:  to.Ptr("Prod"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TagRule = armdynatrace.TagRule{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Dynatrace.Observability/monitors/tagRules"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Dynatrace.Observability/monitors/myMonitor/tagRules/default"),
	// 	Properties: &armdynatrace.MonitoringTagRulesProperties{
	// 		LogRules: &armdynatrace.LogRules{
	// 			FilteringTags: []*armdynatrace.FilteringTag{
	// 				{
	// 					Name: to.Ptr("Environment"),
	// 					Action: to.Ptr(armdynatrace.TagActionInclude),
	// 					Value: to.Ptr("Prod"),
	// 				},
	// 				{
	// 					Name: to.Ptr("Environment"),
	// 					Action: to.Ptr(armdynatrace.TagActionExclude),
	// 					Value: to.Ptr("Dev"),
	// 			}},
	// 			SendAADLogs: to.Ptr(armdynatrace.SendAADLogsStatusEnabled),
	// 			SendActivityLogs: to.Ptr(armdynatrace.SendActivityLogsStatusEnabled),
	// 			SendSubscriptionLogs: to.Ptr(armdynatrace.SendSubscriptionLogsStatusEnabled),
	// 		},
	// 		MetricRules: &armdynatrace.MetricRules{
	// 			FilteringTags: []*armdynatrace.FilteringTag{
	// 				{
	// 					Name: to.Ptr("Environment"),
	// 					Action: to.Ptr(armdynatrace.TagActionInclude),
	// 					Value: to.Ptr("Prod"),
	// 			}},
	// 		},
	// 		ProvisioningState: to.Ptr(armdynatrace.ProvisioningStateSucceeded),
	// 	},
	// 	SystemData: &armdynatrace.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-16T04:25:21.040Z"); return t}()),
	// 		CreatedBy: to.Ptr("alice@microsoft.com"),
	// 		CreatedByType: to.Ptr(armdynatrace.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-16T04:25:21.040Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("alice@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armdynatrace.CreatedByTypeUser),
	// 	},
	// }
}
