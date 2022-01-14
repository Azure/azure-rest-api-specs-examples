Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Falertsmanagement%2Farmalertsmanagement%2Fv0.3.0/sdk/resourcemanager/alertsmanagement/armalertsmanagement/README.md) on how to add the SDK to your project and authenticate.

```go
package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// x-ms-original-file: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/AlertProcessingRules_Create_or_update_add_action_group_all_alerts_in_subscription.json
func ExampleAlertProcessingRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armalertsmanagement.NewAlertProcessingRulesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<alert-processing-rule-name>",
		armalertsmanagement.AlertProcessingRule{
			Location: to.StringPtr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armalertsmanagement.AlertProcessingRuleProperties{
				Description: to.StringPtr("<description>"),
				Actions: []armalertsmanagement.ActionClassification{
					&armalertsmanagement.AddActionGroups{
						ActionType: armalertsmanagement.ActionType("AddActionGroups").ToPtr(),
						ActionGroupIDs: []*string{
							to.StringPtr("/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/ActionGroup1")},
					}},
				Enabled: to.BoolPtr(true),
				Scopes: []*string{
					to.StringPtr("/subscriptions/subId1")},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AlertProcessingRulesClientCreateOrUpdateResult)
}
```
