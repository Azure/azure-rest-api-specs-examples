Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Falertsmanagement%2Farmalertsmanagement%2Fv0.6.0/sdk/resourcemanager/alertsmanagement/armalertsmanagement/README.md) on how to add the SDK to your project and authenticate.

```go
package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/AlertProcessingRules_Create_or_update_add_action_group_all_alerts_in_subscription.json
func ExampleAlertProcessingRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armalertsmanagement.NewAlertProcessingRulesClient("subId1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"alertscorrelationrg",
		"AddActionGroupToSubscription",
		armalertsmanagement.AlertProcessingRule{
			Location: to.Ptr("Global"),
			Tags:     map[string]*string{},
			Properties: &armalertsmanagement.AlertProcessingRuleProperties{
				Description: to.Ptr("Add ActionGroup1 to all alerts in the subscription"),
				Actions: []armalertsmanagement.ActionClassification{
					&armalertsmanagement.AddActionGroups{
						ActionType: to.Ptr(armalertsmanagement.ActionTypeAddActionGroups),
						ActionGroupIDs: []*string{
							to.Ptr("/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/ActionGroup1")},
					}},
				Enabled: to.Ptr(true),
				Scopes: []*string{
					to.Ptr("/subscriptions/subId1")},
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
