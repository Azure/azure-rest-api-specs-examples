package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy/v2"
)

// Generated from example definition: 2026-01-01-preview/listPolicyExemptionsForSubscription.json
func ExampleExemptionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("ae640e6b-ba3e-4256-9d62-2993eecfa6f2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExemptionsClient().NewListPager(&armpolicy.ExemptionsClientListOptions{
		Filter: to.Ptr("atScope()")})
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
		// page = armpolicy.ExemptionsClientListResponse{
		// 	ExemptionListResult: armpolicy.ExemptionListResult{
		// 		Value: []*armpolicy.Exemption{
		// 			{
		// 				Properties: &armpolicy.ExemptionProperties{
		// 					PolicyAssignmentID: to.Ptr("/providers/Microsoft.Management/managementGroups/DevOrg/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
		// 					PolicyDefinitionReferenceIDs: []*string{
		// 						to.Ptr("Limit_Skus"),
		// 					},
		// 					ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryWaiver),
		// 					DisplayName: to.Ptr("Exempt demo cluster"),
		// 					Description: to.Ptr("Exempt demo cluster from limit sku"),
		// 					Metadata: map[string]any{
		// 						"reason": "Temporary exemption for a expensive VM demo",
		// 					},
		// 				},
		// 				SystemData: &armpolicy.SystemData{
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.1075056Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T02:01:01.1075056Z"); return t}()),
		// 				},
		// 				ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyExemptions/TestVMSub"),
		// 				Type: to.Ptr("Microsoft.Authorization/policyExemptions"),
		// 				Name: to.Ptr("TestVMSub"),
		// 			},
		// 			{
		// 				Properties: &armpolicy.ExemptionProperties{
		// 					PolicyAssignmentID: to.Ptr("/providers/Microsoft.Management/managementGroups/DevOrg/providers/Microsoft.Authorization/policyAssignments/LimitPorts"),
		// 					ExemptionCategory: to.Ptr(armpolicy.ExemptionCategoryMitigated),
		// 					DisplayName: to.Ptr("Exempt jump box open ports"),
		// 					Description: to.Ptr("Exempt jump box open ports from limit ports policy"),
		// 					Metadata: map[string]any{
		// 						"reason": "Need to open RDP port to corp net",
		// 					},
		// 				},
		// 				SystemData: &armpolicy.SystemData{
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T01:01:01.1075056Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-01T02:01:01.1075056Z"); return t}()),
		// 				},
		// 				ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyExemptions/TestVNetSub"),
		// 				Type: to.Ptr("Microsoft.Authorization/policyExemptions"),
		// 				Name: to.Ptr("TestVNetSub"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
