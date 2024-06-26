package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6d2438481021a94793b07b226df06d5f3c61d51d/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2021-07-22-preview/examples/listPrometheusRuleGroups.json
func ExamplePrometheusRuleGroupsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrometheusRuleGroupsClient().NewListByResourceGroupPager("promResourceGroup", nil)
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
		// page.PrometheusRuleGroupResourceCollection = armalertsmanagement.PrometheusRuleGroupResourceCollection{
		// 	Value: []*armalertsmanagement.PrometheusRuleGroupResource{
		// 		{
		// 			Type: to.Ptr("Microsoft.AlertsManagement/prometheusRuleGroups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/promResourceGroup/providers/Microsoft.AlertsManagement/prometheusRuleGroups/myPrometheusRuleGroup"),
		// 			SystemData: &armalertsmanagement.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T22:05:09.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("abc@microsoft.com"),
		// 				CreatedByType: to.Ptr(armalertsmanagement.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-12T22:05:09.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("xyz@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armalertsmanagement.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armalertsmanagement.PrometheusRuleGroupProperties{
		// 				Description: to.Ptr("This is the description of the first rule group"),
		// 				Rules: []*armalertsmanagement.PrometheusRule{
		// 					{
		// 						Expression: to.Ptr("histogram_quantile(0.99, sum(rate(jobs_duration_seconds_bucket{service=\"billing-processing\"}[5m])) by (job_type))"),
		// 						Record: to.Ptr("job_type:billing_jobs_duration_seconds:99p5m"),
		// 					},
		// 					{
		// 						Actions: []*armalertsmanagement.PrometheusRuleGroupAction{
		// 							{
		// 								ActionGroupID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2"),
		// 								ActionProperties: map[string]*string{
		// 									"key11": to.Ptr("value11"),
		// 									"key12": to.Ptr("value12"),
		// 								},
		// 						}},
		// 						Alert: to.Ptr("Billing_Processing_Very_Slow"),
		// 						Expression: to.Ptr("job_type:billing_jobs_duration_seconds:99p5m > 30"),
		// 						For: to.Ptr("5m"),
		// 						Labels: map[string]*string{
		// 							"team": to.Ptr("prod"),
		// 						},
		// 						ResolveConfiguration: &armalertsmanagement.PrometheusRuleResolveConfiguration{
		// 							AutoResolved: to.Ptr(true),
		// 							TimeToResolve: to.Ptr("10m"),
		// 						},
		// 						Severity: to.Ptr[int32](2),
		// 				}},
		// 				Scopes: []*string{
		// 					to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/promResourceGroup/providers/microsoft.monitor/accounts/myMonitoringAccount")},
		// 				},
		// 		}},
		// 	}
	}
}
