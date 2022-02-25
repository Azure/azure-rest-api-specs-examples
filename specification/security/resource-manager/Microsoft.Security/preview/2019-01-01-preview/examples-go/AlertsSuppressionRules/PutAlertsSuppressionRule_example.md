Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.4.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/AlertsSuppressionRules/PutAlertsSuppressionRule_example.json
func ExampleAlertsSuppressionRulesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewAlertsSuppressionRulesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<alerts-suppression-rule-name>",
		armsecurity.AlertsSuppressionRule{
			Properties: &armsecurity.AlertsSuppressionRuleProperties{
				AlertType:         to.StringPtr("<alert-type>"),
				Comment:           to.StringPtr("<comment>"),
				ExpirationDateUTC: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-01T19:50:47.083633Z"); return t }()),
				Reason:            to.StringPtr("<reason>"),
				State:             armsecurity.RuleStateEnabled.ToPtr(),
				SuppressionAlertsScope: &armsecurity.SuppressionAlertsScope{
					AllOf: []*armsecurity.ScopeElement{
						{
							AdditionalProperties: map[string]interface{}{
								"in": []interface{}{
									"104.215.95.187",
									"52.164.206.56",
								},
							},
							Field: to.StringPtr("<field>"),
						},
						{
							AdditionalProperties: map[string]interface{}{
								"contains": "POWERSHELL.EXE",
							},
							Field: to.StringPtr("<field>"),
						}},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AlertsSuppressionRulesClientUpdateResult)
}
```
