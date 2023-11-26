package armlogz_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logz/armlogz"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/Monitors_ListByResourceGroup.json
func ExampleMonitorsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogz.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page.MonitorResourceListResponse = armlogz.MonitorResourceListResponse{
		// 	Value: []*armlogz.MonitorResource{
		// 		{
		// 			Name: to.Ptr("myMonitor"),
		// 			Type: to.Ptr("Microsoft.Logz/monitors"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/monitors/myMonitor"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armlogz.MonitorProperties{
		// 				PlanData: &armlogz.PlanData{
		// 					BillingCycle: to.Ptr("Monthly"),
		// 					EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-30T13:14:33.000Z"); return t}()),
		// 					PlanDetails: to.Ptr("logzapitestplan"),
		// 					UsageType: to.Ptr("Committed"),
		// 				},
		// 				ProvisioningState: to.Ptr(armlogz.ProvisioningStateSucceeded),
		// 				UserInfo: &armlogz.UserInfo{
		// 					EmailAddress: to.Ptr("alice@microsoft.com"),
		// 					FirstName: to.Ptr("Alice"),
		// 					LastName: to.Ptr("Bob"),
		// 					PhoneNumber: to.Ptr("123456"),
		// 				},
		// 			},
		// 			Tags: map[string]*string{
		// 				"Environment": to.Ptr("Dev"),
		// 			},
		// 	}},
		// }
	}
}
