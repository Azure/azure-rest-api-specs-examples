package armprometheusrulegroups_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/prometheusrulegroups/armprometheusrulegroups"
)

// Generated from example definition: 2023-03-01/createOrUpdatePrometheusRuleGroup.json
func ExampleClient_CreateOrUpdate_createOrUpdateAPrometheusRuleGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprometheusrulegroups.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().CreateOrUpdate(ctx, "promResourceGroup", "myPrometheusRuleGroup", armprometheusrulegroups.PrometheusRuleGroupResource{
		Location: to.Ptr("East US"),
		Properties: &armprometheusrulegroups.PrometheusRuleGroupProperties{
			Description: to.Ptr("This is the description of the following rule group"),
			ClusterName: to.Ptr("myClusterName"),
			Enabled:     to.Ptr(true),
			Interval:    to.Ptr("PT10M"),
			Rules: []*armprometheusrulegroups.PrometheusRule{
				{
					Expression: to.Ptr("histogram_quantile(0.99, sum(rate(jobs_duration_seconds_bucket{service=\"billing-processing\"}[5m])) by (job_type))"),
					Labels: map[string]*string{
						"team": to.Ptr("prod"),
					},
					Record: to.Ptr("job_type:billing_jobs_duration_seconds:99p5m"),
				},
				{
					Actions: []*armprometheusrulegroups.PrometheusRuleGroupAction{
						{
							ActionGroupID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourcegroups/myrg/providers/microsoft.insights/actiongroups/myactiongroup"),
							ActionProperties: map[string]*string{
								"key11": to.Ptr("value11"),
								"key12": to.Ptr("value12"),
							},
						},
						{
							ActionGroupID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourcegroups/myrg/providers/microsoft.insights/actiongroups/myotheractiongroup"),
							ActionProperties: map[string]*string{
								"key21": to.Ptr("value21"),
								"key22": to.Ptr("value22"),
							},
						},
					},
					Alert: to.Ptr("Billing_Processing_Very_Slow"),
					Annotations: map[string]*string{
						"annotationName1": to.Ptr("annotationValue1"),
					},
					Enabled:    to.Ptr(true),
					Expression: to.Ptr("job_type:billing_jobs_duration_seconds:99p5m > 30"),
					For:        to.Ptr("PT5M"),
					Labels: map[string]*string{
						"team": to.Ptr("prod"),
					},
					ResolveConfiguration: &armprometheusrulegroups.PrometheusRuleResolveConfiguration{
						AutoResolved:  to.Ptr(true),
						TimeToResolve: to.Ptr("PT10M"),
					},
					Severity: to.Ptr[int32](2),
				},
			},
			Scopes: []*string{
				to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroup/providers/microsoft.monitor/accounts/myAzureMonitorWorkspace"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armprometheusrulegroups.ClientCreateOrUpdateResponse{
	// 	PrometheusRuleGroupResource: &armprometheusrulegroups.PrometheusRuleGroupResource{
	// 		Type: to.Ptr("Microsoft.AlertsManagement/prometheusRuleGroups"),
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/promResourceGroup/providers/Microsoft.AlertsManagement/prometheusRuleGroups/myPrometheusRuleGroup"),
	// 		Location: to.Ptr("East US"),
	// 		Properties: &armprometheusrulegroups.PrometheusRuleGroupProperties{
	// 			Description: to.Ptr("This is the description of the following rule group"),
	// 			ClusterName: to.Ptr("myClusterName"),
	// 			Enabled: to.Ptr(true),
	// 			Interval: to.Ptr("PT10M"),
	// 			Rules: []*armprometheusrulegroups.PrometheusRule{
	// 				{
	// 					Expression: to.Ptr("histogram_quantile(0.99, sum(rate(jobs_duration_seconds_bucket{service=\"billing-processing\"}[5m])) by (job_type))"),
	// 					Labels: map[string]*string{
	// 						"team": to.Ptr("prod"),
	// 					},
	// 					Record: to.Ptr("job_type:billing_jobs_duration_seconds:99p5m"),
	// 				},
	// 				{
	// 					Actions: []*armprometheusrulegroups.PrometheusRuleGroupAction{
	// 						{
	// 							ActionGroupID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourcegroups/myrg/providers/microsoft.insights/actiongroups/myactiongroup"),
	// 							ActionProperties: map[string]*string{
	// 								"key11": to.Ptr("value11"),
	// 								"key12": to.Ptr("value12"),
	// 							},
	// 						},
	// 						{
	// 							ActionGroupID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourcegroups/myrg/providers/microsoft.insights/actiongroups/myotheractiongroup"),
	// 							ActionProperties: map[string]*string{
	// 								"key21": to.Ptr("value21"),
	// 								"key22": to.Ptr("value22"),
	// 							},
	// 						},
	// 					},
	// 					Alert: to.Ptr("Billing_Processing_Very_Slow"),
	// 					Annotations: map[string]*string{
	// 						"annotationName1": to.Ptr("annotationValue1"),
	// 					},
	// 					Enabled: to.Ptr(true),
	// 					Expression: to.Ptr("job_type:billing_jobs_duration_seconds:99p5m > 30"),
	// 					For: to.Ptr("PT5M"),
	// 					Labels: map[string]*string{
	// 						"team": to.Ptr("prod"),
	// 					},
	// 					ResolveConfiguration: &armprometheusrulegroups.PrometheusRuleResolveConfiguration{
	// 						AutoResolved: to.Ptr(true),
	// 						TimeToResolve: to.Ptr("PT10M"),
	// 					},
	// 					Severity: to.Ptr[int32](2),
	// 				},
	// 			},
	// 			Scopes: []*string{
	// 				to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroup/providers/microsoft.monitor/accounts/myAzureMonitorWorkspace"),
	// 			},
	// 		},
	// 	},
	// }
}
