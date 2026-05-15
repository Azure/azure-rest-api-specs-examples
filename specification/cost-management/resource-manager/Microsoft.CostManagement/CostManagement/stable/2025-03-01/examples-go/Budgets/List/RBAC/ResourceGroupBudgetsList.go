package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v3"
)

// Generated from example definition: 2025-03-01/Budgets/List/RBAC/ResourceGroupBudgetsList.json
func ExampleBudgetsClient_NewListPager_resourceGroupBudgetsList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBudgetsClient().NewListPager("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG", nil)
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
		// page = armcostmanagement.BudgetsClientListResponse{
		// 	BudgetsListResult: armcostmanagement.BudgetsListResult{
		// 		Value: []*armcostmanagement.Budget{
		// 			{
		// 				Name: to.Ptr("TestBudget0"),
		// 				Type: to.Ptr("Microsoft.CostManagement/budgets"),
		// 				ETag: to.Ptr(azcore.ETag("\"1d34d012214157f\"")),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.CostManagement/budgets/TestBudget0"),
		// 				Properties: &armcostmanagement.BudgetProperties{
		// 					Amount: to.Ptr[float32](100.65),
		// 					Category: to.Ptr(armcostmanagement.CategoryTypeCost),
		// 					CurrentSpend: &armcostmanagement.CurrentSpend{
		// 						Amount: to.Ptr[float32](80.89),
		// 						Unit: to.Ptr("USD"),
		// 					},
		// 					Filter: &armcostmanagement.BudgetFilter{
		// 						And: []*armcostmanagement.BudgetFilterProperties{
		// 							{
		// 								Dimensions: &armcostmanagement.BudgetComparisonExpression{
		// 									Name: to.Ptr("ResourceId"),
		// 									Operator: to.Ptr(armcostmanagement.BudgetOperatorTypeIn),
		// 									Values: []*string{
		// 										to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/MSVM2"),
		// 										to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/platformcloudplatformGeneric1"),
		// 									},
		// 								},
		// 							},
		// 							{
		// 								Tags: &armcostmanagement.BudgetComparisonExpression{
		// 									Name: to.Ptr("category"),
		// 									Operator: to.Ptr(armcostmanagement.BudgetOperatorTypeIn),
		// 									Values: []*string{
		// 										to.Ptr("Dev"),
		// 										to.Ptr("Prod"),
		// 									},
		// 								},
		// 							},
		// 							{
		// 								Tags: &armcostmanagement.BudgetComparisonExpression{
		// 									Name: to.Ptr("department"),
		// 									Operator: to.Ptr(armcostmanagement.BudgetOperatorTypeIn),
		// 									Values: []*string{
		// 										to.Ptr("engineering"),
		// 										to.Ptr("sales"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 					Notifications: map[string]*armcostmanagement.Notification{
		// 						"Actual_GreaterThanOrEqualTo_90_Percent": &armcostmanagement.Notification{
		// 							ContactEmails: []*string{
		// 								to.Ptr("johndoe@contoso.com"),
		// 								to.Ptr("janesmith@contoso.com"),
		// 							},
		// 							ContactGroups: []*string{
		// 								to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/microsoft.insights/actionGroups/SampleActionGroup"),
		// 							},
		// 							ContactRoles: []*string{
		// 								to.Ptr("Contributor"),
		// 								to.Ptr("Reader"),
		// 							},
		// 							Enabled: to.Ptr(true),
		// 							Operator: to.Ptr(armcostmanagement.BudgetNotificationOperatorTypeGreaterThanOrEqualTo),
		// 							Threshold: to.Ptr[float32](90),
		// 							ThresholdType: to.Ptr(armcostmanagement.ThresholdTypeActual),
		// 						},
		// 						"Actual_GreaterThan_80_Percent": &armcostmanagement.Notification{
		// 							ContactEmails: []*string{
		// 								to.Ptr("johndoe@contoso.com"),
		// 								to.Ptr("janesmith@contoso.com"),
		// 							},
		// 							ContactRoles: []*string{
		// 								to.Ptr("Contributor"),
		// 								to.Ptr("Reader"),
		// 							},
		// 							Enabled: to.Ptr(true),
		// 							Operator: to.Ptr(armcostmanagement.BudgetNotificationOperatorTypeGreaterThan),
		// 							Threshold: to.Ptr[float32](80),
		// 						},
		// 					},
		// 					TimeGrain: to.Ptr(armcostmanagement.TimeGrainTypeMonthly),
		// 					TimePeriod: &armcostmanagement.BudgetTimePeriod{
		// 						EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-31T00:00:00Z"); return t}()),
		// 						StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-01T00:00:00Z"); return t}()),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("TestBudget1"),
		// 				Type: to.Ptr("Microsoft.CostManagement/budgets"),
		// 				ETag: to.Ptr(azcore.ETag("\"1d34d012214157f\"")),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.CostManagement/budgets/TestBudget1"),
		// 				Properties: &armcostmanagement.BudgetProperties{
		// 					Amount: to.Ptr[float32](200.65),
		// 					Category: to.Ptr(armcostmanagement.CategoryTypeCost),
		// 					CurrentSpend: &armcostmanagement.CurrentSpend{
		// 						Amount: to.Ptr[float32](120.89),
		// 						Unit: to.Ptr("USD"),
		// 					},
		// 					Filter: &armcostmanagement.BudgetFilter{
		// 					},
		// 					Notifications: map[string]*armcostmanagement.Notification{
		// 						"Actual_GreaterThanOrEqualTo_60_Percent": &armcostmanagement.Notification{
		// 							ContactEmails: []*string{
		// 								to.Ptr("johndoe@contoso.com"),
		// 								to.Ptr("janesmith@contoso.com"),
		// 							},
		// 							ContactGroups: []*string{
		// 								to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/microsoft.insights/actionGroups/SampleActionGroup"),
		// 							},
		// 							ContactRoles: []*string{
		// 								to.Ptr("Contributor"),
		// 								to.Ptr("Reader"),
		// 							},
		// 							Enabled: to.Ptr(true),
		// 							Operator: to.Ptr(armcostmanagement.BudgetNotificationOperatorTypeGreaterThanOrEqualTo),
		// 							Threshold: to.Ptr[float32](60),
		// 							ThresholdType: to.Ptr(armcostmanagement.ThresholdTypeActual),
		// 						},
		// 						"Actual_GreaterThan_40_Percent": &armcostmanagement.Notification{
		// 							ContactEmails: []*string{
		// 								to.Ptr("johndoe@contoso.com"),
		// 								to.Ptr("janesmith@contoso.com"),
		// 							},
		// 							ContactRoles: []*string{
		// 								to.Ptr("Contributor"),
		// 								to.Ptr("Reader"),
		// 							},
		// 							Enabled: to.Ptr(true),
		// 							Operator: to.Ptr(armcostmanagement.BudgetNotificationOperatorTypeGreaterThan),
		// 							Threshold: to.Ptr[float32](40),
		// 						},
		// 					},
		// 					TimeGrain: to.Ptr(armcostmanagement.TimeGrainTypeMonthly),
		// 					TimePeriod: &armcostmanagement.BudgetTimePeriod{
		// 						EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-31T00:00:00Z"); return t}()),
		// 						StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-01T00:00:00Z"); return t}()),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
