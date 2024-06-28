package armappcomplianceautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_CreateOrUpdate.json
func ExampleReportClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReportClient().BeginCreateOrUpdate(ctx, "testReportName", armappcomplianceautomation.ReportResource{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReportResource = armappcomplianceautomation.ReportResource{
	// 	Name: to.Ptr("testReportName"),
	// 	Type: to.Ptr("Microsfot.AppComplianceAutomation/reports"),
	// 	ID: to.Ptr("/provider/Microsfot.AppComplianceAutomation/reports/testReportName"),
	// 	SystemData: &armappcomplianceautomation.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
	// 		CreatedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		CreatedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		LastModifiedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
	// 	},
	// 	Properties: &armappcomplianceautomation.ReportProperties{
	// 		CertRecords: []*armappcomplianceautomation.CertSyncRecord{
	// 			{
	// 				CertificationStatus: to.Ptr("CertIngestion"),
	// 				Controls: []*armappcomplianceautomation.ControlSyncRecord{
	// 					{
	// 						ControlID: to.Ptr("Operational_Security_10"),
	// 						ControlStatus: to.Ptr("Approved"),
	// 				}},
	// 				IngestionStatus: to.Ptr("EvidenceResubmitted"),
	// 				OfferGUID: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 		}},
	// 		ComplianceStatus: &armappcomplianceautomation.ReportComplianceStatus{
	// 			M365: &armappcomplianceautomation.OverviewStatus{
	// 				FailedCount: to.Ptr[int32](0),
	// 				ManualCount: to.Ptr[int32](0),
	// 				PassedCount: to.Ptr[int32](0),
	// 			},
	// 		},
	// 		LastTriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T05:00:00.000Z"); return t}()),
	// 		NextTriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T05:00:00.000Z"); return t}()),
	// 		OfferGUID: to.Ptr("00000000-0000-0000-0000-000000000001,00000000-0000-0000-0000-000000000002"),
	// 		ProvisioningState: to.Ptr(armappcomplianceautomation.ProvisioningStateSucceeded),
	// 		Resources: []*armappcomplianceautomation.ResourceMetadata{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService"),
	// 				ResourceOrigin: to.Ptr(armappcomplianceautomation.ResourceOriginAzure),
	// 				ResourceType: to.Ptr("Microsoft.SignalRService/SignalR"),
	// 			},
	// 			{
	// 				AccountID: to.Ptr("000000000000"),
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/acat-aws/providers/microsoft.security/securityconnectors/acatawsconnector/securityentitydata/aws-iam-user-testuser"),
	// 				ResourceOrigin: to.Ptr(armappcomplianceautomation.ResourceOriginAWS),
	// 				ResourceType: to.Ptr("iam.user"),
	// 		}},
	// 		Status: to.Ptr(armappcomplianceautomation.ReportStatusActive),
	// 		StorageInfo: &armappcomplianceautomation.StorageInfo{
	// 			AccountName: to.Ptr("testStorageAccount"),
	// 			Location: to.Ptr("East US"),
	// 			ResourceGroup: to.Ptr("testResourceGroup"),
	// 			SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 		Subscriptions: []*string{
	// 			to.Ptr("00000000-0000-0000-0000-000000000000")},
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			TimeZone: to.Ptr("GMT Standard Time"),
	// 			TriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T05:00:00.000Z"); return t}()),
	// 		},
	// 	}
}
