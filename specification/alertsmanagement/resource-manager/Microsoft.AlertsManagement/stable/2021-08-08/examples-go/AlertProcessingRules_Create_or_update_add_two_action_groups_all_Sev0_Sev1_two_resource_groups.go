package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/AlertProcessingRules_Create_or_update_add_two_action_groups_all_Sev0_Sev1_two_resource_groups.json
func ExampleAlertProcessingRulesClient_CreateOrUpdate_createOrUpdateARuleThatAddsTwoActionGroupsToAllSev0AndSev1AlertsInTwoResourceGroups() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armalertsmanagement.NewAlertProcessingRulesClient("subId1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "alertscorrelationrg", "AddActionGroupsBySeverity", armalertsmanagement.AlertProcessingRule{
		Location: to.Ptr("Global"),
		Tags:     map[string]*string{},
		Properties: &armalertsmanagement.AlertProcessingRuleProperties{
			Description: to.Ptr("Add AGId1 and AGId2 to all Sev0 and Sev1 alerts in these resourceGroups"),
			Actions: []armalertsmanagement.ActionClassification{
				&armalertsmanagement.AddActionGroups{
					ActionType: to.Ptr(armalertsmanagement.ActionTypeAddActionGroups),
					ActionGroupIDs: []*string{
						to.Ptr("/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/AGId1"),
						to.Ptr("/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/AGId2")},
				}},
			Conditions: []*armalertsmanagement.Condition{
				{
					Field:    to.Ptr(armalertsmanagement.FieldSeverity),
					Operator: to.Ptr(armalertsmanagement.OperatorEquals),
					Values: []*string{
						to.Ptr("sev0"),
						to.Ptr("sev1")},
				}},
			Enabled: to.Ptr(true),
			Scopes: []*string{
				to.Ptr("/subscriptions/subId1/resourceGroups/RGId1"),
				to.Ptr("/subscriptions/subId1/resourceGroups/RGId2")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
