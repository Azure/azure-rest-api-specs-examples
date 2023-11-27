package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/Budget.json
func ExampleBudgetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBudgetsClient().Get(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", "TestBudget", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Budget = armconsumption.Budget{
	// 	Name: to.Ptr("TestBudget"),
	// 	Type: to.Ptr("Microsoft.Consumption/budgets"),
	// 	ETag: to.Ptr("\"1d34d012214157f\""),
	// 	ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Consumption/budgets/TestBudget"),
	// 	Properties: &armconsumption.BudgetProperties{
	// 		Amount: to.Ptr[float64](100.65),
	// 		Category: to.Ptr(armconsumption.CategoryTypeCost),
	// 		CurrentSpend: &armconsumption.CurrentSpend{
	// 			Amount: to.Ptr[float64](80.89),
	// 			Unit: to.Ptr("USD"),
	// 		},
	// 		Filter: &armconsumption.BudgetFilter{
	// 			And: []*armconsumption.BudgetFilterProperties{
	// 				{
	// 					Dimensions: &armconsumption.BudgetComparisonExpression{
	// 						Name: to.Ptr("ResourceId"),
	// 						Operator: to.Ptr(armconsumption.BudgetOperatorTypeIn),
	// 						Values: []*string{
	// 							to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/MSVM2"),
	// 							to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/platformcloudplatformGeneric1")},
	// 						},
	// 					},
	// 					{
	// 						Tags: &armconsumption.BudgetComparisonExpression{
	// 							Name: to.Ptr("category"),
	// 							Operator: to.Ptr(armconsumption.BudgetOperatorTypeIn),
	// 							Values: []*string{
	// 								to.Ptr("Dev"),
	// 								to.Ptr("Prod")},
	// 							},
	// 						},
	// 						{
	// 							Tags: &armconsumption.BudgetComparisonExpression{
	// 								Name: to.Ptr("department"),
	// 								Operator: to.Ptr(armconsumption.BudgetOperatorTypeIn),
	// 								Values: []*string{
	// 									to.Ptr("engineering"),
	// 									to.Ptr("sales")},
	// 								},
	// 						}},
	// 					},
	// 					Notifications: map[string]*armconsumption.Notification{
	// 						"Actual_GreaterThan_80_Percent": &armconsumption.Notification{
	// 							ContactEmails: []*string{
	// 								to.Ptr("johndoe@contoso.com"),
	// 								to.Ptr("janesmith@contoso.com")},
	// 								ContactGroups: []*string{
	// 									to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/microsoft.insights/actionGroups/SampleActionGroup")},
	// 									ContactRoles: []*string{
	// 										to.Ptr("Contributor"),
	// 										to.Ptr("Reader")},
	// 										Enabled: to.Ptr(true),
	// 										Operator: to.Ptr(armconsumption.OperatorTypeGreaterThan),
	// 										Threshold: to.Ptr[float64](80),
	// 										ThresholdType: to.Ptr(armconsumption.ThresholdTypeActual),
	// 									},
	// 								},
	// 								TimeGrain: to.Ptr(armconsumption.TimeGrainTypeMonthly),
	// 								TimePeriod: &armconsumption.BudgetTimePeriod{
	// 									EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-31T00:00:00.000Z"); return t}()),
	// 									StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-01T00:00:00.000Z"); return t}()),
	// 								},
	// 							},
	// 						}
}
