package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6d2438481021a94793b07b226df06d5f3c61d51d/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/AlertProcessingRules_Create_or_update_add_action_group_all_alerts_in_subscription.json
func ExampleAlertProcessingRulesClient_CreateOrUpdate_createOrUpdateARuleThatAddsAnActionGroupToAllAlertsInASubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertProcessingRulesClient().CreateOrUpdate(ctx, "alertscorrelationrg", "AddActionGroupToSubscription", armalertsmanagement.AlertProcessingRule{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AlertProcessingRule = armalertsmanagement.AlertProcessingRule{
	// 	Name: to.Ptr("AddActionGroupToSubscription"),
	// 	Type: to.Ptr("Microsoft.AlertsManagement/actionRules"),
	// 	ID: to.Ptr("/subscriptions/subId1/resourceGroups/alertscorrelationrg/providers/Microsoft.AlertsManagement/actionRules/AddActionGroupToSubscription"),
	// 	Location: to.Ptr("Global"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armalertsmanagement.AlertProcessingRuleProperties{
	// 		Description: to.Ptr("Add ActionGroup1 to all alerts in the subscription"),
	// 		Actions: []armalertsmanagement.ActionClassification{
	// 			&armalertsmanagement.AddActionGroups{
	// 				ActionType: to.Ptr(armalertsmanagement.ActionTypeAddActionGroups),
	// 				ActionGroupIDs: []*string{
	// 					to.Ptr("/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/ActionGroup1")},
	// 			}},
	// 			Enabled: to.Ptr(true),
	// 			Scopes: []*string{
	// 				to.Ptr("/subscriptions/subId1")},
	// 			},
	// 			SystemData: &armalertsmanagement.SystemData{
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-12T22:05:09.000Z"); return t}()),
	// 				CreatedBy: to.Ptr("abc@microsoft.com"),
	// 				CreatedByType: to.Ptr(armalertsmanagement.CreatedByTypeUser),
	// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-13T16:15:34.000Z"); return t}()),
	// 				LastModifiedBy: to.Ptr("xyz@microsoft.com"),
	// 				LastModifiedByType: to.Ptr(armalertsmanagement.CreatedByTypeUser),
	// 			},
	// 		}
}
