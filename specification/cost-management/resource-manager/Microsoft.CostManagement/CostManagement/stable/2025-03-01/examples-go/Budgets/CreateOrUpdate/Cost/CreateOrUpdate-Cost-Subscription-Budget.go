package armcostmanagement_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v3"
)

// Generated from example definition: 2025-03-01/Budgets/CreateOrUpdate/Cost/CreateOrUpdate-Cost-Subscription-Budget.json
func ExampleBudgetsClient_CreateOrUpdate_createOrUpdateCostSubscriptionBudget() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBudgetsClient().CreateOrUpdate(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", "TestBudget", armcostmanagement.Budget{
		ETag: to.Ptr(azcore.ETag("\"1d34d016a593709\"")),
		Properties: &armcostmanagement.BudgetProperties{
			Amount:   to.Ptr[float32](100.65),
			Category: to.Ptr(armcostmanagement.CategoryTypeCost),
			Filter: &armcostmanagement.BudgetFilter{
				And: []*armcostmanagement.BudgetFilterProperties{
					{
						Dimensions: &armcostmanagement.BudgetComparisonExpression{
							Name:     to.Ptr("ResourceId"),
							Operator: to.Ptr(armcostmanagement.BudgetOperatorTypeIn),
							Values: []*string{
								to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/MSVM2"),
								to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/platformcloudplatformGeneric1"),
							},
						},
					},
					{
						Tags: &armcostmanagement.BudgetComparisonExpression{
							Name:     to.Ptr("category"),
							Operator: to.Ptr(armcostmanagement.BudgetOperatorTypeIn),
							Values: []*string{
								to.Ptr("Dev"),
								to.Ptr("Prod"),
							},
						},
					},
					{
						Tags: &armcostmanagement.BudgetComparisonExpression{
							Name:     to.Ptr("department"),
							Operator: to.Ptr(armcostmanagement.BudgetOperatorTypeIn),
							Values: []*string{
								to.Ptr("engineering"),
								to.Ptr("sales"),
							},
						},
					},
				},
			},
			Notifications: map[string]*armcostmanagement.Notification{
				"Actual_GreaterThan_80_Percent": {
					ContactEmails: []*string{
						to.Ptr("johndoe@contoso.com"),
						to.Ptr("janesmith@contoso.com"),
					},
					ContactGroups: []*string{
						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/microsoft.insights/actionGroups/SampleActionGroup"),
					},
					ContactRoles: []*string{
						to.Ptr("Contributor"),
						to.Ptr("Reader"),
					},
					Enabled:       to.Ptr(true),
					Locale:        to.Ptr(armcostmanagement.CultureCodeEnUs),
					Operator:      to.Ptr(armcostmanagement.BudgetNotificationOperatorTypeGreaterThan),
					Threshold:     to.Ptr[float32](80),
					ThresholdType: to.Ptr(armcostmanagement.ThresholdTypeActual),
				},
			},
			TimeGrain: to.Ptr(armcostmanagement.TimeGrainTypeMonthly),
			TimePeriod: &armcostmanagement.BudgetTimePeriod{
				EndDate:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-31T00:00:00Z"); return t }()),
				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-01T00:00:00Z"); return t }()),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcostmanagement.BudgetsClientCreateOrUpdateResponse{
	// 	Budget: armcostmanagement.Budget{
	// 		Name: to.Ptr("TestBudget"),
	// 		Type: to.Ptr("Microsoft.CostManagement/budgets"),
	// 		ETag: to.Ptr(azcore.ETag("\"1d34d012214157f\"")),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CostManagement/budgets/TestBudget"),
	// 		Properties: &armcostmanagement.BudgetProperties{
	// 			Amount: to.Ptr[float32](100.65),
	// 			Category: to.Ptr(armcostmanagement.CategoryTypeCost),
	// 			CurrentSpend: &armcostmanagement.CurrentSpend{
	// 				Amount: to.Ptr[float32](80.89),
	// 				Unit: to.Ptr("USD"),
	// 			},
	// 			Filter: &armcostmanagement.BudgetFilter{
	// 				And: []*armcostmanagement.BudgetFilterProperties{
	// 					{
	// 						Dimensions: &armcostmanagement.BudgetComparisonExpression{
	// 							Name: to.Ptr("ResourceId"),
	// 							Operator: to.Ptr(armcostmanagement.BudgetOperatorTypeIn),
	// 							Values: []*string{
	// 								to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/MSVM2"),
	// 								to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/platformcloudplatformGeneric1"),
	// 							},
	// 						},
	// 					},
	// 					{
	// 						Tags: &armcostmanagement.BudgetComparisonExpression{
	// 							Name: to.Ptr("category"),
	// 							Operator: to.Ptr(armcostmanagement.BudgetOperatorTypeIn),
	// 							Values: []*string{
	// 								to.Ptr("Dev"),
	// 								to.Ptr("Prod"),
	// 							},
	// 						},
	// 					},
	// 					{
	// 						Tags: &armcostmanagement.BudgetComparisonExpression{
	// 							Name: to.Ptr("department"),
	// 							Operator: to.Ptr(armcostmanagement.BudgetOperatorTypeIn),
	// 							Values: []*string{
	// 								to.Ptr("engineering"),
	// 								to.Ptr("sales"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Notifications: map[string]*armcostmanagement.Notification{
	// 				"Actual_GreaterThan_80_Percent": &armcostmanagement.Notification{
	// 					ContactEmails: []*string{
	// 						to.Ptr("johndoe@contoso.com"),
	// 						to.Ptr("janesmith@contoso.com"),
	// 					},
	// 					ContactGroups: []*string{
	// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/microsoft.insights/actionGroups/SampleActionGroup"),
	// 					},
	// 					ContactRoles: []*string{
	// 						to.Ptr("Contributor"),
	// 						to.Ptr("Reader"),
	// 					},
	// 					Enabled: to.Ptr(true),
	// 					Locale: to.Ptr(armcostmanagement.CultureCodeEnUs),
	// 					Operator: to.Ptr(armcostmanagement.BudgetNotificationOperatorTypeGreaterThan),
	// 					Threshold: to.Ptr[float32](80),
	// 					ThresholdType: to.Ptr(armcostmanagement.ThresholdTypeActual),
	// 				},
	// 			},
	// 			TimeGrain: to.Ptr(armcostmanagement.TimeGrainTypeMonthly),
	// 			TimePeriod: &armcostmanagement.BudgetTimePeriod{
	// 				EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-31T00:00:00Z"); return t}()),
	// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-01T00:00:00Z"); return t}()),
	// 			},
	// 		},
	// 	},
	// }
}
