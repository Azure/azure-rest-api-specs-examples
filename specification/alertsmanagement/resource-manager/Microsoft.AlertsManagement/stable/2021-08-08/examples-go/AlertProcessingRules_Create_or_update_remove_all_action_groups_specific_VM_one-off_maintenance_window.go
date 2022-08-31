package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/AlertProcessingRules_Create_or_update_remove_all_action_groups_specific_VM_one-off_maintenance_window.json
func ExampleAlertProcessingRulesClient_CreateOrUpdate_createOrUpdateARuleThatRemovesAllActionGroupsFromAlertsOnASpecificVmDuringAOneOffMaintenanceWindow18002000AtASpecificDatePacificStandardTime() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armalertsmanagement.NewAlertProcessingRulesClient("subId1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "alertscorrelationrg", "RemoveActionGroupsMaintenanceWindow", armalertsmanagement.AlertProcessingRule{
		Location: to.Ptr("Global"),
		Tags:     map[string]*string{},
		Properties: &armalertsmanagement.AlertProcessingRuleProperties{
			Description: to.Ptr("Removes all ActionGroups from all Alerts on VMName during the maintenance window"),
			Actions: []armalertsmanagement.ActionClassification{
				&armalertsmanagement.RemoveAllActionGroups{
					ActionType: to.Ptr(armalertsmanagement.ActionTypeRemoveAllActionGroups),
				}},
			Enabled: to.Ptr(true),
			Schedule: &armalertsmanagement.Schedule{
				EffectiveFrom:  to.Ptr("2021-04-15T18:00:00"),
				EffectiveUntil: to.Ptr("2021-04-15T20:00:00"),
				TimeZone:       to.Ptr("Pacific Standard Time"),
			},
			Scopes: []*string{
				to.Ptr("/subscriptions/subId1/resourceGroups/RGId1/providers/Microsoft.Compute/virtualMachines/VMName")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
