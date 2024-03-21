package armnewrelicobservability_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/newrelic/armnewrelicobservability"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-01-01/examples/MonitoredSubscriptions_Update.json
func ExampleMonitoredSubscriptionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnewrelicobservability.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMonitoredSubscriptionsClient().BeginUpdate(ctx, "myResourceGroup", "myMonitor", armnewrelicobservability.ConfigurationNameDefault, armnewrelicobservability.MonitoredSubscriptionProperties{}, nil)
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
	// res.MonitoredSubscriptionProperties = armnewrelicobservability.MonitoredSubscriptionProperties{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("NewRelic.Observability/monitors/monitoredSubscriptions"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/NewRelic.Observability/monitors/myMonitor/monitoredSubscriptions/default"),
	// 	Properties: &armnewrelicobservability.SubscriptionList{
	// 		MonitoredSubscriptionList: []*armnewrelicobservability.MonitoredSubscription{
	// 			{
	// 				Status: to.Ptr(armnewrelicobservability.StatusActive),
	// 				SubscriptionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000"),
	// 				TagRules: &armnewrelicobservability.MonitoringTagRulesProperties{
	// 					LogRules: &armnewrelicobservability.LogRules{
	// 						FilteringTags: []*armnewrelicobservability.FilteringTag{
	// 							{
	// 								Name: to.Ptr("saokgpjvdlorciqbjmjxazpee"),
	// 								Action: to.Ptr(armnewrelicobservability.TagActionInclude),
	// 								Value: to.Ptr("sarxrqsxouhdjwsrqqicbeirdb"),
	// 						}},
	// 						SendAADLogs: to.Ptr(armnewrelicobservability.SendAADLogsStatusEnabled),
	// 						SendActivityLogs: to.Ptr(armnewrelicobservability.SendActivityLogsStatusEnabled),
	// 						SendSubscriptionLogs: to.Ptr(armnewrelicobservability.SendSubscriptionLogsStatusEnabled),
	// 					},
	// 					MetricRules: &armnewrelicobservability.MetricRules{
	// 						FilteringTags: []*armnewrelicobservability.FilteringTag{
	// 							{
	// 								Name: to.Ptr("saokgpjvdlorciqbjmjxazpee"),
	// 								Action: to.Ptr(armnewrelicobservability.TagActionInclude),
	// 								Value: to.Ptr("sarxrqsxouhdjwsrqqicbeirdb"),
	// 						}},
	// 						UserEmail: to.Ptr("test@testing.com"),
	// 					},
	// 					ProvisioningState: to.Ptr(armnewrelicobservability.ProvisioningStateSucceeded),
	// 				},
	// 			},
	// 			{
	// 				Status: to.Ptr(armnewrelicobservability.StatusFailed),
	// 				SubscriptionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000001"),
	// 				TagRules: &armnewrelicobservability.MonitoringTagRulesProperties{
	// 					LogRules: &armnewrelicobservability.LogRules{
	// 						FilteringTags: []*armnewrelicobservability.FilteringTag{
	// 							{
	// 								Name: to.Ptr("saokgpjvdlorciqbjmjxazpee"),
	// 								Action: to.Ptr(armnewrelicobservability.TagActionInclude),
	// 								Value: to.Ptr("sarxrqsxouhdjwsrqqicbeirdb"),
	// 						}},
	// 						SendAADLogs: to.Ptr(armnewrelicobservability.SendAADLogsStatusEnabled),
	// 						SendActivityLogs: to.Ptr(armnewrelicobservability.SendActivityLogsStatusEnabled),
	// 						SendSubscriptionLogs: to.Ptr(armnewrelicobservability.SendSubscriptionLogsStatusEnabled),
	// 					},
	// 					MetricRules: &armnewrelicobservability.MetricRules{
	// 						FilteringTags: []*armnewrelicobservability.FilteringTag{
	// 							{
	// 								Name: to.Ptr("saokgpjvdlorciqbjmjxazpee"),
	// 								Action: to.Ptr(armnewrelicobservability.TagActionInclude),
	// 								Value: to.Ptr("sarxrqsxouhdjwsrqqicbeirdb"),
	// 						}},
	// 						UserEmail: to.Ptr("test@testing.com"),
	// 					},
	// 					ProvisioningState: to.Ptr(armnewrelicobservability.ProvisioningStateSucceeded),
	// 				},
	// 		}},
	// 	},
	// }
}
