package armpolicyinsights_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: 2024-10-01/PolicyEvents_FilterAndGroupByWithAggregate.json
func ExamplePolicyEventsClient_NewListQueryResultsForSubscriptionPager_filterAndGroupWithAggregate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPolicyEventsClient().NewListQueryResultsForSubscriptionPager(armpolicyinsights.PolicyEventsResourceTypeDefault, "fffedd8f-ffff-fffd-fffd-fffed2f84852", &armpolicyinsights.PolicyEventsClientListQueryResultsForSubscriptionOptions{
		Apply:  to.Ptr("groupby((PolicyAssignmentId, PolicyDefinitionId, PolicyDefinitionAction, ResourceId), aggregate($count as NumEvents))"),
		Filter: to.Ptr("PolicyDefinitionAction eq 'audit' or PolicyDefinitionAction eq 'deny'"),
		From:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-05T18:00:00Z"); return t }()),
		Top:    to.Ptr[int32](2)})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armpolicyinsights.PolicyEventsClientListQueryResultsForSubscriptionResponse{
		// 	PolicyEventsQueryResults: armpolicyinsights.PolicyEventsQueryResults{
		// 		ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default"),
		// 		ODataCount: to.Ptr[int32](2),
		// 		Value: []*armpolicyinsights.PolicyEvent{
		// 			{
		// 				ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default/$entity"),
		// 				PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/3f3c4330183b4e218fe6fd29"),
		// 				PolicyDefinitionAction: to.Ptr("audit"),
		// 				PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/24813039-7534-408a-9842-eb99f45721b1"),
		// 				ResourceID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup/providers/microsoft.servicefabric/clusters/myCluster/applications/resourcescachemonitor/services/myService"),
		// 				AdditionalProperties: map[string]any{
		// 				"NumEvents": 1,
		// 			},
		// 			},
		// 			{
		// 				ODataContext: to.Ptr("https://management.azure.com/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/Microsoft.PolicyInsights/policyEvents/$metadata#default/$entity"),
		// 				PolicyAssignmentID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policyassignments/d6be6bb37e5f4333baa95c2a"),
		// 				PolicyDefinitionAction: to.Ptr("audit"),
		// 				PolicyDefinitionID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/providers/microsoft.authorization/policydefinitions/5948d091-78b7-4d3b-a404-cc6a0329b0c6"),
		// 				ResourceID: to.Ptr("/subscriptions/fffedd8f-ffff-fffd-fffd-fffed2f84852/resourcegroups/myResourceGroup/providers/microsoft.servicefabric/clusters/myCluster/applications/resourcescachemonitor/services/myService"),
		// 				AdditionalProperties: map[string]any{
		// 				"NumEvents": 1,
		// 			},
		// 			},
		// 		},
		// 	},
		// }
	}
}
