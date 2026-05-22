package armelastic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic/v3"
)

// Generated from example definition: 2025-06-01/MonitoredSubscriptions_CreateorUpdate.json
func ExampleMonitoredSubscriptionsClient_BeginCreateorUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelastic.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMonitoredSubscriptionsClient().BeginCreateorUpdate(ctx, "myResourceGroup", "myMonitor", "default", armelastic.MonitoredSubscriptionProperties{}, nil)
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
	// res = armelastic.MonitoredSubscriptionsClientCreateorUpdateResponse{
	// 	MonitoredSubscriptionProperties: &armelastic.MonitoredSubscriptionProperties{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.Elastic/monitors/monitoredSubscriptions"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Elastic/monitors/myMonitor/monitoredSubscriptions/default"),
	// 		Properties: &armelastic.SubscriptionList{
	// 			MonitoredSubscriptionList: []*armelastic.MonitoredSubscription{
	// 				{
	// 					Status: to.Ptr(armelastic.StatusActive),
	// 					SubscriptionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000"),
	// 					TagRules: &armelastic.MonitoringTagRulesProperties{
	// 						LogRules: &armelastic.LogRules{
	// 							FilteringTags: []*armelastic.FilteringTag{
	// 								{
	// 									Name: to.Ptr("saokgpjvdlorciqbjmjxazpee"),
	// 									Action: to.Ptr(armelastic.TagActionInclude),
	// 									Value: to.Ptr("sarxrqsxouhdjwsrqqicbeirdb"),
	// 								},
	// 							},
	// 							SendAADLogs: to.Ptr(true),
	// 							SendActivityLogs: to.Ptr(true),
	// 							SendSubscriptionLogs: to.Ptr(true),
	// 						},
	// 						ProvisioningState: to.Ptr(armelastic.ProvisioningStateAccepted),
	// 					},
	// 				},
	// 				{
	// 					Status: to.Ptr(armelastic.StatusFailed),
	// 					SubscriptionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000001"),
	// 					TagRules: &armelastic.MonitoringTagRulesProperties{
	// 						LogRules: &armelastic.LogRules{
	// 							FilteringTags: []*armelastic.FilteringTag{
	// 								{
	// 									Name: to.Ptr("saokgpjvdlorciqbjmjxazpee"),
	// 									Action: to.Ptr(armelastic.TagActionInclude),
	// 									Value: to.Ptr("sarxrqsxouhdjwsrqqicbeirdb"),
	// 								},
	// 							},
	// 							SendAADLogs: to.Ptr(true),
	// 							SendActivityLogs: to.Ptr(true),
	// 							SendSubscriptionLogs: to.Ptr(true),
	// 						},
	// 						ProvisioningState: to.Ptr(armelastic.ProvisioningStateAccepted),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
