package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v3"
)

// Generated from example definition: 2025-03-01/CostAllocationRulesList.json
func ExampleCostAllocationRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCostAllocationRulesClient().NewListPager("100", nil)
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
		// page = armcostmanagement.CostAllocationRulesClientListResponse{
		// 	CostAllocationRuleList: armcostmanagement.CostAllocationRuleList{
		// 		Value: []*armcostmanagement.CostAllocationRuleDefinition{
		// 			{
		// 				Name: to.Ptr("testRule"),
		// 				Type: to.Ptr("Microsoft.CostManagement/costAllocationRules"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/100/providers/Microsoft.CostManagement/costAllocationRules/testRule"),
		// 				Properties: &armcostmanagement.CostAllocationRuleProperties{
		// 					Description: to.Ptr("This is a testRule"),
		// 					CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-18T22:21:51.1287144Z"); return t}()),
		// 					Status: to.Ptr(armcostmanagement.RuleStatusNotActive),
		// 					UpdatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-18T22:21:51.1287144Z"); return t}()),
		// 					Details: &armcostmanagement.CostAllocationRuleDetails{
		// 						SourceResources: []*armcostmanagement.SourceCostAllocationResource{
		// 							{
		// 								Name: to.Ptr("ResourceGroupName"),
		// 								ResourceType: to.Ptr(armcostmanagement.CostAllocationResourceTypeDimension),
		// 								Values: []*string{
		// 									to.Ptr("sampleRG"),
		// 								},
		// 							},
		// 						},
		// 						TargetResources: []*armcostmanagement.TargetCostAllocationResource{
		// 							{
		// 								Name: to.Ptr("ResourceGroupName"),
		// 								PolicyType: to.Ptr(armcostmanagement.CostAllocationPolicyTypeFixedProportion),
		// 								ResourceType: to.Ptr(armcostmanagement.CostAllocationResourceTypeDimension),
		// 								Values: []*armcostmanagement.CostAllocationProportion{
		// 									{
		// 										Name: to.Ptr("destinationRG"),
		// 										Percentage: to.Ptr[float32](50),
		// 									},
		// 									{
		// 										Name: to.Ptr("destinationRG2"),
		// 										Percentage: to.Ptr[float32](50),
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("testRule2"),
		// 				Type: to.Ptr("Microsoft.CostManagement/costAllocationRules"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/100/providers/Microsoft.CostManagement/costAllocationRules/testRule2"),
		// 				Properties: &armcostmanagement.CostAllocationRuleProperties{
		// 					Description: to.Ptr("This is a second test Rule"),
		// 					CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-18T22:21:51.1287144Z"); return t}()),
		// 					Status: to.Ptr(armcostmanagement.RuleStatusActive),
		// 					UpdatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-18T22:21:51.1287144Z"); return t}()),
		// 					Details: &armcostmanagement.CostAllocationRuleDetails{
		// 						SourceResources: []*armcostmanagement.SourceCostAllocationResource{
		// 							{
		// 								Name: to.Ptr("SubscriptionId"),
		// 								ResourceType: to.Ptr(armcostmanagement.CostAllocationResourceTypeDimension),
		// 								Values: []*string{
		// 									to.Ptr("2A002F2D-536F-4D7C-90DA-3D0BAE879B0E"),
		// 								},
		// 							},
		// 						},
		// 						TargetResources: []*armcostmanagement.TargetCostAllocationResource{
		// 							{
		// 								Name: to.Ptr("category"),
		// 								PolicyType: to.Ptr(armcostmanagement.CostAllocationPolicyTypeFixedProportion),
		// 								ResourceType: to.Ptr(armcostmanagement.CostAllocationResourceTypeTag),
		// 								Values: []*armcostmanagement.CostAllocationProportion{
		// 									{
		// 										Name: to.Ptr("devops"),
		// 										Percentage: to.Ptr[float32](100),
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("testRule3"),
		// 				Type: to.Ptr("Microsoft.CostManagement/costAllocationRules"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/100/providers/Microsoft.CostManagement/costAllocationRules/testRule3"),
		// 				Properties: &armcostmanagement.CostAllocationRuleProperties{
		// 					Description: to.Ptr("This is a third test Rule"),
		// 					CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-18T22:21:51.1287144Z"); return t}()),
		// 					Status: to.Ptr(armcostmanagement.RuleStatusActive),
		// 					UpdatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-18T22:21:51.1287144Z"); return t}()),
		// 					Details: &armcostmanagement.CostAllocationRuleDetails{
		// 						SourceResources: []*armcostmanagement.SourceCostAllocationResource{
		// 							{
		// 								Name: to.Ptr("category"),
		// 								ResourceType: to.Ptr(armcostmanagement.CostAllocationResourceTypeTag),
		// 								Values: []*string{
		// 									to.Ptr("devops"),
		// 								},
		// 							},
		// 						},
		// 						TargetResources: []*armcostmanagement.TargetCostAllocationResource{
		// 							{
		// 								Name: to.Ptr("ResourceGroupName"),
		// 								PolicyType: to.Ptr(armcostmanagement.CostAllocationPolicyTypeFixedProportion),
		// 								ResourceType: to.Ptr(armcostmanagement.CostAllocationResourceTypeDimension),
		// 								Values: []*armcostmanagement.CostAllocationProportion{
		// 									{
		// 										Name: to.Ptr("ResourceGroup"),
		// 										Percentage: to.Ptr[float32](55.55),
		// 									},
		// 									{
		// 										Name: to.Ptr("ResourceGroupSecond"),
		// 										Percentage: to.Ptr[float32](44.45),
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
