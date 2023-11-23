package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/preview/2022-07-01-preview/examples/listPolicyExemptionsForManagementGroup.json
func ExampleExemptionsClient_NewListForManagementGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExemptionsClient().NewListForManagementGroupPager("DevOrg", &armpolicy.ExemptionsClientListForManagementGroupOptions{Filter: to.Ptr("atScope()")})
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
		// page.ExemptionListResult = armpolicy.ExemptionListResult{
		// 	Value: []*armpolicy.Exemption{
		// 		{
		// 			Name: to.Ptr("ResearchBudgetExemption"),
		// 			Type: to.Ptr("Microsoft.Authorization/policyExemptions"),
		// 			ID: to.Ptr("/providers/Microsoft.Management/managementGroups/DevOrg/providers/Microsoft.Authorization/policyExemptions/ResearchBudgetExemption"),
		// 			Properties: &armpolicy.ExemptionProperties{
		// 				Description: to.Ptr("Exempt demo cluster from limit sku"),
		// 				DisplayName: to.Ptr("Exempt demo cluster"),
		// 				ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryWaiver),
		// 				Metadata: map[string]any{
		// 					"reason": "Temporary exemption for a expensive VM demo",
		// 				},
		// 				PolicyAssignmentID: to.Ptr("/providers/Microsoft.Management/managementGroups/HardwareDivision/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
		// 				PolicyDefinitionReferenceIDs: []*string{
		// 					to.Ptr("Limit_Skus")},
		// 				},
		// 				SystemData: &armpolicy.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.107Z"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T02:01:01.107Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("VNetIsMonitored"),
		// 				Type: to.Ptr("Microsoft.Authorization/policyExemptions"),
		// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/DevOrg/providers/Microsoft.Authorization/policyExemptions/VNetIsMonitored"),
		// 				Properties: &armpolicy.ExemptionProperties{
		// 					Description: to.Ptr("Exempt jump box open ports from limit ports policy"),
		// 					DisplayName: to.Ptr("Exempt jump box open ports"),
		// 					ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryMitigated),
		// 					Metadata: map[string]any{
		// 						"reason": "Need to open RDP port to corp net",
		// 					},
		// 					PolicyAssignmentID: to.Ptr("/providers/Microsoft.Management/managementGroups/HardwareDivision/providers/Microsoft.Authorization/policyAssignments/LimitPorts"),
		// 				},
		// 				SystemData: &armpolicy.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.107Z"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T02:01:01.107Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 				},
		// 		}},
		// 	}
	}
}
