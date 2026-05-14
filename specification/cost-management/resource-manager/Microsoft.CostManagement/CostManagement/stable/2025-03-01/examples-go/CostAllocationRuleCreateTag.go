package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v3"
)

// Generated from example definition: 2025-03-01/CostAllocationRuleCreateTag.json
func ExampleCostAllocationRulesClient_CreateOrUpdate_costAllocationRulesCreateTag() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCostAllocationRulesClient().CreateOrUpdate(ctx, "100", "testRule", armcostmanagement.CostAllocationRuleDefinition{
		Properties: &armcostmanagement.CostAllocationRuleProperties{
			Description: to.Ptr("This is a testRule"),
			Status:      to.Ptr(armcostmanagement.RuleStatusActive),
			Details: &armcostmanagement.CostAllocationRuleDetails{
				SourceResources: []*armcostmanagement.SourceCostAllocationResource{
					{
						Name:         to.Ptr("category"),
						ResourceType: to.Ptr(armcostmanagement.CostAllocationResourceTypeTag),
						Values: []*string{
							to.Ptr("devops"),
						},
					},
				},
				TargetResources: []*armcostmanagement.TargetCostAllocationResource{
					{
						Name:         to.Ptr("ResourceGroupName"),
						PolicyType:   to.Ptr(armcostmanagement.CostAllocationPolicyTypeFixedProportion),
						ResourceType: to.Ptr(armcostmanagement.CostAllocationResourceTypeDimension),
						Values: []*armcostmanagement.CostAllocationProportion{
							{
								Name:       to.Ptr("destinationRG"),
								Percentage: to.Ptr[float32](33.33),
							},
							{
								Name:       to.Ptr("destinationRG2"),
								Percentage: to.Ptr[float32](33.33),
							},
							{
								Name:       to.Ptr("destinationRG3"),
								Percentage: to.Ptr[float32](33.34),
							},
						},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcostmanagement.CostAllocationRulesClientCreateOrUpdateResponse{
	// 	CostAllocationRuleDefinition: armcostmanagement.CostAllocationRuleDefinition{
	// 		Name: to.Ptr("testRule"),
	// 		Type: to.Ptr("Microsoft.CostManagement/costAllocationRules"),
	// 		ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/100/providers/Microsoft.CostManagement/costAllocationRules/testRule"),
	// 		Properties: &armcostmanagement.CostAllocationRuleProperties{
	// 			Description: to.Ptr("This is a testRule"),
	// 			CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-18T22:21:51.1287144Z"); return t}()),
	// 			Status: to.Ptr(armcostmanagement.RuleStatus("Creating")),
	// 			UpdatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-18T22:21:51.1287144Z"); return t}()),
	// 			Details: &armcostmanagement.CostAllocationRuleDetails{
	// 				SourceResources: []*armcostmanagement.SourceCostAllocationResource{
	// 					{
	// 						Name: to.Ptr("ResourceGroupName"),
	// 						ResourceType: to.Ptr(armcostmanagement.CostAllocationResourceTypeDimension),
	// 						Values: []*string{
	// 							to.Ptr("sampleRG"),
	// 						},
	// 					},
	// 				},
	// 				TargetResources: []*armcostmanagement.TargetCostAllocationResource{
	// 					{
	// 						Name: to.Ptr("ResourceGroupName"),
	// 						PolicyType: to.Ptr(armcostmanagement.CostAllocationPolicyTypeFixedProportion),
	// 						ResourceType: to.Ptr(armcostmanagement.CostAllocationResourceTypeDimension),
	// 						Values: []*armcostmanagement.CostAllocationProportion{
	// 							{
	// 								Name: to.Ptr("destinationRG"),
	// 								Percentage: to.Ptr[float32](50),
	// 							},
	// 							{
	// 								Name: to.Ptr("destinationRG2"),
	// 								Percentage: to.Ptr[float32](50),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
