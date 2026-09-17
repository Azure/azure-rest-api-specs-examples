package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-10-01-preview/alertRules/GetMicrosoftSecurityIncidentCreationAlertRule.json
func ExampleAlertRulesClient_Get_getAMicrosoftSecurityIncidentCreationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertRulesClient().Get(ctx, "myRg", "myWorkspace", "microsoftSecurityIncidentCreationRuleExample", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.AlertRulesClientGetResponse{
	// 	AlertRuleClassification: &armsecurityinsights.MicrosoftSecurityIncidentCreationAlertRule{
	// 		Name: to.Ptr("microsoftSecurityIncidentCreationRuleExample"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/alertRules"),
	// 		Etag: to.Ptr("\"260097e0-0000-0d00-0000-5d6fa88f0000\""),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/microsoftSecurityIncidentCreationRuleExample"),
	// 		Kind: to.Ptr(armsecurityinsights.AlertRuleKindMicrosoftSecurityIncidentCreation),
	// 		Properties: &armsecurityinsights.MicrosoftSecurityIncidentCreationAlertRuleProperties{
	// 			DisplayName: to.Ptr("testing displayname"),
	// 			Enabled: to.Ptr(true),
	// 			LastModifiedUTC: to.Ptr(time.Date(2019, time.September, 4, 12, 5, 35, 729631100, time.UTC)),
	// 			ProductFilter: to.Ptr(armsecurityinsights.MicrosoftSecurityProductNameMicrosoftCloudAppSecurity),
	// 		},
	// 	},
	// }
}
