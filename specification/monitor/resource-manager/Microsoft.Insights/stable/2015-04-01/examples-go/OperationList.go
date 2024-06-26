package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2015-04-01/examples/OperationList.json
func ExampleOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsClient().List(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationListResult = armmonitor.OperationListResult{
	// 	Value: []*armmonitor.Operation{
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/Operations/Read"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Operations read"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Operations"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/MetricDefinitions/Read"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Metric definitions read"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Metric Definitions"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/Metrics/Read"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Metrics read"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Metrics"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/AlertRules/Write"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Alert Rule write"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Alert Rules"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/AlertRules/Delete"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Alert Rule delete"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Alert Rules"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/AlertRules/Read"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Alert Rule read"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Alert Rules"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/MetricAlerts/Write"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Metric alert write"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Metric alerts"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/MetricAlerts/Delete"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Metric alert delete"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Metric alerts"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/MetricAlerts/Read"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Metric alert read"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Metric alerts"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/AutoscaleSettings/Write"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Autoscale Setting write"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Autoscale"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/AutoscaleSettings/Delete"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Autoscale Setting delete"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Autoscale"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/AutoscaleSettings/Read"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Autoscale Setting read"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Autoscale"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/AlertRules/Incidents/Read"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Alert Rule Incidents read"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Alert Rule Incident resource"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/AutoscaleSettings/providers/Microsoft.Insights/MetricDefinitions/Read"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Metric definitions read"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Metric Definitions"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/ActionGroups/Write"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Action group write"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Action groups"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/ActionGroups/Delete"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Action group delete"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Action groups"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/ActionGroups/Read"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Action group read"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Action groups"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/ActivityLogAlerts/Write"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Activity log alert read"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Activity log alert"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/ActivityLogAlerts/Delete"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Activity log alert delete"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Activity log alert"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/ActivityLogAlerts/Read"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Activity log alert read"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Activity log alert"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/ActivityLogAlerts/Activated/Action"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Activity Log Alert Activated"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Activity Log Alert"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/EventCategories/Read"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Event category read"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Event category"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/eventtypes/values/Read"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Event types management values read"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Events"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/eventtypes/digestevents/Read"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Event types management digest read"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Digest events"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/DiagnosticSettings/Write"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Diagnostic settings write"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Diagnostic settings"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/DiagnosticSettings/Delete"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Diagnostic settings delete"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Diagnostic settings"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/DiagnosticSettings/Read"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Diagnostic settings read"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Diagnostic settings"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/ExtendedDiagnosticSettings/Write"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Extended Diagnostic settings write"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Extended Diagnostic settings"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/ExtendedDiagnosticSettings/Delete"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Extended Diagnostic settings delete"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Extended Diagnostic settings"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/ExtendedDiagnosticSettings/Read"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Extended Diagnostic settings read"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Extended Diagnostic settings"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/LogProfiles/Write"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Log profile write"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Log Profiles"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/LogProfiles/Delete"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Log profile delete"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Log Profiles"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/LogProfiles/Read"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Log profile read"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Log Profiles"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/LogDefinitions/Read"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Log Definitions read"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Log Definitions"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/AutoscaleSettings/Scaleup/Action"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Autoscale scale up operation"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Autoscale"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/AutoscaleSettings/Scaledown/Action"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Autoscale scale down operation"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Autoscale"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/AlertRules/Activated/Action"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Alert Rule activated"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Alert Rules"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/AlertRules/Resolved/Action"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Alert Rule resolved"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Alert Rules"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/AlertRules/Throttled/Action"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Alert Rule throttled"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Alert Rules"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/Register/Action"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Register Microsoft.Insights"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Microsoft.Insights"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/Components/Write"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Application insights component write"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Application insights components"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/Components/Delete"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Application insights component delete"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Application insights components"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/Components/Read"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Application insights component read"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Application insights components"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/Webtests/Write"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Webtest write"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Web tests"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/Webtests/Delete"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Webtest delete"),
	// 				Provider: to.Ptr("Microsoft Monitoring Insights"),
	// 				Resource: to.Ptr("Web tests"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/Workbooks/Read"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Workbooks read"),
	// 				Provider: to.Ptr("Microsoft Application Insights"),
	// 				Resource: to.Ptr("Workbooks"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/Workbooks/Write"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Workbooks write"),
	// 				Provider: to.Ptr("Microsoft Application Insights"),
	// 				Resource: to.Ptr("Workbooks"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/Workbooks/Delete"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Workbooks delete"),
	// 				Provider: to.Ptr("Microsoft Application Insights"),
	// 				Resource: to.Ptr("Workbooks"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Insights/Workbooks/Read"),
	// 			Display: &armmonitor.OperationDisplay{
	// 				Operation: to.Ptr("Workbooks read"),
	// 				Provider: to.Ptr("Microsoft Application Insights"),
	// 				Resource: to.Ptr("Workbooks"),
	// 			},
	// 	}},
	// }
}
