package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/Diagnostics_ExecuteSiteAnalysisSlot.json
func ExampleDiagnosticsClient_ExecuteSiteAnalysisSlot_executeSiteSlotAnalysis() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDiagnosticsClient().ExecuteSiteAnalysisSlot(ctx, "Sample-WestUSResourceGroup", "SampleApp", "availability", "apprestartanalyses", "staging", &armappservice.DiagnosticsClientExecuteSiteAnalysisSlotOptions{StartTime: nil,
		EndTime:   nil,
		TimeGrain: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiagnosticAnalysis = armappservice.DiagnosticAnalysis{
	// 	Name: to.Ptr("apprestartanalysis"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/Sample-WestUSResourceGroup/providers/Microsoft.Web/sites/SampleApp/slots/staging/diagnostics/availability/analyses/apprestartanalyses"),
	// 	Properties: &armappservice.DiagnosticAnalysisProperties{
	// 		AbnormalTimePeriods: []*armappservice.AbnormalTimePeriod{
	// 			{
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-06T22:50:00.000Z"); return t}()),
	// 				Events: []*armappservice.DetectorAbnormalTimePeriod{
	// 					{
	// 						Type: to.Ptr(armappservice.IssueTypeServiceIncident),
	// 						EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-06T22:21:48.000Z"); return t}()),
	// 						Message: to.Ptr("Your application process was restarted as application environment variables changed. This can most likely occur due to update in app settings or swap operation. This event occurred multiple times during the day."),
	// 						MetaData: [][]*armappservice.NameValuePair{
	// 							[]*armappservice.NameValuePair{
	// 								{
	// 									Name: to.Ptr("feature"),
	// 									Value: to.Ptr("auditlogs"),
	// 								},
	// 								{
	// 									Name: to.Ptr("displayedName"),
	// 									Value: to.Ptr("Check Audit Logs"),
	// 						}}},
	// 						Priority: to.Ptr[float64](0),
	// 						Solutions: []*armappservice.Solution{
	// 						},
	// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-06T22:21:50.000Z"); return t}()),
	// 				}},
	// 				Solutions: []*armappservice.Solution{
	// 				},
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-05T22:50:00.000Z"); return t}()),
	// 		}},
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-06T22:50:00.000Z"); return t}()),
	// 		NonCorrelatedDetectors: []*armappservice.DetectorDefinition{
	// 		},
	// 		Payload: []*armappservice.AnalysisData{
	// 			{
	// 				Data: [][]*armappservice.NameValuePair{
	// 				},
	// 				Metrics: []*armappservice.DiagnosticMetricSet{
	// 					{
	// 						Name: to.Ptr("All Application Stop Events"),
	// 						EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-06T22:50:00.000Z"); return t}()),
	// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-05T22:50:00.000Z"); return t}()),
	// 						TimeGrain: to.Ptr("00:05:00"),
	// 						Values: []*armappservice.DiagnosticMetricSample{
	// 							{
	// 								RoleInstance: to.Ptr("RD00155D3C15BE"),
	// 								Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-06T00:00:00.000Z"); return t}()),
	// 								Total: to.Ptr[float64](2),
	// 							},
	// 							{
	// 								IsAggregated: to.Ptr(true),
	// 								Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-06T00:00:00.000Z"); return t}()),
	// 								Total: to.Ptr[float64](2),
	// 							},
	// 							{
	// 								RoleInstance: to.Ptr("RD00155D3C15C1"),
	// 								Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-06T00:10:00.000Z"); return t}()),
	// 								Total: to.Ptr[float64](2),
	// 							},
	// 							{
	// 								RoleInstance: to.Ptr("RD00155D3BE0FB"),
	// 								Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-06T00:10:00.000Z"); return t}()),
	// 								Total: to.Ptr[float64](2),
	// 							},
	// 							{
	// 								IsAggregated: to.Ptr(true),
	// 								Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-06T00:10:00.000Z"); return t}()),
	// 								Total: to.Ptr[float64](4),
	// 						}},
	// 					},
	// 					{
	// 						Name: to.Ptr("User Events"),
	// 						EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-06T22:50:00.000Z"); return t}()),
	// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-05T22:50:00.000Z"); return t}()),
	// 						TimeGrain: to.Ptr("00:05:00"),
	// 						Unit: to.Ptr(""),
	// 						Values: []*armappservice.DiagnosticMetricSample{
	// 							{
	// 								IsAggregated: to.Ptr(true),
	// 								Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-06T22:20:00.000Z"); return t}()),
	// 								Total: to.Ptr[float64](3),
	// 							},
	// 							{
	// 								RoleInstance: to.Ptr("RD00155D3C09FC"),
	// 								Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-06T22:20:00.000Z"); return t}()),
	// 								Total: to.Ptr[float64](1),
	// 							},
	// 							{
	// 								RoleInstance: to.Ptr("RD00155D3C2ADC"),
	// 								Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-06T22:20:00.000Z"); return t}()),
	// 								Total: to.Ptr[float64](1),
	// 							},
	// 							{
	// 								RoleInstance: to.Ptr("RD00155D3C214E"),
	// 								Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-06T22:20:00.000Z"); return t}()),
	// 								Total: to.Ptr[float64](1),
	// 						}},
	// 				}},
	// 				Source: to.Ptr("workerprocessrecycle"),
	// 		}},
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-05T22:50:00.000Z"); return t}()),
	// 	},
	// }
}
