package armappcomplianceautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82ea406b73d671269217053d7ef336450d860345/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/preview/2022-11-16-preview/examples/Snapshots_List.json
func ExampleSnapshotsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSnapshotsClient().NewListPager("testReportName", &armappcomplianceautomation.SnapshotsClientListOptions{SkipToken: to.Ptr("1"),
		Top:                   to.Ptr[int32](100),
		Select:                nil,
		ReportCreatorTenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		OfferGUID:             to.Ptr("00000000-0000-0000-0000-000000000000"),
	})
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
		// page.SnapshotResourceList = armappcomplianceautomation.SnapshotResourceList{
		// 	Value: []*armappcomplianceautomation.SnapshotResource{
		// 		{
		// 			Name: to.Ptr("testSnapshot"),
		// 			Type: to.Ptr("Microsfot.AppComplianceAutomation/reports/snapshots"),
		// 			ID: to.Ptr("/provider/Microsfot.AppComplianceAutomation/reports/testReportName/snapshots/testSnapshot"),
		// 			SystemData: &armappcomplianceautomation.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
		// 				CreatedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				CreatedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				LastModifiedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
		// 			},
		// 			Properties: &armappcomplianceautomation.SnapshotProperties{
		// 				ComplianceResults: []*armappcomplianceautomation.ComplianceResult{
		// 					{
		// 						Categories: []*armappcomplianceautomation.Category{
		// 							{
		// 								CategoryName: to.Ptr("Operational Security"),
		// 								CategoryStatus: to.Ptr(armappcomplianceautomation.CategoryStatusHealthy),
		// 								CategoryType: to.Ptr(armappcomplianceautomation.CategoryTypePartiallyAutomated),
		// 								ControlFamilies: []*armappcomplianceautomation.ControlFamily{
		// 									{
		// 										Controls: []*armappcomplianceautomation.Control{
		// 											{
		// 												Assessments: []*armappcomplianceautomation.Assessment{
		// 													{
		// 														Name: to.Ptr("AssessmentName"),
		// 														Description: to.Ptr("Assessment Description"),
		// 														IsPass: to.Ptr(armappcomplianceautomation.IsPassTrue),
		// 														PolicyID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 														Remediation: to.Ptr("Remediation"),
		// 														ResourceList: []*armappcomplianceautomation.AssessmentResource{
		// 															{
		// 																Reason: to.Ptr("N/A"),
		// 																ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService"),
		// 																ResourceStatus: to.Ptr(armappcomplianceautomation.ResourceStatusHealthy),
		// 																StatusChangeDate: to.Ptr("2022-03-04"),
		// 														}},
		// 														Severity: to.Ptr(armappcomplianceautomation.AssessmentSeverityHigh),
		// 												}},
		// 												ControlDescription: to.Ptr("Provide demonstrable evidence that all member of the incident response team have completed annual training or a table top exercise"),
		// 												ControlDescriptionHyperLink: to.Ptr("https://doc.microsoft.com"),
		// 												ControlFullName: to.Ptr("Provide demonstrable evidence that all member of the incident response team have completed annual training or a table top exercise"),
		// 												ControlID: to.Ptr("Operational Security#75"),
		// 												ControlShortName: to.Ptr("Provide demonstrable evidence that all member of the incident response team have completed annual training or a table top exercise"),
		// 												ControlStatus: to.Ptr(armappcomplianceautomation.ControlStatusPassed),
		// 												ControlType: to.Ptr(armappcomplianceautomation.ControlTypeManual),
		// 										}},
		// 										FamilyName: to.Ptr("Incident Response"),
		// 										FamilyStatus: to.Ptr(armappcomplianceautomation.ControlFamilyStatusHealthy),
		// 										FamilyType: to.Ptr(armappcomplianceautomation.ControlFamilyTypePartiallyAutomated),
		// 								}},
		// 						}},
		// 						ComplianceName: to.Ptr("M365"),
		// 				}},
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T15:33:59.160Z"); return t}()),
		// 				ID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ProvisioningState: to.Ptr(armappcomplianceautomation.ProvisioningStateSucceeded),
		// 				ReportProperties: &armappcomplianceautomation.ReportProperties{
		// 					ComplianceStatus: &armappcomplianceautomation.ReportComplianceStatus{
		// 						M365: &armappcomplianceautomation.OverviewStatus{
		// 							FailedCount: to.Ptr[int32](0),
		// 							ManualCount: to.Ptr[int32](0),
		// 							PassedCount: to.Ptr[int32](0),
		// 						},
		// 					},
		// 					ID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					LastTriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T15:33:59.160Z"); return t}()),
		// 					NextTriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T15:33:59.160Z"); return t}()),
		// 					OfferGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					ProvisioningState: to.Ptr(armappcomplianceautomation.ProvisioningStateSucceeded),
		// 					ReportName: to.Ptr("testReportName"),
		// 					Resources: []*armappcomplianceautomation.ResourceMetadata{
		// 						{
		// 							ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService"),
		// 							Tags: map[string]*string{
		// 								"key1": to.Ptr("value1"),
		// 							},
		// 					}},
		// 					Status: to.Ptr(armappcomplianceautomation.ReportStatusActive),
		// 					Subscriptions: []*string{
		// 						to.Ptr("00000000-0000-0000-0000-000000000000")},
		// 						TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						TimeZone: to.Ptr("GMT Standard Time"),
		// 						TriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T15:33:59.160Z"); return t}()),
		// 					},
		// 					ReportSystemData: &armappcomplianceautomation.SystemData{
		// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
		// 						CreatedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						CreatedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
		// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
		// 						LastModifiedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						LastModifiedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
		// 					},
		// 					SnapshotName: to.Ptr("testSnapshot"),
		// 				},
		// 		}},
		// 	}
	}
}
