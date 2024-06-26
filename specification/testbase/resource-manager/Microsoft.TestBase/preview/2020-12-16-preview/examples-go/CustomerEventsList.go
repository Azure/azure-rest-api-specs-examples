package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/CustomerEventsList.json
func ExampleCustomerEventsClient_NewListByTestBaseAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtestbase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomerEventsClient().NewListByTestBaseAccountPager("contoso-rg1", "contoso-testBaseAccount1", nil)
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
		// page.CustomerEventListResult = armtestbase.CustomerEventListResult{
		// 	Value: []*armtestbase.CustomerEventResource{
		// 		{
		// 			Name: to.Ptr("ExecutionFailure"),
		// 			Type: to.Ptr("Microsoft.TestBase/testBaseAccounts/customerEvents"),
		// 			ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg1/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount1/customerEvents/ExecutionFailure"),
		// 			Properties: &armtestbase.CustomerEventProperties{
		// 				EventName: to.Ptr("ExecutionFailure"),
		// 				Receivers: []*armtestbase.NotificationEventReceiver{
		// 					{
		// 						ReceiverType: to.Ptr("SubscriptionRole"),
		// 						ReceiverValue: &armtestbase.NotificationReceiverValue{
		// 							SubscriptionReceiverValue: &armtestbase.SubscriptionReceiverValue{
		// 								Role: to.Ptr("Contributor"),
		// 								SubscriptionID: to.Ptr("test-subscription-id"),
		// 								SubscriptionName: to.Ptr("test-subscription"),
		// 							},
		// 						},
		// 					},
		// 					{
		// 						ReceiverType: to.Ptr("DistributionGroup"),
		// 						ReceiverValue: &armtestbase.NotificationReceiverValue{
		// 							DistributionGroupListReceiverValue: &armtestbase.DistributionGroupListReceiverValue{
		// 								DistributionGroups: []*string{
		// 									to.Ptr("test@microsoft.com")},
		// 								},
		// 							},
		// 					}},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("WeeklySummary"),
		// 				Type: to.Ptr("Microsoft.TestBase/testBaseAccounts/customerEvents"),
		// 				ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg1/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount1/customerEvents/WeeklySummary"),
		// 				Properties: &armtestbase.CustomerEventProperties{
		// 					EventName: to.Ptr("WeeklySummary"),
		// 					Receivers: []*armtestbase.NotificationEventReceiver{
		// 						{
		// 							ReceiverType: to.Ptr("UserObjects"),
		// 							ReceiverValue: &armtestbase.NotificationReceiverValue{
		// 								UserObjectReceiverValue: &armtestbase.UserObjectReceiverValue{
		// 									UserObjectIDs: []*string{
		// 										to.Ptr("245245245245325"),
		// 										to.Ptr("365365365363565")},
		// 									},
		// 								},
		// 						}},
		// 					},
		// 			}},
		// 		}
	}
}
