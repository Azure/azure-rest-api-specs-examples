package armalertprocessingrules_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertprocessingrules/armalertprocessingrules"
)

// Generated from example definition: 2021-08-08/AlertProcessingRules_Create_or_update_add_action_group_all_alerts_in_subscription.json
func ExampleClient_CreateOrUpdate_createOrUpdateARuleThatAddsAnActionGroupToAllAlertsInASubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertprocessingrules.NewClientFactory("subId1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().CreateOrUpdate(ctx, "alertscorrelationrg", "AddActionGroupToSubscription", armalertprocessingrules.AlertProcessingRule{
		Location: to.Ptr("Global"),
		Properties: &armalertprocessingrules.AlertProcessingRuleProperties{
			Description: to.Ptr("Add ActionGroup1 to all alerts in the subscription"),
			Actions: []armalertprocessingrules.ActionClassification{
				&armalertprocessingrules.AddActionGroups{
					ActionGroupIDs: []*string{
						to.Ptr("/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/ActionGroup1"),
					},
					ActionType: to.Ptr(armalertprocessingrules.ActionTypeAddActionGroups),
				},
			},
			Enabled: to.Ptr(true),
			Scopes: []*string{
				to.Ptr("/subscriptions/subId1"),
			},
		},
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armalertprocessingrules.ClientCreateOrUpdateResponse{
	// 	AlertProcessingRule: &armalertprocessingrules.AlertProcessingRule{
	// 		Name: to.Ptr("AddActionGroupToSubscription"),
	// 		Type: to.Ptr("Microsoft.AlertsManagement/actionRules"),
	// 		ID: to.Ptr("/subscriptions/subId1/resourceGroups/alertscorrelationrg/providers/Microsoft.AlertsManagement/actionRules/AddActionGroupToSubscription"),
	// 		Location: to.Ptr("Global"),
	// 		Properties: &armalertprocessingrules.AlertProcessingRuleProperties{
	// 			Description: to.Ptr("Add ActionGroup1 to all alerts in the subscription"),
	// 			Actions: []armalertprocessingrules.ActionClassification{
	// 				&armalertprocessingrules.AddActionGroups{
	// 					ActionGroupIDs: []*string{
	// 						to.Ptr("/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/ActionGroup1"),
	// 					},
	// 					ActionType: to.Ptr(armalertprocessingrules.ActionTypeAddActionGroups),
	// 				},
	// 			},
	// 			Enabled: to.Ptr(true),
	// 			Scopes: []*string{
	// 				to.Ptr("/subscriptions/subId1"),
	// 			},
	// 		},
	// 		SystemData: &armalertprocessingrules.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-12T22:05:09Z"); return t}()),
	// 			CreatedBy: to.Ptr("abc@microsoft.com"),
	// 			CreatedByType: to.Ptr(armalertprocessingrules.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-13T16:15:34Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("xyz@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armalertprocessingrules.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
