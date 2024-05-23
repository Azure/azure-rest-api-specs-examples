package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/ListBySecurityConnectorGovernanceRules_example.json
func ExampleGovernanceRulesClient_NewListPager_listGovernanceRulesBySecurityConnectorScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGovernanceRulesClient().NewListPager("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/gcpResourceGroup/providers/Microsoft.Security/securityConnectors/gcpconnector", nil)
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
		// page.GovernanceRuleList = armsecurity.GovernanceRuleList{
		// 	Value: []*armsecurity.GovernanceRule{
		// 		{
		// 			Name: to.Ptr("ad9a8e26-29d9-4829-bb30-e597a58cdbb8"),
		// 			Type: to.Ptr("Microsoft.Security/governanceRules"),
		// 			ID: to.Ptr("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/gcpResourceGroup/providers/Microsoft.Security/securityConnectors/gcpconnector/providers/Microsoft.Security/governanceRules/ad9a8e26-29d9-4829-bb30-e597a58cdbb8"),
		// 			Properties: &armsecurity.GovernanceRuleProperties{
		// 				Description: to.Ptr("A rule on critical GCP recommendations"),
		// 				ConditionSets: []any{
		// 					map[string]any{
		// 						"conditions":[]any{
		// 							map[string]any{
		// 								"operator": "In",
		// 								"property": "$.AssessmentKey",
		// 								"value": "[\"b1cd27e0-4ecc-4246-939f-49c426d9d72f\", \"fe83f80b-073d-4ccf-93d9-6797eb870201\"]",
		// 							},
		// 						},
		// 				}},
		// 				DisplayName: to.Ptr("Admin's GCP rule"),
		// 				ExcludedScopes: []*string{
		// 				},
		// 				GovernanceEmailNotification: &armsecurity.GovernanceRuleEmailNotification{
		// 					DisableManagerEmailNotification: to.Ptr(false),
		// 					DisableOwnerEmailNotification: to.Ptr(false),
		// 				},
		// 				IncludeMemberScopes: to.Ptr(false),
		// 				IsDisabled: to.Ptr(false),
		// 				IsGracePeriod: to.Ptr(true),
		// 				Metadata: &armsecurity.GovernanceRuleMetadata{
		// 					CreatedBy: to.Ptr("c23b5354-ff0a-4b2a-9f92-6f144effd936"),
		// 					CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-10T08:31:26.799Z"); return t}()),
		// 					UpdatedBy: to.Ptr("c23b5354-ff0a-4b2a-9f92-6f144effd936"),
		// 					UpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-10T08:31:26.799Z"); return t}()),
		// 				},
		// 				OwnerSource: &armsecurity.GovernanceRuleOwnerSource{
		// 					Type: to.Ptr(armsecurity.GovernanceRuleOwnerSourceTypeManually),
		// 					Value: to.Ptr("user@contoso.com"),
		// 				},
		// 				RemediationTimeframe: to.Ptr("7.00:00:00"),
		// 				RulePriority: to.Ptr[int32](100),
		// 				RuleType: to.Ptr(armsecurity.GovernanceRuleTypeIntegrated),
		// 				SourceResourceType: to.Ptr(armsecurity.GovernanceRuleSourceResourceTypeAssessments),
		// 				TenantID: to.Ptr("f0b6d37b-e4bc-4719-9291-c066c3194f23"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("4106f43c-6d82-4fc8-a92c-dcfe50799d1d"),
		// 			Type: to.Ptr("Microsoft.Security/governanceRules"),
		// 			ID: to.Ptr("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/gcpResourceGroup/providers/Microsoft.Security/securityConnectors/gcpconnector/providers/Microsoft.Security/governanceRules/4106f43c-6d82-4fc8-a92c-dcfe50799d1d"),
		// 			Properties: &armsecurity.GovernanceRuleProperties{
		// 				Description: to.Ptr("A rule on critical GCP recommendations"),
		// 				ConditionSets: []any{
		// 					map[string]any{
		// 						"conditions":[]any{
		// 							map[string]any{
		// 								"operator": "Equals",
		// 								"property": "$.Metadata.Severity",
		// 								"value": "Low",
		// 							},
		// 						},
		// 				}},
		// 				DisplayName: to.Ptr("GCP Admin's rule"),
		// 				ExcludedScopes: []*string{
		// 				},
		// 				GovernanceEmailNotification: &armsecurity.GovernanceRuleEmailNotification{
		// 					DisableManagerEmailNotification: to.Ptr(false),
		// 					DisableOwnerEmailNotification: to.Ptr(false),
		// 				},
		// 				IncludeMemberScopes: to.Ptr(false),
		// 				IsDisabled: to.Ptr(false),
		// 				IsGracePeriod: to.Ptr(true),
		// 				Metadata: &armsecurity.GovernanceRuleMetadata{
		// 					CreatedBy: to.Ptr("c23b5354-ff0a-4b2a-9f92-6f144effd936"),
		// 					CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-10T08:31:26.799Z"); return t}()),
		// 					UpdatedBy: to.Ptr("c23b5354-ff0a-4b2a-9f92-6f144effd936"),
		// 					UpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-10T08:31:26.799Z"); return t}()),
		// 				},
		// 				OwnerSource: &armsecurity.GovernanceRuleOwnerSource{
		// 					Type: to.Ptr(armsecurity.GovernanceRuleOwnerSourceTypeManually),
		// 					Value: to.Ptr("user@contoso.com"),
		// 				},
		// 				RemediationTimeframe: to.Ptr("7.00:00:00"),
		// 				RulePriority: to.Ptr[int32](200),
		// 				RuleType: to.Ptr(armsecurity.GovernanceRuleTypeIntegrated),
		// 				SourceResourceType: to.Ptr(armsecurity.GovernanceRuleSourceResourceTypeAssessments),
		// 				TenantID: to.Ptr("f0b6d37b-e4bc-4719-9291-c066c3194f23"),
		// 			},
		// 	}},
		// }
	}
}
