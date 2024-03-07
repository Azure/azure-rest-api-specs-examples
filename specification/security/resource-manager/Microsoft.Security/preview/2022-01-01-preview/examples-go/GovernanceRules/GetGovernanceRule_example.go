package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/GetGovernanceRule_example.json
func ExampleGovernanceRulesClient_Get_getAGovernanceRuleOverSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGovernanceRulesClient().Get(ctx, "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23", "ad9a8e26-29d9-4829-bb30-e597a58cdbb8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GovernanceRule = armsecurity.GovernanceRule{
	// 	Name: to.Ptr("ad9a8e26-29d9-4829-bb30-e597a58cdbb8"),
	// 	Type: to.Ptr("Microsoft.Security/governanceRules"),
	// 	ID: to.Ptr("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/governanceRules/ad9a8e26-29d9-4829-bb30-e597a58cdbb8"),
	// 	Properties: &armsecurity.GovernanceRuleProperties{
	// 		Description: to.Ptr("A rule for critical recommendations"),
	// 		ConditionSets: []any{
	// 			map[string]any{
	// 				"conditions":[]any{
	// 					map[string]any{
	// 						"operator": "In",
	// 						"property": "$.AssessmentKey",
	// 						"value": "[\"b1cd27e0-4ecc-4246-939f-49c426d9d72f\", \"fe83f80b-073d-4ccf-93d9-6797eb870201\"]",
	// 					},
	// 				},
	// 		}},
	// 		DisplayName: to.Ptr("Admin's rule"),
	// 		ExcludedScopes: []*string{
	// 		},
	// 		GovernanceEmailNotification: &armsecurity.GovernanceRuleEmailNotification{
	// 			DisableManagerEmailNotification: to.Ptr(false),
	// 			DisableOwnerEmailNotification: to.Ptr(false),
	// 		},
	// 		IncludeMemberScopes: to.Ptr(false),
	// 		IsDisabled: to.Ptr(false),
	// 		IsGracePeriod: to.Ptr(true),
	// 		Metadata: &armsecurity.GovernanceRuleMetadata{
	// 			CreatedBy: to.Ptr("c23b5354-ff0a-4b2a-9f92-6f144effd936"),
	// 			CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-10T08:31:26.799Z"); return t}()),
	// 			UpdatedBy: to.Ptr("c23b5354-ff0a-4b2a-9f92-6f144effd936"),
	// 			UpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-10T08:31:26.799Z"); return t}()),
	// 		},
	// 		OwnerSource: &armsecurity.GovernanceRuleOwnerSource{
	// 			Type: to.Ptr(armsecurity.GovernanceRuleOwnerSourceTypeManually),
	// 			Value: to.Ptr("user@contoso.com"),
	// 		},
	// 		RemediationTimeframe: to.Ptr("7.00:00:00"),
	// 		RulePriority: to.Ptr[int32](200),
	// 		RuleType: to.Ptr(armsecurity.GovernanceRuleTypeIntegrated),
	// 		SourceResourceType: to.Ptr(armsecurity.GovernanceRuleSourceResourceTypeAssessments),
	// 		TenantID: to.Ptr("f0b6d37b-e4bc-4719-9291-c066c3194f23"),
	// 	},
	// }
}
