package armappcomplianceautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Snapshot_List.json
func ExampleSnapshotClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSnapshotClient().NewListPager("testReportName", &armappcomplianceautomation.SnapshotClientListOptions{SkipToken: to.Ptr("1"),
		Top:                   to.Ptr[int32](100),
		Select:                nil,
		Filter:                nil,
		Orderby:               nil,
		OfferGUID:             to.Ptr("00000000-0000-0000-0000-000000000001"),
		ReportCreatorTenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
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
		// page.SnapshotResourceListResult = armappcomplianceautomation.SnapshotResourceListResult{
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
		// 								CategoryStatus: to.Ptr(armappcomplianceautomation.CategoryStatusPassed),
		// 								ControlFamilies: []*armappcomplianceautomation.ControlFamily{
		// 									{
		// 										ControlFamilyName: to.Ptr("Incident Response"),
		// 										ControlFamilyStatus: to.Ptr(armappcomplianceautomation.ControlFamilyStatusPassed),
		// 										Controls: []*armappcomplianceautomation.Control{
		// 											{
		// 												ControlDescription: to.Ptr("Provide demonstrable evidence that all member of the incident response team have completed annual training or a table top exercise"),
		// 												ControlDescriptionHyperLink: to.Ptr("https://aka.ms/acat/m365cert/operational/control73"),
		// 												ControlFullName: to.Ptr("Provide demonstrable evidence that all member of the incident response team have completed annual training or a table top exercise"),
		// 												ControlID: to.Ptr("Operational_Security_75"),
		// 												ControlName: to.Ptr("Provide demonstrable evidence that all member of the incident response team have completed annual training or a table top exercise"),
		// 												ControlStatus: to.Ptr(armappcomplianceautomation.ControlStatusPassed),
		// 												Responsibilities: []*armappcomplianceautomation.Responsibility{
		// 													{
		// 														EvidenceFiles: []*string{
		// 															to.Ptr("https://management.azure.com/providers/Microsoft.AppComplianceAutomation/reports/reportABC/fileName?api-version=2024-06-27")},
		// 															FailedResourceCount: to.Ptr[int32](0),
		// 															Guidance: to.Ptr("Please upload the screen capture file to ACAT service."),
		// 															Justification: to.Ptr("Here is my evidence files"),
		// 															RecommendationList: []*armappcomplianceautomation.Recommendation{
		// 																{
		// 																	RecommendationID: to.Ptr("failed_reason_1"),
		// 																	RecommendationShortName: to.Ptr("Invalid TLS Config"),
		// 																	RecommendationSolutions: []*armappcomplianceautomation.RecommendationSolution{
		// 																		{
		// 																			IsRecommendSolution: to.Ptr(armappcomplianceautomation.IsRecommendSolutionTrue),
		// 																			RecommendationSolutionContent: to.Ptr("Setting minimal TLS version to 1.2 improves security by ensuring your SQL Managed Instance can only be accessed from clients using TLS 1.2. Using versions of TLS less than 1.2 is not recommended since they have well documented security vulnerabilities"),
		// 																			RecommendationSolutionIndex: to.Ptr("1"),
		// 																	}},
		// 																},
		// 																{
		// 																	RecommendationID: to.Ptr("failed_reason_2"),
		// 																	RecommendationShortName: to.Ptr("Invalid AWS TLS Config"),
		// 																	RecommendationSolutions: []*armappcomplianceautomation.RecommendationSolution{
		// 																		{
		// 																			IsRecommendSolution: to.Ptr(armappcomplianceautomation.IsRecommendSolutionTrue),
		// 																			RecommendationSolutionContent: to.Ptr("Open the AWS related service, and set its TLS version to 1.2 or higher version."),
		// 																			RecommendationSolutionIndex: to.Ptr("1"),
		// 																	}},
		// 															}},
		// 															ResourceList: []*armappcomplianceautomation.ResponsibilityResource{
		// 																{
		// 																	RecommendationIDs: []*string{
		// 																		to.Ptr("failed_reason_1")},
		// 																		ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService"),
		// 																		ResourceOrigin: to.Ptr(armappcomplianceautomation.ResourceOriginAzure),
		// 																		ResourceStatus: to.Ptr(armappcomplianceautomation.ResourceStatusUnhealthy),
		// 																		ResourceStatusChangeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-12T16:17:20.150Z"); return t}()),
		// 																		ResourceType: to.Ptr("Microsoft.SignalRService/SignalR"),
		// 																	},
		// 																	{
		// 																		AccountID: to.Ptr("000000000000"),
		// 																		RecommendationIDs: []*string{
		// 																			to.Ptr("failed_reason_2")},
		// 																			ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/acat-aws/providers/microsoft.security/securityconnectors/acatawsconnector/securityentitydata/aws-iam-user-testuser"),
		// 																			ResourceOrigin: to.Ptr(armappcomplianceautomation.ResourceOriginAWS),
		// 																			ResourceStatus: to.Ptr(armappcomplianceautomation.ResourceStatusUnhealthy),
		// 																			ResourceStatusChangeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-12T16:17:20.150Z"); return t}()),
		// 																			ResourceType: to.Ptr("iam.user"),
		// 																	}},
		// 																	ResponsibilityDescription: to.Ptr("Restrict access to the Kubernetes Service Management API by granting API access only to IP addresses in specific ranges. It is recommended to limit access to authorized IP ranges to ensure that only applications from allowed networks can access the cluster."),
		// 																	ResponsibilityEnvironment: to.Ptr(armappcomplianceautomation.ResponsibilityEnvironmentAzure),
		// 																	ResponsibilityID: to.Ptr("authorized_ip_ranges_should_be_defined_on_kubernetes_services"),
		// 																	ResponsibilitySeverity: to.Ptr(armappcomplianceautomation.ResponsibilitySeverityHigh),
		// 																	ResponsibilityStatus: to.Ptr(armappcomplianceautomation.ResponsibilityStatusPassed),
		// 																	ResponsibilityTitle: to.Ptr("Authorized IP ranges should be defined on Kubernetes Services"),
		// 																	ResponsibilityType: to.Ptr(armappcomplianceautomation.ResponsibilityTypeAutomated),
		// 																	TotalResourceCount: to.Ptr[int32](1),
		// 															}},
		// 													}},
		// 											}},
		// 									}},
		// 									ComplianceName: to.Ptr("M365"),
		// 							}},
		// 							CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T15:33:59.160Z"); return t}()),
		// 							ProvisioningState: to.Ptr(armappcomplianceautomation.ProvisioningStateSucceeded),
		// 							ReportProperties: &armappcomplianceautomation.ReportProperties{
		// 								CertRecords: []*armappcomplianceautomation.CertSyncRecord{
		// 									{
		// 										CertificationStatus: to.Ptr("CertIngestion"),
		// 										Controls: []*armappcomplianceautomation.ControlSyncRecord{
		// 											{
		// 												ControlID: to.Ptr("Operational_Security_10"),
		// 												ControlStatus: to.Ptr("Approved"),
		// 										}},
		// 										IngestionStatus: to.Ptr("EvidenceResubmitted"),
		// 										OfferGUID: to.Ptr("00000000-0000-0000-0000-000000000001"),
		// 								}},
		// 								ComplianceStatus: &armappcomplianceautomation.ReportComplianceStatus{
		// 									M365: &armappcomplianceautomation.OverviewStatus{
		// 										FailedCount: to.Ptr[int32](0),
		// 										ManualCount: to.Ptr[int32](0),
		// 										NotApplicableCount: to.Ptr[int32](0),
		// 										PassedCount: to.Ptr[int32](0),
		// 										PendingCount: to.Ptr[int32](0),
		// 									},
		// 								},
		// 								Errors: []*string{
		// 								},
		// 								LastTriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T15:00:00.000Z"); return t}()),
		// 								NextTriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T15:00:00.000Z"); return t}()),
		// 								OfferGUID: to.Ptr("00000000-0000-0000-0000-000000000001,00000000-0000-0000-0000-000000000002"),
		// 								ProvisioningState: to.Ptr(armappcomplianceautomation.ProvisioningStateSucceeded),
		// 								Resources: []*armappcomplianceautomation.ResourceMetadata{
		// 									{
		// 										ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService"),
		// 										ResourceOrigin: to.Ptr(armappcomplianceautomation.ResourceOriginAzure),
		// 										ResourceType: to.Ptr("Microsoft.SignalRService/SignalR"),
		// 									},
		// 									{
		// 										AccountID: to.Ptr("000000000000"),
		// 										ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/acat-aws/providers/microsoft.security/securityconnectors/acatawsconnector/securityentitydata/aws-iam-user-testuser"),
		// 										ResourceOrigin: to.Ptr(armappcomplianceautomation.ResourceOriginAWS),
		// 										ResourceType: to.Ptr("iam.user"),
		// 								}},
		// 								Status: to.Ptr(armappcomplianceautomation.ReportStatusActive),
		// 								StorageInfo: &armappcomplianceautomation.StorageInfo{
		// 									AccountName: to.Ptr("testStorageAccount"),
		// 									Location: to.Ptr("East US"),
		// 									ResourceGroup: to.Ptr("testResourceGroup"),
		// 									SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 								},
		// 								Subscriptions: []*string{
		// 									to.Ptr("00000000-0000-0000-0000-000000000000")},
		// 									TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 									TimeZone: to.Ptr("GMT Standard Time"),
		// 									TriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T15:00:00.000Z"); return t}()),
		// 								},
		// 								ReportSystemData: &armappcomplianceautomation.SystemData{
		// 									CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
		// 									CreatedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 									CreatedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
		// 									LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
		// 									LastModifiedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 									LastModifiedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
		// 								},
		// 								SnapshotName: to.Ptr("testSnapshot"),
		// 							},
		// 					}},
		// 				}
	}
}
