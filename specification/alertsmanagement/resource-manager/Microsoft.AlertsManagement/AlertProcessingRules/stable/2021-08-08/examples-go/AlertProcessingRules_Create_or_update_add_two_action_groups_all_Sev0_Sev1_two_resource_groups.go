package armalertprocessingrules_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertprocessingrules/armalertprocessingrules"
)

// Generated from example definition: 2021-08-08/AlertProcessingRules_Create_or_update_add_two_action_groups_all_Sev0_Sev1_two_resource_groups.json
func ExampleClient_CreateOrUpdate_createOrUpdateARuleThatAddsTwoActionGroupsToAllSev0AndSev1AlertsInTwoResourceGroups() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertprocessingrules.NewClientFactory("subId1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().CreateOrUpdate(ctx, "alertscorrelationrg", "AddActionGroupsBySeverity", armalertprocessingrules.AlertProcessingRule{
		Location: to.Ptr("Global"),
		Properties: &armalertprocessingrules.AlertProcessingRuleProperties{
			Description: to.Ptr("Add AGId1 and AGId2 to all Sev0 and Sev1 alerts in these resourceGroups"),
			Actions: []armalertprocessingrules.ActionClassification{
				&armalertprocessingrules.AddActionGroups{
					ActionGroupIDs: []*string{
						to.Ptr("/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/AGId1"),
						to.Ptr("/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/AGId2"),
					},
					ActionType: to.Ptr(armalertprocessingrules.ActionTypeAddActionGroups),
				},
			},
			Conditions: []*armalertprocessingrules.Condition{
				{
					Field:    to.Ptr(armalertprocessingrules.FieldSeverity),
					Operator: to.Ptr(armalertprocessingrules.OperatorEquals),
					Values: []*string{
						to.Ptr("sev0"),
						to.Ptr("sev1"),
					},
				},
			},
			Enabled: to.Ptr(true),
			Scopes: []*string{
				to.Ptr("/subscriptions/subId1/resourceGroups/RGId1"),
				to.Ptr("/subscriptions/subId1/resourceGroups/RGId2"),
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
	// 		Name: to.Ptr("AddActionGroupsBySeverity"),
	// 		Type: to.Ptr("Microsoft.AlertsManagement/actionRules"),
	// 		ID: to.Ptr("/subscriptions/subId1/resourceGroups/alertscorrelationrg/providers/Microsoft.AlertsManagement/actionRules/AddActionGroupsBySeverity"),
	// 		Location: to.Ptr("Global"),
	// 		Properties: &armalertprocessingrules.AlertProcessingRuleProperties{
	// 			Description: to.Ptr("Add AGId1 and AGId2 to all Sev0 and Sev1 alerts in these resourceGroups"),
	// 			Actions: []armalertprocessingrules.ActionClassification{
	// 				&armalertprocessingrules.AddActionGroups{
	// 					ActionGroupIDs: []*string{
	// 						to.Ptr("/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/AGId1"),
	// 						to.Ptr("/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/AGId2"),
	// 					},
	// 					ActionType: to.Ptr(armalertprocessingrules.ActionTypeAddActionGroups),
	// 				},
	// 			},
	// 			Conditions: []*armalertprocessingrules.Condition{
	// 				{
	// 					Field: to.Ptr(armalertprocessingrules.FieldSeverity),
	// 					Operator: to.Ptr(armalertprocessingrules.OperatorEquals),
	// 					Values: []*string{
	// 						to.Ptr("sev0"),
	// 						to.Ptr("sev1"),
	// 					},
	// 				},
	// 			},
	// 			Enabled: to.Ptr(true),
	// 			Scopes: []*string{
	// 				to.Ptr("/subscriptions/subId1/resourceGroups/RGId1"),
	// 				to.Ptr("/subscriptions/subId1/resourceGroups/RGId2"),
	// 			},
	// 		},
	// 		SystemData: &armalertprocessingrules.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T22:05:09Z"); return t}()),
	// 			CreatedBy: to.Ptr("abc@microsoft.com"),
	// 			CreatedByType: to.Ptr(armalertprocessingrules.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-13T22:05:09Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("xyz@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armalertprocessingrules.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
