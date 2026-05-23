package armelastic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic/v3"
)

// Generated from example definition: 2025-06-01/MonitoredSubscriptions_List.json
func ExampleMonitoredSubscriptionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelastic.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
		// page = armelastic.MonitoredSubscriptionsClientListResponse{
		// 	MonitoredSubscriptionPropertiesList: armelastic.MonitoredSubscriptionPropertiesList{
		// 		Value: []*armelastic.MonitoredSubscriptionProperties{
		// 			{
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.Elastic/monitors/monitoredSubscriptions"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Elastic/monitors/myMonitor/monitoredSubscriptions/default"),
		// 				Properties: &armelastic.SubscriptionList{
		// 					MonitoredSubscriptionList: []*armelastic.MonitoredSubscription{
		// 						{
		// 							Status: to.Ptr(armelastic.StatusActive),
		// 							SubscriptionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000"),
		// 							TagRules: &armelastic.MonitoringTagRulesProperties{
		// 								LogRules: &armelastic.LogRules{
		// 									FilteringTags: []*armelastic.FilteringTag{
		// 										{
		// 											Name: to.Ptr("saokgpjvdlorciqbjmjxazpee"),
		// 											Action: to.Ptr(armelastic.TagActionInclude),
		// 											Value: to.Ptr("sarxrqsxouhdjwsrqqicbeirdb"),
		// 										},
		// 									},
		// 									SendAADLogs: to.Ptr(true),
		// 									SendActivityLogs: to.Ptr(true),
		// 									SendSubscriptionLogs: to.Ptr(true),
		// 								},
		// 								ProvisioningState: to.Ptr(armelastic.ProvisioningStateAccepted),
		// 							},
		// 						},
		// 						{
		// 							Status: to.Ptr(armelastic.StatusFailed),
		// 							SubscriptionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000001"),
		// 							TagRules: &armelastic.MonitoringTagRulesProperties{
		// 								LogRules: &armelastic.LogRules{
		// 									FilteringTags: []*armelastic.FilteringTag{
		// 										{
		// 											Name: to.Ptr("saokgpjvdlorciqbjmjxazpee"),
		// 											Action: to.Ptr(armelastic.TagActionInclude),
		// 											Value: to.Ptr("sarxrqsxouhdjwsrqqicbeirdb"),
		// 										},
		// 									},
		// 									SendAADLogs: to.Ptr(true),
		// 									SendActivityLogs: to.Ptr(true),
		// 									SendSubscriptionLogs: to.Ptr(true),
		// 								},
		// 								ProvisioningState: to.Ptr(armelastic.ProvisioningStateAccepted),
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
