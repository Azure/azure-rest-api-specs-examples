Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.7.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurity_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/AlertsSuppressionRules/PutAlertsSuppressionRule_example.json
func ExampleAlertsSuppressionRulesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewAlertsSuppressionRulesClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"dismissIpAnomalyAlerts",
		armsecurity.AlertsSuppressionRule{
			Properties: &armsecurity.AlertsSuppressionRuleProperties{
				AlertType:         to.Ptr("IpAnomaly"),
				Comment:           to.Ptr("Test VM"),
				ExpirationDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-01T19:50:47.083633Z"); return t }()),
				Reason:            to.Ptr("FalsePositive"),
				State:             to.Ptr(armsecurity.RuleStateEnabled),
				SuppressionAlertsScope: &armsecurity.SuppressionAlertsScope{
					AllOf: []*armsecurity.ScopeElement{
						{
							AdditionalProperties: map[string]interface{}{
								"in": []interface{}{
									"104.215.95.187",
									"52.164.206.56",
								},
							},
							Field: to.Ptr("entities.ip.address"),
						},
						{
							AdditionalProperties: map[string]interface{}{
								"contains": "POWERSHELL.EXE",
							},
							Field: to.Ptr("entities.process.commandline"),
						}},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
