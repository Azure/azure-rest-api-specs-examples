package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/AlertsSuppressionRules/GetAlertsSuppressionRule_example.json
func ExampleAlertsSuppressionRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertsSuppressionRulesClient().Get(ctx, "dismissIpAnomalyAlerts", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AlertsSuppressionRule = armsecurity.AlertsSuppressionRule{
	// 	Name: to.Ptr("dismissIpAnomalyAlerts"),
	// 	Type: to.Ptr("Microsoft.Security/alertsSuppressionRules"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/alertsSuppressionRules/dismissIpAnomalyAlerts"),
	// 	Properties: &armsecurity.AlertsSuppressionRuleProperties{
	// 		AlertType: to.Ptr("IpAnomaly"),
	// 		Comment: to.Ptr("Test VM"),
	// 		ExpirationDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-01T19:50:47.083Z"); return t}()),
	// 		LastModifiedUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-31T19:50:47.083Z"); return t}()),
	// 		Reason: to.Ptr("FalsePositive"),
	// 		State: to.Ptr(armsecurity.RuleStateEnabled),
	// 		SuppressionAlertsScope: &armsecurity.SuppressionAlertsScope{
	// 			AllOf: []*armsecurity.ScopeElement{
	// 				{
	// 					AdditionalProperties: map[string]any{
	// 						"in": []any{
	// 							"104.215.95.187",
	// 							"52.164.206.56",
	// 						},
	// 					},
	// 					Field: to.Ptr("entities.ip.address"),
	// 				},
	// 				{
	// 					AdditionalProperties: map[string]any{
	// 						"contains": "POWERSHELL.EXE",
	// 					},
	// 					Field: to.Ptr("entities.process.commandline"),
	// 			}},
	// 		},
	// 	},
	// }
}
