package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2024-03-01-preview/getMetricAlertQuery.json
func ExampleMetricAlertsClient_Get_getAQueryBasedAlertRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMetricAlertsClient().Get(ctx, "gigtest", "chiricutin", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmonitor.MetricAlertsClientGetResponse{
	// 	MetricAlertResource: armmonitor.MetricAlertResource{
	// 		Type: to.Ptr("Microsoft.Insights/metricAlerts"),
	// 		ID: to.Ptr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/providers/microsoft.insights/metricalerts/chiricutin"),
	// 		Identity: &armmonitor.Identity{
	// 			Type: to.Ptr(armmonitor.IdentityTypeUserAssigned),
	// 			TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 			UserAssignedIdentities: map[string]*armmonitor.UserIdentityProperties{
	// 				"/subscriptions/2f1a501a-6e1d-4f37-a445-462d7f8a563d/resourceGroups/AdisTest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/msi-test-euap": &armmonitor.UserIdentityProperties{
	// 					ClientID: to.Ptr("778f4d04-3c60-4622-a839-5cf05866c983"),
	// 					PrincipalID: to.Ptr("669dd76d-cde8-4dc3-b882-1de566b0c628"),
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armmonitor.MetricAlertProperties{
	// 			Description: to.Ptr("This is the description of the rule1"),
	// 			Actions: []*armmonitor.MetricAlertAction{
	// 				{
	// 					ActionGroupID: to.Ptr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2"),
	// 					WebHookProperties: map[string]*string{
	// 						"key11": to.Ptr("value11"),
	// 						"key12": to.Ptr("value12"),
	// 					},
	// 				},
	// 			},
	// 			Criteria: &armmonitor.PromQLCriteria{
	// 				AllOf: []armmonitor.MultiPromQLCriteriaClassification{
	// 					&armmonitor.StaticPromQLCriteria{
	// 						Name: to.Ptr("Metric1"),
	// 						CriterionType: to.Ptr(armmonitor.CriterionTypeStaticThresholdCriterion),
	// 						Query: to.Ptr("floor(minute(vector(time())) / 10 ) % 2 == 1"),
	// 					},
	// 				},
	// 				ODataType: to.Ptr(armmonitor.OdatatypeMicrosoftAzureMonitorPromQLCriteria),
	// 			},
	// 			Enabled: to.Ptr(true),
	// 			EvaluationFrequency: to.Ptr("PT1M"),
	// 			ResolveConfiguration: &armmonitor.ResolveConfiguration{
	// 				AutoResolved: to.Ptr(true),
	// 				TimeToResolve: to.Ptr("PT10M"),
	// 			},
	// 			Scopes: []*string{
	// 				to.Ptr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/microsoft.monitor/accounts/gigwadme"),
	// 			},
	// 			Severity: to.Ptr[int32](3),
	// 		},
	// 		Tags: map[string]*string{
	// 			"hidden-link:/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/leoalerttest": to.Ptr("Resource"),
	// 		},
	// 	},
	// }
}
