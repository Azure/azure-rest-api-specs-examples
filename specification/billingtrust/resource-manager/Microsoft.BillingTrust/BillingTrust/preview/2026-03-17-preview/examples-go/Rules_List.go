package armbillingtrust_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingtrust/armbillingtrust"
)

// Generated from example definition: 2026-03-17-preview/Rules_List.json
func ExampleRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingtrust.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRulesClient().NewListPager("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/billing-edu-rg/providers/Microsoft.Program/educationEnrollments/default", nil)
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
		// page = armbillingtrust.RulesClientListResponse{
		// 	RuleListResult: armbillingtrust.RuleListResult{
		// 		Value: []*armbillingtrust.Rule{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/billing-edu-rg/providers/Microsoft.Program/educationEnrollments/default/providers/Microsoft.BillingTrust/assessments/default/rules/Qualify_Edu"),
		// 				Name: to.Ptr("Qualify_Edu"),
		// 				Type: to.Ptr("Microsoft.BillingTrust/assessments/rules"),
		// 				Properties: &armbillingtrust.EduQualificationRuleProperties{
		// 					Kind: to.Ptr(armbillingtrust.RuleKindEduQualification),
		// 					EvaluationState: to.Ptr(armbillingtrust.RuleStateRunning),
		// 					ProvisioningState: to.Ptr(armbillingtrust.ProvisioningStateSucceeded),
		// 					Domains: []*armbillingtrust.DomainEntry{
		// 						{
		// 							DomainNames: []*string{
		// 								to.Ptr("students.contoso.edu"),
		// 							},
		// 							TenantID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 							State: to.Ptr(armbillingtrust.DomainEntryStatePending),
		// 						},
		// 					},
		// 				},
		// 				SystemData: &armbillingtrust.SystemData{
		// 					CreatedBy: to.Ptr("edu-admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armbillingtrust.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-17T10:00:00.000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("edu-admin@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armbillingtrust.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-17T10:05:00.000Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
